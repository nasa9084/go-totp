package totp_test

import (
	"testing"
	"time"

	"github.com/nasa9084/go-totp"
)

var generator = totp.Generator{
	TimeStep:     30,
	Base32Secret: "12345678901234567890",
	Digit:        8,
}

func TestGenerate(t *testing.T) {
	candidates := []struct {
		time     int64
		expected int64
	}{
		{59, 94287082},
		{1111111109, 7081804},
		{1111111111, 14050471},
		{1234567890, 89005924},
		{2000000000, 69279037},
		{20000000000, 65353130},
	}
	for _, c := range candidates {
		generator.StartTime = time.Now().Unix() - c.time
		otp := generator.Generate()
		if c.expected != otp {
			t.Errorf("%d != %d", otp, c.expected)
			return
		}
	}
}
