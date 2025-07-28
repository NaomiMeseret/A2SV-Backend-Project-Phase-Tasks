package usecases

import (
	"errors"
	domain "task_manager/Domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// MockTaskRepo is a fake repository for testing TaskUsecase

type MockTaskRepo struct {
    tasks map[string]*domain.Task
}

func (m *MockTaskRepo) CreateTask(t *domain.Task) error {
    if t.Title == "" {
        return errors.New("task title cannot be empty")
    }
    m.tasks[t.ID] = t
    return nil
}
func (m *MockTaskRepo) GetTaskByID(id string) (*domain.Task, error) {
    if task, ok := m.tasks[id]; ok {
        return task, nil
    }
    return nil, errors.New("task not found")
}
func (m *MockTaskRepo) GetTasksByUserID(userID string) ([]*domain.Task, error) {
    var result []*domain.Task
    for _, t := range m.tasks {
        if t.UserID == userID {
            result = append(result, t)
        }
    }
    return result, nil
}
func (m *MockTaskRepo) UpdateTask(t *domain.Task) error {
    if t.ID == "" {
        return errors.New("task ID cannot be empty")
    }
    m.tasks[t.ID] = t
    return nil
}
func (m *MockTaskRepo) DeleteTask(id string) error {
    if _, ok := m.tasks[id]; ok {
        delete(m.tasks, id)
        return nil
    }
    return errors.New("task not found")
}

// TaskUsecaseTestSuite groups all related tests for TaskUsecase
type TaskUsecaseTestSuite struct {
    suite.Suite
    repo    *MockTaskRepo
    usecase domain.ITaskUsecase
}

// SetupTest runs before each test, giving a clean state
func (suite *TaskUsecaseTestSuite) SetupTest() {
    suite.repo = &MockTaskRepo{tasks: make(map[string]*domain.Task)}
    suite.usecase = NewTaskUsecase(suite.repo)
}

// Test Creating a  vaild task
func (suite *TaskUsecaseTestSuite) TestCreateTask_Success() {
    task := &domain.Task{ID: "1", Title: "Test Task", UserID: "001"}
    err := suite.usecase.CreateTask(task)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), "Test Task", suite.repo.tasks["1"].Title)
}

//Test creating a task with empy title
func (suite *TaskUsecaseTestSuite) TestCreateTask_EmptyTitle() {
    task := &domain.Task{ID: "2", Title: "", UserID: "001"}
    err := suite.usecase.CreateTask(task)
    assert.Error(suite.T(), err)
    assert.Equal(suite.T(), "task title cannot be empty", err.Error())
}

// test getting a task by ID
func (suite *TaskUsecaseTestSuite) TestGetTaskByID_Success() {
    task := &domain.Task{ID: "3", Title: "Read", UserID: "002"}
    suite.repo.tasks["3"] = task
    result, err := suite.usecase.GetTaskByID("3")
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), "Read", result.Title)
}

//Test updating a task
func (suite *TaskUsecaseTestSuite) TestUpdateTask_Success() {
    task := &domain.Task{ID: "4", Title: "Old", UserID: "003"}
    suite.repo.tasks["4"] = task
    updated := &domain.Task{ID: "4", Title: "New", UserID: "user3"}
    err := suite.usecase.UpdateTask(updated)
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), "New", suite.repo.tasks["4"].Title)
}

//Test deleting a task
func (suite *TaskUsecaseTestSuite) TestDeleteTask_Success() {
    task := &domain.Task{ID: "5", Title: "Delete", UserID: "004"}
    suite.repo.tasks["5"] = task
    err := suite.usecase.DeleteTask("5")
    assert.NoError(suite.T(), err)
    _, exists := suite.repo.tasks["5"]
    assert.False(suite.T(), exists)
}

func (suite *TaskUsecaseTestSuite) TestGetTaskByID_NotFound() {
    _, err := suite.usecase.GetTaskByID("999")
    assert.Error(suite.T(), err)
    assert.Equal(suite.T(), "task not found", err.Error())
}

//Run the test suite
func TestTaskUsecaseTestSuite(t *testing.T) {
    suite.Run(t, new(TaskUsecaseTestSuite))
} 

