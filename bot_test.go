package telegram_test

import (
    "net/http"
    "net/url"
    "os"
    "testing"

    "github.com/temoon/go-telegram-bots-api"
)

func TestNewBot(t *testing.T) {
    newBot(t)
}

func TestBot_GetMe(t *testing.T) {
    bot := newBot(t)
    me, err := bot.GetMe()

    if err != nil {
        t.Error(err)
    }

    if !me.IsBot {
        t.Error("Expected to be a bot")
    }

    t.Log("ID:", me.ID)
    t.Log("Username:", me.Username)
    t.Log("First name:", me.FirstName)
    t.Log("Last name:", me.LastName)
    t.Log("Canguage code:", me.LanguageCode)
}

func newBot(t *testing.T) telegram.BotAPI {
    token := os.Getenv("TOKEN")

    if token == "" {
        t.Skip("No token provided")
    }

    rawProxyURL := os.Getenv("HTTP_PROXY")
    httpClient := new(http.Client)

    if rawProxyURL != "" {
        proxyURL, err := url.Parse(rawProxyURL)

        if err != nil {
            panic(err)
        }

        httpClient.Transport = &http.Transport{
            Proxy: http.ProxyURL(proxyURL),
        }
    }

    opts := &telegram.BotOpts{
        Token:  token,
        Client: httpClient,
    }

    return telegram.NewBot(opts)
}
