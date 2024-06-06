package user

import (
	"context"
	"database/sql"

	mysqlc "github.com/trann/e_commerce/internal/model"
	"github.com/trann/e_commerce/internal/types"
)

type Store struct {
	db      *sql.DB
	queries *mysqlc.Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		queries: mysqlc.New(db),
	}
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	dbUser, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	user := convertDBUserToTypeUser(dbUser)
	return &user, nil
}

// func (s *Store) CreateUser(user types.User) error {
// 	ctx := context.Background()
// 	return s.queries.CreateUser(ctx, user.FirstName, user.LastName, user.Email, user.Password)
// }

func (s *Store) GetUserByID(ctx context.Context, id int) (*types.User, error) {
	userDB, err := s.queries.GetUserByID(ctx, uint32(id))
	if err != nil {
		return nil, err
	}
	user := convertDBUserToTypeUser(userDB)
	return &user, nil
}
func (s *Store) CreateUser(ctx context.Context, user types.User) error {

	err := s.queries.CreateUser(ctx, mysqlc.CreateUserParams{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

// convertDBUserToTypeUser converts a db.User to a types.User
func convertDBUserToTypeUser(dbUser mysqlc.User) types.User {
	return types.User{
		ID:        int(dbUser.ID),
		FirstName: dbUser.Firstname,
		LastName:  dbUser.Lastname,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: dbUser.Createdat,
	}
}
