package sms

import "errors"

var (
	ErrEmptyPhoneNumbers = errors.New("PhoneNumbers must be not empty")
	ErrLimitPhoneNumbers = errors.New("PhoneNumbers max size 1000")
	ErrEmptySignName     = errors.New("SignName must be not empty")
	ErrEmptyTemplateCode = errors.New("TemplateCode must be not empty")
)
