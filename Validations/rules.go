package Validations

import (
	"regexp"

	"github.com/bykovme/gotrans"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

/*
	Set your own validation rules here to accessable from your controller
*/

func RequiredRule() validation.Rule {
	return validation.Required.Error(gotrans.T("required"))
}

func MinMaxRule() validation.Rule {
	return validation.Length(3, 50).Error(gotrans.T("min_max"))
}

func Email() validation.Rule {
	return is.Email.Error(gotrans.T("email"))
}

func NumericValue() validation.Rule {
	return validation.Match(regexp.MustCompile("^[0-9]{5}$")).Error(gotrans.T("numeric"))
}
