package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	pb "userService/generated/user"
)

type NewUser struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *NewUser {
	return &NewUser{Db: db}
}

func (U *NewUser) GetUser(userId *pb.UserId) (*pb.User, error) {
	if U.Db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	if userId == nil || userId.UserId == "" {
		return nil, fmt.Errorf("userId is nil or empty")
	}

	user := pb.User{}
	err := U.Db.QueryRow(`SELECT user_id, username, email, created_at
       FROM Users
       WHERE user_id = $1 AND deleted_at IS NULL`, userId.UserId).
		Scan(&user.UserId, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (U *NewUser) UpdateUser(user *pb.User) (*pb.Status, error) {
	query := `UPDATE Users SET user_id = $1`
	arr := []interface{}{user.UserId}
	var param []string
	if len(user.Username) > 0 {
		arr = append(arr, user.Username)
		param = append(param, "username")
		query += ", username = :username"
	}
	if len(user.Email) > 0 {
		arr = append(arr, user.Email)
		param = append(param, "email")
		query += ", email = :email"
	}
	if len(user.PasswordHash) > 0 {
		arr = append(arr, user.PasswordHash)
		param = append(param, "passwordHash")
		query += ", password_hash = :passwordHash"
	}

	n := 2
	for _, j := range param {
		query = strings.Replace(query, ":"+j, "$"+strconv.Itoa(n), 1)
		n++
	}

	query += fmt.Sprintf(" WHERE deleted_at is null and user_id = $%d", n)
	arr = append(arr, user.UserId)
	_, err := U.Db.Exec(query, arr...)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (U *NewUser) DeleteUser(userId *pb.UserId) (*pb.Status, error) {
	_, err := U.Db.Exec(`UPDATE Users 
       SET deleted_at = $1
       WHERE user_id = $2`, time.Now(), userId.UserId)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (U *NewUser) GetUserProfile(userId *pb.UserId) (*pb.UserProfile, error) {
	userProfile := pb.UserProfile{}
	err := U.Db.QueryRow("SELECT user_id, full_name, bio, location, avatar_url FROM User_profiles WHERE user_id = $1", userId.UserId).
		Scan(&userProfile.UserId, &userProfile.FullName, &userProfile.Bio, &userProfile.Location, &userProfile.AvatarUrl)
	if err != nil {
		return nil, err
	}

	return &userProfile, nil
}

func (U *NewUser) UpdateUserProfile(user *pb.UserProfile) (*pb.Status, error) {
	query := `UPDATE User_profiles SET user_id = $1`
	arr := []interface{}{user.UserId}
	var param []string
	if len(user.FullName) > 0 {
		arr = append(arr, user.FullName)
		param = append(param, "full_name")
		query += ", full_name = :full_name"
	}
	if len(user.Bio) > 0 {
		arr = append(arr, user.Bio)
		param = append(param, "bio")
		query += ", bio = :bio"
	}
	if len(user.Location) > 0 {
		arr = append(arr, user.Location)
		param = append(param, "location")
		query += ", location = :location"
	}

	if len(user.AvatarUrl) > 0 {
		arr = append(arr, user.AvatarUrl)
		param = append(param, "avatar_url")
		query += ", avatar_url = :avatar_url"
	}

	n := 2
	for _, j := range param {
		query = strings.Replace(query, ":"+j, "$"+strconv.Itoa(n), 1)
		n++
	}

	query += fmt.Sprintf(" WHERE user_id = $%d", n)
	arr = append(arr, user.UserId)

	_, err := U.Db.Exec(query, arr...)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}


func(U *NewUser) CreateUser(user *pb.User)(*pb.Status, error){
	_, err := U.Db.Exec(`INSERT INTO Users(username, email, password_hash) Values($1, $2, $3)`,
					user.Username, user.Email, user.PasswordHash)
	if err != nil{
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}