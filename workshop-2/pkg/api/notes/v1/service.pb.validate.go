// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/notes/v1/service.proto

package notes

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on UpdateNoteByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateNoteByIDResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNoteByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateNoteByIDResponseMultiError, or nil if none found.
func (m *UpdateNoteByIDResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNoteByIDResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateNoteByIDResponseMultiError(errors)
	}

	return nil
}

// UpdateNoteByIDResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateNoteByIDResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateNoteByIDResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNoteByIDResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNoteByIDResponseMultiError) AllErrors() []error { return m }

// UpdateNoteByIDResponseValidationError is the validation error returned by
// UpdateNoteByIDResponse.Validate if the designated constraints aren't met.
type UpdateNoteByIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNoteByIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNoteByIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNoteByIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNoteByIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNoteByIDResponseValidationError) ErrorName() string {
	return "UpdateNoteByIDResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNoteByIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNoteByIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNoteByIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNoteByIDResponseValidationError{}

// Validate checks the field values on UpdateNoteByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateNoteByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNoteByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateNoteByIDRequestMultiError, or nil if none found.
func (m *UpdateNoteByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNoteByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NoteId

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteByIDRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteByIDRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteByIDRequestValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateNoteByIDRequestMultiError(errors)
	}

	return nil
}

// UpdateNoteByIDRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateNoteByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateNoteByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNoteByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNoteByIDRequestMultiError) AllErrors() []error { return m }

// UpdateNoteByIDRequestValidationError is the validation error returned by
// UpdateNoteByIDRequest.Validate if the designated constraints aren't met.
type UpdateNoteByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNoteByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNoteByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNoteByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNoteByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNoteByIDRequestValidationError) ErrorName() string {
	return "UpdateNoteByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNoteByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNoteByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNoteByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNoteByIDRequestValidationError{}

// Validate checks the field values on DeleteNoteByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteNoteByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteNoteByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteNoteByIDRequestMultiError, or nil if none found.
func (m *DeleteNoteByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNoteByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NoteId

	if len(errors) > 0 {
		return DeleteNoteByIDRequestMultiError(errors)
	}

	return nil
}

// DeleteNoteByIDRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteNoteByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteNoteByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNoteByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNoteByIDRequestMultiError) AllErrors() []error { return m }

// DeleteNoteByIDRequestValidationError is the validation error returned by
// DeleteNoteByIDRequest.Validate if the designated constraints aren't met.
type DeleteNoteByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNoteByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNoteByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNoteByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNoteByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNoteByIDRequestValidationError) ErrorName() string {
	return "DeleteNoteByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNoteByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNoteByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNoteByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNoteByIDRequestValidationError{}

// Validate checks the field values on DeleteNoteByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteNoteByIDResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteNoteByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteNoteByIDResponseMultiError, or nil if none found.
func (m *DeleteNoteByIDResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNoteByIDResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteNoteByIDResponseMultiError(errors)
	}

	return nil
}

// DeleteNoteByIDResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteNoteByIDResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteNoteByIDResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNoteByIDResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNoteByIDResponseMultiError) AllErrors() []error { return m }

// DeleteNoteByIDResponseValidationError is the validation error returned by
// DeleteNoteByIDResponse.Validate if the designated constraints aren't met.
type DeleteNoteByIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNoteByIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNoteByIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNoteByIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNoteByIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNoteByIDResponseValidationError) ErrorName() string {
	return "DeleteNoteByIDResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNoteByIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNoteByIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNoteByIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNoteByIDResponseValidationError{}

// Validate checks the field values on GetNoteByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetNoteByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNoteByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNoteByIDRequestMultiError, or nil if none found.
func (m *GetNoteByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNoteByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NoteId

	// no validation rules for SomeQueryParam

	if len(errors) > 0 {
		return GetNoteByIDRequestMultiError(errors)
	}

	return nil
}

// GetNoteByIDRequestMultiError is an error wrapping multiple validation errors
// returned by GetNoteByIDRequest.ValidateAll() if the designated constraints
// aren't met.
type GetNoteByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNoteByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNoteByIDRequestMultiError) AllErrors() []error { return m }

// GetNoteByIDRequestValidationError is the validation error returned by
// GetNoteByIDRequest.Validate if the designated constraints aren't met.
type GetNoteByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNoteByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNoteByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNoteByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNoteByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNoteByIDRequestValidationError) ErrorName() string {
	return "GetNoteByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetNoteByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNoteByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNoteByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNoteByIDRequestValidationError{}

// Validate checks the field values on GetNoteByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetNoteByIDResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNoteByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNoteByIDResponseMultiError, or nil if none found.
func (m *GetNoteByIDResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNoteByIDResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetNoteByIDResponseMultiError(errors)
	}

	return nil
}

// GetNoteByIDResponseMultiError is an error wrapping multiple validation
// errors returned by GetNoteByIDResponse.ValidateAll() if the designated
// constraints aren't met.
type GetNoteByIDResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNoteByIDResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNoteByIDResponseMultiError) AllErrors() []error { return m }

// GetNoteByIDResponseValidationError is the validation error returned by
// GetNoteByIDResponse.Validate if the designated constraints aren't met.
type GetNoteByIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNoteByIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNoteByIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNoteByIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNoteByIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNoteByIDResponseValidationError) ErrorName() string {
	return "GetNoteByIDResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetNoteByIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNoteByIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNoteByIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNoteByIDResponseValidationError{}

// Validate checks the field values on NoteInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NoteInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NoteInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NoteInfoMultiError, or nil
// if none found.
func (m *NoteInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *NoteInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTitle()); l < 3 || l > 140 {
		err := NoteInfoValidationError{
			field:  "Title",
			reason: "value length must be between 3 and 140 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_NoteInfo_Title_Pattern.MatchString(m.GetTitle()) {
		err := NoteInfoValidationError{
			field:  "Title",
			reason: "value does not match regex pattern \"^[a-zA-Z]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 10 || l > 1000 {
		err := NoteInfoValidationError{
			field:  "Content",
			reason: "value length must be between 10 and 1000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return NoteInfoMultiError(errors)
	}

	return nil
}

// NoteInfoMultiError is an error wrapping multiple validation errors returned
// by NoteInfo.ValidateAll() if the designated constraints aren't met.
type NoteInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoteInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoteInfoMultiError) AllErrors() []error { return m }

// NoteInfoValidationError is the validation error returned by
// NoteInfo.Validate if the designated constraints aren't met.
type NoteInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoteInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoteInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoteInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoteInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoteInfoValidationError) ErrorName() string { return "NoteInfoValidationError" }

// Error satisfies the builtin error interface
func (e NoteInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNoteInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoteInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoteInfoValidationError{}

var _NoteInfo_Title_Pattern = regexp.MustCompile("^[a-zA-Z]+$")

// Validate checks the field values on Note with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Note) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Note with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NoteMultiError, or nil if none found.
func (m *Note) ValidateAll() error {
	return m.validate(true)
}

func (m *Note) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NoteId

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NoteMultiError(errors)
	}

	return nil
}

// NoteMultiError is an error wrapping multiple validation errors returned by
// Note.ValidateAll() if the designated constraints aren't met.
type NoteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoteMultiError) AllErrors() []error { return m }

// NoteValidationError is the validation error returned by Note.Validate if the
// designated constraints aren't met.
type NoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoteValidationError) ErrorName() string { return "NoteValidationError" }

// Error satisfies the builtin error interface
func (e NoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoteValidationError{}

// Validate checks the field values on SaveNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SaveNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveNoteRequestMultiError, or nil if none found.
func (m *SaveNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetInfo() == nil {
		err := SaveNoteRequestValidationError{
			field:  "Info",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveNoteRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveNoteRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveNoteRequestValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SaveNoteRequestMultiError(errors)
	}

	return nil
}

// SaveNoteRequestMultiError is an error wrapping multiple validation errors
// returned by SaveNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type SaveNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveNoteRequestMultiError) AllErrors() []error { return m }

// SaveNoteRequestValidationError is the validation error returned by
// SaveNoteRequest.Validate if the designated constraints aren't met.
type SaveNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveNoteRequestValidationError) ErrorName() string { return "SaveNoteRequestValidationError" }

// Error satisfies the builtin error interface
func (e SaveNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveNoteRequestValidationError{}

// Validate checks the field values on SaveNoteResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SaveNoteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveNoteResponseMultiError, or nil if none found.
func (m *SaveNoteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveNoteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NoteId

	if len(errors) > 0 {
		return SaveNoteResponseMultiError(errors)
	}

	return nil
}

// SaveNoteResponseMultiError is an error wrapping multiple validation errors
// returned by SaveNoteResponse.ValidateAll() if the designated constraints
// aren't met.
type SaveNoteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveNoteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveNoteResponseMultiError) AllErrors() []error { return m }

// SaveNoteResponseValidationError is the validation error returned by
// SaveNoteResponse.Validate if the designated constraints aren't met.
type SaveNoteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveNoteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveNoteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveNoteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveNoteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveNoteResponseValidationError) ErrorName() string { return "SaveNoteResponseValidationError" }

// Error satisfies the builtin error interface
func (e SaveNoteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveNoteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveNoteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveNoteResponseValidationError{}

// Validate checks the field values on ListNotesResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListNotesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListNotesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListNotesResponseMultiError, or nil if none found.
func (m *ListNotesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListNotesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNotes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListNotesResponseValidationError{
						field:  fmt.Sprintf("Notes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListNotesResponseValidationError{
						field:  fmt.Sprintf("Notes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListNotesResponseValidationError{
					field:  fmt.Sprintf("Notes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListNotesResponseMultiError(errors)
	}

	return nil
}

// ListNotesResponseMultiError is an error wrapping multiple validation errors
// returned by ListNotesResponse.ValidateAll() if the designated constraints
// aren't met.
type ListNotesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListNotesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListNotesResponseMultiError) AllErrors() []error { return m }

// ListNotesResponseValidationError is the validation error returned by
// ListNotesResponse.Validate if the designated constraints aren't met.
type ListNotesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNotesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNotesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNotesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNotesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNotesResponseValidationError) ErrorName() string {
	return "ListNotesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListNotesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNotesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNotesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNotesResponseValidationError{}
