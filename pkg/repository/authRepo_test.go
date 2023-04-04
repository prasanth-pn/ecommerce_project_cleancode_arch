package repository

import (
	"clean/pkg/domain"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

//test register user

func Test_Register(t *testing.T) {
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

//test finduser

func Test_Finduser(t *testing.T) {
	testUser := domain.UserResponse{
		ID:         1,
		First_Name: "testuserfirstname",
		Last_Name:  "testuserlastname",
		Email:      "testemail",
		Gender:     "male",
		Password:   "passoword",
		Phone:      "testphone",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error in mock find user db:%v", err)
	}
	defer db.Close()

	mockQuery := "SELECT user_id,first_name,last_name,email,gender,password,phone FROM users WHERE email=\\$1;"
	mock.ExpectQuery(mockQuery).WithArgs(testUser.Email).WillReturnRows(sqlmock.NewRows([]string{"user_id", "first_name", "last_name", "email", "gender", "password", "phone"}).
		AddRow(testUser.ID, testUser.First_Name, testUser.Last_Name, testUser.Email, testUser.Gender, testUser.Password, testUser.Phone))

	authRepo := authDatabase{DB: db}
	expected, ActualErr := authRepo.FindUser(testUser.Email)

	assert.NoError(t, ActualErr)
	assert.Equal(t, testUser, expected)

}

//test adminRegister

func Test_AdminRegister(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error register admin in DB:%v", err)
	}
	defer db.Close()
	//	mockQuery := "INSERT INTO admins \\(user_name,password\\)VALUES\\(\\$1,\\$2\\);"
	//test admin
	testAdmin := domain.Admins{
		UserName: "testuser",
		Password: "testpassword",
	}

	mock.ExpectQuery("INSERT INTO admins\\(user_name,password\\)VALUES\\(\\$1,\\$2\\);").WithArgs(testAdmin.UserName, testAdmin.Password).WillReturnRows()
	authRepo := NewAuthRepository(db)
	actualErr := authRepo.AdminRegister(testAdmin)
	fmt.Println("the error is :", actualErr)
	assert.NoError(t, actualErr)

}
