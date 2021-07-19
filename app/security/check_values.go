package security

import (
	"errors"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"regexp"
)

var numbers = regexp.MustCompile(`^[0-9]+$`)

var words = regexp.MustCompile(`^[ !#*+,\-./0-9:;A-Z\\_a-z~А-Яа-я]+$`)

var lines = regexp.MustCompile(`^[\t\n !#*+,\-./0-9:;A-Z\\_a-z~А-Яа-я]+$`)

func CheckValueMayBeEmpty(name, value string) error {
	if "" == value {
		return nil
	}
	return CheckValue(name, value)
}

func CheckValue(name, value string) error {
	if words.MatchString(value) {
		return nil
	}
	return errors.New(
		" Error for the field " + name + " in a value: `" + value +
			"`. For value allowed only characters (A-Z a-z А-Я а-я)," +
			" the space and numbers")
}

func CheckNumericValue(name, value string) error {
	if numbers.MatchString(value) {
		return nil
	}
	return errors.New(
		" Error for the field " + name + " in a value: `" + value +
			"`. For value allowed only characters (A-Z a-z А-Я а-я)," +
			" the space and numbers")
}

func CheckLines(name, line string) error {
	if "" == line || lines.MatchString(line) {
		return nil
	}
	return errors.New(" Error for the field " + name + " in a value: `" + line +
		"`. For value allowed only characters (A-Z a-z А-Я а-я)," +
		" the space and numbers")
}

func CheckSignIn(signIn *domain.Signin) error {

	err := CheckValue("Username", signIn.Username)

	if err != nil {
		return err
	}
	err = CheckValue("Password", signIn.Password)

	if err != nil {
		return err
	}
	err = CheckValueMayBeEmpty("Firstname", signIn.Name)

	if err != nil {
		return err
	}
	err = CheckValueMayBeEmpty("Surname", signIn.Surname)

	if err != nil {
		return err
	}
	err = CheckNumericValue("Age", signIn.Age)

	if err != nil {
		return err
	}
	err = CheckValue("City", signIn.City)

	if err != nil {
		return err
	}
	err = CheckLines("Interests", signIn.Interests)

	if err != nil {
		return err
	}
	return nil
}
