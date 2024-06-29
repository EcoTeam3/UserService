package postgres

import (
	"database/sql"
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
