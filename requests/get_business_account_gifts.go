package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type GetBusinessAccountGifts struct {
	BusinessConnectionId        string
	ExcludeFromBlockchain       *bool
	ExcludeLimitedNonUpgradable *bool
	ExcludeLimitedUpgradable    *bool
	ExcludeSaved                *bool
	ExcludeUnique               *bool
	ExcludeUnlimited            *bool
	ExcludeUnsaved              *bool
	Limit                       *int64
	Offset                      *string
	SortByPrice                 *bool
}

func (r *GetBusinessAccountGifts) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.OwnedGifts)
	err = b.CallMethod(ctx, "getBusinessAccountGifts", r, response)
	return
}

func (r *GetBusinessAccountGifts) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

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

	if r.ExcludeSaved != nil {
		if *r.ExcludeSaved {
			values["exclude_saved"] = "1"
		} else {
			values["exclude_saved"] = "0"
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

	if r.ExcludeUnsaved != nil {
		if *r.ExcludeUnsaved {
			values["exclude_unsaved"] = "1"
		} else {
			values["exclude_unsaved"] = "0"
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

func (r *GetBusinessAccountGifts) GetFiles() (files map[string]io.Reader) {
	return
}
