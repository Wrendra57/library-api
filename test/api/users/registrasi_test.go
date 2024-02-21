package users

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestRegiterUserSuccess(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "testingname")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.NotNil(t, responseBody["data"], err)
}

// test not valid email
func TestRegisterNotValidEmail(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "testingname")
	writer.WriteField("email", "testinggmail.com")
	writer.WriteField("password", "1234")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Email must be a valid email address", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test empty email
func TestRegisterEmptyEmail(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "testingname")
	writer.WriteField("email", "")
	writer.WriteField("password", "1234")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Email is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test duplicate email
func TestRegisterDuplicateEmail(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "testingname")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
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

	body2 := new(bytes.Buffer)
	writer2 := multipart.NewWriter(body2)
	writer2.WriteField("name", "testingname")
	writer2.WriteField("email", "testing@gmail.com")
	writer2.WriteField("password", "1234")
	writer2.WriteField("gender", "male")
	writer2.WriteField("telp", "62812345678")
	writer2.WriteField("birthdate", "2006-01-02")
	writer2.WriteField("address", "testingkota")

	part2, err := writer2.CreateFormFile("foto", "testimage.png")
	helper.PanicIfError(err)

	_, err = io.Copy(part2, file)
	helper.PanicIfError(err)

	writer2.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	request2 := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body2)
	request2.Header.Set("Content-Type", writer2.FormDataContentType())

	recorder1 := httptest.NewRecorder()
	recorder2 := httptest.NewRecorder()

	router.ServeHTTP(recorder1, request)
	router.ServeHTTP(recorder2, request2)

	response := recorder2.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Email already exists", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test mac char email
func TestRegisterMaxCharEmail(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "testingname")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
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

	body2 := new(bytes.Buffer)
	writer2 := multipart.NewWriter(body2)
	writer2.WriteField("name", "testingname")
	writer2.WriteField("email", "Dolore aliqua ad laborum velit adipisicing culpa. Ex esse proident nulla amet adipisicing velit ea labore eu eu ipsum et cillum. Qui qui ipsum mollit tempor. Anim quis esse ipsum ullamco esse. Aute fugiat occaecat qui voluptate ut. Anim adipisicing consequat ut dolor proident deserunt anim do aute ut ex in minim.@gmail.com")
	writer2.WriteField("password", "1234")
	writer2.WriteField("gender", "male")
	writer2.WriteField("telp", "62812345678")
	writer2.WriteField("birthdate", "2006-01-02")
	writer2.WriteField("address", "testingkota")

	part2, err := writer2.CreateFormFile("foto", "testimage.png")
	helper.PanicIfError(err)

	_, err = io.Copy(part2, file)
	helper.PanicIfError(err)

	writer2.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	request2 := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body2)
	request2.Header.Set("Content-Type", writer2.FormDataContentType())

	recorder1 := httptest.NewRecorder()
	recorder2 := httptest.NewRecorder()

	router.ServeHTTP(recorder1, request)
	router.ServeHTTP(recorder2, request2)

	response := recorder2.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Email must be maximum 100 characters long", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test empty name
func TestRegisterEmptyName(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Name is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test max char name
func TestRegisterMaxCharName(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Voluptate velit aliquip laborum in dolor tempor. Nisi labore mollit Lorem et eu sunt duis esse. Id aliquip anim est minim ipsum laborum eu occaecat anim ad. Ex irure nostrud anim reprehenderit nisi occaecat do in velit.")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Name must be maximum 100 characters long", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test minimum name 3 char
func TestRegisterMinCharName(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "te")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Name must be at least 3 characters long", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

func TestRegisterEmptyPassword(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "tettt")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Password is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test max char password
func TestRegisterMaxCharPassword(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "tettt")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "Fugiat cillum tempor ipsum incididunt mollit dolor dolore adipisicing. Labore dolor reprehenderit non fugiat est magna. Cillum ipsum sint fugiat minim. Fugiat minim qui sint nostrud sunt.")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Password must be maximum 100 characters long", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test empty gender
func TestRegisterEmptyGender(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "tettt")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
	writer.WriteField("gender", "")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Gender is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test not valid gender (male, female)
func TestRegisterNotValidGender(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "tettt")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
	writer.WriteField("gender", "none")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Gender must be a male female", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test empty telp
func TestRegisterEmptyTelp(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "tettt")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
	writer.WriteField("gender", "male")
	writer.WriteField("telp", "")
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Telp is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test empty birthdate
func TestRegisterEmptyBirthdate(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "tettt")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
	writer.WriteField("gender", "male")
	writer.WriteField("telp", "0821324565")
	writer.WriteField("birthdate", "")
	writer.WriteField("address", "testingkota")

	file, err := os.Open("../../asset/testimage.png")
	helper.PanicIfError(err)
	defer file.Close()

	part, err := writer.CreateFormFile("foto", "testimage.png")
	helper.PanicIfError(err)

	_, err = io.Copy(part, file)
	helper.PanicIfError(err)

	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Birthdate is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test not valid format birtdate
func TestRegisterNotValidBirthdate(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "tettt")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
	writer.WriteField("gender", "male")
	writer.WriteField("telp", "0821324565")
	writer.WriteField("birthdate", "3333-22-345")
	writer.WriteField("address", "testingkota")

	file, err := os.Open("../../asset/testimage.png")
	helper.PanicIfError(err)
	defer file.Close()

	part, err := writer.CreateFormFile("foto", "testimage.png")
	helper.PanicIfError(err)

	_, err = io.Copy(part, file)
	helper.PanicIfError(err)

	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Birthdate must be format YYYY-MM-DD", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test empty addres
func TestRegisterEmptyAddress(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "tettt")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
	writer.WriteField("gender", "male")
	writer.WriteField("telp", "0821324565")
	writer.WriteField("birthdate", "2006-09-11")
	writer.WriteField("address", "")

	file, err := os.Open("../../asset/testimage.png")
	helper.PanicIfError(err)
	defer file.Close()

	part, err := writer.CreateFormFile("foto", "testimage.png")
	helper.PanicIfError(err)

	_, err = io.Copy(part, file)
	helper.PanicIfError(err)

	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Address is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

// test empty foto
func TestRegisterEmptyFoto(t *testing.T) {
	db := test.SetupTestDB()
	test.DeleteUser(db)
	router := test.SetupRouter(db)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "tettt")
	writer.WriteField("email", "testing@gmail.com")
	writer.WriteField("password", "1234")
	writer.WriteField("gender", "male")
	writer.WriteField("telp", "0821324565")
	writer.WriteField("birthdate", "2006-09-11")
	writer.WriteField("address", "")

	// file, err := os.Open("../../asset/testimage.png")
	// helper.PanicIfError(err)
	// defer file.Close()

	// part, err := writer.CreateFormFile("foto", "testimage.png")
	// helper.PanicIfError(err)

	// _, err = io.Copy(part, file)
	// helper.PanicIfError(err)

	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	bodyResp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyResp, &responseBody)

	// fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Foto is required", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}
