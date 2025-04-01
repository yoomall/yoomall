package jobs

import "github.com/lazyfury/pulse/framework"

type JobServer struct {
	Cron *framework.Cron
}

func NewJobServer(cron *framework.Cron, job1 *Job1) *JobServer {
	cron.AddJob(job1)
	return &JobServer{
		Cron: cron,
	}
}

func (j *JobServer) Start() {
	j.Cron.Start()
}
