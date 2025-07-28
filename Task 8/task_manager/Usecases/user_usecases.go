package usecases

import (
	"errors"
	domain "task_manager/Domain"
	"unicode"
)

//userUsecase implements the UserUsecase interface form the domain layer.
type userUsecase struct{
	repo domain .IUserRepository
	passwordService domain.IPasswordService
}

//NewUserUsecase creates a new UserUsecase with the given repository
func NewUserUsecase(repo domain.IUserRepository ,passwordService domain.IPasswordService) domain.IUserUsecase{
	return &userUsecase{repo: repo ,
	passwordService: passwordService,}
}

func isAllLowerCase(s string)bool{
	for _,r:=range s{
		if unicode.IsUpper(r){
			return false
		}
	}
	return true
}
func (u *userUsecase)RegisterUser(user *domain.User)error{
	if !isAllLowerCase(user.Email){
		return errors.New("email must be all lower case")
	}
	if user.Email == ""{
		return errors.New("email cannot be empty")
	}
	if user.Password == ""{
		return errors.New("password cannot be empty")
	}
	if len(user.Password)<4{
		return errors.New("password must be at least 4 characters long")
	}
	exists , err := u.repo.UserExists(user.Email)
	if err!=nil{
		return err
	}
	if exists{
		return errors.New("email already taken")
	}
	if counter, err:=u.repo.CountUsers();err==nil &&counter ==0{
		user.Role = "admin"
	}else{
		user.Role = "user"
	}
	// Hash the password before storing
	hashed , err :=u .passwordService.HashPassword(user.Password)
	if err !=nil{
		return err
	}
	user.Password = hashed
	return u.repo.CreateUser(user)
}
func (u *userUsecase) LoginUser(email, password string)(*domain.User , error){
	if !isAllLowerCase(email){
		return nil , errors.New("email must be all lowercase")
	}
	if email  == "" || password == ""{
		return nil , errors.New("email and password cannot be empty")

	}
	user , err := u.repo.GetUserByEmail(email)
	if err !=nil{
		return nil , err
	}
	// Use passwordService to check password
	if !u.passwordService.CheckPasswordHash(password , user.Password){
		return nil , errors.New("invalid password")
	}
	return user , nil
}

// GetUserByEmail fetches a user by their email address.
func (u *userUsecase)GetUserByEmail(email string) (*domain.User , error){
	return u.repo.GetUserByEmail(email)
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

// PromoteUser sets a user's role to admin by id.
func (u *userUsecase) PromoteUser(id string) error{
	return u.repo.PromoteUser(id)

}



