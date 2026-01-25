# Telegram Bots API

Go-библиотека для работы с [Telegram Bot API](https://core.telegram.org/bots/api). Реализует полную типизацию всех методов и типов API версии 9.3.0.

## Установка

```bash
go get github.com/temoon/telegram-bots-api
```

Требуется Go 1.24 или выше.

## Быстрый старт

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
	// Создание бота
	bot := telegram.NewBot(telegram.BotOpts{
		Token: "YOUR_BOT_TOKEN",
	})

	// Отправка сообщения
	req := &requests.SendMessage{
		ChatId: telegram.NewChatId(123456789, ""),
		Text:   "Привет!",
	}

	response, err := req.Call(context.Background(), bot)
	if err != nil {
		log.Fatal(err)
	}

	msg := response.(*telegram.Message)
	fmt.Printf("Сообщение отправлено: %d\n", msg.MessageId)
}
```

## Конфигурация бота

### Базовая настройка

```go
bot := telegram.NewBot(telegram.BotOpts{
	Token: "YOUR_BOT_TOKEN",
})
```

### С пользовательским HTTP-клиентом и таймаутом

```go
import (
	"net/http"
	"time"
)

bot := telegram.NewBot(telegram.BotOpts{
	Token:   "YOUR_BOT_TOKEN",
	Timeout: 30 * time.Second,
	Client: &http.Client{
		Timeout: 30 * time.Second,
		// Кастомный транспорт, прокси и т.д.
	},
})
```

### Тестовое окружение

Для работы с [тестовым окружением Telegram](https://core.telegram.org/bots/webapps#testing-mini-apps):

```go
bot := telegram.NewBot(telegram.BotOpts{
	Token: "YOUR_BOT_TOKEN",
	Env:   telegram.EnvTest, // Использует api.telegram.org/bot{token}/test/
})
```

## Примеры использования

### Получение обновлений (Long Polling)

```go
func pollUpdates(ctx context.Context, bot *telegram.Bot) error {
	var offset int64 = 0
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
			log.Printf("Ошибка получения обновлений: %v", err)
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
	if update.Message != nil {
		log.Printf("Новое сообщение от %s: %s",
			update.Message.From.FirstName,
			*update.Message.Text)
	}
}
```

### Установка Webhook

```go
req := &requests.SetWebhook{
	Url: "https://example.com/webhook",
	AllowedUpdates: []string{
		telegram.UpdatesAllowedMessage,
		telegram.UpdatesAllowedCallbackQuery,
	},
	SecretToken: ptr("your-secret-token"),
}

response, err := req.Call(ctx, bot)
if err != nil {
	log.Fatal(err)
}

if *response.(*bool) {
	log.Println("Webhook установлен")
}

// Вспомогательная функция для создания указателей
func ptr[T any](v T) *T {
	return &v
}
```

### Отправка сообщения с форматированием

```go
parseMode := telegram.ParseModeHTML

req := &requests.SendMessage{
	ChatId:    telegram.NewChatId(chatId, ""),
	Text:      "<b>Жирный</b>, <i>курсив</i>, <code>код</code>",
	ParseMode: &parseMode,
}

response, err := req.Call(ctx, bot)
```

Поддерживаемые режимы форматирования:
- `telegram.ParseModeHTML`
- `telegram.ParseModeMarkdown`
- `telegram.ParseModeMarkdownV2`

Подробнее: [Formatting options](https://core.telegram.org/bots/api#formatting-options)

### Отправка сообщения с Inline-клавиатурой

```go
callbackData := "button_clicked"

req := &requests.SendMessage{
	ChatId: telegram.NewChatId(chatId, ""),
	Text:   "Выберите действие:",
	ReplyMarkup: telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{Text: "Кнопка 1", CallbackData: &callbackData},
				{Text: "Кнопка 2", CallbackData: ptr("button_2")},
			},
			{
				{Text: "Открыть ссылку", Url: ptr("https://telegram.org")},
			},
		},
	},
}

response, err := req.Call(ctx, bot)
```

### Обработка Callback Query

```go
func handleCallbackQuery(ctx context.Context, bot *telegram.Bot, query *telegram.CallbackQuery) {
	// Отвечаем на callback (убираем "часики" с кнопки)
	answerReq := &requests.AnswerCallbackQuery{
		CallbackQueryId: query.Id,
		Text:            ptr("Действие выполнено!"),
	}
	answerReq.Call(ctx, bot)

	// Обрабатываем данные
	if query.Data != nil {
		switch *query.Data {
		case "button_clicked":
			// Обработка нажатия
		}
	}
}
```

### Отправка фото

#### По URL

```go
req := &requests.SendPhoto{
	ChatId:  telegram.NewChatId(chatId, ""),
	Photo:   telegram.NewInputFile("https://example.com/photo.jpg", nil, ""),
	Caption: ptr("Описание фото"),
}

response, err := req.Call(ctx, bot)
```

#### Загрузка файла

```go
file, err := os.Open("photo.jpg")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

req := &requests.SendPhoto{
	ChatId:  telegram.NewChatId(chatId, ""),
	Photo:   telegram.NewInputFile("", file, "photo.jpg"),
	Caption: ptr("Загруженное фото"),
}

response, err := req.Call(ctx, bot)
```

#### По file_id (повторная отправка)

```go
req := &requests.SendPhoto{
	ChatId: telegram.NewChatId(chatId, ""),
	Photo:  telegram.NewInputFile("AgACAgIAAxkBAAI...", nil, ""),
}

response, err := req.Call(ctx, bot)
```

### Отправка документа

```go
file, _ := os.Open("document.pdf")
defer file.Close()

req := &requests.SendDocument{
	ChatId:   telegram.NewChatId(chatId, ""),
	Document: telegram.NewInputFile("", file, "document.pdf"),
	Caption:  ptr("Важный документ"),
}

response, err := req.Call(ctx, bot)
```

### Отправка в канал по username

```go
req := &requests.SendMessage{
	ChatId: telegram.NewChatId(0, "@channelname"),
	Text:   "Сообщение в канал",
}

response, err := req.Call(ctx, bot)
```

### Редактирование сообщения

```go
req := &requests.EditMessageText{
	ChatId:    telegram.NewChatId(chatId, ""),
	MessageId: &messageId,
	Text:      "Обновлённый текст",
}

response, err := req.Call(ctx, bot)
```

### Удаление сообщения

```go
req := &requests.DeleteMessage{
	ChatId:    telegram.NewChatId(chatId, ""),
	MessageId: messageId,
}

response, err := req.Call(ctx, bot)
```

### Отображение статуса "печатает..."

```go
req := &requests.SendChatAction{
	ChatId: telegram.NewChatId(chatId, ""),
	Action: telegram.ChatActionTyping,
}

req.Call(ctx, bot)
```

Доступные действия: `ChatActionTyping`, `ChatActionUploadPhoto`, `ChatActionRecordVideo` и другие.

### Получение информации о боте

```go
req := &requests.GetMe{}
response, err := req.Call(ctx, bot)
if err != nil {
	log.Fatal(err)
}

me := response.(*telegram.User)
fmt.Printf("Бот: @%s (%s)\n", *me.Username, me.FirstName)
```

### Reply-клавиатура

```go
req := &requests.SendMessage{
	ChatId: telegram.NewChatId(chatId, ""),
	Text:   "Выберите вариант:",
	ReplyMarkup: telegram.ReplyKeyboardMarkup{
		Keyboard: [][]telegram.KeyboardButton{
			{
				{Text: "Вариант 1"},
				{Text: "Вариант 2"},
			},
			{
				{Text: "Отмена"},
			},
		},
		ResizeKeyboard:  ptr(true),
		OneTimeKeyboard: ptr(true),
	},
}

response, err := req.Call(ctx, bot)
```

### Удаление Reply-клавиатуры

```go
req := &requests.SendMessage{
	ChatId: telegram.NewChatId(chatId, ""),
	Text:   "Клавиатура убрана",
	ReplyMarkup: telegram.ReplyKeyboardRemove{
		RemoveKeyboard: true,
	},
}

response, err := req.Call(ctx, bot)
```

### Ответ на сообщение (Reply)

```go
req := &requests.SendMessage{
	ChatId: telegram.NewChatId(chatId, ""),
	Text:   "Это ответ на ваше сообщение",
	ReplyParameters: &telegram.ReplyParameters{
		MessageId: originalMessageId,
	},
}

response, err := req.Call(ctx, bot)
```

## Обработка ошибок

Ошибки API возвращаются как обычные Go-ошибки:

```go
response, err := req.Call(ctx, bot)
if err != nil {
	// Ошибка содержит описание от Telegram API
	// Например: "Bad Request: chat not found"
	log.Printf("Ошибка API: %v", err)
	return
}
```

### Использование контекста для таймаутов

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

response, err := req.Call(ctx, bot)
if err != nil {
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Превышен таймаут запроса")
	} else {
		log.Printf("Ошибка: %v", err)
	}
}
```

### Использование контекста для отмены

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
	// Отмена запроса через 5 секунд
	time.Sleep(5 * time.Second)
	cancel()
}()

response, err := req.Call(ctx, bot)
if err != nil {
	if ctx.Err() == context.Canceled {
		log.Println("Запрос отменён")
	}
}
```

## Типы данных

### ChatId

Универсальный тип для идентификации чата — поддерживает как числовой ID, так и username:

```go
// По числовому ID
chatId := telegram.NewChatId(123456789, "")

// По username канала/группы
chatId := telegram.NewChatId(0, "@channelname")
```

### InputFile

Универсальный тип для отправки файлов:

```go
// По URL
file := telegram.NewInputFile("https://example.com/file.jpg", nil, "")

// По file_id (для повторной отправки)
file := telegram.NewInputFile("AgACAgIAAxkBAAI...", nil, "")

// Загрузка файла
reader, _ := os.Open("file.jpg")
file := telegram.NewInputFile("", reader, "file.jpg")
```

## Константы

Библиотека предоставляет константы для всех значений API:

```go
// Режимы форматирования
telegram.ParseModeHTML
telegram.ParseModeMarkdown
telegram.ParseModeMarkdownV2

// Типы чатов
telegram.ChatTypePrivate
telegram.ChatTypeGroup
telegram.ChatTypeSupergroup
telegram.ChatTypeChannel

// Действия в чате
telegram.ChatActionTyping
telegram.ChatActionUploadPhoto
telegram.ChatActionRecordVideo
// ... и другие

// Типы обновлений для фильтрации
telegram.UpdatesAllowedMessage
telegram.UpdatesAllowedCallbackQuery
telegram.UpdatesAllowedInlineQuery
// ... и другие
```

## Полезные ссылки

- [Telegram Bot API Documentation](https://core.telegram.org/bots/api)
- [Создание бота через @BotFather](https://core.telegram.org/bots#botfather)
- [Formatting options](https://core.telegram.org/bots/api#formatting-options)
- [Inline keyboards](https://core.telegram.org/bots/features#inline-keyboards)

## Генерация кода

Большая часть кода библиотеки генерируется автоматически из официальной документации Telegram Bot API с помощью [telegram-bots-api-generator](https://github.com/temoon/telegram-bots-api-generator).

Автоматически генерируемые файлы:
- `types.go` — определения типов
- `requests/*.go` — реализации методов API
- `requests/*_test.go` — тесты методов

## Лицензия

См. файл [LICENSE](LICENSE).
