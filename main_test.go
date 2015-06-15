package main

import "testing"

func TestModifyLine(t *testing.T) {
	prefix := "_"

	tests := []struct {
		line string
		out  string
	}{
		{"whatever", "whatever"},
		{"	aggregate_retailer_products_test.go:23: message", "	_aggregate_retailer_products_test.go:23: message"},
		{"aggregate_retailer_products_test.go:23: message", "_aggregate_retailer_products_test.go:23: message"},
	}

	for _, tst := range tests {
		actual := modifyLine(prefix, tst.line)
		if actual != tst.out {
			t.Errorf("%q\nexpected\n%s\ngot\n%s", tst.line, tst.out, actual)
		}
	}
}
