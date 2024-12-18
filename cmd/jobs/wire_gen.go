// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"yoomall/modules/jobs"
	"yoomall/yoo"
)

// Injectors from wire.go:

func NewCron() *jobs.JobServer {
	cron := yoo.NewCorn()
	job1 := jobs.NewJob1()
	jobServer := jobs.NewJobServer(cron, job1)
	return jobServer
}
