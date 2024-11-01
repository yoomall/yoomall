//go:build wireinject
// +build wireinject

package main

import (
	"yoomall/src/apps/jobs"

	"github.com/google/wire"
)

func NewCron() *jobs.JobServer {
	wire.Build(jobs.NewCorn, jobs.NewJob1, jobs.NewJobServer)
	return &jobs.JobServer{}
}
