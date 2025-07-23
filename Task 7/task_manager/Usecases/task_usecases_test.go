package usecases

import (
	domain "task_manager/Domain"
	"testing"
)

type MockTaskRepo struct {
	CreateTaskFunc      func(*domain.Task) error
	GetTaskByIDFunc     func(string) (*domain.Task, error)
	GetTasksByUserIDFunc func(string) ([]*domain.Task, error)
	UpdateTaskFunc      func(*domain.Task) error
	DeleteTaskFunc      func(string) error
}

func (m *MockTaskRepo) CreateTask(t *domain.Task) error {
	if m.CreateTaskFunc != nil {
		return m.CreateTaskFunc(t)
	}
	return nil
}
func (m *MockTaskRepo) GetTaskByID(id string) (*domain.Task, error) {
	if m.GetTaskByIDFunc != nil {
		return m.GetTaskByIDFunc(id)
	}
	return nil, nil
}
func (m *MockTaskRepo) GetTasksByUserID(userID string) ([]*domain.Task, error) {
	if m.GetTasksByUserIDFunc != nil {
		return m.GetTasksByUserIDFunc(userID)
	}
	return nil, nil
}
func (m *MockTaskRepo) UpdateTask(t *domain.Task) error {
	if m.UpdateTaskFunc != nil {
		return m.UpdateTaskFunc(t)
	}
	return nil
}
func (m *MockTaskRepo) DeleteTask(id string) error {
	if m.DeleteTaskFunc != nil {
		return m.DeleteTaskFunc(id)
	}
	return nil
}

func TestCreateTask(t *testing.T) {
	repo := &MockTaskRepo{
		CreateTaskFunc: func(t *domain.Task) error { return nil },
	}
	usecase := NewTaskUsecase(repo)
	task := &domain.Task{Title: "Test"}
	err := usecase.CreateTask(task)
	if err != nil {
		t.Error("should not get error")
	}
}

func TestCreateTask_EmptyTitle(t *testing.T) {
	repo := &MockTaskRepo{}
	usecase := NewTaskUsecase(repo)
	task := &domain.Task{Title: ""}
	err := usecase.CreateTask(task)
	if err == nil {
		t.Error("should get error for empty title")
	}
}

func TestGetTaskByID(t *testing.T) {
	repo := &MockTaskRepo{
		GetTaskByIDFunc: func(id string) (*domain.Task, error) {
			return &domain.Task{ID: id, Title: "Test"}, nil
		},
	}
	usecase := NewTaskUsecase(repo)
	task, err := usecase.GetTaskByID("1")
	if err != nil || task.ID != "1" {
		t.Error("should get task with ID 1")
	}
}

func TestUpdateTask(t *testing.T) {
	repo := &MockTaskRepo{
		UpdateTaskFunc: func(t *domain.Task) error { return nil },
	}
	usecase := NewTaskUsecase(repo)
	task := &domain.Task{ID: "1", Title: "Update"}
	err := usecase.UpdateTask(task)
	if err != nil {
		t.Error("should not get error")
	}
}

func TestDeleteTask(t *testing.T) {
	repo := &MockTaskRepo{
		DeleteTaskFunc: func(id string) error { return nil },
	}
	usecase := NewTaskUsecase(repo)
	err := usecase.DeleteTask("1")
	if err != nil {
		t.Error("should not get error")
	}
} 