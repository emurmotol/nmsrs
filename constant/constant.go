package constant

type ctxKey string

const (
	UserCtxKey ctxKey = "user"
	// RegistrantCtxKey ctxKey = "registrant"
	AuthUserCtxKey  ctxKey = "authUser"
	TokenAuthCtxKey ctxKey = "tokenAuth"
)
