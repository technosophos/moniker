package moniker

import (
	"testing"
)

func TestNamer(t *testing.T) {
	n := New().(*defaultNamer)
	n.Descriptor = []string{"foo"}
	n.Noun = []string{"bar"}

	if name := n.Name(); name != "foo bar" {
		t.Fatalf("Got %s", name)
	}

	if name := n.NameSep("$"); name != "foo$bar" {
		t.Fatalf("Got %s", name)
	}

	if name := n.NameWithOptions("$", "f", false); name != "foo$bar" {
		t.Fatalf("Got %s", name)
	}	
}

func TestAlliterations(t *testing.T){
	n := New().(*defaultNamer)
	n.Descriptor = []string{"foo"}
	n.Noun = []string{"bar", "far"}

	if name := n.NameWithOptions(" ", "f", true); name != "foo far" {
		t.Fatalf("Got %s", name)
	}

	if name := n.NameWithOptions("$", "f", true); name != "foo$far" {
		t.Fatalf("Got %s", name)
	}	
}
