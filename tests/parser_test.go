package tests

import (
	"github.com/artemkakun/skillshot-tgbot/parser"
	"testing"
)

func TestLinksCount(t *testing.T) {
	links := parser.GetVacanciesLinksList()
	if len(links) != 30 {
		t.Errorf("Count was incorrect, got: %d, want: %d.", len(links), 30)
	}
}
