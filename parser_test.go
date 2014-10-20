package semver

import (
	"fmt"
	"testing"
)

var items = map[itemType]string{
	0: "itemVersion",
	1: "itemOperator",
	2: "itemSet",
	3: "itemRange",
	4: "itemError",
	5: "itemEOF",
}

var constraints = []string{
	"1.0 || >=2.5.0 || 5.0.0 - 7.2.3",
	"~1.2.3",
	"^4.5.2-alpha.1",
	"=2.3.2",
	"<=1.2.3",
	"5.3.5||4.3.5",
	"5.3.5 ||4.3.5",
	"5.3.5|| 4.3.5",
	"5.3.5 4.3.5",
}

func TestParser(t *testing.T) {
	for _, c := range constraints {
		_, ch := lex("test", c)
		for {
			s, ok := <-ch
			if ok != false {
				fmt.Printf("%v: '%v' \n", items[s.typ], s)
			} else {
				break
			}
		}
	}
}