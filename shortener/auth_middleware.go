package shortener

import (
	"net/http"
)

type authRoundTripper struct {
	token string
	next  http.RoundTripper
}

func NewAuthRoundTripper(token string) *authRoundTripper {
	return &authRoundTripper{
		token: token,
		next:  http.DefaultTransport,
	}
}

func (t *authRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add(AUTH_HEADER, t.token)
	return t.next.RoundTrip(r)
}
