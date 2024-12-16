package validate

import (
	"testing"
)

func Test_phone(t *testing.T) {
	testTable := []struct {
		name   string
		phone  any
		expect error
	}{
		{
			name:   "Empty string",
			phone:  "",
			expect: PhoneErrEmptyString,
		},
		{
			name:   "Number",
			phone:  1,
			expect: PhoneErrMustBeString,
		},
		{
			name:   "Bool: true",
			phone:  true,
			expect: PhoneErrMustBeString,
		},
		{
			name:   "Bool: false",
			phone:  false,
			expect: PhoneErrMustBeString,
		},
		{
			name:   "Byte slice",
			phone:  []byte("test byte"),
			expect: PhoneErrMustBeString,
		},
		{
			name:   "Struct",
			phone:  struct{}{},
			expect: PhoneErrMustBeString,
		},
		{
			name:   "Channel",
			phone:  make(chan int),
			expect: PhoneErrMustBeString,
		},
		{
			name:   "Nil",
			phone:  nil,
			expect: PhoneErrMustBeString,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := Phone(testCase.phone)

			// Сравниваем ожидаемую ошибку с результатом
			if err != nil && testCase.expect != nil && err.Error() != testCase.expect.Error() {
				t.Errorf("Expected error: %v, but got: %v", testCase.expect, err)
			}

			// Если ошибка не ожидается
			if err == nil && testCase.expect != nil {
				t.Errorf("Expected error: %v, but got: nil", testCase.expect)
			}

			// Если ошибка ожидается
			if err != nil && testCase.expect == nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
		})
	}
}
