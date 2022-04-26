package database

import (
	"context"
	"database/sql"
	"go-rest-ws/models"
	"log"

	_ "github.com/lib/pq" // Es necessario para que funcione la conexion con postgres
)

/*
	sql: NO se tuvo que descargar ya que es parte de Go

	Pero hay que descargar el paquete de postgresql
	$ go get github.com/lib/pq
*/
type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

// El contexto es para poder hacer un track de la app
func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	// row := repo.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE id = $1", id)
	// user := &models.User{}
	// err := row.Scan(&user.Id, &user.Email, &user.Password)
	// if err != nil {
	// 	return nil, err
	// }
	// return user, nil

	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	defer func() {
		// Cerrar la conexion cuando se termine de usar
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
