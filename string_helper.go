package frasgohelper

import (
	"strings"
)

func ReplaceBlankNull(str string) string {
	if str == "" {
		str = "-"
	}

	return str
}

func AnonymousMobilePhone(mobilePhone string) string {
	mobilePhone = strings.Repeat("*", len(mobilePhone)-3) + mobilePhone[len(mobilePhone)-3:]

	return mobilePhone
}

func AnonymousEmail(email string) string {
	string_slice := strings.Split(email, "@")

	email = strings.Repeat("*", len(string_slice[0])-3) + string_slice[0][len(string_slice[0])-3:] + "@" + string_slice[1]

	return email
}
