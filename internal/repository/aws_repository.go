package repository

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
)

type IAwsRepository interface {
	SendEmail(ctx context.Context, email *entity.Email) error
	SaveAtS3(ctx context.Context, path, filename string, file []byte) (string, error)
}

type AwsRepository struct {
	sess *session.Session
}

func NewAwsRepository(sess *session.Session) *AwsRepository {
	return &AwsRepository{
		sess: sess,
	}
}

// Given path without "/" and filename to create a complete path.
func (ar *AwsRepository) SaveAtS3(ctx context.Context, path, filename string, file []byte) (string, error) {

	conf := config.Get()

	svc := s3.New(ar.sess)

	filename = fmt.Sprintf("%s/%s.png", path, filename)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(conf.Cloud.BucketName),
		Key:         aws.String(filename), // Maintain the same filename in the bucket
		Body:        bytes.NewReader(file),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", conf.Cloud.BucketName, filename)

	return url, nil
}

func (ar *AwsRepository) SendEmail(ctx context.Context, email *entity.Email) error {

	conf := config.Get()

	svc := ses.New(ar.sess)

	emailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(email.Recipient)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(email.Body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(email.Subject),
			},
		},
		Source: aws.String(conf.Cloud.Source),
	}

	_, err := svc.SendEmail(emailInput)

	if err != nil {
		return err
	}

	return nil

}
