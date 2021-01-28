package scm

import (
	"github.com/pkg/errors"
	"net/url"
	"path"
)

// UrlJoin joins the given url parts.
func UrlJoin(params ...string) (string, error) {
	if len(params) < 1 {
		return "", nil
	}
	u, err := url.Parse(params[0])
	if err != nil {
		return "", errors.Wrapf(err, "failed to parse URL parts: %v", params)
	}
	params[0] = u.Path
	if len(params) == 1 {
		return params[0], nil
	}
	u.Path = path.Join(params...)
	return u.String(), nil
}
