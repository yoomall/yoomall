package jobs

type JobServer struct {
	Cron *Cron
}

func NewJobServer(cron *Cron, job1 *Job1) *JobServer {
	cron.AddJob(job1)
	return &JobServer{
		Cron: cron,
	}
}

func (j *JobServer) Start() {
	j.Cron.Start()
}
