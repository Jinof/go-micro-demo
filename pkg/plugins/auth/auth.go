package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter/v2"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/micro/cli/v2"
	"github.com/micro/micro/v2/plugin"
)

type Auth struct {
	secret       string
	enforcer     *casbin.Enforcer
	useMysql     bool
	model        string
	adapter      string
	mysqlAdapter string
}

func (a *Auth) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "secret",
			Usage:   "Token secret e.g `mySecret`",
			EnvVars: []string{"SECRET"},
		},
		&cli.StringFlag{
			Name:  "casbin_model",
			Usage: "Casbin model config file",
			Value: "./conf/casbin_model.conf",
		},
		&cli.StringFlag{
			Name:  "casbin_file_adapter",
			Usage: "Casbin adapter e.g `path/to/casbin_policy.csv`",
			Value: "./conf/casbin_policy.csv",
		},
		&cli.StringFlag{
			Name:  "casbin_adapter_use_mysql",
			Usage: "Casbin adapter use mysql e.g `true/false`",
			Value: "true",
		},
		&cli.StringFlag{
			Name:  "casbin_adapter_mysql",
			Usage: "Casbin adapter mysql uri e.g `root:123456@tcp(127.0.0.1:3306)/`",
			Value: "mysql:123456@tcp(localhost:3306)/",
		},
	}
}

func (a *Auth) Commands() []*cli.Command {
	return nil
}

func (a *Auth) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return a.LoginHandler(h)
	}
}

// LoginHandler checkout token and get username.
func (a *Auth) LoginHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		login := path[len(path)-5:]
		if login == "auth" {
			h.ServeHTTP(w, r)
			return
		}
		token := r.Header.Get("authorization")
		fmt.Println(token)
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
		// 随机生成角色
		// 真实场景下可以从token种取role
		// role := GetRoleFromToken(token)
		role := []string{"public", "admin", "alice", "bob", "cathy"}
		random := func() int {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			return r.Intn(4)
		}

		method := r.Method
		if allowed, err := a.enforcer.Enforce(role[random()], path, method); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if !allowed {
			http.Error(w, fmt.Sprintf("对不起%s没有对%s的%s权限", role[random()], path, method), http.StatusForbidden)
			return
		}

		r.Header.Set("username", strconv.Itoa(username))
		r.Header.Set("authorization", "Bearer "+token)
		h.ServeHTTP(w, r)
	}
}

// ParesToken xh是学号\学工号
func (a *Auth) ParesToken(tokenString string) (int, error) {
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

func (a *Auth) Init(ctx *cli.Context) error {
	a.secret = ctx.String("secret")
	a.model = ctx.String("casbin_model")
	a.adapter = ctx.String("casbin_file_adapter")
	a.useMysql = ctx.Bool("casbin_adapter_use_mysql")
	a.mysqlAdapter = ctx.String("casbin_adapter_mysql")

	var ef *casbin.Enforcer
	var err error
	if a.useMysql {
		ef, err = NewEnforcer(a.model, "", a.mysqlAdapter)
		if err != nil {
			panic(err)
		}
	} else {
		ef, err = NewEnforcer(a.model, a.adapter, "")
		if err != nil {
			panic(err)
		}
	}
	a.enforcer = ef
	// TODO 添加消息推送，及时LoadPolicy
	// 目前方案:
	//    1. 轮询mysql
	//    2. 写mysql插件, 自动通知
	if err = a.enforcer.LoadPolicy(); err != nil {
		panic(err)
	}
	a.enforcer.EnableAutoSave(true)
	ok, err := a.enforcer.AddPolicy("public", "/user/call", "GET")
	if err != nil {
		panic(err)
	}
	if !ok {
		return errors.New("存在重复的policy")
	}
	return nil
}

func NewEnforcer(model, fileAdapter, mysqlAdapter string) (*casbin.Enforcer, error) {
	if fileAdapter != "" {
		return casbin.NewEnforcer(model, fileAdapter)
	}
	adapter, err := xormadapter.NewAdapter("mysql", mysqlAdapter)
	if err != nil {
		panic(err)
	}
	return casbin.NewEnforcer(model, adapter)
}

func (a *Auth) String() string {
	return "Auth"
}

func NewPlugin() plugin.Plugin {
	return &Auth{}
}
