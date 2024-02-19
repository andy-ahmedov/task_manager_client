package commands

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/andy-ahmedov/task_manager_client/internal/domain"
	grpc_client "github.com/andy-ahmedov/task_manager_client/internal/transport/grpc"
	"github.com/andy-ahmedov/task_manager_client/internal/transport/rabbitmq"
	"github.com/spf13/cobra"
)

func UpdateHandle(client *grpc_client.Client, cmd *cobra.Command, args []string, broker *rabbitmq.Broker) {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")
	status, _ := cmd.Flags().GetString("status")

	id := GetID(args[0])
	upd := initUpdateTaskInput(name, description, status)
	err := client.Update(context.Background(), id, upd)

	item := NewItem("update", err)

	broker.SendToQueue(item)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The task has been successfully updated!")
}

func initUpdateTaskInput(name, description, status string) *domain.UpdateTaskInput {
	upd := &domain.UpdateTaskInput{
		Name:        nil,
		Description: nil,
		Status:      nil,
	}

	if name != "" {
		upd.Name = &name
	}
	if description != "" {
		upd.Description = &description
	}
	if status != "" {
		upd.Status = &status
	}

	return upd
}

func GetID(args string) int64 {
	tmp, err := strconv.Atoi(args)
	if err != nil {
		log.Fatal(err)
	}

	id := int64(tmp)

	return id
}

func NewItem(action string, err error) domain.LogItem {
	item := domain.LogItem{Action: action}

	if err == nil {
		item.Status = "OK!"
	} else {
		item.Status = fmt.Sprintf("Error: %v", err)
	}

	return item
}
