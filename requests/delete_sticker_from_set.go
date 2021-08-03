package requests

type DeleteStickerFromSet struct {
	Sticker string
}

func (r *DeleteStickerFromSet) IsMultipart() bool {
	return false
}

func (r *DeleteStickerFromSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker"] = r.Sticker

	return
}
