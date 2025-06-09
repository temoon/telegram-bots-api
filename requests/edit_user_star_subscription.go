package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type EditUserStarSubscription struct {
	IsCanceled              bool
	TelegramPaymentChargeId string
	UserId                  int64
}

func (r *EditUserStarSubscription) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "editUserStarSubscription", r, response)
	return
}

func (r *EditUserStarSubscription) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.IsCanceled {
		values["is_canceled"] = "1"
	} else {
		values["is_canceled"] = "0"
	}

	values["telegram_payment_charge_id"] = r.TelegramPaymentChargeId

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *EditUserStarSubscription) GetFiles() (files map[string]io.Reader) {
	return
}
