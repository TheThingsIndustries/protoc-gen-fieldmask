// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package otherpackage

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// ValidateFields checks the field values on Embed with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Embed) ValidateFields(paths ...string) error {
	if len(paths) > 0 {
		return fmt.Errorf("message Embed has no fields, but paths %s were specified", paths)
	}
	return nil
}

// EmbedValidationError is the validation error returned by
// Embed.ValidateFields if the designated constraints aren't met.
type EmbedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmbedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmbedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmbedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmbedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmbedValidationError) ErrorName() string { return "EmbedValidationError" }

// Error satisfies the builtin error interface
func (e EmbedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmbed.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmbedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmbedValidationError{}
