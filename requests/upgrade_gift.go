package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type UpgradeGift struct {
	BusinessConnectionId string
	OwnedGiftId          string
	KeepOriginalDetails  *bool
	StarCount            *int64
}

func (r *UpgradeGift) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "upgradeGift", r, response)
	return
}

func (r *UpgradeGift) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	values["owned_gift_id"] = r.OwnedGiftId

	if r.KeepOriginalDetails != nil {
		if *r.KeepOriginalDetails {
			values["keep_original_details"] = "1"
		} else {
			values["keep_original_details"] = "0"
		}
	}

	if r.StarCount != nil {
		values["star_count"] = strconv.FormatInt(*r.StarCount, 10)
	}

	return
}

func (r *UpgradeGift) GetFiles() (files map[string]io.Reader) {
	return
}
