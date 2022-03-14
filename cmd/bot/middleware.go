package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"

	tele "gopkg.in/telebot.v3"
)

func (app *application) recoverPanic(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		defer func() {
			if err := recover(); err != nil {
				app.serverErrorResponse(c, fmt.Errorf("%s", err))
			}
		}()
		return next(c)
	}
}

func (app *application) rateLimit(next tele.HandlerFunc) tele.HandlerFunc {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[int64]*client)
	)

	go func() {
		for {
			time.Sleep(time.Minute)

			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()
	return func(c tele.Context) error {
		if app.config.limiter.enabled {

			id := c.Sender().ID

			mu.Lock()

			if _, found := clients[id]; !found {
				clients[id] = &client{limiter: rate.NewLimiter(rate.Limit(app.config.limiter.rps), app.config.limiter.burst)}
			}

			clients[id].lastSeen = time.Now()

			if !clients[id].limiter.Allow() {
				mu.Unlock()
				app.rateLimitExceededResponse(c)
				return nil
			}

			mu.Unlock()

			return next(c)
		} else {
			return next(c)
		}
	}
}
