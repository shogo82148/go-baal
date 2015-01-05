package baal

import (
	"testing"
)

func TestScan(t *testing.T) {
	scanner := new(Scanner)
	scanner.Init("namespace Data{ entity Hoge{Fuga: !list of integer}}")
	for {
		tok, lit, pos, err := scanner.Scan()
		t.Logf("%d\t%s\t%v\t%v\n", tok, lit, pos, err)
		if tok == EOF {
			break
		}
	}
}
