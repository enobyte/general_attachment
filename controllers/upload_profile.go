package controllers

import (
	"fmt"
	"general_attachment/models"
	"general_attachment/utils"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var UploadFile = func(w http.ResponseWriter, r *http.Request) {
	user := &models.AttachmentModel{}

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("claim_doc")
	typeAttach := r.FormValue("type")
	payorCode := r.FormValue("payor_code")
	typeFile := r.FormValue("type_file")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	tempFile, err := ioutil.TempFile("/usr/share/nginx/html/claimrembursement", strings.Trim(payorCode, " ")+"_*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Chmod(0755)
	if err != nil {
		log.Println(err)
	}
	tempFile.Write(fileBytes)

	user = models.InsertMetaAttachment(typeAttach, filepath.Base(tempFile.Name()), payorCode, typeFile)
	resp := utils.Message(utils.SuccesReq(), true, "data uploaded")
	resp["data"] = user
	utils.Respond(w, resp)
}
