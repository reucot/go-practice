package err_practice

import (
	"errors"
	"strings"
)

// 1) вставьте определение типа для []error
// 2) определите метод Error для вашего типа, который будет выводить
//    все ошибки слайса
// 3) реализуйте функцию MyCheck
//
// ...

type SliceError []error

func (e SliceError) Error() string {
	var out string
	for _, err := range e {
		out += err.Error() + ";"
	}
	return strings.TrimRight(out, `;`)
}

func MyCheck(input string) error {
	var (
		err      SliceError
		spaces   int
		hasDigit bool
	)
	if len([]rune(input)) >= 20 {
		err = SliceError{errors.New(`line is too long`)}
	}
	for _, ch := range input {
		if ch == ' ' {
			spaces++
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}
	if hasDigit {
		err = append(err, errors.New(`found numbers`))
	}
	if spaces != 2 {
		err = append(err, errors.New(`no two spaces`))
	}
	if len(err) == 0 {
		return nil
	}
	return err
}
