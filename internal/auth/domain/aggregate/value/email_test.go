package value

import (
	"errors"
	"testing"

	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEmail_NewEmailFromInput(t *testing.T) {
	Convey("Given an email input", t, func() {
		Convey("When the email is valid", func() {
			Convey("Then it should create Email successfully", func() {
				email, err := NewEmailBuilder("test@example.com").Build()

				So(err, ShouldBeNil)
				So(email, ShouldNotBeNil)
				So(email.Value(), ShouldEqual, "test@example.com")
				So(email.Type(), ShouldEqual, constants.IdentityTypeEmail)
			})

			Convey("And it should handle special characters", func() {
				testCases := []struct {
					input    string
					expected string
				}{
					{"test+tag@example.com", "test+tag@example.com"},
					{"first.last@example.com", "first.last@example.com"},
					{"user-name@example.com", "user-name@example.com"},
				}

				for _, tc := range testCases {
					Convey("For input with "+tc.input, func() {
						email, err := NewEmailBuilder(tc.input).Build()
						So(err, ShouldBeNil)
						So(email, ShouldNotBeNil)
						So(email.Value(), ShouldEqual, tc.expected)
					})
				}
			})
		})

		Convey("When the email is invalid", func() {
			Convey("Then it should return an error", func() {
				testCases := []struct {
					input       string
					errContains error
				}{
					{"not-an-email", errno.EmailFormatError},
					{"testexample.com", errno.EmailFormatError},
					{"", errno.EmailFormatError},
					{"@example.com", errno.EmailFormatError},
					{"test@", errno.EmailFormatError},
				}

				for _, tc := range testCases {
					Convey("For invalid input: "+tc.input, func() {
						email, err := NewEmailBuilder(tc.input).Build()

						So(err, ShouldNotBeNil)
						So(email, ShouldBeNil)
						So(errors.Is(err, tc.errContains), ShouldBeTrue)
					})
				}
			})
		})
	})
}
