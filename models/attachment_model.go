package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type AttachmentModel struct {
	IdAttachment string         `json:"id_attachment"`
	PayorCode    string         `json:"payor_code"`
	Createat     time.Time      `json:"createat"`
	TypeFile     string         `json:"type_file"`
	File         postgres.Jsonb `json:"file"`
}

type FileData struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

func InsertMetaAttachment(id_attach string, payorCode string, typeFile string, fileName string) *AttachmentModel {
	baseUrl := "http://103.107.103.56/claimrembursement/"
	attachmentModel := &AttachmentModel{}
	fileData := []FileData{{Filename: fileName, Url: baseUrl + fileName}}
	outputFile, err := json.MarshalIndent(fileData, "", " ")
	if err != nil {
		print(err)
	}
	fmt.Println(string(outputFile))
	GetDB().Raw("select * from fileattachment.insert_attachment('" + id_attach + "','" + payorCode + "', '" + typeFile + "', '" + string(outputFile) + "')").Find(&attachmentModel)
	return attachmentModel
}
