// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chaos/experimentation/v1/fields.proto

package experimentationv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _fields_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Fields with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Fields) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetFields() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FieldsValidationError{
					field:  fmt.Sprintf("Fields[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FieldsValidationError is the validation error returned by Fields.Validate if
// the designated constraints aren't met.
type FieldsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldsValidationError) ErrorName() string { return "FieldsValidationError" }

// Error satisfies the builtin error interface
func (e FieldsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFields.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldsValidationError{}

// Validate checks the field values on TextField with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TextField) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Label

	// no validation rules for Value

	return nil
}

// TextFieldValidationError is the validation error returned by
// TextField.Validate if the designated constraints aren't met.
type TextFieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TextFieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TextFieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TextFieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TextFieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TextFieldValidationError) ErrorName() string { return "TextFieldValidationError" }

// Error satisfies the builtin error interface
func (e TextFieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTextField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TextFieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TextFieldValidationError{}