package tree_sitter_vue_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_vue "github.com/tree-sitter/tree-sitter-vue/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_vue.Language())
	if language == nil {
		t.Errorf("Error loading Vue grammar")
	}
}
