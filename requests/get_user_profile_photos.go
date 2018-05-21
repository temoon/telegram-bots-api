package requests

import (
    "strconv"
)

type GetUserProfilePhotos struct {
    UserID uint32
    Offset uint32
    Limit  uint32
}

func (r *GetUserProfilePhotos) IsMultipart() bool {
    return false
}

func (r *GetUserProfilePhotos) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["user_id"] = strconv.FormatUint(uint64(r.UserID), 10)

    if r.Offset != 0 {
        values["offset"] = strconv.FormatUint(uint64(r.Offset), 10)
    }

    if r.Limit != 0 {
        values["limit"] = strconv.FormatUint(uint64(r.Limit), 10)
    }

    return
}
