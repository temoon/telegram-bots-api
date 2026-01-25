package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type RefundStarPayment struct {
	TelegramPaymentChargeId string
	UserId                  int64
}

func (r *RefundStarPayment) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "refundStarPayment", r, response)
	return
}

func (r *RefundStarPayment) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["telegram_payment_charge_id"] = r.TelegramPaymentChargeId

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *RefundStarPayment) GetFiles() (files map[string]io.Reader) {
	return
}
