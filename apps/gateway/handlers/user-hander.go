package handlers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tpmanc/gateway/db"
	"github.com/tpmanc/gateway/helpers"
	"github.com/tpmanc/gateway/models"
	"github.com/tpmanc/gateway/repositories"
	"github.com/tpmanc/gateway/requests"
	"github.com/tpmanc/gateway/services"
	"net/http"
)

var mySigningKey = []byte("secret")

func getService() services.UserServiceInterface {
	repository := repositories.GetUserRepository(db.Get())
	return services.GetUserService(repository)
}

type LoginResponse struct {
	Id uint
	Username string
	Token string
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := requests.ParseLoginRequest(r)

	service := getService()
	user := service.Login(req)
	if user == nil {
		errInfo := map[string] string {
			"message": "Invalid login credentials. Please try again",
		}
		helpers.ResponseError(w, errInfo)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	token.Claims = claims
	// token.Claims["name"] = "Ado Kukic"
	// token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, _ := token.SignedString(mySigningKey)

	response := &LoginResponse{
		Id: user.ID,
		Username: user.Username,
		Token: tokenString,
	}
	helpers.RespondWithJSON(w, http.StatusOK, response)
}

type SignupResponse struct {
	User *models.User
}
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	req := requests.ParseSignupRequest(r)

	service := getService()
	user := service.Signup(req)

	response := &SignupResponse{
		User: user,
	}
	helpers.RespondWithJSON(w, http.StatusOK, response)
}
