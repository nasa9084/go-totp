package totp

import (
	"time"

	"github.com/nasa9084/go-hotp"
)

// Generator generates Time-based One-Time Password
type Generator struct {
	TimeStep  uint64    // X in RFC6238
	StartTime int64  // T0 in RFC6238, default 0 is OK
	Secret    string // shared secret for generate hotp
	Digit     int
}

// Generate OTP
func (g *Generator) Generate() int64 {
	if g.TimeStep == 0 {
		g.TimeStep = 30
	}
	now := time.Now().Unix()
	t := (now - g.StartTime) / int64(g.TimeStep)
	h := hotp.Generator{
		Secret:  g.Secret,
		Digit:   g.Digit,
		Counter: uint64(t),
	}
	return h.Generate()
}
