package jobs

import (
	"fmt"
	"time"

	"github.com/suborbital/reactr/rt"
)

type timed struct{}

// Run runs a generic job
func (t timed) Run(job rt.Job, ctx *rt.Ctx) (interface{}, error) {
	time.Sleep(time.Second * 30)

	j := job.Data().(*Job)
	j.Result = &Result{
		Message: fmt.Sprintf("Finished job at %s", time.Now().String()),
	}

	return fmt.Sprintf("finished %s", job.String()), nil
}

// OnChange is called when Reactr starts or stops a worker to handle jobs,
// and allows the Runnable to set up before receiving jobs or tear down if needed.
func (t timed) OnChange(change rt.ChangeEvent) error {
	return nil
}
