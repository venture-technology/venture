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
}

type AwsRepository struct {
	sess *session.Session
}

func NewAwsRepository(sess *session.Session) *AwsRepository {
	return &AwsRepository{
		sess: sess,
	}
}

func (ar *AwsRepository) SaveImageOnAWSBucket(ctx context.Context, image []byte, filename string) (string, error) {

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
