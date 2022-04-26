package repository

import (
	"context"
	"go-rest-ws/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	Close() error
}

/*
	Inversion de dependencia

	Los codigos deben de esstar basardo en abstraccion y no en cosas concretas

	Concreta:
	El codigo se vuelve complicado con cada cambio que se haga en la base de datos
	Handler - GetUserByIdPostgres ... ... ... (Toda la logica del codigo)
	Handler - GetUserByIdMongoDB ... ... ... (Toda la logica del codigo)

	Absatraccion:
	Es mejor tener un funcion de como se va a implementar y no el
	comportamiento interno de la base de datos
*/

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
