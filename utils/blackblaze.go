package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type UploadURLType struct {
	AuthorizationToken string `json:"authorizationToken"`
	UploadURL          string `json:"uploadUrl"`
}

func GetB2AuthToken() (string, error) {
	username := "5640111a3b80"
	password := "00397a4a6a46c8de77b793cf0853a0ceb18a62a29a"

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.backblazeb2.com/b2api/v3/b2_authorize_account", nil)

	if err != nil {
		fmt.Println("Some error in getting auth token for B2")
		return "", err
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Some error in getting auth token for B2")
		return "", err
	}

	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Some error in getting auth token for B2")
		return "", err
	}

	fmt.Println("bodyText", string(bodyText))

	type ResponseData struct {
		AuthorizationToken string `json:"authorizationToken"`
	}

	response := ResponseData{}
	if err := json.Unmarshal(bodyText, &response); err != nil {
		fmt.Println("Some error in getting auth token for B2")
		return "", err
	}

	fmt.Println(response)

	return response.AuthorizationToken, nil
}

func GetB2UploadURL(token string) (UploadURLType, error) {
	bucketID := "85862470016141ba933b0810"

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api003.backblazeb2.com/b2api/v3/b2_get_upload_url?bucketId="+bucketID, nil)

	if err != nil {
		fmt.Println("Some error in getting upload URL for B2")
		return UploadURLType{AuthorizationToken: "", UploadURL: ""}, nil
	}

	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Some error in getting upload URL for B2")
		return UploadURLType{AuthorizationToken: "", UploadURL: ""}, nil
	}

	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Some error in getting upload URL for B2")
		return UploadURLType{AuthorizationToken: "", UploadURL: ""}, nil
	}

	response := UploadURLType{}

	if err := json.Unmarshal(bodyText, &response); err != nil {
		fmt.Println("Some error in getting upload URL for B2")
		return UploadURLType{AuthorizationToken: "", UploadURL: ""}, nil
	}

	return UploadURLType{AuthorizationToken: response.AuthorizationToken, UploadURL: response.UploadURL}, nil
}

func UploadB2File(file *multipart.FileHeader, uploadURL UploadURLType) (string, error) {
	fileContent, err := file.Open()
	if err != nil {
		fmt.Println("Some error in uploading file to B2")
		return "", nil
	}

	defer fileContent.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", file.Filename)

	if err != nil {
		fmt.Println("Some error in uploading file to B2")
		return "", nil
	}

	_, err = io.Copy(part, fileContent)

	if err != nil {
		fmt.Println("Some error in uploading file to B2")
		return "", nil
	}

	err = writer.Close()

	if err != nil {
		fmt.Println("Some error in uploading file to B2")
		return "", nil
	}

	client := http.Client{}

	req, err := http.NewRequest("POST", uploadURL.UploadURL, &body)

	if err != nil {
		fmt.Println("Some error in uploading file to B2")
		return "", nil
	}

	uniqueID := uuid.New()
	fileName := strings.Replace(uniqueID.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]

	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", uploadURL.AuthorizationToken)
	req.Header.Add("X-Bz-File-Name", fmt.Sprintf("%s.%s", fileName, fileExt))
	req.Header.Add("X-Bz-Content-Sha1", "do_not_verify")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Some error in uploading file to B2")
		return "", nil
	}

	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Some error in uploading file to B2")
		return "", nil
	}

	type FileURLType struct {
		FileName string `json:"fileName"`
	}

	response := FileURLType{}
	if err := json.Unmarshal(bodyText, &response); err != nil {
		fmt.Println("Some error in uploading file to B2")
		return "", nil
	}

	return response.FileName, nil
}
