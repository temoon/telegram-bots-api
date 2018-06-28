package requests_test

import (
    "bytes"
    "encoding/json"
    "io"
    "os"
    "runtime"
    "strconv"
    "testing"

    "go-telegram-bots-api"
    "go-telegram-bots-api/requests"
)

func TestAddStickerToSet_IsMultipart(t *testing.T) {
    req := &requests.AddStickerToSet{}

    if !req.IsMultipart() {
        t.Fail()
    }
}

func TestAddStickerToSet_GetValues(t *testing.T) {
    maskPosition := &telegram.MaskPosition{}

    _, filename, _, ok := runtime.Caller(0)

    if !ok {
        t.Fail()
    }

    inputFile, err := os.Open(filename)

    if err != nil {
        t.Error(err)
    }

    reqs := []requests.AddStickerToSet{
        {UserID: 0, Name: "Name", PNGSticker: "http://example.com/sticker.png", Emojis: "Emojis", MaskPosition: maskPosition},
        {UserID: 0, Name: "Name", PNGSticker: inputFile, Emojis: "Emojis", MaskPosition: maskPosition},
    }

    for _, req := range reqs {
        if !req.IsMultipart() {
            t.Fail()
        }

        values, err := req.GetValues()

        if err != nil {
            t.Error(err)
        }

        if value, ok := values["user_id"]; !ok {
            t.Error("missing user_id")
        } else if value != strconv.FormatUint(uint64(req.UserID), 10) {
            t.Error("invalid user_id")
        }

        if value, ok := values["name"]; !ok {
            t.Error("missing name")
        } else if value != req.Name {
            t.Error("invalid name")
        }

        if value, ok := values["png_sticker"]; !ok {
            t.Error("missing png_sticker")
        } else {
            switch inputFile := value.(type) {
            case string:
                if inputFile != req.PNGSticker {
                    t.Error("invalid png_sticker")
                }
            case io.Reader:
                var a []byte
                var b []byte

                if _, err := io.ReadFull(req.PNGSticker.(io.Reader), a); err != nil {
                    t.Error(err)
                }

                if _, err := io.ReadFull(inputFile, b); err != nil {
                    t.Error(err)
                }

                if bytes.Compare(a, b) != 0 {
                    t.Error("invalid png_sticker")
                }
            default:
                t.Error("invalid png_sticker")
            }
        }

        if value, ok := values["emojis"]; !ok {
            t.Error("missing emojis")
        } else if value != req.Emojis {
            t.Error("invalid emojis")
        }

        if value, ok := values["mask_position"]; ok {
            data := new(telegram.MaskPosition)

            if err := json.Unmarshal([]byte(value.(string)), data); err != nil {
                t.Error(err)
            }
        }
    }

    req := requests.AddStickerToSet{}

    if _, err := req.GetValues(); err == nil {
        t.Fail()
    }
}
