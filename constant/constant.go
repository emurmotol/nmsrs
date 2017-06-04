package constant

type ctxKey string

const (
	UserCtxKey      ctxKey = "user"
	AuthUserCtxKey  ctxKey = "authUser"
	TokenAuthCtxKey ctxKey = "tokenAuth"
)
