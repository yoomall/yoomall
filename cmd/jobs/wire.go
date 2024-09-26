//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"lazyfury.github.com/yoomall-server/jobs"
)

func NewCron() *jobs.JobServer {
	wire.Build(jobs.NewCorn, jobs.NewJob1, jobs.NewJobServer)
	return &jobs.JobServer{}
}
