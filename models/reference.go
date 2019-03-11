package models

type Reference struct {
	ID          int    `json:"id" xml:"id"`
	Orig        string `json:"orig" xml:"orig"`
	Destination string `json:"destination" xml:"destination"`
}

