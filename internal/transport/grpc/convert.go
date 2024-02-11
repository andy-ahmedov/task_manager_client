package grpc

import (
	"github.com/andy-ahmedov/task_manager_client/internal/domain"
	"github.com/andy-ahmedov/task_manager_grpc_api/api"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ConverteTime(task *api.Task) domain.Task {
	newTask := domain.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		Created_at:  task.CreatedAt.AsTime(),
	}

	return newTask
}

func ConvertTasks(apiTasks []*api.Task) []domain.Task {
	tasks := make([]domain.Task, 0)

	for _, apiTask := range apiTasks {
		task := ConverteTime(apiTask)
		tasks = append(tasks, task)
	}

	return tasks
}

func ConvertToApiShortTask(upd *domain.UpdateTaskInput) *api.ShortTask {
	apiTask := &api.ShortTask{
		Name:        nil,
		Description: nil,
		Status:      nil,
	}

	if upd.Name != nil {
		apiTask.Name = wrapperspb.String("")
		apiTask.Name.Value = *upd.Name
	}
	if upd.Description != nil {
		apiTask.Description = wrapperspb.String("")
		apiTask.Description.Value = *upd.Description
	}
	if upd.Status != nil {
		apiTask.Status = wrapperspb.String("")
		apiTask.Status.Value = *upd.Status
	}

	return apiTask
}
