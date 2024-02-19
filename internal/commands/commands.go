package commands

import (
	"context"
	"fmt"
	"log"

	grpc_client "github.com/andy-ahmedov/task_manager_client/internal/transport/grpc"
	"github.com/andy-ahmedov/task_manager_client/internal/transport/rabbitmq"
	"github.com/spf13/cobra"
)

func CommandUpdate(client *grpc_client.Client, cmd *cobra.Command, broker *rabbitmq.Broker) *cobra.Command {
	var update = &cobra.Command{
		Use:   "update id --name [name] --description [description] --status [status]",
		Short: "Update an existing task",
		Run: func(cmd *cobra.Command, args []string) {
			UpdateHandle(client, cmd, args, broker)
		},
	}

	update.Flags().String("name", "", "Name of the task")
	update.Flags().String("description", "", "Description of the task")
	update.Flags().String("status", "", "Status of the task")

	return update
}

func CommandCreate(client *grpc_client.Client, cmd *cobra.Command, broker *rabbitmq.Broker) *cobra.Command {
	var cmdCreate = &cobra.Command{
		Use:   "create [name] [description] [status]",
		Short: "Create a new task",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			err := client.Create(context.Background(), args[0], args[1], args[2])
			if err != nil {
				log.Fatal(err)
			}

			item := NewItem("create", err)

			broker.SendToQueue(item)

			fmt.Printf("Creating task with name %s, description %s and status %s\n", args[0], args[1], args[2])
		},
	}

	return cmdCreate
}

func CommandDelete(client *grpc_client.Client, cmd *cobra.Command, broker *rabbitmq.Broker) *cobra.Command {
	var cmdDelete = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client.Delete(context.Background(), GetID(args[0]))

			item := NewItem("delete", nil)

			broker.SendToQueue(item)

			fmt.Printf("Deleting task with id %s\n", args[0])
		},
	}

	return cmdDelete
}

func CommandGet(client *grpc_client.Client, cmd *cobra.Command, broker *rabbitmq.Broker) *cobra.Command {
	var cmdGet = &cobra.Command{
		Use:   "get [id]",
		Short: "Get a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			task, err := client.Get(context.Background(), GetID(args[0]))
			if err != nil {
				log.Fatal(err)
			}

			item := NewItem("get", err)

			broker.SendToQueue(item)

			fmt.Printf("Getting task with id %s:\n%v\n", args[0], task)
		},
	}

	return cmdGet
}

func CommandGetAll(client *grpc_client.Client, cmd *cobra.Command, broker *rabbitmq.Broker) *cobra.Command {
	var cmdGetAll = &cobra.Command{
		Use:   "getall",
		Short: "Get all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := client.GetAll(context.TODO())
			if err != nil {
				log.Fatal(err)
			}

			item := NewItem("getall", err)

			broker.SendToQueue(item)
			fmt.Println("Getting all tasks:")
			for _, task := range tasks {
				fmt.Println(task)
			}
		},
	}

	return cmdGetAll
}
