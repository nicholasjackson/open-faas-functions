package function

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Request defines the function request payload
type Request struct {
	Filename    string `json:"filename"`
	Permissions string `json:"permissions"`
	DataBase64  string `json:"data_base64"`
}

// Handle a serverless request
func Handle(req []byte) string {
	awsAccessKey := os.Getenv("accessKeyID")
	awsSecret := os.Getenv("secretAccessKey")
	bucketName := os.Getenv("bucket")
	region := os.Getenv("region")

	request := Request{}
	err := json.Unmarshal(req, &request)
	if err != nil {
		return "Error: invalid request object"
	}

	data, err := base64.StdEncoding.DecodeString(request.DataBase64)
	if err != nil {
		return "Invalid data, not base64"
	}

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, "")

	_, err = creds.Get()
	if err != nil {
		return fmt.Sprintf("bad credentials: %s", err)
	}

	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	sess := session.New(cfg)
	svc := s3manager.NewUploader(sess)

	var result *s3manager.UploadOutput
	result, err = svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(request.Filename),
		Body:   bytes.NewReader(data),
		ACL:    aws.String(request.Permissions),
	})

	if err != nil {
		return fmt.Sprintf("Error uploading file: %s", err)
	}

	return result.Location
}
