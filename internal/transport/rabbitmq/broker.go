package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/andy-ahmedov/task_manager_client/internal/config"
	"github.com/andy-ahmedov/task_manager_client/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Broker struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
	Q    *amqp.Queue
	Log  *logrus.Logger
}

func NewBroker(cfg *config.Broker, logg *logrus.Logger) *Broker {
	conn, err := ConnectToTCP(cfg, logg)
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()

	ch, err := CreateChannel(conn, logg)
	if err != nil {
		log.Fatal(err)
	}
	// defer ch.Close()

	q, err := DeclareQueue(ch, "LogItemsQueue", logg)
	if err != nil {
		log.Fatal(err)
	}

	return &Broker{
		Conn: conn,
		Ch:   ch,
		Q:    &q,
		Log:  logg,
	}
}

func (b *Broker) SendToQueue(item domain.LogItem) error {
	body := item
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		b.Log.Error("Failed to encode body into JSON")
		return err
	}

	err = b.Ch.Publish(
		"",
		b.Q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         bodyBytes,
		})
	if err != nil {
		return err
	}

	b.Log.Infof(" [x] Sent %s", body)
	defer b.Conn.Close()
	defer b.Ch.Close()

	return nil
}
