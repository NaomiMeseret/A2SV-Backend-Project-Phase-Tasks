package usecases

import (
	"errors"
	domain "task_manager/Domain"
)

//userUsecase implements the UserUsecase interface form the domain layer.
type userUsecase struct{
	repo domain .UserRepository
}

//NewUserUsecase creates a new UserUsecase with the given repository
func NewUserUsecase(repo domain.UserRepository)domain.UserUsecase{
	return &userUsecase{repo: repo}
}

func (u *userUsecase)RegisterUser(user *domain.User)error{
	if user.UserName == ""{
		return errors.New("username cannot be empty")
	}
	if user.Password == ""{
		return errors.New("password cannot be empty")
	}
	if len(user.Password)<4{
		return errors.New("password must be at least 4 characters long")
	}
	existing  , _ := u.repo.GetUserByUsername(user.UserName)
	if existing!=nil{
		return errors.New("username already taken")
	}
	if counter, err:=u.repo.CountUsers();err==nil &&counter ==0{
		user.Role = "admin"
	}else{
		user.Role = "user"
	}
	return u.repo.CreateUser(user)
}

func (u *userUsecase) LoginUser(username, password string)(*domain.User , error){
	if username  == "" || password == ""{
		return nil , errors.New("username and password cannot be empty")

	}
	user , err := u.repo.GetUserByUsername(username)
	if err !=nil{
		return nil , err
	}
	return user , nil
}

func (u *userUsecase)GetUserByID(id string)(*domain.User , error){
	return u.repo.GetUserByID(id)
}

func (u *userUsecase)GetAllUsers()([]*domain.User , error){
	users ,err:=u.repo.GetAllUsers()
	if err!=nil{
		return nil, err
	}
	for _ , user:=range users{
		user.Password =""
	}
	return users , nil
}

func (u *userUsecase) PromoteUser(id string) error{
	return u.repo.PromoteUser(id)

}



