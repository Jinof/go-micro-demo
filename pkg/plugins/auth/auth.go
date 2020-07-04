package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"

	"github.com/micro/cli/v2"
	"github.com/micro/micro/v2/plugin"
)

type auth struct {
	prefix []string
	secret string
}

func (a *auth) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "path_prefix",
			Usage:   "Comma separated list of path prefixes to strip before continuing with request e.g /api,/foo,/bar",
			EnvVars: []string{"PATH_PREFIX"},
		},
		&cli.StringFlag{
			Name:    "secret",
			Usage:   "Token secret e.g `mySecret`",
			EnvVars: []string{"SECRET"},
		},
	}
}

func (a *auth) Commands() []*cli.Command {
	return nil
}

func (a *auth) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return a.PrefixHandler(a.LoginHandler(h))
	}
}

func (a *auth) PrefixHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// strip auth if we have a match
		for _, prefix := range a.prefix {
			if strings.HasPrefix(r.URL.Path, prefix) {
				r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
				break
			}
		}

		// serve request
		h.ServeHTTP(w, r)
	}
}

// LoginHandler checkout token and get username.
func (a *auth) LoginHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("authorization")
		username, err := a.ParesToken(token)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			encoder := json.NewEncoder(w)
			if err := encoder.Encode(map[string]interface{}{
				"status":  0,
				"message": "403 Forbidden",
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		fmt.Println(username)
		r.Header.Set("username", strconv.Itoa(username))
		r.Header.Set("authorization", "Bearer "+token)
		h.ServeHTTP(w, r)
	}
}

// ParesToken xh是学号\学工号
func (a *auth) ParesToken(tokenString string) (int, error) {
	secretKey := []byte(a.secret)
	kv := strings.SplitAfter(tokenString, " ")
	if len(kv) < 2 {
		return 0, errors.New("403 Forbidden")
	}
	token, err := jwt.Parse(kv[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method %v ", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		xh := claims["xh"].(string)
		userId, err := strconv.Atoi(xh)
		if err != nil {
			return 0, err
		}
		return userId, nil
	} else {
		return 0, err
	}
}

func (a *auth) Init(ctx *cli.Context) error {
	if prefix := ctx.String("path_prefix"); len(prefix) > 0 {
		a.prefix = append(a.prefix, strings.Split(prefix, ",")...)
	}
	secret := ctx.String("secret")
	a.secret = secret
	return nil
}

func (a *auth) String() string {
	return "auth"
}

func NewPlugin() plugin.Plugin {
	return &auth{}
}
