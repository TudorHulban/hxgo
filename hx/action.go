package hx

import (
	"github.com/TudorHulban/hxgo/helpers"
)

// TODO: move method to HXMethod.
type HXMethod string

type Action struct {
	Method                   string // if not provided will switch to new tab.
	Endpoint                 string
	CSSIDValidationDisable   string
	CSSIDValidationPasswords [3]string

	Swaps   []string // includes CSS ID sanitization
	Require []string // includes CSS ID sanitization
	Sends   []string // includes CSS ID sanitization

	Enable  []string // includes CSS ID sanitization
	Disable []string // includes CSS ID sanitization

	Show []string
	Hide []string

	OnChangeEnable []string // includes CSS ID sanitization

	LengthMax uint16
	LengthMin uint16
}

func (a Action) HavePasswordsValues() bool {
	return len(a.CSSIDValidationPasswords[0]) > 0 && len(a.CSSIDValidationPasswords[1]) > 0 && len(a.CSSIDValidationPasswords[2]) > 0
}

func (a Action) PasswordsIDsSanitized() string {
	return helpers.SanitizeCSSId(a.CSSIDValidationPasswords[0]) +
		`,` +
		helpers.SanitizeCSSId(a.CSSIDValidationPasswords[1]) +
		`,` +
		helpers.SanitizeCSSId(a.CSSIDValidationPasswords[2])
}
