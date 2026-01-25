package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type GetUserGifts struct {
	UserId                      int64
	ExcludeFromBlockchain       *bool
	ExcludeLimitedNonUpgradable *bool
	ExcludeLimitedUpgradable    *bool
	ExcludeUnique               *bool
	ExcludeUnlimited            *bool
	Limit                       *int64
	Offset                      *string
	SortByPrice                 *bool
}

func (r *GetUserGifts) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.OwnedGifts)
	err = b.CallMethod(ctx, "getUserGifts", r, response)
	return
}

func (r *GetUserGifts) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.ExcludeFromBlockchain != nil {
		if *r.ExcludeFromBlockchain {
			values["exclude_from_blockchain"] = "1"
		} else {
			values["exclude_from_blockchain"] = "0"
		}
	}

	if r.ExcludeLimitedNonUpgradable != nil {
		if *r.ExcludeLimitedNonUpgradable {
			values["exclude_limited_non_upgradable"] = "1"
		} else {
			values["exclude_limited_non_upgradable"] = "0"
		}
	}

	if r.ExcludeLimitedUpgradable != nil {
		if *r.ExcludeLimitedUpgradable {
			values["exclude_limited_upgradable"] = "1"
		} else {
			values["exclude_limited_upgradable"] = "0"
		}
	}

	if r.ExcludeUnique != nil {
		if *r.ExcludeUnique {
			values["exclude_unique"] = "1"
		} else {
			values["exclude_unique"] = "0"
		}
	}

	if r.ExcludeUnlimited != nil {
		if *r.ExcludeUnlimited {
			values["exclude_unlimited"] = "1"
		} else {
			values["exclude_unlimited"] = "0"
		}
	}

	if r.Limit != nil {
		values["limit"] = strconv.FormatInt(*r.Limit, 10)
	}

	if r.Offset != nil {
		values["offset"] = *r.Offset
	}

	if r.SortByPrice != nil {
		if *r.SortByPrice {
			values["sort_by_price"] = "1"
		} else {
			values["sort_by_price"] = "0"
		}
	}

	return
}

func (r *GetUserGifts) GetFiles() (files map[string]io.Reader) {
	return
}
