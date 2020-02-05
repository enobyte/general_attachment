package models

import "time"

type AttachmentModel struct {
	IdAttachment string    `json:"id_attachment"`
	Type         string    `json:"type"`
	Url          string    `json:"url"`
	PayorCode    string    `json:"payor_code"`
	Createat     time.Time `json:"createat"`
	TypeFile     string    `json:"type_file"`
}

func InsertMetaAttachment(typeAttach string, url string, payorCode string, typeFile string) *AttachmentModel {
	baseUrl := "http://103.107.103.56/claimrembursement/"
	attachmentModel := &AttachmentModel{}
	GetDB().Raw("select * from fileattachment.insert_attachment('" + typeAttach + "','" + baseUrl + url + "','" + payorCode + "', '" + typeFile + "')").Find(&attachmentModel)
	return attachmentModel
}
