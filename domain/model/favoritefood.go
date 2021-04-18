package model

import "gorm.io/gorm"

type (
	// FavoriteFood は食べ物の情報を格納
	FavoriteFood struct {
		gorm.Model `json:"-"`
		Food       string
		ProfileID  uint `json:"-"` // json:"-" はレスポンスから隠蔽したいときに使う(https://stackoverflow.com/questions/17306358/removing-fields-from-struct-or-hiding-them-in-json-response)
	}
)
