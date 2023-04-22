package template_method

import (
	"fmt"
	"testing"
)

func TestOtp(t *testing.T) {
	smsOTP := &sms{}
	o := otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println()

	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}
