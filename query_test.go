package goquery

import (
	"testing"
)

func TestIs(t *testing.T) {
	sel := Doc().Find(".footer p:nth-child(1)")
	if !sel.Is("p") {
		t.Error("Expected .footer p:nth-child(1) to be p.")
	}
}

func TestIsPositional(t *testing.T) {
	sel := Doc().Find(".footer p:nth-child(2)")
	if !sel.Is("p:nth-child(2)") {
		t.Error("Expected .footer p:nth-child(2) to be p:nth-child(2).")
	}
}

func TestIsPositionalNot(t *testing.T) {
	sel := Doc().Find(".footer p:nth-child(1)")
	if sel.Is("p:nth-child(2)") {
		t.Error("Expected .footer p:nth-child(1) NOT to be p:nth-child(2).")
	}
}

func TestIsFunction(t *testing.T) {
	ok := Doc().Find("div").IsFunction(func(i int, s *Selection) bool {
		return s.HasClass("container-fluid")
	})

	if !ok {
		t.Error("Expected some div to have a container-fluid class.")
	}
}

func TestIsSelection(t *testing.T) {
	sel := Doc().Find("div")
	sel2 := Doc().Find(".pvk-gutter")

	if !sel.IsSelection(sel2) {
		t.Error("Expected some div to have a pvk-gutter class.")
	}
}

func TestIsSelectionNot(t *testing.T) {
	sel := Doc().Find("div")
	sel2 := Doc().Find("a")

	if sel.IsSelection(sel2) {
		t.Error("Expected some div NOT to be an anchor.")
	}
}

func TestIsNodes(t *testing.T) {
	sel := Doc().Find("div")
	sel2 := Doc().Find(".footer")

	if !sel.IsNodes(sel2.Nodes[0]) {
		t.Error("Expected some div to have a footer class.")
	}
}

func TestHasClass(t *testing.T) {
	sel := Doc().Find("div")
	if !sel.HasClass("span12") {
		t.Error("Expected at least one div to have class span12.")
	}
}

func TestHasClassNone(t *testing.T) {
	sel := Doc().Find("h2")
	if sel.HasClass("toto") {
		t.Error("Expected h1 to have no class.")
	}
}

func TestHasClassNotFirst(t *testing.T) {
	sel := Doc().Find(".alert")
	if !sel.HasClass("alert-error") {
		t.Error("Expected .alert to also have class .alert-error.")
	}
}

func TestDocContains(t *testing.T) {
	sel := Doc().Find("h1")
	if !Doc().Contains(sel.Nodes[0]) {
		t.Error("Expected document to contain H1 tag.")
	}
}

func TestSelContains(t *testing.T) {
	sel := Doc().Find(".row-fluid")
	sel2 := Doc().Find("a[ng-click]")
	if !sel.Contains(sel2.Nodes[0]) {
		t.Error("Expected .row-fluid to contain a[ng-click] tag.")
	}
}

func TestSelNotContains(t *testing.T) {
	sel := Doc().Find("a.link")
	sel2 := Doc().Find("span")
	if sel.Contains(sel2.Nodes[0]) {
		t.Error("Expected a.link to NOT contain span tag.")
	}
}
