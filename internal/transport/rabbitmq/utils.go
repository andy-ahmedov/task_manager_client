package rabbitmq

import (
	"fmt"

	"github.com/andy-ahmedov/task_manager_client/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

func ConnectToTCP(cfg *config.Broker, logg *logrus.Logger) (*amqp.Connection, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		logg.Error("Failed to connect to RabbitMQ")
		return nil, err
	}
	return conn, err
}

func CreateChannel(conn *amqp.Connection, logg *logrus.Logger) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		logg.Error("Failed to open a channel")
		return nil, err
	}
	return ch, err
}

func DeclareQueue(ch *amqp.Channel, name string, logg *logrus.Logger) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logg.Error("Failed to declare a queue")
		return q, err
	}
	return q, nil
}
