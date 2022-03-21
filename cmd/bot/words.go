package main

import (
	tele "gopkg.in/telebot.v3"
)

func (app *application) showWordDesc(bot *tele.Bot) {
	bot.Handle(tele.OnText, func(c tele.Context) error {
		var (
			text = c.Text()
			// user = c.Sender()
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
		return c.Send(response, tele.ModeHTML)
	})
}

func (app *application) listUserWords(bot *tele.Bot) {
	bot.Handle("/words", func(c tele.Context) error {
		return c.Send("To do")
	})
}
