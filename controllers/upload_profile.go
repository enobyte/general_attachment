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

	_ = r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("claim_doc")
	id := r.FormValue("id")
	//fileName := r.FormValue("file_name")
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
	tempFile, err := ioutil.TempFile("/usr/share/nginx/html/claimrembursement", strings.Split(handler.Filename, ".")[0]+"_*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	_ = tempFile.Chmod(0755)
	if err != nil {
		log.Println(err)
	}
	_, err = tempFile.Write(fileBytes)

	if err != nil {
		resp := utils.Message(utils.BadReq(), false, "uploaded failed", nil)
		utils.Respond(w, resp)
	} else {
		user = models.InsertMetaAttachment(id, payorCode, typeFile, filepath.Base(tempFile.Name()))
		resp := utils.Message(utils.SuccesReq(), true, "data uploaded", user)
		utils.Respond(w, resp)
	}

}
