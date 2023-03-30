package jenkins

type Jenkins struct {
	url string	`json:"url"`
	user string	`json:"user"`
	password string	`json:"password"`
}


// func (jenkins *Jenkins)Init() error {

// }