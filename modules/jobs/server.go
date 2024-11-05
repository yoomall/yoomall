package jobs

import "yoomall/yoo"

type JobServer struct {
	Cron *yoo.Cron
}

func NewJobServer(cron *yoo.Cron, job1 *Job1) *JobServer {
	cron.AddJob(job1)
	return &JobServer{
		Cron: cron,
	}
}

func (j *JobServer) Start() {
	j.Cron.Start()
}
