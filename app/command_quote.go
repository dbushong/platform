// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"strings"

	"github.com/mattermost/platform/model"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
)

type QuoteProvider struct {
}

const (
	CMD_QUOTE = "quote"
)

func init() {
	RegisterCommandProvider(&QuoteProvider{})
}

func (me *QuoteProvider) GetTrigger() string {
	return CMD_QUOTE
}

func (me *QuoteProvider) GetCommand(T goi18n.TranslateFunc) *model.Command {
	return &model.Command{
		Trigger:          CMD_QUOTE,
		AutoComplete:     true,
		AutoCompleteDesc: T("api.command_quote.desc"),
		AutoCompleteHint: T("api.command_quote.hint"),
		DisplayName:      T("api.command_quote.name"),
	}
}

func (me *QuoteProvider) DoCommand(args *model.CommandArgs, message string) *model.CommandResponse {
	if len(message) == 0 {
		return &model.CommandResponse{Text: args.T("api.command_quote.message.app_error"), ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL}
	}
	rmsg := "> " + strings.Join(strings.Split(message, "\n"), "\n> ")
	return &model.CommandResponse{ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL, Text: rmsg}
}
