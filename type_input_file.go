package telegram

import (
	"encoding/json"
	"io"
)

type InputFile struct {
	fileId   string
	file     io.Reader
	fileName string
}

//goland:noinspection GoUnusedExportedFunction
func NewInputFile(fileId string, file io.Reader, fileName string) (inputFile InputFile) {
	if len(fileId) > 0 {
		inputFile.fileId = fileId
	} else if file != nil && len(fileName) > 0 {
		inputFile.file = file
		inputFile.fileName = fileName
	}

	return
}

//goland:noinspection GoMixedReceiverTypes
func (f InputFile) MarshalJSON() (data []byte, err error) {
	return json.Marshal(f.String())
}

//goland:noinspection GoMixedReceiverTypes
func (f *InputFile) UnmarshalJSON(data []byte) (err error) {
	var fileId string
	if err = json.Unmarshal(data, &fileId); err != nil {
		return
	}

	f.fileId = fileId

	return
}

//goland:noinspection GoMixedReceiverTypes
func (f *InputFile) String() string {
	if f.fileId != "" {
		return f.fileId
	} else if f.HasFile() {
		return "attach://" + f.GetFormFieldName()
	}

	return ""
}

//goland:noinspection GoMixedReceiverTypes
func (f *InputFile) GetValue() interface{} {
	if len(f.fileId) > 0 {
		return f.fileId
	} else if f.HasFile() {
		return f.file
	}

	return ""
}

//goland:noinspection GoMixedReceiverTypes
func (f *InputFile) HasFile() bool {
	return f.file != nil && len(f.fileName) > 0
}

//goland:noinspection GoMixedReceiverTypes
func (f *InputFile) GetFormFieldName() string {
	return "file_" + f.fileName
}

//goland:noinspection GoMixedReceiverTypes
func (f *InputFile) GetFile() io.Reader {
	return f.file
}
