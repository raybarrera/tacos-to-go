package main

import "testing"

func TestPragmaMapToDbUrl(t *testing.T) {
	pragmas := map[string]string{
		"t1": "v1",
		"t2": "v2",
	}

	url := PragmaMapToDbUrl(pragmas)
	expected := "?_pragma=t1(v1)&_pragma=t2(v2)"
	if url != expected {
		t.Errorf("Expected %s, got %s", expected, url)
	}
}
