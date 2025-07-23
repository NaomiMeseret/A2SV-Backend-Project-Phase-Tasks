package domain

type Task struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Date string `json:"date"` //due-date eg:2025-03-04
	Status string `json:"status"`
	UserID string `json:"user_id"`

}

type User struct{
	ID string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type TaskRepository interface{
	CreateTask(task *Task)error
	GetTaskByID(id string)(*Task , error)
	GetTasksByUserID (userID string) ([]*Task , error)
	UpdateTask(task *Task) error
	DeleteTask (id string)error
}

type UserRepository interface{
	CreateUser(user *User) error
	GetUserByUsername (username string)(*User , error)
	GetUserByID(id string)(*User , error)
	CountUsers()(int , error)
	GetAllUsers()([]*User , error)
	PromoteUser(id string)error
}

type TaskUsecase interface{
	CreateTask(task *Task)error
	GetTaskByID(id string)(*Task , error)
	GetTasksByUserID (userID string) ([]*Task , error)
	UpdateTask(task *Task) error
	DeleteTask (id string)error

}
type UserUsecase interface{
	RegisterUser(user *User)error
	LoginUser(username , password string) (*User , error)
	GetUserByID(id string) (*User , error)
	GetAllUsers() ([]*User, error)
    PromoteUser(id string) error

}
