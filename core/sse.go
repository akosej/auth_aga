package core

import (
	"bufio"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/valyala/fasthttp"
)

type SSEEvent struct {
	Event string
	Data  string
}

var clients = make(map[string]*bufio.Writer)

func SSEHandler(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")
	claims := c.Context().Value("claims").(*jwt.RegisteredClaims)
	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		clients[claims.ID] = w
		for {
			err := w.Flush()
			if err != nil {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}))
	return nil
}

func SendSSEEventToUser(username string, event SSEEvent) {
	if writer, exists := clients[username]; exists {
		writer.WriteString(fmt.Sprintf("event: %s\ndata: %s\n\n", event.Event, event.Data))
		writer.Flush()
	}
}
