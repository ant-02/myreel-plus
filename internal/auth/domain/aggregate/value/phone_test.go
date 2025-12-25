package value

import (
	"errors"
	"testing"

	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPhone_NewPhoneFromInput(t *testing.T) {
	Convey("Given an phone input", t, func() {
		Convey("When the phone is valid", func() {
			Convey("Then it should create Phone successfully", func() {
				testCases := []struct {
					input    string
					expected string
				}{
					{"13912345678", "13912345678"},
					{"13012345678", "13012345678"},
					{"13312345678", "13312345678"},
					{"17001234567", "17001234567"},
				}

				for _, tc := range testCases {
					Convey("For input with "+tc.input, func() {
						phone, err := NewPhoneBuilder(tc.input).Build()
						So(err, ShouldBeNil)
						So(phone, ShouldNotBeNil)
						So(phone.Value(), ShouldEqual, tc.expected)
						So(phone.Type(), ShouldEqual, constants.IdentityTypePhone)
					})
				}
			})
		})

		Convey("When the phone is invalid", func() {
			Convey("Then it should return an error", func() {
				testCases := []struct {
					input       string
					errContains error
				}{
					{"13800138", errno.PhoneFormatError},
					{"13800138000123", errno.PhoneFormatError},
					{"", errno.PhoneFormatError},
					{"一三八零零一三八零零零", errno.PhoneFormatError},
					{"13800_38000", errno.PhoneFormatError},
				}

				for _, tc := range testCases {
					Convey("For invalid input: "+tc.input, func() {
						phone, err := NewPhoneBuilder(tc.input).Build()

						So(err, ShouldNotBeNil)
						So(phone, ShouldBeNil)
						So(errors.Is(err, tc.errContains), ShouldBeTrue)
					})
				}
			})
		})
	})
}
