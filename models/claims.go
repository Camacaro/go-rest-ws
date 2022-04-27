package models

import "github.com/golang-jwt/jwt"

/*
	Todas las propiedades de jwt.StandardClaims seran
	pasadas a la estructura de AppClaims
*/
type AppClaims struct {
	UserId             string `json:"userId"`
	jwt.StandardClaims        // Composicion sobre la herencia
}
