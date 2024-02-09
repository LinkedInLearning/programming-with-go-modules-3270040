package search

import (
	"testing"
)

func TestByDesc(t *testing.T) {
	tests := map[string]struct {
		params Params
		want   []string
	}{
		"fruit": {
			Params{Include: []string{"fruit"}},
			[]string{"🍇", "🍈", "🍉", "🍊", "🍋"},
		},
		"cat": {
			Params{Include: []string{"cat"}},
			[]string{"🐱"},
		},
		"animal faces": {
			Params{Include: []string{"face"}, Exclude: []string{"smile", "laugh", "grin", "upside-down"}},
			[]string{"🐵", "🐶", "🐱", "🐯", "🦊"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := All(test.params)
			found := 0
			for _, emoji := range result {
				for _, ewant := range test.want {
					if ewant == emoji.Unicode {
						found++
					}
				}
			}
			if found != len(test.want) {
				t.Error("Did find all expected emojis")
			}
		})
	}
}

func TestTag(t *testing.T) {
	tests := map[string]struct {
		tags []string
		want []string
	}{
		"fruit": {
			[]string{"fruit"},
			[]string{"🍇", "🍈", "🍉", "🍊", "🍋"},
		},
		"cat": {
			[]string{"cat"},
			[]string{"🐱"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ByTags(test.tags...)
			found := 0
			for _, emoji := range result {
				for _, ewant := range test.want {
					if ewant == emoji.Unicode {
						found++
					}
				}
			}
			if found != len(test.want) {
				t.Error("Did find all expected emojis")
			}
		})
	}
}
