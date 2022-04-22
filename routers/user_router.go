package routers

import (
	"encoding/json"
	"errors"
	"time"

	"net/http"
	"os"
	"strings"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/tapiaw38/resources-api/jwt"
	"github.com/tapiaw38/resources-api/models/user"
)

type UserRouter struct {
	Storage user.Storage
}

var userFind user.User

// ValidateToken validates the token
func (ur *UserRouter) ValidateToken(tk string) (*user.Claim, bool, user.User, error) {

	u := user.User{}

	myKey := []byte(os.Getenv("JWT_SECRET"))

	claims := &user.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, u, errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwtGo.ParseWithClaims(tk, claims, func(token *jwtGo.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		user, find := ur.Storage.CheckUser(claims.Email)

		if find {
			userFind = user
		}

		return claims, find, user, nil

	}

	if !tkn.Valid {
		return claims, false, u, errors.New("invalid token")
	}

	return claims, false, u, err

}

// Get the user from the token
func GetUser() user.User {
	return userFind
}

// Login handles the request to login a user
func (ur *UserRouter) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var u user.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "The email or password are invalid "+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "The user email is required", 400)
		return
	}

	foundUser, exist := ur.Login(u.Email, u.Password)

	if !exist {
		http.Error(w, "The email or password are invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(foundUser)

	if err != nil {
		http.Error(w, "An error occurred while generating the token"+err.Error(), 400)
		return
	}

	response := user.LoginResponse{
		User: user.User{
			ID:        foundUser.ID,
			FirstName: foundUser.FirstName,
			LastName:  foundUser.LastName,
			Username:  foundUser.Username,
			Email:     foundUser.Email,
			Picture:   foundUser.Picture,
			IsActive:  foundUser.IsActive,
			IsAdmin:   foundUser.IsAdmin,
		},
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	//coquie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

// Login is function to return the user
func (ur *UserRouter) Login(email string, password string) (user.User, bool) {

	u, find := ur.Storage.CheckUser(email)

	if !find {
		return u, false
	}

	if !u.PasswordMatch(password) {
		return u, false
	}

	return u, true
}

// CreateUserHandler handles the request to create a new user
func (ur *UserRouter) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var u user.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	user, err := ur.Storage.CreateUser(ctx, &u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user in database "+err.Error(), 400)
		return
	}

	u.PasswordHash = ""

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

// GetUsersHandler handles the request to get all users
func (ur *UserRouter) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	users, err := ur.Storage.GetUsers(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get users "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

// GetUserByIdHandler handles the request to get a user by id
func (ur *UserRouter) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	user, err := ur.Storage.GetUserById(ctx, id)

	if err != nil {
		http.Error(w, "No record found with that id "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

// GetUserByUsernameHandler handles the request to get a user by username
func (ur *UserRouter) GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	username := r.URL.Query().Get("username")

	user, err := ur.Storage.GetUserByUsername(ctx, username)

	if err != nil {
		http.Error(w, "No record found with that username "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

// UpdateUserHandler handles the request to update a user
func (ur UserRouter) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	var u user.User

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	user, err := ur.Storage.UpdateUser(ctx, id, u)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a user in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// DeleteUserHandler handles the request to delete a user
func (ur UserRouter) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err := ur.Storage.DeleteUser(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred while trying to delete a record from the database"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
