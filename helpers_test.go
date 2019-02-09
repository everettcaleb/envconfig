package envconfig

import "testing"

var parseFriendlyBoolTests = []struct {
	example        string
	expected       bool
	errShouldBeNil bool
}{
	{"yes", true, false},
	{"true", true, false},
	{"y", true, false},
	{"t", true, false},
	{"1", true, false},
	{"no", false, false},
	{"false", false, false},
	{"n", false, false},
	{"f", false, false},
	{"0", false, false},

	{"yEs", true, false},
	{"tRue", true, false},
	{"Y", true, false},
	{"T", true, false},
	{"1", true, false},
	{"nO", false, false},
	{"fAlSe", false, false},
	{"N", false, false},
	{"F", false, false},
	{"0", false, false},

	{"BLAH", false, true},
	{"foo", false, true},
}

func TestParseFriendlyBool(t *testing.T) {
	for _, v := range parseFriendlyBoolTests {
		b, err := parseFriendlyBool(v.example)
		if b != v.expected || (err != nil && !v.errShouldBeNil) {
			t.Errorf("Expected value %v for input %s with error %v", v.expected, v.example, err)
		}
	}
}
