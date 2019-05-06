package main

import (
	"flag"
	"testing"
)

func TestTranslate(t *testing.T) {
	t.Run("translate test", func(t *testing.T) {
		tests := []struct {
			source string
			target string
			text   string
			want   string
		}{
			{"ja", "en", "こんにちは", "Hello"},
			{"ja", "en", "お腹が空いた", "I am hungry"},
			{"en", "ja", "Hello", "こんにちは"},
			{"en", "ja", "I'm hungry", "お腹が空きました"},
		}

		for _, test := range tests {
			flag.CommandLine.Set("source", test.source)
			flag.CommandLine.Set("target", test.target)
			flag.CommandLine.Set("text", test.text)

			result, err := translate()
			if err != nil {
				t.Errorf("cannnot translate %s", err)
			}

			if result != test.want {
				t.Errorf("result is wrong. want=%s, actual=%s", test.want, result)
			}
		}

	})
}
