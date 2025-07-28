package infrastructure

import(
	"golang.org/x/crypto/bcrypt"
)
// PasswordService provides methods for hashing and checking passwords. 
type PasswordService struct{}

// NewPasswordService creates a new PasswordService instance.
func NewPasswordService()*PasswordService{
	return &PasswordService{}
}


func (ps *PasswordService) HashPassword(password string) (string , error){
	hash , err:=bcrypt.GenerateFromPassword([]byte(password) , bcrypt.DefaultCost)
	if err != nil{
		return "" , err
	}
	return string(hash) , nil
}

// CheckPasswordHash compares a plain password with a hashed password.
func (ps *PasswordService) CheckPasswordHash(password , hash string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}