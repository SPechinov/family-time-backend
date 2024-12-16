package validate

import (
	"testing"
)

func Test_email(t *testing.T) {
	testTable := []struct {
		name   string
		email  any
		expect error
	}{
		{
			name:   "Format: qwe@qwe.qwe",
			email:  "qwe@qwe.qwe",
			expect: nil,
		},
		{
			name:   "Format: qweqwe.qwe@qwe.qwe",
			email:  "qweqwe.qwe@qwe.qwe",
			expect: nil,
		},
		{
			name:   "Format: qweqwe_qwe@qwe.qwe",
			email:  "qweqwe_qwe@qwe.qwe",
			expect: nil,
		},
		{
			name:   "Format: qweqwe-qwe@qwe.qwe",
			email:  "qweqwe-qwe@qwe.qwe",
			expect: nil,
		},
		{
			name:   "Format: qweqwe+qwe@qwe.qwe",
			email:  "qweqwe+qwe@qwe.qwe",
			expect: nil,
		},
		{
			name:   "Format: qweqwe%qwe@qwe.qwe",
			email:  "qweqwe%qwe@qwe.qwe",
			expect: nil,
		},
		{
			name:   "Format: qweqwe@qwe.qwe.qwe",
			email:  "qweqwe@qwe.qwe.qwe",
			expect: nil,
		},
		{
			name:   "Format: qweqwe@qwe-qwe.qwe",
			email:  "qweqwe@qwe-qwe.qwe",
			expect: nil,
		},
		{
			name:   "Format: qweqwe@123qwe.qwe",
			email:  "qweqwe@q123qwe.qwe",
			expect: nil,
		},
		{
			name:   "Format: qwe@qwe.qw",
			email:  "qwe@qwe.qw",
			expect: nil,
		},
		{
			name:   "Format: qwe@.qwe.qw",
			email:  "qwe@.qwe.qw",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe@-qwe.qw",
			email:  "qwe@-qwe.qw",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe.@qwe.qw",
			email:  "qwe.@qwe.qw",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe_@qwe.qw",
			email:  "qwe_@qwe.qw",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe%@qwe.qw",
			email:  "qwe%@qwe.qw",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe+@qwe.qw",
			email:  "qwe+@qwe.qw",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe-@qwe.qw",
			email:  "qwe-@qwe.qw",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe@qwe.q",
			email:  "qwe@qwe.q",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: @.",
			email:  "@.",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: @.qwe",
			email:  "@.qwe",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe@qwe.",
			email:  "qwe@qwe.",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe@qwe",
			email:  "qwe@qwe",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qweqwe.qwe",
			email:  "qweqwe.qwe",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: @qwe",
			email:  "@qwe",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe@",
			email:  "qwe@",
			expect: EmailErrInvalid,
		},
		{
			name:   "Format: qwe",
			email:  "qwe",
			expect: EmailErrInvalid,
		},
		{
			name:   "Empty string",
			email:  "",
			expect: EmailErrEmptyString,
		},
		{
			name:   "Number",
			email:  1,
			expect: EmailErrMustBeString,
		},
		{
			name:   "Bool: true",
			email:  true,
			expect: EmailErrMustBeString,
		},
		{
			name:   "Bool: false",
			email:  false,
			expect: EmailErrMustBeString,
		},
		{
			name:   "Byte slice",
			email:  []byte("test byte"),
			expect: EmailErrMustBeString,
		},
		{
			name:   "Struct",
			email:  struct{}{},
			expect: EmailErrMustBeString,
		},
		{
			name:   "Channel",
			email:  make(chan int),
			expect: EmailErrMustBeString,
		},
		{
			name:   "Nil",
			email:  nil,
			expect: EmailErrMustBeString,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := Email(testCase.email)

			if err != nil && testCase.expect != nil && err.Error() != testCase.expect.Error() {
				t.Errorf("Expected error: %v, but got: %v", testCase.expect, err)
			}

			if err == nil && testCase.expect != nil {
				t.Errorf("Expected error: %v, but got: nil", testCase.expect)
			}

			if err != nil && testCase.expect == nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
		})
	}
}
