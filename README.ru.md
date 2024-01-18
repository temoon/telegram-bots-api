# Telegram Bots API: Реализация на Go

О возможностях Telegram Bots можно прочитать в [этой документации](https://core.telegram.org/bots/features). Для получения исчерпывающей информации о типах и методах API прочтите [официальную документацию Telegram Bots API](https://core.telegram.org/bots/api).

Эта реализация полностью повторяет описание из официальной документации по API и основана на данных автоматизированного разбора HTML страницы. Опечатки в названиях методов и типов, гипотетически, полностью исключены. Ошибки типизации полей и возвращаемых значений методов возможны, т.к. информация в исходной документации, к сожалению, не полностью унифицирована.

Автоматизированный разбор страницы документации выполнен с помощью [этого инструмента](https://github.com/temoon/telegram-bots-api-generator).

## Соглашения

### Версионирование

Каждый релиз помечается тегом в формате `v[MAJOR].[API].[MINOR]` (например, `v1.700.2`), где:
* `MAJOR` - главная версия реализации этой библиотеки;
* `API` - версия Telegram Bots API (например, `700` => `7.0` или `622` => `6.2.2`);
* `MINOR` - версия исправления в пределах версии API (т.е. функционально ничего не меняется, но будет улучшена стабильность, скорость или ошибка).

### Совместимость

К сожалению, пока идет активная разработка, я не готов гарантировать полную совместимость в пределах главной версии. Но стоит отметить, что в последнее время действительно мало несовместимых изменений.

## Установка

В каталоге вашего Go проекта выполните команду:

```shell
go get github.com/temoon/telegram-bots-api
```

## Использование

### Инициализация

Для выполнения запросов в Telegram API нужно создать и настроить объект `telegram.Bot`:

```go
package main

import (
    "github.com/temoon/telegram-bots-api"
    "os"
    "time"
)

func main() {
    botOpts := telegram.BotOpts{
        Token:   os.Getenv("TELEGRAM_TOKEN"),
        Client:  nil,
        Timeout: time.Minute,
        Env:     telegram.EnvProduction,
    }

    bot := telegram.NewBot(&botOpts)
}
```

### Настройка

Настройки передаются в объекте `telegram.BotOpts`:

```go
type BotOpts struct {
    Token   string
    Client  *http.Client
    Timeout time.Duration
    Env     string
}
```

* `Token` - авторизационный токен, может быть получен у [BotFather](https://core.telegram.org/bots/features#botfather);
* `Client` - возможность использовать для отправки запросов `http.Client`, настроенный конкретно под ваши задачи (если передано `nil`, используются настройки по умолчанию, за исключением `Timeout` == `BotOpts.Timeout`);
* `Timeout` - используется для передачи времени ожидания в `http.Client`;
* `Env` - возможность выполнять запросы в [тестовое окружение](https://core.telegram.org/bots/features#creating-a-bot-in-the-test-environment), может принимать значения:
  * `const EnvProduction = "prod"` (или по умолчанию, `""`) - промышленное окружение;
  * `const EnvTest = "test"` - тестовое окружение.

### Выполнение запроса

Для понимания механики выполнения запросов в Telegram Bots API через эту реализацию на Go удобно рассматривать шаги на основе простого запроса, [получения информации о пользователе](https://core.telegram.org/bots/api#getme). За имплементацию этого запроса отвечает объект `requests.GetMe`.

Все другие объекты-методы также расположены в пространстве имён `requests`.

Чтобы выполнить запрос, нужно:
1. Создать объект запроса, заполнить передаваемые аргументы:
   ```go
   req := requests.GetMe{}
   ```
2. Вызвать метод `Call` объекта запроса, передать в него контекст и ссылку на объект `telegram.Bot`:
   ```go
   req.Call(context.Background(), bot)
   ```
3. Выполнить преобразование типа вернувшегося значения:
   ```go
   user := res.(*telegram.User)
   ```

#### Полный пример:

```go
package main

import (
    "context"
    "github.com/temoon/telegram-bots-api"
    "github.com/temoon/telegram-bots-api/requests"
    "log"
    "os"
    "time"
)

func main() {
    botOpts := telegram.BotOpts{
        Token:   os.Getenv("TELEGRAM_TOKEN"),
        Timeout: time.Minute,
    }

    bot := telegram.NewBot(&botOpts)

    req := requests.GetMe{}

    var err error
    var res interface{}
    if res, err = req.Call(context.Background(), bot); err != nil {
        log.Fatalln(err)
    }

    if user, ok := res.(*telegram.User); ok {
        if user.Username != nil {
            log.Println(*user.Username)
        }
    }
}
```
