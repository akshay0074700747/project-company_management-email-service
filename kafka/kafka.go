package kafka

import (
	"fmt"
	"log"
	"os"

	"github.com/akshay0074700747/email-service/email"
	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	Mailer *email.SMTPConfig
)

func Getmailer(mail *email.SMTPConfig) {
	Mailer = mail
}

func StartServing() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "consumers",
		"auto.offset.reset": "smallest"})
	if err != nil {
		log.Fatal(err)
	}
	topic := "Emailsender"
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		fmt.Println(err)
	}
	for {
		ev := consumer.Poll(10)
		switch e := ev.(type) {
		case *kafka.Message:
			go func() {
				log.Println(string(e.Value))
				log.Println("above is the value")
				deserialized, err := DeSerialize(e.Value)
				if err != nil {
					log.Println(err)
				}
				log.Println("about to assert...")
				res := deserialized.(*pb.Email)
				fmt.Println(res)
				if err = Mailer.SendMessage(res.GetReciever(), res.GetMessage()); err != nil {
					log.Println(err)
				}
			}()
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
		}
	}
}

func DeSerialize(raw []byte) (m protoreflect.ProtoMessage, err error) {

	var res pb.Email
	err = proto.Unmarshal(raw, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
