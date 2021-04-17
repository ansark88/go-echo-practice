// Package config はhttps://github.com/HON9LIN/go-echo-boilerplate/blob/master/config/database.go を参考にした
package config

import "os"

// Database はデータベースのconfigのインターフェイス
type Database interface{}

// PsqlDbConnection はDB configの中身
type PsqlDbConnection struct {
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string
}

// DatabaseConfig はconfigの中身を格納する
type DatabaseConfig struct {
	Db PsqlDbConnection
}

// DatabaseNew はconfigのインスタンスを環境変数を使って作成
func DatabaseNew() Database {
	return &DatabaseConfig{
		Db: PsqlDbConnection{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbDatabase: os.Getenv("DB_DATABASE"),
			DbUsername: os.Getenv("DB_USERNAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}
}
