package handler

import (
	"bytes"
	"clean/pkg/common/response"
	"clean/pkg/mock"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_VerifyOtp(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	c := mock.NewMockAuthUseCase(ctl)
	authHandler := NewAuthHandler(c, nil)
	testcase := []struct {
		name             string
		email            string
		code             string
		beforetest       func(authusecas *mock.MockAuthUseCase)
		expectedcode     int
		expectedresponse response.Response
		expectedErr      error
	}{
		{
			name:  "success",
			email: "emaildotcome",
			code :"1234",
			beforetest: func(authusecase *mock.MockAuthUseCase) {
				authusecase.EXPECT().VerifyUserOtp("emaildotcome", 1234).Return(nil)
				authusecase.EXPECT().UpdateUserStatus("emaildotcome").Return(nil)

			},
			expectedcode: 200,
			expectedresponse: response.Response{
				Status:  true,
				Message: "success",
				Errors:  nil,
				Data:    "emaildotcome",
			},
			expectedErr: errors.New("eroroor"),
		},
	}
	for _, test := range testcase {
		t.Run(test.name, func(t *testing.T) {
			test.beforetest(c)
			gin := gin.New()
			rec := httptest.NewRecorder()
			gin.GET("/verify/otp", authHandler.VerifyUserOtp)
			var body []byte
			req := httptest.NewRequest("GET", "/verify/otp", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			//set a query parameter named "query " with a value of "param"
			q := req.URL.Query()
			q.Add("email", test.email)
			q.Add("code", test.code)
			req.URL.RawQuery=q.Encode()
			gin.ServeHTTP(rec,req)
			var acutal response.Response
			err:=json.Unmarshal(rec.Body.Bytes(),&acutal)
			assert.NoError(t,err)
			assert.Equal(t,test.expectedcode,rec.Code)
			assert.Equal(t,test.expectedresponse.Message,acutal.Message)


		})
	}
}

// 		},
// 		{
// 			name:  "test sucsess response",
// 			email: "ali",
// 			code:  "54321",
// 			beforeTest: func(userUsecase mock.MockAuthUseCase) {
// 				userUsecase.EXPECT().WorkerVerifyAccount("ali", 54321).Return(errors.New("usecase error"))
// 			},
// 			expectCode: 422,
// 			expectResponse: response.Response{
// 				Status:  false,
// 				Message: "Error while verifing worker mail",
// 				Errors:  []interface{}{"usecase error"},
// 				Data:    nil,
// 			},
// 			expectErr: errors.New("usecase error"),
// 		},
// 	}

// 	for _, tt := range testData {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.beforeTest(*c)

// 			gin := gin.New()
// 			rec := httptest.NewRecorder()

// 			gin.GET("/verify/account", authHandler.WorkerVerifyAccount)

// 			var body []byte
// 			req := httptest.NewRequest("GET", "/verify/account", bytes.NewBuffer(body))
// 			req.Header.Set("Content-Type", "application/json")

// 			// Set a query parameter named "gury" with a value of "param"
// 			q := req.URL.Query()
// 			q.Add("email", tt.email)
// 			q.Add("code", tt.code)
// 			req.URL.RawQuery = q.Encode()

// 			gin.ServeHTTP(rec, req)

// 			var actual response.Response
// 			err := json.Unmarshal(rec.Body.Bytes(), &actual)
// 			assert.NoError(t, err)

// 			assert.Equal(t, tt.expectCode, rec.Code)
// 			assert.Equal(t, tt.expectResponse.Status,actual.Status)
// 			assert.Equal(t,tt.expectResponse.Message,actual.Message)
// 			assert.Equal(t,tt.expectResponse.Errors,actual.Errors)
// 			assert.Equal(t,tt.expectResponse.Data,actual.Data)

// 		})
// 	}
// }
