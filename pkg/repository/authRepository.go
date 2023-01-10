package repository

import (
	domain "clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	"context"
	"database/sql"
	"fmt"
	//"errors"
	//"fmt"
)

type authDatabase struct {
	DB *sql.DB
}

func NewAuthRepository(DB *sql.DB) interfaces.AuthRepository {
	return &authDatabase{
		DB: DB,
	}
}

// ------------------register------------------------------------------------------------------
func (c *authDatabase) Register(ctx context.Context, user domain.Users) (int, error) {
	var User_id int
	query := `INSERT INTO users(
                  first_name,
                  last_name,
                  email,
                  gender,
                  phone,
                  password,
				status)
    VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING user_id;`
	err := c.DB.QueryRow(query,
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Gender,
		user.Phone,
		user.Password,
		user.Status).Scan(&User_id)
	return User_id, err

}

// --------------------------------find user-----------------------------------------------------
func (c *authDatabase) FindUser(email string) (domain.UserResponse, error) {
	var user domain.UserResponse

	query := `SELECT
    user_id,
first_name,
last_name,
email,
gender,
password,
phone
FROM users
WHERE email=$1;`

	err := c.DB.QueryRow(query, email).Scan(&user.ID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Gender,
		&user.Password,
		&user.Phone,
	)
	fmt.Println(err, email, "usdderreopooooooooojjjjj")
	return user, err
}

// --------------------------adminRegister---------------------------------
func (c *authDatabase) AdminRegister(ctx context.Context, admin domain.Admins) error {
	query := `INSERT INTO 
admins (user_name,password)
VALUES ($1,$2);`
	err := c.DB.QueryRow(query, admin.UserName,
		admin.Password,
	).Err()
	return err
}
func (c *authDatabase) FindAdmin(username string) (domain.AdminResponse, error) {
	//log.Println("username of admin", username)
	var admin domain.AdminResponse
	query :=
		`SELECT
id,
user_name,
password
FROM admins WHERE user_name=$1;`
	err := c.DB.QueryRow(query, username).Scan(&admin.ID,
		&admin.UserName,
		&admin.Password)
	return admin, err

}
