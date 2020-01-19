package entry

import "testing"

func TestVerify(t *testing.T) {
	a := &AuthenticationModule{}
	token := a.GenerateToken()
	result := a.VerifyToken(token)
	if result {
		println("OK")
	}
}
