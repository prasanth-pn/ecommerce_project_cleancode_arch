package usecase

import (
	"fmt"
	"log"
	"os"
	"time"

	"clean/pkg/usecase/interfaces"

	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	User_Id    uint
	First_Name string
	Email      string
	jwt.StandardClaims
}
type jwtService struct {
	SecretKey string
}

func NewJWTService() interfaces.JWTService {
	return &jwtService{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}

func (j *jwtService) GenerateToken(User_Id uint, First_Name, Email string) string {
	fmt.Println("first ghost",User_Id,First_Name,Email,"firstghost\n ")
	//j.SecretKey = "kofmjlksdjgmklsdjml"
	claims := &SignedDetails{
		User_Id:    User_Id,
		First_Name: First_Name,
		Email:      Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	fmt.Println("sdfgfdgdsgfssdf", j.SecretKey, "deghfhfhdfhgfghhihihihihi")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("\n\n\n tokensigned token", signedToken, j.SecretKey)
	return signedToken

}

func (j *jwtService) VerifyToken(signedToken string) (bool, *SignedDetails) {
	claims := &SignedDetails{}
	token, _ := j.GetTokenFromString(signedToken, claims)
	if token.Valid {
		if t := claims.Valid(); t == nil {
			return true, claims
		}
	}
	return false, claims

}

func (j *jwtService) GetTokenFromString(signedToken string, claims *SignedDetails) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:#{token.Header['alg']}")
		}
		return []byte(j.SecretKey), nil
	})

}
