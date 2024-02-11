package grpc

import (
	"context"

	"github.com/andy-ahmedov/task_manager_client/internal/domain"
	"github.com/andy-ahmedov/task_manager_grpc_api/api"
)

func (c *Client) Create(ctx context.Context, name, des, status string) error {

	req := api.CreateRequest{
		Name:        name,
		Description: des,
		Status:      status,
	}
	_, err := c.TaskServiceClient.Create(ctx, &req)

	return err
}

func (c *Client) Get(ctx context.Context, id int64) (domain.Task, error) {
	req := api.GetRequest{
		ID: id,
	}

	task, err := c.TaskServiceClient.Get(ctx, &req)
	if err != nil {
		return domain.Task{}, err
	}

	return ConverteTime(task.Task), nil
}

func (c *Client) GetAll(ctx context.Context) ([]domain.Task, error) {
	req := api.GetAllRequest{}

	task, err := c.TaskServiceClient.GetAll(ctx, &req)
	if err != nil {
		return nil, err
	}

	tasks := ConvertTasks(task.Tasks)

	return tasks, nil
}

func (c *Client) Delete(ctx context.Context, id int64) error {
	req := api.DeleteRequest{
		ID: id,
	}

	_, err := c.TaskServiceClient.Delete(ctx, &req)

	return err
}

func (c *Client) Update(ctx context.Context, id int64, upd *domain.UpdateTaskInput) error {
	apiTask := ConvertToApiShortTask(upd)

	req := api.UpdateRequest{
		ID:   id,
		Task: apiTask,
	}

	_, err := c.TaskServiceClient.Update(ctx, &req)

	return err
}
