package security

import (
	"errors"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"regexp"
)

var words = regexp.MustCompile(`^[ !#*+,\-./0-9:;A-Z\\_a-z~А-Яа-я]+$`)

var lines = regexp.MustCompile(`^[\t\n !#*+,\-./0-9:;A-Z\\_a-z~А-Яа-я]+$`)

func CheckValue(value string) error {
	if words.MatchString(value) {
		return nil
	}
	return errors.New(" error in value: " + value)
}

func CheckLines(line string) error {
	if lines.MatchString(line) {
		return nil
	}
	return errors.New(" error in lines: " + line)
}

func CheckSignIn(signIn *domain.Signin) error {

	err := CheckValue(signIn.Username)

	if err != nil {
		return err
	}
	err = CheckValue(signIn.Name)

	if err != nil {
		return err
	}
	err = CheckValue(signIn.Surname)

	if err != nil {
		return err
	}
	err = CheckValue(signIn.Age)

	if err != nil {
		return err
	}
	err = CheckValue(signIn.Sex)

	if err != nil {
		return err
	}
	err = CheckValue(signIn.City)

	if err != nil {
		return err
	}
	err = CheckLines(signIn.Interests)

	if err != nil {
		return err
	}
	return nil
}
