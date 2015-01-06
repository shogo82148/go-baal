package baal

import (
	"testing"
)

func TestParse(t *testing.T) {
	scanner := new(Scanner)
	scanner.Init(`
namespace Data.Hoge import Hoge.Fuga.* {

  /#
    hoge
  #/
  abstract entity Foo
  += FooBar
  {
    /# これはテストなのです #/
		Test: ! list of integer;
  }

  service HogeHoge {
    Hoge: <= !integer => !"Fo\u0020o\n";
  }
}
`)
	n, e := Parse(scanner)
	t.Logf("%v", n)
	t.Logf("%v", e)
}
