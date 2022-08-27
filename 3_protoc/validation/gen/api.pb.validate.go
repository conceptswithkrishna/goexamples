// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api.proto

package validation

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

// Validate checks the field values on SearchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// SearchRequestMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *SearchRequest) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Query

	// no validation rules for PageNumber

	if m.GetResultPerPage() > 10 {
		err := SearchRequestValidationError{
			field:  "ResultPerPage",
			reason: "value must be less than or equal to 10",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SearchRequestMultiError(errors)
	}
	return nil
}

// SearchRequestMultiError is an error wrapping multiple validation errors
// returned by SearchRequest.Validate(true) if the designated constraints
// aren't met.
type SearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchRequestMultiError) AllErrors() []error { return m }

// SearchRequestValidationError is the validation error returned by
// SearchRequest.Validate if the designated constraints aren't met.
type SearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchRequestValidationError) ErrorName() string { return "SearchRequestValidationError" }

// Error satisfies the builtin error interface
func (e SearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchRequestValidationError{}
