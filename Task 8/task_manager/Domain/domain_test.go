package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTaskStruct checks that the Task struct can be created and fields are set correctly
func TestTaskStruct(t *testing.T) {
	task := Task{
		ID:          "1",
		Title:       "Test Task",
		Description: "This is a test task",
		Date:        "2025-03-04",
		Status:      "pending",
		UserID:      "001",
	}
	assert.Equal(t, "1", task.ID)
	assert.Equal(t, "Test Task", task.Title)
	assert.Equal(t, "This is a test task", task.Description)
	assert.Equal(t, "2025-03-04", task.Date)
	assert.Equal(t, "pending", task.Status)
	assert.Equal(t, "001", task.UserID)
}

// TestUserStruct checks that the User struct can be created and fields are set correctly
func TestUserStruct(t *testing.T) {
	user := User{
		ID:       "1",
		Email:    "naomi@email.com",
		Password: "pass123",
		Role:     "admin",
	}
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "naomi@email.com", user.Email)
	assert.Equal(t, "pass123", user.Password)
	assert.Equal(t, "admin", user.Role)
} 
