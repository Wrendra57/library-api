package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/test"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateSucces(t *testing.T) {
	db := test.SetupTestDB()
	// test.DeleteUser(db)
	router := test.SetupRouter(db)
	email := "testing@gmail.com"
	password := "1234"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "testingname")
	writer.WriteField("email", email)
	writer.WriteField("password", password)
	writer.WriteField("gender", "male")
	writer.WriteField("telp", "62812345678")
	writer.WriteField("birthdate", "2006-01-02")
	writer.WriteField("address", "testingkota")

	file, err := os.Open("../../asset/testimage.png")
	helper.PanicIfError(err)
	defer file.Close()

	part, err := writer.CreateFormFile("foto", "testimage.png")
	helper.PanicIfError(err)

	_, err = io.Copy(part, file)
	helper.PanicIfError(err)
	writer.Close()

	requestRegisterUser := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	requestRegisterUser.Header.Set("Content-Type", writer.FormDataContentType())

	recorderRegister := httptest.NewRecorder()

	router.ServeHTTP(recorderRegister, requestRegisterUser)

	loginData := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonData, err := json.Marshal(loginData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	requestLogin := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/login", bytes.NewBuffer(jsonData))
	requestLogin.Header.Set("Content-Type", "application/json")

	recorderLogin := httptest.NewRecorder()
	router.ServeHTTP(recorderLogin, requestLogin)

	response := recorderLogin.Result()

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)
	token, ok := responseBody["data"].(map[string]interface{})["token"].(string)
	if !ok {
		fmt.Println("Token not found in the response.")
		return
	}
	fmt.Println(token)

	tests := []struct {
		name_test       string
		request         string
		token           string
		expected_code   int
		expected_status string
		expected_data   string
	}{
		// {
		// 	name_test:       "Success",
		// 	request:         "GET",
		// 	token:           token,
		// 	expected_code:   200,
		// 	expected_status: "OK",
		// 	expected_data:   "success",
		// },
		// {
		// 	name_test:       "NotValidToken",
		// 	request:         "GET",
		// 	token:           token + "f",
		// 	expected_code:   401,
		// 	expected_status: "OK",
		// 	expected_data:   "failed",
		// },
		{
			name_test:       "NotValidToken",
			request:         "GET",
			token:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RpbmdAZ21haWwuY29tIiwiZXhwIjoxNzA1NDYxNTY0LCJpZCI6MTA1LCJsZXZlbCI6Im1lbWJlciJ9.Tv13ESqu6NqG8GyQuA__VkqtYf7i5pt40GO7shVNwWE",
			expected_code:   401,
			expected_status: "OK",
			expected_data:   "failed",
		},
	}

	for _, test := range tests {
		t.Run(test.name_test, func(t *testing.T) {
			requestLogin := httptest.NewRequest(test.request, "http://localhost:8001/api/user", nil)
			requestLogin.Header.Set("Authorization", "Bearer "+test.token)

			recorderLogin := httptest.NewRecorder()
			router.ServeHTTP(recorderLogin, requestLogin)

			response := recorderLogin.Result()
			// assert.Equal(t, test.expected_code, response.StatusCode)

			bodyResp, _ := io.ReadAll(response.Body)
			var responseBody map[string]interface{}
			json.Unmarshal(bodyResp, &responseBody)

			fmt.Println(responseBody)
			assert.Equal(t, test.expected_code, int(responseBody["code"].(float64)))
			assert.Equal(t, test.expected_status, responseBody["status"])
			if test.expected_data == "success" {
				assert.NotNil(t, responseBody["data"])
			} else {
				assert.Nil(t, responseBody["data"])
			}
		})
	}

}
