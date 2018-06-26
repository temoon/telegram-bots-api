package telegram

import (
    "net/http"
    "net/url"
    "os"
    "testing"
)

func TestNewBot(t *testing.T) {
    rawProxyURL := os.Getenv("HTTP_PROXY")

    httpClient := &http.Client{}

    if rawProxyURL != "" {
        proxyURL, err := url.Parse(rawProxyURL)

        if err != nil {
            t.Error(err)
        }

        httpClient.Transport = &http.Transport{
            Proxy: http.ProxyURL(proxyURL),
        }
    }

    opts := &BotOpts{
        Token:  os.Getenv("TOKEN"),
        Client: httpClient,
    }

    NewBot(opts)
}
