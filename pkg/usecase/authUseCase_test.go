package usecase

import (
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/config"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/mock"
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func Test_SendVerificationEmail(t *testing.T) {

}
func Test_FindAdmin(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	c := mock.NewMockAuthRepository(ctr)
	authUseCase := NewAuthUseCase(c, config.NewMailConfig(), config.Config{})
	testauth := []struct {
		name        string
		userName    string
		bftest      func(authRepo *mock.MockAuthRepository)
		expectError error
	}{
		{
			name:     "test success",
			userName: "prasanth",
			bftest: func(authRepo *mock.MockAuthRepository) {
				authRepo.EXPECT().FindAdmin("prasanth").Return(domain.AdminResponse{
					UserName: "prasanth",
					Password: "12345",
				}, nil)
			},
			expectError: nil,
		},
		{
			name:     "user already exists",
			userName: "prasanth",
			bftest: func(authRepo *mock.MockAuthRepository) {
				authRepo.EXPECT().FindAdmin("prasanth").Return(domain.AdminResponse{
					UserName: "prasanth",
					Password: "12345",
				}, errors.New("user already exists"))
			},
			expectError: errors.New("user already exists"),
		},
	}
	for _, test := range testauth {
		t.Run(test.name, func(t *testing.T) {
			test.bftest(c)
			actualUser, err := authUseCase.FindAdmin(test.userName)
			assert.Equal(t, test.expectError, err)
			if err == nil {
				assert.Equal(t, test.userName, actualUser.UserName)
			}
		})
	}
}
func Test_FindUserById(t *testing.T) {
	users := domain.Users{
		First_Name: "prasanth",
		Last_Name:  "pn",
		Email:      "prasanthpn68@gmail.com",
	}
	ctlr := gomock.NewController(t)
	defer ctlr.Finish()
	c := mock.NewMockAuthRepository(ctlr)
	authUseCase := NewAuthUseCase(c, config.NewMailConfig(), config.Config{})
	testauth := []struct {
		name        string
		user_id     int
		beforetest  func(authRepo *mock.MockAuthRepository)
		expectedErr error
	}{
		//{
			name:    "success",
			user_id: 1,
			beforetest: func(authRepo *mock.MockAuthRepository) {
				authRepo.EXPECT().FindUserById(uint(1)).Return(domain.Users{
					First_Name: "prasanth",
					Last_Name:  "pn",
					Email:      "prasanthpn68@gmail.com",
				}, nil)

			},
			expectedErr: nil,
		},
	}
	for _, test := range testauth {
		test.beforetest(c)
		actualUser, err := authUseCase.FindUserById(uint(test.user_id))
		assert.Equal(t, test.expectedErr, err)
		if err == nil {
			assert.Equal(t, users, actualUser)
		}

	}
}
func Test_FindUser(t *testing.T) {
	user := domain.UserResponse{
		First_Name: "prasanthpn",
		Email:      "prasanthpn68@gmail.com",
	}
	ctl := gomock.NewController(t)
	c := mock.NewMockAuthRepository(ctl)
	authUseCase := NewAuthUseCase(c, config.NewMailConfig(), config.Config{})
	testcase := []struct {
		name        string
		email       string
		beforetest  func(authRepo *mock.MockAuthRepository)
		expectedErr error
	}{
		{
			name:  "success",
			email: "prasanthpn68@gmail.com",
			beforetest: func(authRepo *mock.MockAuthRepository) {
				authRepo.EXPECT().FindUser("prasanthpn68@gmail.com").Return(user, nil)

			},
			expectedErr: nil,
		},
	}
	for _, test := range testcase {
		test.beforetest(c)
		actualUser, err := authUseCase.FindUser(user.Email)
		assert.Equal(t, test.expectedErr, err)
		if err != nil {
			assert.Equal(t, user, actualUser)

		}
	}

}
func Test_BlockUnblockUser(t *testing.T) {
	ctl := gomock.NewController(t)
	c := mock.NewMockAuthRepository(ctl)
	authUsecase := NewAuthUseCase(c, config.NewMailConfig(), config.Config{})
	testcase := []struct {
		name        string
		user_id     uint
		value       bool
		beforetest  func(autheusecase *mock.MockAuthRepository)
		ExpectedErr error
	}{
		{
			name:    "unblocked",
			user_id: 1,
			value:   true,
			beforetest: func(autheusecase *mock.MockAuthRepository) {
				autheusecase.EXPECT().BlockUnblockUser(uint(1), true).Return(nil)
			},
			ExpectedErr: nil,
		},
		{
			name:    "blocked",
			user_id: 1,
			value:   false,
			beforetest: func(authUseCase *mock.MockAuthRepository) {
				authUseCase.EXPECT().BlockUnblockUser(uint(1), false).Return(nil)

			},
			ExpectedErr: nil,
		},
	}
	for _, test := range testcase {
		test.beforetest(c)
		actualErr := authUsecase.BlockUnblockUser(test.user_id, test.value)
		assert.Equal(t, test.ExpectedErr, actualErr)
	}

}

// userRegister
func Test_Register(t *testing.T) {
	user := domain.Users{
		First_Name: "prasanth",
		Email:      "prasanthpn68@gmail.com",
	}
	// userresp := domain.UserResponse{
	// 	First_Name: "prasanth",
	// 	Email:      "prasanthpn68@gmail.com",
	// }

	ctl := gomock.NewController(t)
	c := mock.NewMockAuthRepository(ctl)
	defer ctl.Finish()

	authUseCase := NewAuthUseCase(c, config.NewMailConfig(), config.Config{})

	testcases := []struct {
		name        string
		testUser    domain.Users
		beforeTest  func(authRepo *mock.MockAuthRepository)
		expectedErr error
	}{
		{
			name:     "user already exists",
			testUser: user,
			beforeTest: func(authRepo *mock.MockAuthRepository) {
				authRepo.EXPECT().FindUser("prasanthpn68@gmail.com").Return(domain.UserResponse{
					First_Name: "prasanth",
					Email:      "prasanthpn68@gmail.com",
				}, nil)
			},
			expectedErr:errors.New("user already exists"),
		},
		{
			name:     "register success",
			testUser: user,
			beforeTest: func(authRepo *mock.MockAuthRepository) {
				authRepo.EXPECT().FindUser("prasanthpn68@gmail.com").Return(domain.UserResponse{}, nil)
				authRepo.EXPECT().Register(user).Return(1,nil)
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		tc.beforeTest(c)

		users, actualErr := authUseCase.Register(user)
		assert.Equal(t, tc.expectedErr, actualErr)
		 assert.Equal(t, user, users)
	}
}