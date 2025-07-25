package domain

type Task struct{
	ID   string 
	Title   string 
	Description   string 
	Date   string  //due-date eg:2025-03-04
	Status   string 
	//UserID is used just to know which admin created which task 
	UserID   string 

}

type User struct{
	ID   string 
	Email  string 
	Password   string 
	Role   string 
}

type TaskRepository interface{
	CreateTask(task *Task)error
	//Gets a single task by its unique ID
	GetTaskByID(id string)(*Task , error)
	//Get all tasks that belong to a specific user(by userID)
	GetTasksByUserID (userID string) ([]*Task , error)
	UpdateTask(task *Task) error
	DeleteTask (id string)error
}

type UserRepository interface{
	CreateUser(user *User) error
	UserExists (email string) (bool , error)
	GetUserByEmail(email string)(*User , error)
	GetUserByID(id string)(*User , error)
	CountUsers()(int , error)//Return the total number of users
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
	GetUserByEmail(email string)(*User , error)
	GetUserByID(id string) (*User , error)
	GetAllUsers() ([]*User, error)
    PromoteUser(id string) error

}
