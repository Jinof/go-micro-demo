package login

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

type login struct {
	secret string
}

func (l *login) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "secret",
			Usage:   "Token secret e.g `mySecret`",
			EnvVars: []string{"SECRET"},
		},
	}
}

func (l *login) Commands() []*cli.Command {
	return nil
}

func (l *login) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return l.Router(h)
	}
}

func (l *login) Router(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("login") != "" {
			l.LoginHandler(h)
			return
		}
		h.ServeHTTP(w, r)
	}
}

// LoginHandler checkout token and get username.
func (l *login) LoginHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("authorization")
		username, err := l.ParesToken(token)
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
func (l *login) ParesToken(tokenString string) (int, error) {
	secretKey := []byte(l.secret)
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

func (l *login) Init(ctx *cli.Context) error {
	secret := ctx.String("secret")
	l.secret = secret
	return nil
}

func (l *login) String() string {
	return "login"
}

func NewPlugin() plugin.Plugin {
	return &login{}
}
