# Telegram Bots API

A Go library for working with the [Telegram Bot API](https://core.telegram.org/bots/api). Implements all methods and types from Telegram Bots API version 9.4.0.

## Installation

```bash
go get github.com/temoon/telegram-bots-api
```

Requires Go 1.24 or higher.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	telegram "github.com/temoon/telegram-bots-api"
	"github.com/temoon/telegram-bots-api/requests"
)

func main() {
	// Create bot
	bot := telegram.NewBot(telegram.BotOpts{
		Token: "YOUR_BOT_TOKEN",
	})

	// Get bot information
	req := &requests.GetMe{}

	response, err := req.Call(context.Background(), bot)
	if err != nil {
		log.Fatal(err)
	}

	me := response.(*telegram.User)
	fmt.Printf("Bot: @%s (%s)\n", *me.Username, me.FirstName)
}
```

## Bot Configuration

### Basic Setup

```go
bot := telegram.NewBot(telegram.BotOpts{
	Token: "YOUR_BOT_TOKEN",
})
```

### With Custom HTTP Client and Timeout

```go
bot := telegram.NewBot(telegram.BotOpts{
	Token:   "YOUR_BOT_TOKEN",
	Timeout: 30 * time.Second,
	Client: &http.Client{
		Timeout: 30 * time.Second,
		// Custom transport, proxy, etc.
	},
})
```

### Test Environment

To work with the [Telegram test environment](https://core.telegram.org/bots/webapps#testing-mini-apps):

```go
bot := telegram.NewBot(telegram.BotOpts{
	Token: "YOUR_BOT_TOKEN",
	Env:   telegram.EnvTest, // Uses api.telegram.org/bot{token}/test/
})
```

## Usage Examples

### Getting Updates (Long Polling)

```go
package main

import (
	"context"
	"log"
	"time"

	telegram "github.com/temoon/telegram-bots-api"
	"github.com/temoon/telegram-bots-api/requests"
)

func main() {
	bot := telegram.NewBot(telegram.BotOpts{
		Token: "YOUR_BOT_TOKEN",
	})

	ctx := context.Background()
	if err := pollUpdates(ctx, bot); err != nil {
		log.Fatal(err)
	}
}

func pollUpdates(ctx context.Context, bot *telegram.Bot) error {
	offset := int64(0)
	timeout := int64(30)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req := &requests.GetUpdates{
			Offset:  &offset,
			Timeout: &timeout,
			AllowedUpdates: []string{
				telegram.UpdatesAllowedMessage,
				telegram.UpdatesAllowedCallbackQuery,
			},
		}

		response, err := req.Call(ctx, bot)
		if err != nil {
			log.Printf("Error getting updates: %v", err)
			time.Sleep(time.Second)
			continue
		}

		updates := response.(*[]telegram.Update)
		for _, update := range *updates {
			offset = update.UpdateId + 1
			handleUpdate(ctx, bot, &update)
		}
	}
}

func handleUpdate(ctx context.Context, bot *telegram.Bot, update *telegram.Update) {
	if update.Message != nil && update.Message.Text != nil {
		log.Printf("New message from %s: %s", update.Message.From.FirstName, *update.Message.Text)
	}
}
```

### Setting Up a Webhook

```bash
curl -X POST "https://api.telegram.org/bot<YOUR_BOT_TOKEN>/setWebhook" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/webhook",
    "allowed_updates": ["message", "callback_query"],
    "secret_token": "your-secret-token"
  }'
```

### Handling Webhooks

```go
package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	telegram "github.com/temoon/telegram-bots-api"
	"github.com/temoon/telegram-bots-api/requests"
)

func main() {
	bot := telegram.NewBot(telegram.BotOpts{
		Token: "YOUR_BOT_TOKEN",
	})

	// Use closure to pass bot to handler
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		// Check secret token (if used during setup)
		if r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != "your-secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Parse JSON into Update
		var update telegram.Update
		if err := json.Unmarshal(body, &update); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Handle update
		go handleUpdate(r.Context(), bot, &update)

		// Telegram expects 200 OK immediately
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Starting webhook server on :8443")
	log.Fatal(http.ListenAndServe(":8443", nil))
}

func handleUpdate(ctx context.Context, bot *telegram.Bot, update *telegram.Update) {
	// Handle text messages
	if update.Message != nil && update.Message.Text != nil {
		log.Printf("Received message from %s: %s", update.Message.From.FirstName, *update.Message.Text)

		// Send response
		req := &requests.SendMessage{
			ChatId: telegram.NewChatId(update.Message.Chat.Id, ""),
			Text:   "Got your message!",
		}

		if _, err := req.Call(ctx, bot); err != nil {
			log.Printf("Error sending response: %v", err)
		}
	}
}
```

## Error Handling

API errors are returned as standard Go errors:

```go
response, err := req.Call(ctx, bot)
if err != nil {
	// Error contains description from Telegram API
	// For example: "Bad Request: chat not found"
	log.Printf("API error: %v", err)
	return
}
```

## Special Data Types

### ChatId

Universal type for chat identification — supports both numeric IDs and usernames:

```go
// By numeric ID
chatId := telegram.NewChatId(123456789, "")

// By channel/group username
chatId := telegram.NewChatId(0, "@channelname")
```

### InputFile

Universal type for sending files:

```go
// By URL
file := telegram.NewInputFile("https://example.com/file.jpg", nil, "")

// By file_id (for resending)
file := telegram.NewInputFile("AgACAgIAAxkBAAI...", nil, "")

// File upload
reader, _ := os.Open("file.jpg")
file := telegram.NewInputFile("", reader, "file.jpg")
```

## Useful Links

- [Telegram Bot API Documentation](https://core.telegram.org/bots/api)
- [Creating a bot via @BotFather](https://core.telegram.org/bots#botfather)
- [Formatting options](https://core.telegram.org/bots/api#formatting-options)
- [Inline keyboards](https://core.telegram.org/bots/features#inline-keyboards)

## Code Generation

Most of the library code is automatically generated from the official Telegram Bot API documentation using [telegram-bots-api-generator](https://github.com/temoon/telegram-bots-api-generator).

Auto-generated files:
- `types.go` — type definitions
- `requests/*.go` — API method implementations
- `requests/*_test.go` — method tests

## License

See [LICENSE](LICENSE) file.
