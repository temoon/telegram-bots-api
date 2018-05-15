package requests

type GetFile struct {
    FileID string
}

func (r *GetFile) IsMultipart() bool {
    return false
}

func (r *GetFile) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["file_id"] = r.FileID

    return
}
