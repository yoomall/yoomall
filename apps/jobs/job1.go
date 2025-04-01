package jobs

import (
	"fmt"
	"time"

	"github.com/lazyfury/pulse/framework"
)

type Job1 struct {
}

var _ framework.IJob = (*Job1)(nil)

func NewJob1() *Job1 {
	return &Job1{}
}

func (j *Job1) Job() {
	t := time.Now().Add(time.Second * 1)

	fmt.Println("job1", t.Format("2006-01-02 15:04:05"))
}

func (j *Job1) Spec() string {
	return "@every 5s"
}
