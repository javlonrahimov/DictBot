package main

import (
	"fmt"

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
			return c.Reply("No result found!!!")
		}

		if len(words) == 0 {
			return c.Reply("No result found!!!")
		}

		response := ""
		response += app.Hashtag(app.Bold(text)) + app.NewLine(2)
		for index , word := range words {
			response += app.Bold(fmt.Sprint(index + 1)) + " 📖 " + word.WordType + app.NewLine(1) +
				app.Italic(app.Clean(word.Definition, "\n")) + app.NewLine(2)
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
			return c.Reply("No result found!!!")
		}

		if len(words) == 0 {
			return c.Reply("No words for you!!!")
		}

		response := ""
		for _, word := range words {
			response += app.Hashtag(app.Bold(word.Word)) + app.NewLine(1)
		}

		c.Send(response, tele.ModeHTML)
		return nil
	})
}

func (app *application) clearUserWords(bot *tele.Bot) {
	bot.Handle("/clear", func(c tele.Context) error {
		user := c.Sender()
		err := app.models.Words.ClearUserWords(user.ID)
		if err != nil {
			app.logger.PrintError(err, nil)
			return c.Reply("Something went wrong...")
		}
		return c.Reply("Deleted!!!")
	})
}
