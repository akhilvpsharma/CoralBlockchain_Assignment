package main
type WebForm struct {
	UserName	string	`json:"userName"`
	Password	string	`json:"password"`
	EmailID		string	`json:"emailId"`
	PhoneNumber	string	`json:"phoneNumber"`
	DateTime string `json:"dateTime"`
}