package repository

import (
	"clean/pkg/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

//test register user

func TestRegister(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error register mock DB: %v", err)
	}
	defer db.Close()
	// mockQuery:="INSERT INTO users\\(first_name,last_name,email,gender,phone,password,created_at\\)VALUES\\(\\$1,\\$2,\\$3,\\$4,\\$5,\\$6,\\$7\\)RETURNING user_id;"
	mockQuery := "INSERT INTO users\\(first_name,last_name,email,gender,phone,password\\)VALUES\\(\\$1,\\$2,\\$3,\\$4,\\$5,\\$6\\)RETURNING user_id;"
	testUser := domain.Users{
		First_Name: "testuserfirstname",
		Last_Name:  "testuserlastname",
		Email:      "testuseremail",
		Gender:     "gender",
		Phone:      "testphone",
		Password:   "testpassword",
	}
	// creat expected result
	expectedID := 1
	mock.ExpectQuery(mockQuery).WithArgs(testUser.First_Name, testUser.Last_Name, testUser.Email, testUser.Gender,
		testUser.Phone, testUser.Password).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(expectedID))
	authRepo := NewAuthRepository(db)

	actualID, actualErr := authRepo.Register(testUser)
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedID, actualID)
}

// INSERT INTO users(
// 	first_name,
// 	last_name,
// 	email,
// 	gender,
// 	phone,
// 	password,created_at)
// VALUES($1,$2,$3,$4,$5,$6,$7)
// RETURNING user_id;`
