package tree_sitter_svelte_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-svelte"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_svelte.Language())
	if language == nil {
		t.Errorf("Error loading Svelte grammar")
	}
}
