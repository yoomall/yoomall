//go:build wireinject
// +build wireinject

package main

import (
	"yoomall/apps/jobs"
	"yoomall/yoo"

	"github.com/google/wire"
)

func NewCron() *jobs.JobServer {
	wire.Build(yoo.NewCorn, jobs.NewJobServer, jobs.NewJob1)
	return &jobs.JobServer{}
}
