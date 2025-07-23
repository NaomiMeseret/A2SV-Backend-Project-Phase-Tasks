package usecases

import (
	domain "task_manager/Domain"
	"testing"
)

type MockUserRepo struct {
	CreateUserFunc        func(*domain.User) error
	GetUserByUsernameFunc func(string) (*domain.User, error)
	GetUserByIDFunc       func(string) (*domain.User, error)
	CountUsersFunc        func() (int, error)
	GetAllUsersFunc       func() ([]*domain.User, error)
	PromoteUserFunc       func(string) error
}

func (m *MockUserRepo) CreateUser(u *domain.User) error {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(u)
	}
	return nil
}
func (m *MockUserRepo) GetUserByUsername(username string) (*domain.User, error) {
	if m.GetUserByUsernameFunc != nil {
		return m.GetUserByUsernameFunc(username)
	}
	return nil, nil
}
func (m *MockUserRepo) GetUserByID(id string) (*domain.User, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(id)
	}
	return nil, nil
}
func (m *MockUserRepo) CountUsers() (int, error) {
	if m.CountUsersFunc != nil {
		return m.CountUsersFunc()
	}
	return 0, nil
}
func (m *MockUserRepo) GetAllUsers() ([]*domain.User, error) {
	if m.GetAllUsersFunc != nil {
		return m.GetAllUsersFunc()
	}
	return nil, nil
}
func (m *MockUserRepo) PromoteUser(id string) error {
	if m.PromoteUserFunc != nil {
		return m.PromoteUserFunc(id)
	}
	return nil
}

func TestRegisterUser(t *testing.T) {
	repo := &MockUserRepo{
		CreateUserFunc: func(u *domain.User) error { return nil },
		CountUsersFunc: func() (int, error) { return 0, nil },
	}
	usecase := NewUserUsecase(repo)
	user := &domain.User{UserName: "Naomi", Password: "pass2196"}
	err := usecase.RegisterUser(user)
	if err != nil {
		t.Error("should not get error")
	}
}

func TestRegisterUser_EmptyUsername(t *testing.T) {
	repo := &MockUserRepo{}
	usecase := NewUserUsecase(repo)
	user := &domain.User{UserName: "", Password: "pass2196"}
	err := usecase.RegisterUser(user)
	if err == nil {
		t.Error("should get error for empty username")
	}
}

func TestLoginUser(t *testing.T) {
	repo := &MockUserRepo{
		GetUserByUsernameFunc: func(username string) (*domain.User, error) {
			return &domain.User{UserName: username, Password: "pass2196"}, nil
		},
	}
	usecase := NewUserUsecase(repo)
	_, err := usecase.LoginUser("Naomi", "pass2196")
	if err != nil {
		t.Error("should not get error")
	}
}

func TestGetUserByID(t *testing.T) {
	repo := &MockUserRepo{
		GetUserByIDFunc: func(id string) (*domain.User, error) {
			return &domain.User{ID: id, UserName: "Naomi"}, nil
		},
	}
	usecase := NewUserUsecase(repo)
	user, err := usecase.GetUserByID("1")
	if err != nil || user.ID != "1" {
		t.Error("should get user with ID 1")
	}
}

func TestPromoteUser(t *testing.T) {
	repo := &MockUserRepo{
		PromoteUserFunc: func(id string) error { return nil },
	}
	usecase := NewUserUsecase(repo)
	err := usecase.PromoteUser("1")
	if err != nil {
		t.Error("should not get error")
	}
}

func TestGetAllUsers(t *testing.T) {
	repo := &MockUserRepo{
		GetAllUsersFunc: func() ([]*domain.User, error) {
			return []*domain.User{{ID: "1", UserName: "Naomi"}}, nil
		},
	}
	usecase := NewUserUsecase(repo)
	users, err := usecase.GetAllUsers()
	if err != nil || len(users) != 1 {
		t.Error("should get 1 user")
	}
} 