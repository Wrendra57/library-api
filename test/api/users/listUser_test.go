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

func TestListUserSucces(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	email := "admin@gmail.com"
	password := "1234"

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
	assert.Equal(t, 200, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println("login", responseBody)
	token, ok := responseBody["data"].(map[string]interface{})["token"].(string)
	if !ok {
		fmt.Println("Token not found in the response.")
		return
	}

	requestListUser := httptest.NewRequest(http.MethodGet, "http://localhost:8001/api/users", nil)
	requestListUser.Header.Set("Authorization", "Bearer "+token)

	recorderListUser := httptest.NewRecorder()
	router.ServeHTTP(recorderListUser, requestListUser)

	responseListUser := recorderListUser.Result()

	bodyResp2, _ := io.ReadAll(responseListUser.Body)
	var responseBody2 map[string]interface{}
	json.Unmarshal(bodyResp2, &responseBody2)

	assert.Equal(t, 200, int(responseBody2["code"].(float64)))
	assert.Equal(t, "OK", responseBody2["status"])
	assert.NotNil(t, responseBody2["data"])
}

func TestListUserWithoutToken(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)
	requestListUser := httptest.NewRequest(http.MethodGet, "http://localhost:8001/api/users", nil)

	recorderListUser := httptest.NewRecorder()
	router.ServeHTTP(recorderListUser, requestListUser)

	responseListUser := recorderListUser.Result()

	bodyResp, _ := io.ReadAll(responseListUser.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "unauthorized", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

func TestListUserAnotherAdminrole(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
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

	bodyRespLogin, _ := io.ReadAll(response.Body)
	var respLogin map[string]interface{}
	json.Unmarshal(bodyRespLogin, &respLogin)

	// fmt.Println("login", respLogin)
	token, ok := respLogin["data"].(map[string]interface{})["token"].(string)
	if !ok {
		fmt.Println("Token not found in the response.")
		return
	}

	requestListUser := httptest.NewRequest(http.MethodGet, "http://localhost:8001/api/users", nil)
	requestListUser.Header.Set("Authorization", "Bearer "+token)

	recorderListUser := httptest.NewRecorder()
	router.ServeHTTP(recorderListUser, requestListUser)

	responseListUser := recorderListUser.Result()

	bodyResp, _ := io.ReadAll(responseListUser.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "unauthorized", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

func TestListUserInvalidToken(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	email := "admin@gmail.com"
	password := "1234"

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
	assert.Equal(t, 200, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println("login", responseBody)
	token, ok := responseBody["data"].(map[string]interface{})["token"].(string)
	if !ok {
		fmt.Println("Token not found in the response.")
		return
	}

	requestListUser := httptest.NewRequest(http.MethodGet, "http://localhost:8001/api/users", nil)
	requestListUser.Header.Set("Authorization", "Bearer "+token+"t")

	recorderListUser := httptest.NewRecorder()
	router.ServeHTTP(recorderListUser, requestListUser)

	responseListUser := recorderListUser.Result()

	bodyResp2, _ := io.ReadAll(responseListUser.Body)
	var responseBody2 map[string]interface{}
	json.Unmarshal(bodyResp2, &responseBody2)

	assert.Equal(t, 401, int(responseBody2["code"].(float64)))
	assert.Equal(t, "token signature is invalid: signature is invalid", responseBody2["status"])
	assert.Nil(t, responseBody2["data"])
}
