package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("CR0WD4VND1NG_S3cR3t_K3y")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
