package interfaces

import (
	domain "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"

	_ "github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(ID uint, First_Name string, Email string) string
	VerifyToken(token string) (bool, *domain.SignedDetails)
	//GetTokenFromString(signedToken string,claims *SignedDetails)(*jwt.Token,error)
}
