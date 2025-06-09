package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type GetBusinessAccountGifts struct {
	BusinessConnectionId string
	ExcludeLimited       *bool
	ExcludeSaved         *bool
	ExcludeUnique        *bool
	ExcludeUnlimited     *bool
	ExcludeUnsaved       *bool
	Limit                *int64
	Offset               *string
	SortByPrice          *bool
}

func (r *GetBusinessAccountGifts) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.OwnedGifts)
	err = b.CallMethod(ctx, "getBusinessAccountGifts", r, response)
	return
}

func (r *GetBusinessAccountGifts) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	if r.ExcludeLimited != nil {
		if *r.ExcludeLimited {
			values["exclude_limited"] = "1"
		} else {
			values["exclude_limited"] = "0"
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
