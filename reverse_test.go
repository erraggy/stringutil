package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		actual := Reverse(c.in)
		if actual != c.expected {
			t.Errorf("Reverse(%q) == %q, expected %q", c.in, actual, c.expected)
		}
	}
}
