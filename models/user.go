package models

import (
    "html"
    "strings"

	"gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"

)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
}


func (u *User) Save() (*User, error) {

	var err error
	err = Db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	//remove spaces in username 
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error){
    var user User
    err := Db.Where("username=?", username).Find(&user).Error
    if err != nil {
        return User{}, err
    }
    return user, nil
}
