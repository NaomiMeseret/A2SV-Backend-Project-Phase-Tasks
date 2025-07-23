package infrastructure

import(
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret_key")

//GenerateJWT creates a jwt token for a given user ID and role
func GenerateJWT (userID , role string)(string , error){
	claims:=jwt.MapClaims{
		"user_id":userID,
		"role":role,
		"exp":time.Now().Add(time.Hour *24).Unix(),
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT (tokenString string )(jwt.MapClaims , error){
	token , err:=jwt.Parse(tokenString ,func(token *jwt.Token) (interface{} ,error){
		if _ , ok :=token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil , jwt.ErrSignatureInvalid
		}
		return jwtSecret , nil
	})
	if err !=nil{
		return nil , err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok &&token.Valid{
		return claims , nil
	}
	return nil , jwt.ErrSignatureInvalid
}