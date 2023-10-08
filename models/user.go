package models

import "time"

type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) TableName() string {
	return "user"
}

type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserRegister struct {
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}

func (u *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = u.ID
	resp["email"] = u.Email
	resp["first_name"] = u.FirstName
	resp["last_name"] = u.LastName
	resp["is_active"] = u.IsActive
	resp["created_at"] = u.CreatedAt
	resp["updated_at"] = u.UpdatedAt
	return resp
}
