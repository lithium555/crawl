package parser

import (
	"testing"
	"reflect"
)

func TestRegextUnits(t *testing.T) {
	for _, c := range cases {
		got := c.re.FindStringSubmatch(c.check)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("unexpected match: %q vs %q", got, c.expect)
			t.Errorf("The string, whitch broken: '%q'", c.check)
		}
	}
}