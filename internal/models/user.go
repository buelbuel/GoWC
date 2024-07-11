package models

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        string
	Email     string
	Username  string
	Password  string
	Admin     bool
	IsActive  bool
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserInterface interface {
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
	GetUserByEmail(email string) (*User, error)
}

type UserModel struct {
	DB     *sql.DB
	Logger echo.Logger
}

func NewUserModel(db *sql.DB, logger echo.Logger) *UserModel {
	return &UserModel{DB: db, Logger: logger}
}

func (model *UserModel) CreateUser(user *User) error {
	user.CreatedAt = time.Now()
	user.IsActive = true

	query := `INSERT INTO public.users (email, username, password, admin, is_active, created_at) 
						VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	model.Logger.Infof("Creating user: %+v", user)
	model.Logger.Infof("Hashed password length: %d", len(user.Password))

	err := model.DB.QueryRow(query, user.Email, user.Username, user.Password, user.Admin, user.IsActive, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		model.Logger.Errorf("Error creating user: %v", err)
		return err
	}

	model.Logger.Infof("User created successfully with ID: %s", user.ID)
	return nil
}

func (model *UserModel) GetUser(id string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, username, password, admin, is_active, created_at, updated_at 
              FROM public.users WHERE id = $1 AND is_deleted = false`

	err := model.DB.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Username, &user.Password,
		&user.Admin, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (model *UserModel) UpdateUser(user *User) error {
	user.UpdatedAt = time.Now()

	query := `UPDATE public.users 
              SET email = $1, username = $2, password = $3, admin = $4, is_active = $5, updated_at = $6 
              WHERE id = $7`

	_, err := model.DB.Exec(query, user.Email, user.Username, user.Password, user.Admin, user.IsActive, user.UpdatedAt, user.ID)
	return err
}

func (model *UserModel) DeleteUser(id string) error {
	query := `UPDATE public.users SET is_deleted = true, deleted_at = now() WHERE id = $1`
	_, err := model.DB.Exec(query, id)
	return err
}

func (model *UserModel) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, username, password, admin, is_active, created_at, updated_at 
						FROM public.users WHERE email = $1 AND is_deleted = false`

	model.Logger.Infof("Fetching user by email: %s", email)

	err := model.DB.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Username, &user.Password,
		&user.Admin, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		model.Logger.Errorf("Error fetching user by email: %v", err)
		return nil, err
	}

	model.Logger.Infof("User found: %+v", user)
	model.Logger.Infof("Hashed password length: %d", len(user.Password))

	return user, nil
}
