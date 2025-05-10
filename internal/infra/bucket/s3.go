package bucket

import (
	"bytes"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

type S3Impl struct {
	bucket *s3.S3
}

func NewS3Impl() *S3Impl {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(viper.GetString("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			viper.GetString("AWS_ACCESS_KEY"),
			viper.GetString("AWS_SECRET_KEY"),
			viper.GetString("AWS_TOKEN"),
		),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	return &S3Impl{
		bucket: s3.New(sess),
	}
}

// Given path without "/" and filename to create a complete path.
func (s3Impl *S3Impl) Save(bucket, path, filename string, file []byte) (string, error) {
	filename = fmt.Sprintf("%s/%s.png", path, filename)

	_, err := s3Impl.bucket.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename), // Maintain the same filename in the bucket
		Body:        bytes.NewReader(file),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, filename)
	return url, nil
}

func (s3Impl *S3Impl) List(bucket, path string) ([]string, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(path), // filter by path
	}

	var links []string
	err := s3Impl.bucket.ListObjectsV2Pages(input, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, obj := range page.Contents {
			publicURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, *obj.Key)
			links = append(links, publicURL)
		}
		return !lastPage
	})
	if err != nil {
		return nil, err
	}

	if len(links) == 0 {
		return links, nil
	}

	// remove the first link from the list, because it is the path itself
	return links[1:], nil
}

func (s3Impl *S3Impl) SaveWithType(bucket, path, filaneme string, file []byte, contentType string) (string, error) {
	filename := fmt.Sprintf("%s/%s", path, filaneme)

	_, err := s3Impl.bucket.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(file),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, filename)
	return url, nil
}

func (s3Impl *S3Impl) PDF() string {
	return "application/pdf"
}

func (s3Impl *S3Impl) HTML() string {
	return "text/html"
}

func (s3Impl *S3Impl) PNG() string {
	return "image/png"
}
