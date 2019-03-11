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

// SelectAllReference selecting with * in all table
func SelectAllReference(reference *[]Reference) (err error) {
	if err = db.Find(reference).Error; err != nil {
		return err
	}

	return nil
}

// SelectReferenceById selecting request data by id and redirect to destination
func SelectReferenceById(reference *Reference, id string) (err error) {
	if err = db.First(reference, id).Error; err != nil {
		return err
	}

	return nil
}

// InsertReference directly to database
// It returns error or not error (nil)
func InsertReference(reference *Reference) (err error) {
	if err = db.Create(reference).Error; err != nil {
		return err
	}

	return nil
}
