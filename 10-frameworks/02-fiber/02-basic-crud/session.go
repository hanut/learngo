package main

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

// Create a new session store
var SessionStore *session.Store

func init() {
	SessionStore = session.New(session.Config{
		Storage: redis.New(redis.Config{
			Host: "localhost",
			Port: 6379,
		}),
	})

}
