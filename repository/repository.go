package repository

import (
	"context"
	"go-rest-ws/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id string, userId string) error
	ListPosts(ctx context.Context, page uint64) ([]*models.Post, error)
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

var implementation Repository

func SetRepository(repository Repository) {
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

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Post) error {
	return implementation.UpdatePost(ctx, post)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implementation.DeletePost(ctx, id, userId)
}

func ListPosts(ctx context.Context, page uint64) ([]*models.Post, error) {
	return implementation.ListPosts(ctx, page)
}
