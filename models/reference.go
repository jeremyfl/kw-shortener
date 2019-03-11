package models

import (
	"kw-shortener/config"

	"github.com/jinzhu/gorm"
)

// Reference basic struct model
type Reference struct {
	gorm.Model
	Orig        string `json:"orig" xml:"orig" form:"orig" query:"orig"`
	Destination string `json:"destination" xml:"destination" form:"destination" query:"destination"`
}

var db = config.Database()

// StoreReference directly to database
// It returns error or not error (nil)
func StoreReference(reference *Reference) (err error) {
	if err = db.Create(reference).Error; err != nil {
		return err
	}

	return nil
}
