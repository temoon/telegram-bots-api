package requests

import (
	"strconv"
)

type GetUserProfilePhotos struct {
	Limit  uint64
	Offset uint64
	UserId uint64
}

func (r *GetUserProfilePhotos) IsMultipart() bool {
	return false
}

func (r *GetUserProfilePhotos) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Limit != 0 {
		values["limit"] = strconv.FormatUint(r.Limit, 10)
	}

	if r.Offset != 0 {
		values["offset"] = strconv.FormatUint(r.Offset, 10)
	}

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
