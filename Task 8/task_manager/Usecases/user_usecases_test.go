package usecases

import (
	"errors"
	domain "task_manager/Domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// MockUserRepo is a fake repository for testing IUserUsecase
type MockUserRepo struct {
	users map[string]*domain.User
	emailIndex map[string]string // email -> id
}

func (m *MockUserRepo) CreateUser(u *domain.User) error {
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	if _, exists := m.emailIndex[u.Email]; exists {
		return errors.New("email already taken")
	}
	m.users[u.ID] = u
	m.emailIndex[u.Email] = u.ID
	return nil
}
func (m *MockUserRepo) GetUserByEmail(email string) (*domain.User, error) {
	if id, ok := m.emailIndex[email]; ok {
		return m.users[id], nil
	}
	return nil, errors.New("user not found")
}
func (m *MockUserRepo) GetUserByID(id string) (*domain.User, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}
func (m *MockUserRepo) CountUsers() (int, error) {
	return len(m.users), nil
}
func (m *MockUserRepo) GetAllUsers() ([]*domain.User, error) {
	var result []*domain.User
	for _, u := range m.users {
		result = append(result, u)
	}
	return result, nil
}
func (m *MockUserRepo) PromoteUser(id string) error {
	if user, ok := m.users[id]; ok {
		user.Role = "admin"
		return nil
	}
	return errors.New("user not found")
}

// UserExists checks if a user exists by email in the mock repo.
func (m *MockUserRepo) UserExists(email string) (bool, error) {
    _, exists := m.emailIndex[email]
    return exists, nil
}

// MockPasswordService is a simple mock for IPasswordService

type MockPasswordService struct{}
func (m *MockPasswordService) HashPassword(password string) (string, error) { 
	return password, nil 
   }
func (m *MockPasswordService) CheckPasswordHash(password, hash string) bool {
	 return password == hash 
	}

// UserUsecaseTestSuite groups all related tests for IUserUsecase

type UserUsecaseTestSuite struct {
	suite.Suite
	repo    *MockUserRepo
	usecase domain.IUserUsecase
}

// SetupTest runs before each test, giving a clean state
func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.repo = &MockUserRepo{
		users: make(map[string]*domain.User),
		emailIndex: make(map[string]string),
	}
	passwordService := &MockPasswordService{}
	suite.usecase = NewUserUsecase(suite.repo, passwordService)
}

// Test registering a valid user
func (suite *UserUsecaseTestSuite) TestRegisterUser_Success() {
	user := &domain.User{ID: "1", Email: "naomi@email.com", Password: "pass123"}
	err := suite.usecase.RegisterUser(user)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "naomi@email.com", suite.repo.users["1"].Email)
}

// Test registering a user with empty email
func (suite *UserUsecaseTestSuite) TestRegisterUser_EmptyEmail() {
	user := &domain.User{ID: "2", Email: "", Password: "pass123"}
	err := suite.usecase.RegisterUser(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "email cannot be empty", err.Error())
}

// Test registering a user with empty password
func (suite *UserUsecaseTestSuite) TestRegisterUser_EmptyPassword() {
	user := &domain.User{ID: "3", Email: "abel@email.com", Password: ""}
	err := suite.usecase.RegisterUser(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "password cannot be empty", err.Error())
}

// Test registering a user with duplicate email
func (suite *UserUsecaseTestSuite) TestRegisterUser_DuplicateEmail() {
	user1 := &domain.User{ID: "4", Email: "abel@email.com", Password: "pass123"}
	_ = suite.usecase.RegisterUser(user1)
	user2 := &domain.User{ID: "5", Email: "abel@email.com", Password: "pass456"}
	err := suite.usecase.RegisterUser(user2)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "email already taken", err.Error())
}
// Test registering a user with password too short
func (suite *UserUsecaseTestSuite) TestRegisterUser_PasswordTooShort() {
    user := &domain.User{Email: "naomi@email.com", Password: "123"} 
    err := suite.usecase.RegisterUser(user)
    suite.Assert().Error(err)
    suite.Assert().Contains(err.Error(), "password must be at least 4 characters long")
}

// Test login with correct credentials
func (suite *UserUsecaseTestSuite) TestLoginUser_Success() {
	user := &domain.User{ID: "6", Email: "naomi@email.com", Password: "pass123"}
	_ = suite.usecase.RegisterUser(user)
	result, err := suite.usecase.LoginUser("naomi@email.com", "pass123")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "naomi@email.com", result.Email)
}

// Test login with wrong email
func (suite *UserUsecaseTestSuite) TestLoginUser_WrongEmail() {
	_, err := suite.usecase.LoginUser("notfound@email.com", "pass123")
	assert.Error(suite.T(), err)
}

// Test login with wrong password
func (suite *UserUsecaseTestSuite) TestLoginUser_WrongPassword() {
	user := &domain.User{ID: "7", Email: "abel@email.com", Password: "pass123"}
	_ = suite.usecase.RegisterUser(user)
	_, err := suite.usecase.LoginUser("abel@email.com", "wrongpass")
	assert.Error(suite.T(), err)
}

// Test promote user
func (suite *UserUsecaseTestSuite) TestPromoteUser_Success() {
	user := &domain.User{ID: "8", Email: "abel@email.com", Password: "pass123", Role: "user"}
	suite.repo.users["8"] = user
	err := suite.usecase.PromoteUser("8")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "admin", suite.repo.users["8"].Role)
}

// Run the test suite
func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
} 