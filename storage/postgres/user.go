package postgres

import (
	"database/sql"
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
								user_id = $1`, id).
							Scan(&user.UserId, &user.Username, &user.Email, &user.CreatedAt)	
	return &user, err			
}