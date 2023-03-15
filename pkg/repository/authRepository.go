package repository

import (
	domain "clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	"database/sql"
	"errors"
	"fmt"
	"time"
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
func (c *authDatabase) Register(user domain.Users) (int, error) {
	var User_id int
	user.Created_At = time.Now()
	query := `INSERT INTO users(
                  first_name,
                  last_name,
                  email,
                  gender,
                  phone,
                  password,created_at)
    VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING user_id;`
	err := c.DB.QueryRow(query,
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Gender,
		user.Phone,
		user.Password,
		user.Created_At).Scan(&User_id)
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
	if user.Email == "" {
		return user, nil
	}

	return user, err
}

// --------------------------adminRegister---------------------------------
func (c *authDatabase) AdminRegister(admin domain.Admins) error {
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
func (c *authDatabase) StoreVerificationDetails(email string, code int) error {
	var time = time.Now()
	query := `INSERT INTO verifications(creat_at,email,code)
	VALUES($1,$2,$3);`
	err := c.DB.QueryRow(query, time, email, code).Err()

	return err
}
func (c *authDatabase) VerifyOtp(email string, code int) error {
	var vr domain.Verification
	query := `SELECT email,code FROM verifications WHERE email=$1 and code=$2;`
	err := c.DB.QueryRow(query, email, code).Scan(&vr.Email, &vr.Code)
	//fmt.Println(err, "sellelnkfjkfkljrrofon")
	if err != nil {
		return errors.New("email or otp incorrect")
	}
	return nil

}
func (c *authDatabase) UpdateUserStatus(email string) error {
	var user domain.Users
	query := `UPDATE users SET verification=$1 WHERE email=$2;`
	err := c.DB.QueryRow(query, true, email).Scan(&user.Verification)
	if err != nil {
		return err
	}
	return nil
}
func (c *authDatabase) FindUserById(user_id uint) (domain.Users, error) {
	var user domain.Users

	query := `SELECT first_name,last_name,email,gender,phone,password,verification,country,
	city,block_status FROM users WHERE user_id=$1;`
	err := c.DB.QueryRow(query, user_id).Scan(&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Gender, &user.Phone, &user.Password, &user.Verification, &user.Country, &user.City, &user.Block_Status)

	fmt.Println(err)
	return user, err
}
func (c *authDatabase) BlockUnblockUser(user_id uint, val bool) error {
	query := `UPDATE users SET block_status=$1 WHERE user_id=$2;`
	err := c.DB.QueryRow(query, val, user_id).Err()
	return err
}
