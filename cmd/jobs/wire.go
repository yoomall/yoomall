//go:build wireinject
// +build wireinject

package main

import (
	"github.com/yoomall/yoomall/apps/jobs"
	"github.com/yoomall/yoomall/yoo"

	"github.com/google/wire"
)

func NewCron() *jobs.JobServer {
	wire.Build(yoo.NewCorn, jobs.NewJobServer, jobs.NewJob1)
	return &jobs.JobServer{}
}
