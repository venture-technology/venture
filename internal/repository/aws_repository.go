package repository

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/venture-technology/venture/config"
)

type IAWSRepository interface {
	SaveImageOnAWSBucket(ctx context.Context, image []byte, filename string) (string, error)
}

type AWSRepository struct {
	sess *session.Session
}

func NewAWSRepository(sess *session.Session) *AWSRepository {
	return &AWSRepository{
		sess: sess,
	}
}

func (ar *AWSRepository) SaveImageOnAWSBucket(ctx context.Context, image []byte, filename string) (string, error) {

	conf := config.Get()

	svc := s3.New(ar.sess)

	filename = fmt.Sprintf("qrcodes/%s.png", filename)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(conf.Cloud.BucketName),
		Key:         aws.String(filename), // Maintain the same filename in the bucket
		Body:        bytes.NewReader(image),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		return "", err
	}

	qrCode := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", conf.Cloud.BucketName, filename)

	return qrCode, nil
}
