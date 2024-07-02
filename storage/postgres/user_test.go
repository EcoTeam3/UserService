package postgres

import (
	"reflect"
	"testing"
	"userService/generated"
	"userService/storage"
)

func UserRepo(t *testing.T) *NewUser {
	db, err := storage.Connect()
	if err != nil {
		t.Error("ERROR : ", err)
		return nil
	}
	userRepo := NewUserRepo(db)
	return userRepo
}

func TestGetUser(t *testing.T) {
	userRepo := UserRepo(t)

	userId := generated.UserId{UserId: ""}

	user, err := userRepo.GetUser(&userId)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	case1 := &generated.User{
		UserId:       "",
		Username:     "john_doe",
		Email:        "john.doe@example.com",
		PasswordHash: "hashed_password_1",
		CreatedAt:    "",
		DeletedAt:    "",
	}
	if !reflect.DeepEqual(user, case1) {
		t.Error("Result : ", case1, "Expected : ", user)
		return
	}
}

func TestUpdateUser(t *testing.T) {
	userRepo := UserRepo(t)

	user := generated.User{
		UserId:       "",
		Username:     "",
		Email:        "",
		PasswordHash: "",
	}
	Result, err := userRepo.UpdateUser(&user)
	if err != nil {
		t.Error("Error : ", err)
		return
	}

	case1 := &generated.Status{
		Status: true,
	}

	if !reflect.DeepEqual(Result, case1) {
		t.Error("Result : ", case1, "Expected : ", Result)
		return
	}
}

func TestDeleteUser(t *testing.T) {
	userRepo := UserRepo(t)

	userId := generated.UserId{
		UserId: "",
	}

	status, err := userRepo.DeleteUser(&userId)

	if err != nil {
		t.Error("ERROR : ", err)
		return
	}
	case1 := &generated.Status{
		Status: true,
	}

	if !reflect.DeepEqual(status, case1) {
		t.Error("Result : ", case1, "Expected : ", status)
		return
	}
}

func TestGetUserProfile(t *testing.T) {
	userRepo := UserRepo(t)

	userId := generated.UserId{
		UserId: "",
	}

	userProfile, err := userRepo.GetUserProfile(&userId)
	if err != nil {
		t.Error("ERROR : ", err)
		return
	}

	case1 := &generated.UserProfile{
		UserId:    "",
		FullName:  "John Doe",
		Bio:       "Software developer and tech enthusiast.",
		Location:  "San Francisco, CA",
		AvatarUrl: "https://example.com/avatar/john_doe.jpg",
	}
	if !reflect.DeepEqual(userProfile, case1) {
		t.Error("Result : ", case1, "Expected : ", userProfile)
		return
	}
}

func TestUpdateUserProfile(t *testing.T) {
	userRepo := UserRepo(t)

	userProfile := generated.UserProfile{
		UserId:    "",
		FullName:  "",
		Bio:       "",
		Location:  "",
		AvatarUrl: "",
	}
	status,err := userRepo.UpdateUserProfile(&userProfile)
	if err != nil{
		t.Error("ERROR : ",err)
		return
	}

	case1 := &generated.Status{
		Status: true,
	}

	if !reflect.DeepEqual(status,case1){
		t.Error("Result : ",case1,"Expected : ",status)
		return
	}
}
