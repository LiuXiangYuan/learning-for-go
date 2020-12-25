package db

import "errors"

type DB interface {
	Get(key int) (string, error)
}

func GetValue(db DB, key int) (string, error) {
	value, err := db.Get(key)
	if err != nil {
		return "", errors.New("fail")
	}
	return value, nil
}
