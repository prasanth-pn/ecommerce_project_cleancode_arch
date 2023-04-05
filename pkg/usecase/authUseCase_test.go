package usecase

import (
	"clean/pkg/config"
	"clean/pkg/domain"
	"clean/pkg/mock"
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
		{
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
