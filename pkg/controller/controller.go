package controller

import "sample-project/cmd/app/config"

type SampleInterface interface {
}

type sample struct {
	cc config.Config
}

func New(cfg config.Config) SampleInterface {
	return &sample{
		cfg,
	}
}
