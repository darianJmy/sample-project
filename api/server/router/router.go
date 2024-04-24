package router

import (
	"sample-project/api/server/router/user"
	"sample-project/cmd/app/options"
)

type RegisterFunc func(o *options.ServerRunOptions)

func InstallRouters(o *options.ServerRunOptions) {
	fs := []RegisterFunc{
		user.NewRouter,
	}

	install(o, fs...)
}

func install(o *options.ServerRunOptions, fs ...RegisterFunc) {
	for _, f := range fs {
		f(o)
	}
}
