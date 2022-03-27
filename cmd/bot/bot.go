package main

import (
	"math/rand"
	"time"

	tele "gopkg.in/telebot.v3"
)

func (app *application) prefs() tele.Settings {
	return tele.Settings{
		Token:  app.config.token,
		Poller: &tele.LongPoller{Timeout: time.Duration(rand.Int31n(int32(app.config.pollerTimeout))) * time.Second},
	}
}

func (app *application) startBot() error {
	b, err := tele.NewBot(app.prefs())
	if err != nil {
		app.logger.PrintFatal(err, nil)
		return err
	}
	b.Use(app.recoverPanic)
	b.Use(app.rateLimit)

	app.showWordDesc(b)
	app.listUserWords(b)
	app.handleStartBot(b)
	app.clearUserWords(b)

	b.Start()
	return nil
}
