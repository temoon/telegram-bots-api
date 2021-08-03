package requests

import (
	"strconv"
)

type SetStickerPositionInSet struct {
	Position uint64
	Sticker  string
}

func (r *SetStickerPositionInSet) IsMultipart() bool {
	return false
}

func (r *SetStickerPositionInSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["position"] = strconv.FormatUint(r.Position, 10)

	values["sticker"] = r.Sticker

	return
}
