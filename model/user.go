package model

import (
	"redis/demo/db"
	"reflect"
	"strconv"
)

type User struct {
	Id       int
	Username string
	Email    string
}

func (user *User) Save() error {

	val := reflect.ValueOf(user).Elem()
	typ := val.Type()

	userKey := "user:" + strconv.Itoa(user.Id)

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldVal := val.Field(i).Interface()

		err := db.Rdb.HSet(db.Ctx, userKey, fieldName, fieldVal).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
