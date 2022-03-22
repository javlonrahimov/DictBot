package main

import (
	tele "gopkg.in/telebot.v3"
)

func (app *application) showWordDesc(bot *tele.Bot) {
	bot.Handle(tele.OnText, func(c tele.Context) error {
		var (
			text = c.Text()
			user = c.Sender()
		)

		words, err := app.models.Words.GetByWord(text)
		if err != nil {
			app.logger.PrintError(err, nil)
			return c.Send("No result found!!!")
		}
		response := ""
		for _, word := range words {
			response += app.Bold(word.Word) + " 📖 " + word.WordType + app.NewLine(1) +
				app.Italic(app.Clean(word.Definition, "\n")) + app.NewLine(2)
		}

		if len(response) == 0 {
			return c.Send("No result found!!!")
		}
		c.Send(response, tele.ModeHTML)
		app.models.Words.AddForUser(user.ID, words[0].ID)
		return nil
	})
}

func (app *application) listUserWords(bot *tele.Bot) {
	bot.Handle("/words", func(c tele.Context) error {
		user := c.Sender()
		words, err := app.models.Words.GetAllForUser(user.ID)
		if err != nil {
			app.logger.PrintError(err, nil)
			return c.Send("No result found!!!")
		}
		response := ""
		for _, word := range words {
			response += app.Bold(word.Word) + " 📖 " + word.WordType + app.NewLine(1) +
				app.Italic(app.Clean(word.Definition, "\n")) + app.NewLine(2)
		}

		if len(response) == 0 {
			return c.Send("No words for you!!!")
		}
		c.Send(response, tele.ModeHTML)
		return nil
	})
}
