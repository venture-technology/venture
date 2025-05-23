package queue

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/value"
)

const (
	BatchSize         = 10
	VisibilityTimeout = 30
	WaitSeconds       = 10
)

type SQSImpl struct {
	Client *sqs.SQS
}

func NewSQSQueue() SQSImpl {
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

	return SQSImpl{
		Client: sqs.New(sess),
	}
}

func (s SQSImpl) SendMessage(queue, message string) error {
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(queue),
	}

	_, err := s.Client.SendMessage(sendMessageInput)
	if err != nil {
		log.Printf("sqs - failed to send message to queue: %v", err)
		return err
	}

	return nil
}

func (s SQSImpl) SendFifoMessage(queue, message, group string) error {
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody:    aws.String(message),
		QueueUrl:       aws.String(queue),
		MessageGroupId: aws.String(group),
	}

	_, err := s.Client.SendMessage(sendMessageInput)
	if err != nil {
		log.Printf("sqs - failed to send message to queue: %v", err)
		return err
	}

	return nil
}

func (s SQSImpl) PullMessages(queue string) ([]*value.CreateMessage, error) {
	msgOutput, err := s.Client.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &queue,
		MaxNumberOfMessages: aws.Int64(BatchSize),
		VisibilityTimeout:   aws.Int64(VisibilityTimeout),
		WaitTimeSeconds:     aws.Int64(WaitSeconds),
	})
	if err != nil {
		return nil, fmt.Errorf("sqs - failed to receive message from queue: %v", err)
	}

	totalMessages := len(msgOutput.Messages)
	messages := make([]*value.CreateMessage, 0, totalMessages)
	msg := make(chan *value.CreateMessage)

	for _, rawMessage := range msgOutput.Messages {
		go func(queue *string, rawMessage *sqs.Message, msg chan<- *value.CreateMessage) {
			msg <- &value.CreateMessage{
				QueueURL:      *queue,
				ReceiptHandle: *rawMessage.ReceiptHandle,
				Body:          *rawMessage.Body,
			}
		}(&queue, rawMessage, msg)
	}

	for i := 0; i < totalMessages; i++ {
		messages = append(messages, <-msg)
	}

	close(msg)
	return messages, nil
}

func (s SQSImpl) DeleteMessage(queue, identifier string) error {
	_, err := s.Client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &queue,
		ReceiptHandle: &identifier,
	})

	if err != nil {
		log.Printf("sqs - failed to delete message from queue: %v", err)
		return err
	}

	return nil
}
