package app

import (
	"testing"

	"github.com/mattermost/platform/model"
)

func TestQuoteProviderDoCommand(t *testing.T) {
	qp := QuoteProvider{}
	args := &model.CommandArgs{
		T: func(s string, args ...interface{}) string { return s },
	}

	for msg, expected := range map[string]string{
		"":           "api.command_quote.message.app_error",
		"foo":        "> foo",
		"foo\nbar":   "> foo\n> bar",
		"foo\nbar\n": "> foo\n> bar\n> ",
	} {
		actual := qp.DoCommand(args, msg).Text
		if actual != expected {
			t.Errorf("expected `%v`, got `%v`", expected, actual)
		}
	}
}
