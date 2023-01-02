package vince

import (
	"reflect"
	"testing"
)

func TestParseBot(t *testing.T) {
	ua := `Googlebot/2.1 (http://www.googlebot.com/bot.html)`
	expect := &botMatch{
		name:         "Googlebot",
		category:     "Search bot",
		url:          "http://www.google.com/bot.html",
		producerName: "Google Inc.",
		producerURL:  "http://www.google.com",
	}
	if !reflect.DeepEqual(parseBotUA(ua), expect) {
		t.Error("failed expectation")
	}
}
