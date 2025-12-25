package value

import (
	"errors"
	"testing"

	"github.com/ant-02/myreel-plus/pkg/errno"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPassword_NewPasswordFromInput(t *testing.T) {
	Convey("Given an password input", t, func() {
		Convey("When the password is valid", func() {
			Convey("Then it should create Password successfully", func() {
				testCases := []struct {
					input    string
					expected string
				}{
					{"123456", "123456"},
					{"abc123456", "abc123456"},
					{"abc123456..@$", "abc123456..@$"},
				}

				for _, tc := range testCases {
					Convey("For input with "+tc.input, func() {
						password, err := NewPasswordBuilder(tc.input).WithEncrypt().Build()
						So(err, ShouldBeNil)
						So(password, ShouldNotBeNil)
						So(password.Compare(tc.expected), ShouldBeTrue)
					})
				}
			})
		})

		Convey("When the password is invalid", func() {
			Convey("Then it should return an error", func() {
				testCases := []struct {
					input       string
					errContains error
				}{
					{"1234", errno.PasswordFormatError},
					{"38475729485720938475928374509283745092837450928374509283745092837450928374509283", errno.PasswordFormatError},
				}

				for _, tc := range testCases {
					Convey("For invalid input: "+tc.input, func() {
						password, err := NewPasswordBuilder(tc.input).Build()
						So(err, ShouldNotBeNil)
						So(password, ShouldBeNil)
						So(errors.Is(err, tc.errContains), ShouldBeTrue)
					})
				}
			})

			Convey("Then it should be false", func() {
				testCases := []struct {
					input    string
					expected string
				}{
					{"12345", "56789"},
					{"abc12345", "54321cba"},
				}

				for _, tc := range testCases {
					Convey("For invalid input: "+tc.input, func() {
						password, err := NewPasswordBuilder(tc.input).WithEncrypt().Build()
						So(err, ShouldBeNil)
						So(password, ShouldNotBeNil)
						So(password.Compare(tc.expected), ShouldBeFalse)
					})
				}
			})
		})
	})
}
