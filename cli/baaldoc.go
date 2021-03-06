package main

import (
	"flag"
	"fmt"
	"github.com/shogo82148/go-baal"
	"gopkg.in/fsnotify.v1"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type BaalSet struct {
	Files            map[string][]baal.Namespace
	Entities         map[string]baal.Entity
	EntityFiles      map[string]string
	DeclarationFiles map[string]string
}

type RenderContext struct {
	namespace string
	nest      int
}

var baalSet BaalSet
var mutex sync.RWMutex

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "listen port")
	flag.IntVar(&port, "p", 8080, "listen port")
	flag.Parse()
	baalDir := flag.Arg(0)
	baalSet = loadBaalFiles(baalDir)

	watch(baalDir)

	http.HandleFunc("/", handler)
	http.HandleFunc("/static/style.css", handlerStyle)
	http.HandleFunc("/static/faced.js", handlerFacedJavascript)
	http.HandleFunc("/declarations", handlerDeclarations)
	http.HandleFunc("/information/", handlerInformation)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func watch(baalDir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(
		baalDir,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				return nil
			}
			if strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}
			if err := watcher.Add(path); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	)

	go func() {
		for {
			ev := <-watcher.Events
			if ev.Op == fsnotify.Chmod {
				continue
			}
			loadBaalFiles(baalDir)
		}
	}()
}

func loadBaalFiles(baalDir string) BaalSet {
	files := map[string][]baal.Namespace{}
	entities := map[string]baal.Entity{}
	entityFiles := map[string]string{}
	declarationFiles := map[string]string{}
	filepath.Walk(
		baalDir,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, ".faced") || strings.HasPrefix(info.Name(), ".") {
				return nil
			}
			rel, _ := filepath.Rel(baalDir, path)

			content, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			scanner := new(baal.Scanner)
			scanner.Init(string(content))
			namespaces, _ := baal.Parse(scanner)
			files[rel] = namespaces
			for _, namespace := range namespaces {
				name := strings.Join(namespace.Name, ".")
				for _, declaration := range namespace.Declarations {
					switch declaration := declaration.(type) {
					case baal.Service:
						declarationFiles[name+"."+declaration.Name] = rel
					case baal.Entity:
						entities[name+"."+declaration.Name] = declaration
						entityFiles[name+"."+declaration.Name] = rel
						declarationFiles[name+"."+declaration.Name] = rel
					}
				}
			}
			return nil
		},
	)

	mutex.Lock()
	defer mutex.Unlock()
	baalSet = BaalSet{
		Files:            files,
		Entities:         entities,
		EntityFiles:      entityFiles,
		DeclarationFiles: declarationFiles,
	}
	return baalSet
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		handlerIndex(w, r)
	} else if file, ok := baalSet.Files[path]; ok {
		handlerFaced(w, r, path, file)
	}
}

func handlerStyle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprint(w, `
pre {
  background: #e9e9e9;
  padding: 10px;
  -webkit-border-radius: 5px;
  -moz-border-radius: 5px;
  border-radius: 5px;
}

.k {
  color: #a71d5d;
}

.en {
  color: #795da3;
}

.c {
  color: #969896;
}

.btn-uncollapse, .btn-collapse{
  color: #969896;
}

.uncollapsed > .text-collapsed {
  display: none;
}
.collapsed > .text-uncollapsed {
  display: none;
}
`)
}

func handlerFacedJavascript(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	fmt.Fprint(w, `
window.addEventListener('load', function () {
  var i;
  var collapse = document.getElementsByClassName("btn-collapse")
  var uncollapse = document.getElementsByClassName("btn-uncollapse")

  for (i = 0; i < collapse.length; i++) {
    init(collapse[i], "collapsed");
  }

  for (i = 0; i < uncollapse.length; i++) {
    init(uncollapse[i], "uncollapsed");
  }

  function init(elem, className) {
    elem.addEventListener('click', function() {
      this.parentElement.parentElement.className = className;
    });
  }
});
`)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()

	var fileList []string
	for name, _ := range baalSet.Files {
		fileList = append(fileList, name)
	}
	sort.Strings(fileList)
	tmpl, err := template.New("index").Parse(`<html>
  <head>
    <title>baaldoc</title>
  </head>
  <body>
    <h1>baaldoc</h1>
    <a href="/declarations">Show all declarations</a>
    <ul>{{ range .FileList }}
      <li><a href="/{{.}}">{{.}}</a></li>
    {{end}}</ul>
  </body>
</html>
`)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"FileList": fileList,
	}
	tmpl.Execute(w, data)
}

func handlerFaced(w http.ResponseWriter, r *http.Request, path string, content []baal.Namespace) {
	mutex.RLock()
	defer mutex.RUnlock()

	tmpl, err := template.New("faced").Parse(`<html>
  <head>
    <title>{{.Path}} - baaldoc</title>
    <link rel="stylesheet" type="text/css" href="/static/style.css">
    <script language="javascript" src="/static/faced.js"></script>
  </head>
  <body>
    <a href="/">HOME</a>
    <h1>{{.Path}}</h1>
    <pre>{{.Content}}</pre>
  </body>
</html>
`)
	if err != nil {
		panic(err)
	}
	ctx := RenderContext{}
	data := map[string]interface{}{
		"Path":    path,
		"Content": template.HTML(ctx.renderNamespaces(content)),
	}
	tmpl.Execute(w, data)
}

func collapsable(defaultClass, textCollapsed, textUncollapsed string) string {
	return fmt.Sprintf(`<span class="%s">`+
		`<span class="text-collapsed"><span class="btn-uncollapse">&#x25B7;</span> %s</span>`+
		`<span class="text-uncollapsed"><span class="btn-collapse">&#x25BD;</span> %s</span>`+
		`</span>`, defaultClass, textCollapsed, textUncollapsed)
}

func uncollapsed(textCollapsed, textUncollapsed string) string {
	return collapsable("uncollapsed", textCollapsed, textUncollapsed)
}

func collapsed(textCollapsed, textUncollapsed string) string {
	return collapsable("collapsed", textCollapsed, textUncollapsed)
}

func keyword(text string) string {
	return fmt.Sprintf(`<span class="k">%s</span>`, text)
}

func reference(text string) string {
	return fmt.Sprintf(`<span class="en">%s</span>`, text)
}

func comment(text string) string {
	return fmt.Sprintf(`<span class="c">%s</span>`, text)
}

func (ctx *RenderContext) indent() string {
	result := ""
	for i := 0; i < ctx.nest; i++ {
		result += "\t"
	}
	return result
}

func (ctx *RenderContext) renderNamespaces(namespaces []baal.Namespace) string {
	result := ""
	for _, namespace := range namespaces {
		name := strings.Join(namespace.Name, ".")
		ctx.namespace = name
		ctx.nest++
		body := ctx.renderDeclarations(namespace.Declarations)
		ctx.nest--
		result += fmt.Sprintf(`<span id="%s">`, name)
		result += ctx.indent() + uncollapsed(
			keyword("namespace")+" "+reference(name)+" {"+comment("...")+ctx.indent()+"}",
			keyword("namespace")+" "+reference(name)+"\n"+ctx.indent()+"{\n"+body+ctx.indent()+"}",
		)
		result += "</span>\n\n"
	}
	return result
}

func (ctx *RenderContext) findEntityByName(name []string) (string, baal.Entity) {
	n := strings.Join(name, ".")
	e, ok := baalSet.Entities[n]
	if !ok {
		n = ctx.namespace + "." + n
		e, ok = baalSet.Entities[n]
	}
	if !ok {
		panic("not found: " + strings.Join(name, "."))
	}
	return n, e
}

func (ctx *RenderContext) renderDeclarations(declarations []baal.Declaration) string {
	result := ""
	for _, declaration := range declarations {
		switch declaration := declaration.(type) {
		case baal.Entity:
			result += ctx.renderEntity(declaration)
		case baal.Service:
			result += ctx.renderService(declaration)
		}
	}
	return result
}

func (ctx *RenderContext) renderEntity(entity baal.Entity) string {
	abstract := ""
	if entity.IsAbstract {
		abstract = keyword("abstract") + " "
	}
	includes := ""
	for _, i := range entity.Includes {
		name, e := ctx.findEntityByName(i)
		ctx.nest++
		ee := ctx.renderEntityForReference(name, e)
		ctx.nest--
		includes += ctx.indent() + "+=" + collapsed(reference(fmt.Sprintf(`<a href="/%s#%s">%s</a>`, baalSet.EntityFiles[name], name, name)), ee) + "\n"
	}
	ctx.nest++
	body := ctx.renderFields(entity.Fields)
	ctx.nest--
	id := ctx.namespace + "." + entity.Name
	name := fmt.Sprintf(`<a href="#%s">%s</a><a href="/information/%s.%s">★</a>`,
		id, entity.Name, ctx.namespace, entity.Name)
	result := fmt.Sprintf(`<span id="%s">`, id)
	result += ctx.renderDocument(entity.Document)
	result += ctx.indent() + uncollapsed(
		abstract+keyword("entity")+" "+reference(name)+" {"+comment("...")+"}",
		abstract+keyword("entity")+" "+reference(name)+"\n"+
			includes+ctx.indent()+"{\n"+body+ctx.indent()+"}",
	)
	result += "</span>\n\n"
	return result
}

func (ctx *RenderContext) renderEntityForReference(id string, entity baal.Entity) string {
	abstract := ""
	if entity.IsAbstract {
		abstract = keyword("abstract") + " "
	}
	ctx.nest++
	body := ctx.renderFields(entity.Fields)
	ctx.nest--
	includes := ""
	for _, i := range entity.Includes {
		name, e := ctx.findEntityByName(i)
		ctx.nest++
		ee := ctx.renderEntityForReference(name, e)
		ctx.nest--
		includes += ctx.indent() + "+=" + collapsed(reference(fmt.Sprintf(`<a href="/%s#%s">%s</a>`, baalSet.EntityFiles[name], name, name)), ee) + "\n"
	}
	result := ""
	if entity.Document != "" {
		result += "\n" + ctx.renderDocument(entity.Document) + ctx.indent()
	}
	name := fmt.Sprintf(`<a href="/%s#%s">%s</a>`, baalSet.EntityFiles[id], id, id)
	result += abstract + keyword("entity") + " " + reference(name) + "\n" + includes + ctx.indent() + "{\n" + body + ctx.indent() + "}"
	return result
}

func (ctx *RenderContext) renderFields(fields []baal.Field) string {
	result := ""
	for _, field := range fields {
		result += "\n" + ctx.renderDocument(field.Document)
		result += ctx.indent() + field.Name + ":\n"
		result += ctx.indent() + "\t" + ctx.renderModifiedType(field.ModifiedType) + "\n"
	}
	return result
}

func (ctx *RenderContext) renderService(service baal.Service) string {
	ctx.nest++
	body := ctx.renderMethods(service)
	ctx.nest--
	id := ctx.namespace + "." + service.Name
	name := fmt.Sprintf(`<a href="#%s">%s</a>`, id, service.Name)
	result := fmt.Sprintf(`<span id="%s">`, id)
	result += ctx.renderDocument(service.Document)
	result += ctx.indent() + uncollapsed(
		keyword("service")+" "+reference(name)+" {"+comment("...")+"}",
		keyword("service")+" "+reference(name)+"\n"+ctx.indent()+"{\n"+body+ctx.indent()+"}",
	)
	result += "</span>\n\n"
	return result
}

func (ctx *RenderContext) renderMethods(service baal.Service) string {
	methods := service.Methods
	result := ""
	for _, method := range methods {
		id := ctx.namespace + "." + service.Name + "." + method.Name
		result += "\n"
		result += fmt.Sprintf(`<span id="%s">`, id)
		result += ctx.renderDocument(method.Document)
		result += ctx.indent() + fmt.Sprintf(`<a href="#%s">%s</a>`, id, method.Name) + ":\n"
		result += ctx.indent() + "\t<=" + ctx.renderModifiedType(method.Request) + "\n"
		result += ctx.indent() + "\t=>" + ctx.renderModifiedType(method.Response) + "\n"
		result += "</span>"
	}
	return result
}

func (ctx *RenderContext) renderModifiedType(mt baal.ModifiedType) string {
	result := ""
	switch mt.Occurrence {
	case baal.REQUIRED:
		result += "!"
	case baal.NULLABLE:
		result += "?"
	}
	for _, iteration := range mt.Iterations {
		switch iteration {
		case baal.LIST_OF:
			result += keyword("list of") + " "
		case baal.DICTIONARY_OF:
			result += keyword("dictionary of") + " "
		}
	}
	result += ctx.renderType(mt.Type)
	return result
}

func (ctx *RenderContext) renderType(t baal.Type) string {
	switch t := t.(type) {
	case baal.PrimitiveType:
		switch t.Tok {
		case baal.BOOLEAN:
			return keyword("boolean")
		case baal.INTEGER8:
			return keyword("byte")
		case baal.INTEGER16:
			return keyword("short")
		case baal.INTEGER32:
			return keyword("integer")
		case baal.INTEGER64:
			return keyword("long")
		case baal.FLOAT32:
			return keyword("float")
		case baal.FLOAT64:
			return keyword("double")
		case baal.DECIMAL32:
			return keyword("decimal32")
		case baal.DECIMAL64:
			return keyword("decimal64")
		case baal.DATE:
			return keyword("date")
		case baal.TIME:
			return keyword("time")
		case baal.TIMESTAMP:
			return keyword("datetime")
		case baal.STRING:
			return keyword("string")
		case baal.BINARY:
			return keyword("binary")
		}
	case baal.PseudoType:
		return "&quot;" + t.Name + "&quot;"
	case baal.ReferenceType:
		name, entity := ctx.findEntityByName(t.Name)
		ctx.nest++
		e := ctx.renderEntityForReference(name, entity)
		ctx.nest--
		return collapsed(reference(fmt.Sprintf(`<a href="/%s#%s">%s</a>`, baalSet.EntityFiles[name], name, name)), e)
	}
	return ""
}

func (ctx *RenderContext) renderDocument(document string) string {
	document = strings.TrimSpace(document)
	if document == "" {
		return ""
	}
	lines := strings.Split(document, "\n")
	nest := ctx.nest
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "</") {
			ctx.nest--
		}
		lines[i] = ctx.indent() + line + "\n"
		if !(strings.Contains(line, "</") || strings.Contains(line, "/>")) && strings.HasPrefix(line, "<") {
			ctx.nest++
		}
	}
	ctx.nest = nest
	document = strings.Join(lines, "")
	return ctx.indent() + comment("/#\n"+template.HTMLEscapeString(document)+ctx.indent()+"#/\n")
}

func handlerDeclarations(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()

	var declarationList []string
	for name, _ := range baalSet.DeclarationFiles {
		declarationList = append(declarationList, name)
	}
	sort.Strings(declarationList)
	t := template.New("declarations")
	t.Funcs(template.FuncMap{
		"declarationFiles": func(s string) string {
			return baalSet.DeclarationFiles[s]
		},
	})
	tmpl, err := t.Parse(`<html>
  <head>
    <title>all declarations - baaldoc</title>
  </head>
  <body>
    <a href="/">HOME</a>
    <h1>all declarations</h1>
    <ul>{{ range .declarationList }}
      <li><a href="/{{. | declarationFiles}}#{{.}}">{{.}}</a><a href="/information/{{.}}">★</a></li>
    {{end}}</ul>
  </body>
</html>
`)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"declarationList": declarationList,
	}
	tmpl.Execute(w, data)
}

func handlerInformation(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/information/")
	if file, ok := baalSet.DeclarationFiles[name]; ok {
		// search references
		var fields []string
		var includes []string
		var methods []string
		for _, namespaces := range baalSet.Files {
			for _, namespace := range namespaces {
				ns := strings.Join(namespace.Name, ".")
				getTypeName := func(t baal.Name) string {
					if len(t) == 1 {
						return ns + t[1]
					} else {
						return strings.Join(t, ".")
					}
				}
				for _, declaration := range namespace.Declarations {
					switch declaration := declaration.(type) {
					case baal.Service:
						for _, method := range declaration.Methods {
							if t, ok := method.Request.Type.(baal.ReferenceType); ok {
								if getTypeName(t.Name) == name {
									methods = append(methods, ns+"."+declaration.Name+"."+method.Name)
								}
							}
							if t, ok := method.Response.Type.(baal.ReferenceType); ok {
								if getTypeName(t.Name) == name {
									methods = append(methods, ns+"."+declaration.Name+"."+method.Name)
								}
							}
						}
					case baal.Entity:

						useAsIncludes := false
						for _, t := range declaration.Includes {
							if getTypeName(t) == name {
								useAsIncludes = true
							}
						}
						if useAsIncludes {
							includes = append(includes, ns+"."+declaration.Name)
						}

						useAsFieldType := false
						for _, field := range declaration.Fields {
							if t, ok := field.ModifiedType.Type.(baal.ReferenceType); ok {
								if getTypeName(t.Name) == name {
									useAsFieldType = true
								}
							}
						}
						if useAsFieldType {
							fields = append(fields, ns+"."+declaration.Name)
						}

					}
				}
			}
		}

		sort.Strings(fields)
		sort.Strings(includes)
		sort.Strings(methods)

		t := template.New("information")
		t.Funcs(template.FuncMap{
			"declarationFiles": func(s string) string {
				return baalSet.DeclarationFiles[s]
			},
			"methodFiles": func(s string) string {
				slice := strings.Split(s, ".")
				return baalSet.DeclarationFiles[strings.Join(slice[:len(slice)-1], ".")]
			},
		})
		tmpl, err := t.Parse(`<html>
  <head>
    <title>{{.Name}} - baaldoc</title>
    <link rel="stylesheet" type="text/css" href="/static/style.css">
  </head>
  <body>
    <a href="/">HOME</a>
    <h1>{{.Name}}</h1>
    <h2>Declaration File</h2>
    <a href="/{{.File}}#{{.Name}}">{{.File}}</a>

    <h2>Using as a Field Type</h2>
    <ul>{{ range .Fields }}
      <li><a href="/{{. | declarationFiles}}#{{.}}">{{.}}</a><a href="/information/{{.}}">★</a></li>
    {{end}}</ul>

    <h2>Include from</h2>
    <ul>{{ range .Includes }}
      <li><a href="/{{. | declarationFiles}}#{{.}}">{{.}}</a><a href="/information/{{.}}">★</a></li>
    {{end}}</ul>

    <h2>Use as a Request/Response Type of Methods</h2>
    <ul>{{ range .Methods }}
      <li><a href="/{{. | methodFiles}}#{{.}}">{{.}}</a></li>
    {{end}}</ul>

  </body>
</html>
`)
		if err != nil {
			panic(err)
		}
		data := map[string]interface{}{
			"Name":     name,
			"File":     file,
			"Fields":   fields,
			"Includes": includes,
			"Methods":  methods,
		}
		tmpl.Execute(w, data)
	}
}
