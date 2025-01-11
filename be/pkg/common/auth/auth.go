package auth

import "context"

type (
	ctxKeyT struct{}

	AuthInfo struct {
		User  string
		Roles []string
	}
)

var ctxKey ctxKeyT = ctxKeyT{}

func Inject(ctx context.Context, auth AuthInfo) context.Context {
	return context.WithValue(ctx, ctxKey, auth)
}

func Extract(ctx context.Context) AuthInfo {
	val := ctx.Value(ctxKey)
	authInfo, ok := val.(AuthInfo)
	if ok {
		return authInfo
	}
	return AuthInfo{}
}

func HasAccess(a AuthInfo, path string) bool {
	if a.User != "" {
		return true
	}
	return false
}
