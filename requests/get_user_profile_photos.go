package requests

import (
    "strconv"
)

type GetUserProfilePhotos struct {
    UserID int
    Offset int
    Limit  int
}

func (r *GetUserProfilePhotos) IsMultipart() bool {
    return false
}

func (r *GetUserProfilePhotos) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["user_id"] = []interface{}{strconv.Itoa(r.UserID)}

    if r.Offset != 0 {
        values["offset"] = []interface{}{strconv.Itoa(r.Offset)}
    }

    if r.Limit != 0 {
        values["limit"] = []interface{}{strconv.Itoa(r.Limit)}
    }

    return
}
