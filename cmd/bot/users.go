package main

import (
	"errors"
	"javlonrahimov/dict/internal/data"

	tele "gopkg.in/telebot.v3"
)

func (app *application) handleStartBot(bot *tele.Bot) {
	bot.Handle("/start", func(c tele.Context) error {
		sender := c.Sender()
		word := &data.User{ID: sender.ID, FirstName: sender.FirstName, LastName: sender.LastName, Username: sender.Username}
		err := app.models.Users.Insert(word)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrDuplicateUser):
				return c.Send("Menu")
			}
		}
		return c.Send("Menu")
	})
}
