package constant

type ctxKey string

const (
	UserCtxKey ctxKey = "user"
	// RegistrantCtxKey ctxKey = "registrant"
	AuthCtxKey      ctxKey = "auth"
	TokenAuthCtxKey ctxKey = "tokenAuth"
)
