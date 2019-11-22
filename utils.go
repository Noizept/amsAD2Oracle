package adsync

import "github.com/nlopes/slack"


func errorCheck(api *slack.Client, e error, channel,resource string) {
	if e != nil {
		api.PostMessage(channel, slack.MsgOptionText(resource + " - " + e.Error(), false))
		panic(e)
	}

}
