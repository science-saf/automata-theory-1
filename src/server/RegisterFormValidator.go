package main

import "regexp"

const (
	UserValid = iota
	UserInvalidNickname
	UserInvalidEmail
	UserWeakPassword
	UserPasswordMismatch
)

type RegisterResult struct {
	errors []int32
}

type RegisterFormValidator struct {
}

type RegexpsContainer struct {
	nicknameRegexp  *regexp.Regexp
	emailRegexp     *regexp.Regexp
	passwordRegexp1 *regexp.Regexp
	passwordRegexp2 *regexp.Regexp
}

var regexpsContainer RegexpsContainer

func (self *RegisterFormValidator) Check(user *SiteUser) RegisterResult {
	self.InitRegexpsContainer()
	result := RegisterResult{
		errors: []int32{},
	}
	if !self.IsNicknameValid(user.nickname) {
		result.errors = append(result.errors, UserInvalidNickname)
	}
	if !self.IsEmailValid(user.email) {
		result.errors = append(result.errors, UserInvalidEmail)
	}
	if !self.IsPasswordValid(user.password1) {
		result.errors = append(result.errors, UserWeakPassword)
	}
	if user.password1 != user.password2 {
		result.errors = append(result.errors, UserPasswordMismatch)
	}
	return result
}

func (self *RegisterFormValidator) InitRegexpsContainer() {
	regexpsContainer = RegexpsContainer{
		nicknameRegexp:  regexp.MustCompile("^[A-z0-9_]+$"),
		emailRegexp:     regexp.MustCompile("(?i)^[A-z0-9_]+@(?:gmail\\.com|yandex\\.ru|mail\\.ru)$"),
		passwordRegexp1: regexp.MustCompile("[A-z]+"),
		passwordRegexp2: regexp.MustCompile("\\d+"),
	}
}

func (self *RegisterFormValidator) IsNicknameValid(nickname string) bool {
	return regexpsContainer.nicknameRegexp.MatchString(nickname)
}

func (self *RegisterFormValidator) IsEmailValid(email string) bool {
	return regexpsContainer.emailRegexp.MatchString(email)
}

func (self *RegisterFormValidator) IsPasswordValid(password string) bool {
	result := regexpsContainer.passwordRegexp1.MatchString(password)
	result = result && regexpsContainer.passwordRegexp2.MatchString(password)
	return result
}
