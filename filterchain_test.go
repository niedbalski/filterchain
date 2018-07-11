package filterchain

import (
	"testing"
	"strings"
)

var (
	TestString = "TESTING"
)

type TestToLowerFilter struct {}
type TestRepeatFilter struct {}

func (filter TestToLowerFilter) Apply(t interface{}) (interface{}) {
	return strings.ToLower(t.(string))
}

func (filter TestRepeatFilter) Apply(t interface{}) (interface{}) {
	return strings.Repeat(t.(string), 3)
}

func TestFilterChain_Apply(t *testing.T) {
	if New(DefaultFilter{}, TestToLowerFilter{}).Apply(TestString) != strings.ToLower(TestString) {
		t.Error("Chain sequence not applied orrectly")
	}
}

func TestFilterChain_Apply2(t *testing.T) {
	if New(DefaultFilter{}, TestToLowerFilter{}, TestRepeatFilter{}).Apply(TestString) != "testingtestingtesting" {
		t.Error("Chain sequence not applied correctly")
	}
}