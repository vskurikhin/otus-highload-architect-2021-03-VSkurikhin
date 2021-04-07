package model

import "fmt"

type Signin struct {
	Username  string
	Password  string
	Name      string
	Surname   string
	Age       string
	Sex       string
	City      string
	Interests string
}

func (s *Signin) String() string {
	return fmt.Sprintf(
		`Signin{Username: %s, Password: ***HIDEN***, Name: %s, Surname: %s, Age: %s, Sex: %s, Interests: %s}`,
		s.Username, s.Name, s.Surname, s.Age, s.Sex, s.Interests,
	)
}
