package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	stage           = os.Getenv("STAGE")
	region          = os.Getenv("REGION")
	s3Bucket        = os.Getenv("S3_BUCKET")
	localstackURLS3 = os.Getenv("LOCALSTACK_URL_S3")
)

func fileGet(path string) (*bytes.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(fileBytes)

	return r, nil
}

func svcGet() *s3.S3 {
	svc := s3.New(session.New(), &aws.Config{
		Region: aws.String(region),
	})
	return svc
}

func svcLocalstackGet() *s3.S3 {
	svc := s3.New(session.New(), &aws.Config{
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(localstackURLS3),
	})
	return svc
}

func main() {
	// 画像ファイル取得
	file, err := fileGet("./450-150.png")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 各設定項目
	input := &s3.PutObjectInput{
		Bucket:               aws.String(s3Bucket),
		ACL:                  aws.String("public-read"),
		ServerSideEncryption: aws.String("AES256"),
		Key:                  aws.String("450-150.png"),
		Body:                 file,
	}

	var svc *s3.S3
	if stage == "prod" {
		// 本番環境 S3
		svc = svcGet()
	} else {
		// 開発環境 Localstack S3
		svc = svcLocalstackGet()
	}

	// アップロード
	if _, err := svc.PutObject(input); err != nil {
		log.Fatal(err.Error())
	}
}
