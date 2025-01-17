package email

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
)

type SesImpl struct {
	sess   *session.Session
	config *config.Config
}

func NewSesImpl(config config.Config) *SesImpl {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(config.Cloud.AccessKey, config.Cloud.SecretKey, config.Cloud.Token),
	})
	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	return &SesImpl{
		sess:   sess,
		config: &config,
	}
}

func (sesImpl *SesImpl) SendEmail(email *entity.Email) error {
	svc := ses.New(sesImpl.sess)

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
		Source: aws.String(sesImpl.config.Cloud.Source),
	}

	_, err := svc.SendEmail(emailInput)

	if err != nil {
		return err
	}

	return nil
}
