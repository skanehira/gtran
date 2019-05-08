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
			{"ja", "en", "お腹が空いた", "I am hungry"},
			{"en", "ja", "I'm hungry", "お腹が空きました"},
		}

		for _, test := range tests {
			result, err := translate(test.text, test.source, test.target)
			if err != nil {
				t.Errorf("cannnot translate %s", err)
			}

			if result != test.want {
				t.Errorf("result is wrong. want=%s, actual=%s", test.want, result)
			}
		}

	})

	t.Run("run test", func(t *testing.T) {
		tests := []struct {
			text string
			args []string
			want int
		}{
			{"", []string{}, -1},
			{"", []string{"hello"}, 0},
			{"hello", []string{}, 0},
		}

		for i, test := range tests {

			if test.text != "" {
				flag.CommandLine.Set("text", test.text)
			}

			result := run(test.args)
			if test.want != result {
				t.Errorf(" %d result is wrong. want=%d, actual=%d", i, test.want, result)
			}
		}
	})
}
