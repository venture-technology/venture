package bucket

import (
	"bytes"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/venture-technology/venture/config"
)

type S3Impl struct {
	sess   *session.Session
	config *config.Config
}

func NewS3Impl(config config.Config) *S3Impl {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(config.Cloud.AccessKey, config.Cloud.SecretKey, config.Cloud.Token),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	return &S3Impl{
		sess:   sess,
		config: &config,
	}
}

// Given path without "/" and filename to create a complete path.
func (s3Impl *S3Impl) Save(path, filename string, file []byte) (string, error) {
	svc := s3.New(s3Impl.sess)

	filename = fmt.Sprintf("%s/%s.png", path, filename)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s3Impl.config.Cloud.BucketName),
		Key:         aws.String(filename), // Maintain the same filename in the bucket
		Body:        bytes.NewReader(file),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s3Impl.config.Cloud.BucketName, filename)

	return url, nil
}

func (s3Impl *S3Impl) List(path string) ([]string, error) {
	svc := s3.New(s3Impl.sess)

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s3Impl.config.Cloud.BucketName),
		Prefix: aws.String(path), // filter by path
	}

	var links []string

	err := svc.ListObjectsV2Pages(input, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, obj := range page.Contents {
			publicURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s3Impl.config.Cloud.BucketName, *obj.Key)
			links = append(links, publicURL)
		}
		return !lastPage
	})

	if err != nil {
		return nil, err
	}

	// remove the first link from the list, because it is the path itself
	return links[1:], nil
}
