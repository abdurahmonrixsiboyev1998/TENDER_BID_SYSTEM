package repository

import (
	"context"
	"database/sql"
	"fmt"
	auth "tender_bid_system/auth/hash"
	"tender_bid_system/auth/token"
	"tender_bid_system/model"

	"github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (model.User, error) {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = hashedPassword
	if user.Role != "client" && user.Role != "bidder" {
		return model.User{}, fmt.Errorf("invalid role")
	}

	query, args, err := squirrel.Insert("users").
		Columns("username", "email", "password", "role").
		Values(user.Username, user.Email, user.Password, user.Role).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return model.User{}, err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	query, args, err := squirrel.Select("id", "username", "email", "password", "role").
		From("users").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return model.User{}, err
	}
	var user model.User
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Login(ctx context.Context, email string, password string) (string, error) {
	query, args, err := squirrel.Select("id", "username", "password").
		From("users").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return "", err
	}
	var userID int
	var username string
	var hash string
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&userID, &username, &hash)
	err = auth.ValidHash(hash, password)
	if err != nil {
		return "", err
	}
	if userID > 0 {
		token, err := token.GenerateToken(email, username)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", err
}
