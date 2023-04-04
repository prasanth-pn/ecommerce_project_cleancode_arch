 package handler

// import (
// 	"bytes"
// 	"clean/pkg/domain"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// )

// type login []struct{}
// func (m mockAuthUseCase) Register(user domain.Users) (domain.Users, error) {
// 	user.User_Id = 1
// 	return user, nil
// }

// func TestAuthHandler_Register(t *testing.T) {
// 	// Create a new AuthHandler instance.
// 	authHandler := &AuthHandler{
// 		authUseCase: mockAuthUseCase{},
// 	}

// 	// Create a new HTTP request.
// 	requestBody := []byte(`{"First_Name": "John", "Last_Name": "Doe", "Email": "johndoe@example.com", "Password": "password"}`)
// 	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Set the Content-Type header to JSON.
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create a new ResponseRecorder to record the response.
// 	rr := httptest.NewRecorder()

// 	// Call the Register method, passing in the ResponseRecorder and the HTTP request.
// 	router := gin.Default()
// 	router.POST("/register", authHandler.Register)
// 	router.ServeHTTP(rr, req)

// 	// Check the status code of the response.
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	// Check the response body.
// 	expectedResponse := `{"success":true,"message":"user registration completed  successfully","data":"welcome John"}`
// 	if rr.Body.String() != expectedResponse {
// 		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponse)
// 	}
// }

// // mockAuthUseCase is a mock implementation of the AuthUseCase interface.
// type mockAuthUseCase struct{}


