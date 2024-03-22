package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/akshay0074700747/email-service/email"
)

type SendMail struct {
	Email   string `json:"Email"`
	Message string `json:"Message"`
}

var (
	Mailer *email.SMTPConfig
)

func Getmailer(mail *email.SMTPConfig) {
	Mailer = mail
}

func StartServing() {
	// consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
	// 	"bootstrap.servers":        "host.docker.internal:9092",
	// 	"group.id":                 "EmailConsumers",
	// 	"auto.offset.reset":        "earliest",
	// 	"allow.auto.create.topics": true})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// topic := "Emailsender"
	// err = consumer.Assign([]kafka.TopicPartition{
	// 	{
	// 		Topic:     &topic,
	// 		Partition: 0,
	// 		Offset:    kafka.OffsetStored,
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for {
	// 	ev := consumer.Poll(10)
	// 	switch e := ev.(type) {
	// 	case *kafka.Message:
	// 		go func() {
	// 			log.Println(string(e.Value))
	// 			log.Println("above is the value")
	// 			var res SendMail
	// 			err := json.Unmarshal(e.Value, &res)
	// 			if err != nil {
	// 				log.Println(err)
	// 			}
	// 			log.Println("about to assert...")
	// 			fmt.Println(res)
	// 			if err = Mailer.SendMessage(res.Email, res.Message); err != nil {
	// 				log.Println(err)
	// 			}
	// 		}()
	// 	case kafka.Error:
	// 		fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
	// 	}
	// }
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Metadata.AllowAutoTopicCreation = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	consumer, err := sarama.NewConsumer([]string{"host.docker.internal:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("Emailsender", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case e := <-partitionConsumer.Messages():
			go func() {
				log.Println(string(e.Value))
				log.Println("above is the value")
				var res SendMail
				err := json.Unmarshal(e.Value, &res)
				if err != nil {
					log.Println(err)
				}
				log.Println("about to assert...")
				fmt.Println(res)
				if err = Mailer.SendMessage(res.Email, res.Message); err != nil {
					log.Println(err)
				}
			}()

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}
