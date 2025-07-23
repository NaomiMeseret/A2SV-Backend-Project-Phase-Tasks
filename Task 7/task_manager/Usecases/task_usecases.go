package usecases

import(
	"errors"
	domain "task_manager/Domain"
)
//taskUsecase implements the TaskUsecase interface from the domain layer
type taskUsecase struct{
	repo domain.TaskRepository
}

//NewTaskusecase creates a new TaskUsecase with the given repository
func NewTaskUsecase (repo domain.TaskRepository)domain.TaskUsecase{
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase)CreateTask (task *domain.Task) error{
	if task.Title == ""{
		return errors.New("task title cannot be empty")
	}
	return u.repo.CreateTask(task)
}
func (u *taskUsecase)GetTaskByID(id string) (*domain.Task , error){
	return u.repo.GetTaskByID(id)
}
func (u *taskUsecase)GetTasksByUserID(userID string)([]*domain.Task , error){
	return u.repo.GetTasksByUserID(userID)
}
func (u *taskUsecase)UpdateTask(task *domain.Task)error{
	if task.ID == ""{
		return errors.New("task ID cannot be empty")
	}
	return u.repo.UpdateTask(task)
}
func (u *taskUsecase)DeleteTask(id string)error{
	if id  == ""{
		return errors.New("task ID cannot be empty")
	}
	return u.repo.DeleteTask(id)
}