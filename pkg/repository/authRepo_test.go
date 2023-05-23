package repository

import (
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
	"fmt"
	"testing"
	"time"

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
func Test_FindAdmin(t *testing.T) {
	testAdmin := domain.AdminResponse{
		ID:       1,
		UserName: "testuser",
		Password: "testpassword",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro in find admin mock :%v", err)
	}
	defer db.Close()
	mockQuery := "SELECT id,user_name,password FROM admins WHERE user_name=\\$1;"
	mock.ExpectQuery(mockQuery).WithArgs(testAdmin.UserName).WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "password"}).AddRow(
		testAdmin.ID, testAdmin.UserName, testAdmin.Password,
	))
	authRepo := &authDatabase{DB: db}
	admin, AcutalErr := authRepo.FindAdmin(testAdmin.UserName)
	assert.NoError(t, AcutalErr)
	assert.Equal(t, testAdmin, admin)

}
func Test_StoreVerificationDetails(t *testing.T) {
	testTime := time.Date(2022, time.March, 1, 10, 0, 0, 0, time.UTC)
	testotp := domain.Verification{
		Creat_At: testTime,
		Email:    "testemail",
		Code:     3456,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error form test storeverification:%v", err)
	}
	defer db.Close()
	mockQuery := "INSERT INTO verifications\\(creat_at,email,code\\)VALUES\\(\\$1,\\$2,\\$3\\);"
	mock.ExpectQuery(mockQuery).WithArgs(testotp.Creat_At, testotp.Email, testotp.Code).WillReturnRows()
	authRepo := NewAuthRepository(db)
	actualError := authRepo.StoreVerificationDetails(testotp.Email, testotp.Code)
	assert.NoError(t, actualError)
}

func Test_VerifyOtp(t *testing.T) {
	testotp := domain.Verification{
		Email: "testemail",
		Code:  1234,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error from very otp :%v", err)
	}
	defer db.Close()
	mockQuery := "SELECT email,code FROM verifications WHERE email=\\$1 and code=\\$2"

	mock.ExpectQuery(mockQuery).WithArgs(testotp.Email, testotp.Code).WillReturnRows()
	authRepo := &authDatabase{DB: db}
	actualErr := authRepo.VerifyOtp(testotp.Email, testotp.Code)
	assert.NoError(t, actualErr)

}
func Test_UpdateUserStatus(t *testing.T){
	db,mock,err:=sqlmock.New()
	if err!=nil{
		t.Fatalf("error in testupdateuserstatus :%v",err)
	}
	testUser:=domain.Users{
		Verification: true,
		Email: "testemail",
	}
	
	defer db.Close()
	mockQuery:="UPDATE users SET verification=\\$1 WHERE email=\\$2;"
	mock.ExpectQuery(mockQuery).WithArgs(testUser.Verification,testUser.Email).WillReturnRows(sqlmock.NewRows([]string{"verification"}).AddRow(testUser.Verification))
	authRepo:=&authDatabase{DB:db}
	actualErr:=authRepo.UpdateUserStatus(testUser.Email)
	assert.NoError(t,actualErr)
}
//test find userbyId
func Test_FindUserById(t *testing.T){
	db,mock,err:=sqlmock.New()
	assert.NoError(t,err)
	defer db.Close()
	testUser:=domain.Users{
		First_Name: "testfname",
		Last_Name: "testlastname",
		Email: "testemail",
		Gender: "testgender",
		Phone: "testphone",
		Verification: true,
		Country: "india",
		City: "palakkad",
		Block_Status: false,
	}
	var user_id int32=1
	mockQuery:= "SELECT first_name,last_name,email,gender,phone,password,verification,country,city,block_status FROM users WHERE user_id=\\$1;"
	mock.ExpectQuery(mockQuery).WithArgs(user_id).WillReturnRows(sqlmock.NewRows([]string{"first_name","last_name","email","gender","phone","password","verification","country","city","block_status"}).AddRow(testUser.First_Name,testUser.Last_Name,testUser.Email,testUser.Gender,testUser.Phone,testUser.Password,testUser.Verification,testUser.Country,testUser.City,testUser.Block_Status))
   authRepo:=&authDatabase{DB:db}
   user,actualErr:=authRepo.FindUserById(uint(user_id))
   assert.NoError(t,actualErr)
   assert.Equal(t,testUser,user)
}
//blockuserstatus
func Test_BlockUnblockUser(t *testing.T){
	db,mock,err:=sqlmock.New()
	assert.NoError(t,err)
	testUser:=domain.Users{
		Block_Status: true,
		User_Id: 1,
	}
	defer db.Close()
	mockQuery:="UPDATE users SET block_status=\\$1 WHERE user_id=\\$2;"
	mock.ExpectQuery(mockQuery).WithArgs(testUser.Block_Status,testUser.User_Id).WillReturnRows()
	authRepo:=&authDatabase{DB:db}
	actualErr:=authRepo.BlockUnblockUser(testUser.User_Id,testUser.Block_Status)
	assert.NoError(t,actualErr)

}