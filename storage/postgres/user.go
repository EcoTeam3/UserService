package postgres

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
	pb "userService/generated"
)


type NewUser struct{
	Db *sql.DB
}

func NewUserRepo(db *sql.DB)*NewUser{
	return &NewUser{Db: db}
}

func(U *NewUser) GetUser(userId *pb.UserId)(*pb.User, error){
	id := userId.UserId
	user := pb.User{}
	err := U.Db.QueryRow(  `SELECT
								user_id, username, email, created_at
							FROM 
								Users
							WHERE 
								user_id = $1 and deleted_at is null`, id).
							Scan(&user.UserId, &user.Username, &user.Email, &user.CreatedAt)	
	return &user, err			
}


func(U *NewUser) UpdateUser(user *pb.User)(*pb.Status, error){
	query := `UPDATE Users SET WHERE deleted_at is null`
	arr := []interface{}{}
	var param []string
	if len(user.Username) > 0{
		arr = append(arr, user.Username)
		param = append(param, "username")
		query += " and username = :username"
	} 
	if len(user.Email) > 0{
		arr = append(arr, user.Email)
		param = append(param, "email")
		query += " and email = :email"
	}
	if len(user.PasswordHash) > 0{
		arr = append(arr, user.PasswordHash)
		param = append(param, "passwordHash")
		query += " and password_hash = :passwordHash"
	}

	for i, j := range param{
		query = strings.Replace(query, ":" + j, "$" + strconv.Itoa(i+1), 1)
	}

	_, err := U.Db.Exec(query, arr...)
	if err != nil{
		return &pb.Status{Status: false}, err 
	}
	return &pb.Status{Status: true}, err
}

func (U *NewUser) DeleteUser(userId *pb.UserId)(*pb.Status, error) {
	id := userId.UserId
	_,err := U.Db.Exec(`UPDATE Users 
							SET
								deleted_at = $1
							WHERE
								user_id = $2`,time.Now(),id)
	if err != nil{
		return &pb.Status{
			Status: false,
		},err
	}
	return &pb.Status{
		Status: true,
	},nil
}

func (U *NewUser) GetUserProfile(userId *pb.UserId)(*pb.UserProfile,error){
	id := userId.UserId
	userProfile := pb.UserProfile{}
	err := U.Db.QueryRow("SELECT * FROM User_Profile WHERE user_id = $1",id).Scan(
		&userProfile.UserId,&userProfile.FullName,&userProfile.Bio,&userProfile.Location,&userProfile.AvatarUrl,
	)
	if err != nil{
		return nil,err
	}
	return &userProfile,nil
}

func(U *NewUser) UpdateUserProfile(user *pb.UserProfile)(*pb.Status, error){
	query := `UPDATE User_profiles SET WHERE true`
	arr := []interface{}{}
	var param []string
	if len(user.FullName) > 0{
		arr = append(arr, user.FullName)
		param = append(param, "full_name")
		query += " and full_name = :full_name"
	} 
	if len(user.Bio) > 0{
		arr = append(arr, user.Bio)
		param = append(param, "bio")
		query += " and bio = :bio"
	}
	if len(user.Location) > 0{
		arr = append(arr, user.Location)
		param = append(param, "location")
		query += " and location = :location"
	}

	if len(user.AvatarUrl) > 0{
		arr = append(arr, user.AvatarUrl)
		param = append(param, "avatar_url")
		query += " and avatar_url = :avatar_url"
	}

	for i, j := range param{
		query = strings.Replace(query, ":" + j, "$" + strconv.Itoa(i+1), 1)
	}

	_, err := U.Db.Exec(query, arr...)
	if err != nil{
		return &pb.Status{Status: false}, err 
	}
	return &pb.Status{Status: true}, err
}





