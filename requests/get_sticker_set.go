package requests

type GetStickerSet struct {
    Name string
}

func (r *GetStickerSet) IsMultipart() bool {
    return false
}

func (r *GetStickerSet) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["name"] = r.Name

    return
}
