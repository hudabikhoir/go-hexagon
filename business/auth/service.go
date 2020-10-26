package auth

import (
	"arkan-jaya/business/auth/spec"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	validator "github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

//Service outgoing port for auth
type Service interface {
	LoginHandler(upsertAuthSpec spec.UpsertAuthSpec) (string, error)
}

//=============== The implementation of those interface put below =======================

type service struct {
	repository Repository
	validate   *validator.Validate
}

//NewService Construct auth service object
func NewService(repository Repository) Service {
	return &service{
		repository,
		validator.New(),
	}
}

//GetUsers Get all auths, return zero array if not match
func (s *service) LoginHandler(upsertAuthSpec spec.UpsertAuthSpec) (string, error) {

	user, err := s.repository.FindUserByUsername(upsertAuthSpec.Username)
	if user == nil || err != nil {
		return "", err
	} else if CheckPasswordHash(upsertAuthSpec.Password, user.Password) {
		// Set custom claims
		claims := &spec.JWTClaimSpec{
			user.ID,
			user.Username,
			user.RoleID,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(viper.GetString("jwtSign")))
		if err != nil {
			return "", err
		}
		return t, nil
	}

	return "", errors.New("Invalid username and password")
	// return auths, err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
