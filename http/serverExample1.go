package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type TestStruct struct {
}

type UserHandler struct {
}

type UserModel struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
}

type ApiResult struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	Error   string      `json:"error"`
}

var users map[int]UserModel = map[int]UserModel{}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/google", http.RedirectHandler("https://google.com", 301))
	mux.Handle("/users", &UserHandler{})

	mux.HandleFunc("/test", test)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: &TestStruct{},
	// }
	server.ListenAndServe()

}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	// w.Write([]byte(`{"value":"test"}`))
	fmt.Fprintf(w, "{\"value\":\"%s\",\"method\":\"%s\"}", "test", r.Method)
}

func (testStruct *TestStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome")

}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) > 0:
		getUser(w, r.URL.Query().Get("id"))
		return
	case r.Method == http.MethodGet:
		getUsers(w, r)
		return
	case r.Method == http.MethodPost:
		createUser(w, r)
		return

	}

}

func getUser(w http.ResponseWriter, id string) {
	i, _ := strconv.Atoi(id)
	user, exists := users[i]
	if !exists {
		setApiResult(w, http.StatusNotFound, nil, errors.New("user not exists"))
		return
	}
	setApiResult(w, http.StatusOK, user, nil)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	setApiResult(w, http.StatusOK, users, nil)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Api-key")
	if authorization != "545655" {
		setApiResult(w, http.StatusUnauthorized, nil, errors.New("invalid api key"))
		return
	}

	var user UserModel

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		setApiResult(w, http.StatusInternalServerError, nil, errors.New("error in decode body"))

		return
	}

	users[user.Id] = user
	setApiResult(w, http.StatusOK, user, nil)

}

func setApiResult(w http.ResponseWriter, statusCode int, result interface{}, errorInfo error) {
	apiResut := ApiResult{}
	if statusCode != http.StatusOK {
		apiResut.Success = false
		apiResut.Error = errorInfo.Error()
	} else {
		apiResut.Success = true
	}

	apiResut.Result = result

	jsonResult, _ := json.Marshal(apiResut)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)

	w.Write(jsonResult)
	return
}
