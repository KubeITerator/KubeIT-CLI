package functions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"kubeitcli/httpd"
	"kubeitcli/httpd/requests"
	"os"
	"path/filepath"
)

func SplitInChunks(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

func UploadToS3(filename string, rclient *httpd.RequestClient) (url string, err error) {

	var multi bool
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", errors.New(fmt.Sprintf("path: %v does not exist", filename))
	}

	fmt.Println("[S3-UPLOAD] Reading file data: " + filename)
	data, err := ioutil.ReadFile(filename)

	if len(data) > 5e+9 {
		multi = true
	} else {
		multi = false
	}

	if err != nil {
		return "", err
	}

	fmt.Println("[S3-UPLOAD] Initializing S3 upload")
	passkey, err := requests.InitS3(filepath.Base(filename), multi, rclient)

	if passkey == "" {
		return "", errors.New("[S3-UPLOAD] initialization failed")
	}

	if err != nil {
		fmt.Println("[S3-UPLOAD] Error in S3 initialization")
		return "", err
	}

	if multi {
		chunks := SplitInChunks(data, 100000000)
		for part, chunk := range chunks {
			fmt.Println(fmt.Sprintf("[S3-UPLOAD] Uploading part: %v/%v", part, len(chunks)))
			upURL, err := requests.GetS3Upload(passkey, rclient)
			if err != nil {
				fmt.Println("[S3-UPLOAD] Error in getting Upload URL")
				return "", err
			}
			err = rclient.S3UploadRequest(upURL, chunk)
			if err != nil {
				fmt.Println("[S3-UPLOAD] Error in S3UploadRequest")
				return "", err
			}
		}
	} else {
		fmt.Println("[S3-UPLOAD] Uploading part: 1/1")
		upURL, err := requests.GetS3Upload(passkey, rclient)
		if err != nil {
			fmt.Println("[S3-UPLOAD] Error in getting Upload URL")
			return "", err
		}
		err = rclient.S3UploadRequest(upURL, data)
		if err != nil {
			fmt.Println("[S3-UPLOAD] Error in S3UploadRequest")
			return "", err
		}
	}

	failed, err := requests.FinishS3Upload(passkey, rclient)

	if err != nil {
		fmt.Println("[S3-UPLOAD] Error in finishing S3 Upload")
		return "", err
	}

	if failed {
		fmt.Println("[S3-UPLOAD] Failed finishing S3 Upload")
		return "", errors.New("failed finishing S3 upload")
	}
	fmt.Println("[S3-UPLOAD] Upload finished")

	return requests.GetDownload(passkey, rclient)

}
