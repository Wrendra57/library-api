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

func TestLoginSucces(t *testing.T) {
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
	assert.Equal(t, 200, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.NotNil(t, responseBody["data"], err)
}

// test empty email
func TestLoginEmptyEmail(t *testing.T) {
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
		"email":    "",
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
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Email is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test not valid email
func TestLoginNotValidEmail(t *testing.T) {
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
		"email":    "testming.com",
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
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Email must be a valid email address", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test  empty password
func TestLoginEmptyPassword(t *testing.T) {
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
		"password": "",
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
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Password is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test not found email
func TestLoginNotFoundEmail(t *testing.T) {
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
		"email":    "test@gmail.com",
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
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Email test@gmail.com not found", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

func TestLoginWrongPassword(t *testing.T) {
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
		"password": "12345",
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
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Password not match", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}
