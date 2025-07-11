package models

type Farmer struct {
	NationalID    int    `json:"national_id"` // 14-digit - must be unique - primary key
	FullName      string `json:"full_name"`   // Must have 4+ parts
	Gender        string `json:"gender"`
	PersonalImage string `json:"personal_image"` // file path
	IDImageFront  string `json:"id_image_front"` // file path
	IDImageBack   string `json:"id_image_back"`  // file path
	PhoneNumber   string `json:"phone_number"`
}
