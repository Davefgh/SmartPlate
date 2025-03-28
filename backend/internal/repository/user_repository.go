package repository

import (
	"vehicle-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository  struct{
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

//create a new user
func (r *UserRepository) Create (user *models.User) error {
	query := `INSERT INTO users (last_name, first_name, middle_name, email, password, role, status, lto_client_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			RETURNING user_id, created, updated`
	return r.db.QueryRow(query, user.LAST_NAME, user.FIRST_NAME, user.MIDDLE_NAME, user.EMAIL, user.PASSWORD, user.ROLE, user.STATUS, user.LTO_CLIENT_ID).Scan(
		&user.USER_ID, &user.CREATED, &user.UPDATED)
}

//get all users
func (r *UserRepository) GetAll() ([]models.User, error){
	var users []models.User
	err := r.db.Select(&users, "SELECT * FROM users ORDER BY user_id")
	return users, err
}

//get user by id
func (r *UserRepository) GetByID(user_id int) (models.User, error){
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE user_id = $1", user_id)
	return user, err
}

// In your repository/user_repository.go
func (r *UserRepository) GetByLTOClientID(ltoClientID string) (models.User, error) {
    var user models.User
    err := r.db.Get(&user, "SELECT * FROM users WHERE lto_client_id = $1", ltoClientID)
    return user, err
}
//get user by email.l
func (r *UserRepository) GetByEmail(email string) (models.User, error){
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	return user, err
}

//update user
func (r *UserRepository) Update(user * models.User) error{
	queryv := `UPDATE users SET last_name = $1, first_name = $2, middle_name = $3, email = $4, password = $5, role = $6, status = $7, lto_client_id = $8
	WHERE user_id = $9`
	_, err := r.db.Exec(queryv, user.LAST_NAME, user.FIRST_NAME, user.MIDDLE_NAME, user.EMAIL, user.PASSWORD, user.ROLE, user.STATUS, user.LTO_CLIENT_ID, user.USER_ID)
	return err
}

//delete by id use
func (r *UserRepository) Delete(user_id int) error{
	_, err := r.db.Exec("DELETE FROM users WHERE user_id = $1", user_id)
	return err
}