package infrastructure

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret_key")

// It implements the domain.IJWTService interface.
type JWTService struct{}

// NewJWTService creates a new JWTService instance.
func NewJWTService() *JWTService {
	return &JWTService{}
}

func (js *JWTService) GenerateJWT(userID, role string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), 
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func (js *JWTService) ValidateJWT(tokenString string) (map[string]interface{}, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Make sure the signing method is HMAC
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return jwtSecret, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return map[string]interface{}(claims), nil
    }
    return nil, jwt.ErrSignatureInvalid
} 