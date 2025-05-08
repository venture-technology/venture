package email

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/entity"
)

type SesImpl struct {
	ses *ses.SES
}

func NewSesImpl() *SesImpl {
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

	return &SesImpl{
		ses: ses.New(sess),
	}
}

func (sesImpl *SesImpl) SendEmail(email *entity.Email) error {
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
		Source: aws.String(viper.GetString("AWS_SES_EMAIL_FROM")),
	}

	_, err := sesImpl.ses.SendEmail(emailInput)

	if err != nil {
		return err
	}

	return nil
}
