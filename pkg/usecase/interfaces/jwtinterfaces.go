package interfaces

import _ "github.com/golang-jwt/jwt/v4"


type JWTService interface{
	GenerateToken(ID uint,First_Name string,Email string)string
   /// VerifyToken(token string)(bool,*SignedDetails)
	//GetTokenFromString(signedToken string,claims *SignedDetails)(*jwt.Token,error)
}