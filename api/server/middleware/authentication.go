package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-project/api/server/httputils"
	"sample-project/tools"
	"strings"
)

func (m *middleware) authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := httputils.NewResponse()

		path := c.Request.URL.Path
		if path == "/users/login" {
			return
		}

		method := c.Request.Method
		token := c.GetHeader("Authorization")
		if token == "" {
			r.SetCode(http.StatusUnauthorized)
			httputils.SetFailed(c, r, errors.New("permission denied"))
			c.Abort()
			return
		}

		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			r.SetCode(http.StatusUnauthorized)
			httputils.SetFailed(c, r, errors.New("permission denied"))
			c.Abort()
			return
		}

		fmt.Println(token)

		fields := strings.Fields(token)
		if len(fields) != 2 {
			r.SetCode(http.StatusUnauthorized)
			httputils.SetFailed(c, r, fmt.Errorf("invalid authorization header format"))
			c.Abort()
			return
		}

		fmt.Println(fields)
		if fields[0] != "Bearer" {
			r.SetCode(http.StatusUnauthorized)
			httputils.SetFailed(c, r, fmt.Errorf("unsupported authorization type"))
			c.Abort()
			return
		}

		fmt.Println("已经过了吧")

		accessToken := fields[1]
		claims, err := tools.ParseToken(accessToken)
		if err != nil {
			r.SetCode(http.StatusUnauthorized)
			httputils.SetFailed(c, r, err)
			c.Abort()
			return
		}

		ok, err := m.enforcer.Enforce(claims.Id, path, method)
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

		fmt.Println("已经过了吧")
	}
}
