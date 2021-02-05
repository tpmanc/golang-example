package middleware

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tpmanc/gateway/helpers"
	"net/http"
	"strings"
)

type AuthToken struct {
	Id uint
	Username string
	jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// notAuth := []string{"/login"} //Список эндпоинтов, для которых не требуется авторизация
		// requestPath := r.URL.Path //текущий путь запроса

		// //проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
		// for _, value := range notAuth {
		// 	if value == requestPath {
		// 		next.ServeHTTP(w, r)
		// 		return
		// 	}
		// }

		// response := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" { // Токен отсутствует
			helpers.ResponseForbidden(w, "Missing auth token")
			return
		}

		splitted := strings.Split(tokenHeader, " ") // Bearer {token}
		if len(splitted) != 2 {
			helpers.ResponseForbidden(w, "Invalid/Malformed auth token")
			return
		}
		tokenPart := splitted[1]
		tk := &AuthToken{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil { // Неправильный токен
			helpers.ResponseForbidden(w, "Malformed authentication token")
			return
		}

		if !token.Valid { //токен недействителен, возможно, не подписан на этом сервере
			helpers.ResponseForbidden(w, "Invalid token")
			return
		}

		// fmt.Println(tk.Id)
		// fmt.Println(tk.Username)

		// //Всё прошло хорошо, продолжаем выполнение запроса
		// fmt.Sprintf("User %", tk.Username) //Полезно для мониторинга
		// ctx := context.WithValue(r.Context(), "user", tk.UserId)
		// r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //передать управление следующему обработчику!
	});
}