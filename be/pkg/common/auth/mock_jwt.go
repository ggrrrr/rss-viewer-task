package auth

import "log/slog"

type mockJwt struct {
}

var _ (Verifier) = (mockJwt)(mockJwt{})

// Verify implements Verifier.
func (m mockJwt) Verify(jwtPayload string) (AuthInfo, error) {
	slog.Warn("MockVerifier.Verify")
	return AuthInfo{
		User: "admin",
	}, nil
}

func NewMockVerifier() mockJwt {
	slog.Warn("NewMockVerifier")
	return mockJwt{}
}
