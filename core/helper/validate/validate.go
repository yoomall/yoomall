package validate

import (
	"regexp"
	"time"
)

type Validate interface {
	GetField() string
	IsValid(val interface{}) bool
	GetMessage() string
}

type DefValidate struct {
	Field    string
	Required bool
	Msg      string
}

var _ Validate = (*DefValidate)(nil)

func (s *DefValidate) IsValid(val interface{}) bool {
	return true
}

func (s *DefValidate) GetField() string {
	return s.Field
}

func (s *DefValidate) GetMessage() string {
	return s.Msg
}

type StringValidate struct {
	*DefValidate
	Field    string
	Required bool
	Msg      string
	Min      int
	Max      int
	Regx     *regexp.Regexp
}

var _ Validate = (*StringValidate)(nil)

func (s *StringValidate) IsValid(val interface{}) bool {
	str, ok := val.(string)
	if !ok {
		return false
	}
	if s.Required && len(str) == 0 {
		return false
	}

	if s.Min > 0 && len(str) < s.Min {
		return false
	}
	if s.Max > 0 && len(str) > s.Max {
		return false
	}
	if s.Regx != nil && !s.Regx.MatchString(str) {
		return false
	}
	return true

}

type NumberValidate struct {
	*DefValidate
	Field    string
	Required bool
	Msg      string
	Min      int
	Max      int
}

var _ Validate = (*NumberValidate)(nil)

func (s *NumberValidate) IsValid(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		if s.Min > 0 && val.(int) < s.Min {
			return false
		}
		if s.Max > 0 && val.(int) > s.Max {
			return false
		}
		return true
	default:
		return false
	}
}

type EmailValidate struct {
	*DefValidate
	Field    string
	Required bool
	Msg      string
}

var _ Validate = (*EmailValidate)(nil)

func (s *EmailValidate) IsValid(val interface{}) bool {
	str, ok := val.(string)
	if !ok {
		return false
	}
	if s.Required && len(str) == 0 {
		return false
	}
	regx := "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"

	return regexp.MustCompile(regx).MatchString(str)
}

type DateValidate struct {
	*DefValidate
	Field    string
	Required bool
	Msg      string
}

var _ Validate = (*DateValidate)(nil)

func (s *DateValidate) IsValid(val interface{}) bool {
	t, ok := val.(time.Time)
	if !ok {
		return false
	}
	if s.Required && t.IsZero() {
		return false
	}
	return true
}

type Validator struct {
	Validates []Validate
}

func (v *Validator) AddValidate(validate Validate) {
	v.Validates = append(v.Validates, validate)
}

func (v *Validator) Validate(data map[string]interface{}) (bool, string) {
	var valid bool
	for _, validate := range v.Validates {
		val := data[validate.GetField()]
		valid = validate.IsValid(val)
		if !valid {
			return valid, validate.GetMessage()
		}
	}
	return valid, ""
}
