package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/mehmetaligok/venom-example-project/src/model"
)

// Repo is the structure that is used to communicate with the db.
type Repo struct {
	db *sql.DB
}

// NewUserRepo returns an instance of the User repo.
func NewUserRepo(dsn string) *Repo {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return &Repo{
		db: db,
	}
}

// InsertUser inserts given user to database
func (repo *Repo) InsertUser(ctx context.Context, user *model.User) error {
	stmt := `insert into "users" ("id", "first_name", "last_name") values($1, $2, $3)`
	_, err := repo.db.ExecContext(ctx, stmt, user.ID, user.FirstName, user.LastName)

	return err
}

// GetUser returns users from database by id
func (repo *Repo) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user := &model.User{}
	err := repo.db.
		QueryRowContext(ctx, "SELECT id,  first_name, last_name FROM users WHERE id=$1;", userID.String()).
		Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	return user, nil
}
