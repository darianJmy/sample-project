package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sample-project/api/server/httputils"
	"sample-project/tools"
	"strconv"
	"strings"
)

func (m *middleware) authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := httputils.NewResponse()

		if os.Getenv("DEBUG") == "true" {
			return
		}

		path := c.Request.URL.Path
		if path == "/users/login" {
			return
		}

		method := c.Request.Method
		auth := c.GetHeader("Authorization")

		if len(auth) == 0 || auth == "" {
			r.SetCode(http.StatusUnauthorized)
			httputils.SetFailed(c, r, errors.New("permission denied"))
			c.Abort()
			return
		}

		auth = strings.TrimLeft(auth, "\"")
		auth = strings.TrimRight(auth, "\"")

		fields := strings.Fields(auth)
		if len(fields) != 2 || fields[0] != "Bearer" {
			r.SetCode(http.StatusUnauthorized)
			httputils.SetFailed(c, r, errors.New("invalid authorization header format"))
			c.Abort()
			return
		}

		claims, err := tools.ParseToken(fields[1])
		if err != nil {
			r.SetCode(http.StatusUnauthorized)
			httputils.SetFailed(c, r, err)
			c.Abort()
			return
		}

		fmt.Println(claims.Id, path, method)

		ok, err := m.enforcer.Enforce(strconv.FormatInt(claims.Id, 10), path, method)
		if err != nil {
			r.SetCode(http.StatusInternalServerError)
			httputils.SetFailed(c, r, errors.New("inner error"))
			c.Abort()
			return
		}

		if !ok {
			r.SetCode(http.StatusUnauthorized)
			httputils.SetFailed(c, r, errors.New("permission denied"))
			c.Abort()
			return
		}
	}
}
