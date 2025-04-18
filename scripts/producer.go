package main

import "fmt"

func main() {
	// conf, err := config.Load("../config/config.yaml")
	// if err != nil {
	// 	log.Fatalf("error loading config: %s", err.Error())
	// }

	// brokers := []string{conf.Uchiha.Address}
	// topic := conf.Uchiha.Queue

	// writer := kafka.Writer{
	// 	Addr:     kafka.TCP(brokers...),
	// 	Topic:    topic,
	// 	Balancer: &kafka.LeastBytes{},
	// }

	// messageText := `{"recipient": "teste@gmail.com", "subject": "envinado mensagem pra fila", "body": "leitura de mensagem"}`

	// message := kafka.Message{
	// 	Key:   []byte("key"),
	// 	Value: []byte(messageText),
	// }

	// err = writer.WriteMessages(context.Background(), message)
	// if err != nil {
	// 	log.Fatalf("Failed to produce message: %v", err)
	// }

	// if err := writer.Close(); err != nil {
	// 	log.Fatalf("Failed to close writer: %v", err)
	// }
	fmt.Println("Message sent to RabbitMQ topic successfully")
}
