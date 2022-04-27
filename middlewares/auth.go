package middlewares

import (
	"fmt"
	"go-rest-ws/models"
	"go-rest-ws/server"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	// AuthMiddleware is a middleware that check if the user is authenticated
	NO_AUTH_MIDDLEWARE = []string{
		"login",
		"signup",
	}
)

func shouldCheckToken(route string) bool {
	for _, routeName := range NO_AUTH_MIDDLEWARE {
		// if route == routeName {
		// 	return false
		// }

		if strings.Contains(route, routeName) {
			return false
		}
	}
	return true
}

/*
	Recibimos una handler y devolvemos un handler
*/
func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if !shouldCheckToken(r.URL.Path) {
				// If the route is not protected, just continue with the next handler
				next.ServeHTTP(w, r)
				return
			}

			// Get the token from the request
			// token, err := r.Cookie("token")
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

			/*
				tokenString: Token del user
				models.AppClaims{}: El como queremos tener la data al descomprimir el token
				func: que verifica si la llave enviada es correcta
			*/
			data, err := jwt.ParseWithClaims(tokenString, models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})

			if err != nil {
				log.Println(err)
				// If there is an error, do not continue.
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			fmt.Printf("Data del token %+v\n", data)

			// If there is no error, continue.
			next.ServeHTTP(w, r)
		})
	}
}
