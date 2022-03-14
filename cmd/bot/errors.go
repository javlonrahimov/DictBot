package main

import tele "gopkg.in/telebot.v3"

func (app *application) serverErrorResponse(ctx tele.Context, err error) {
	app.logger.PrintError(err, nil)
	ctx.Send("Something went wrong!!!")
}

func (app *application) rateLimitExceededResponse(ctx tele.Context) {
	ctx.Send("Rate limit exceeded")
}
