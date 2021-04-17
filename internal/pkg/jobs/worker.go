package jobs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/suborbital/reactr/rt"
)

const workerKey = "genericWorkerCount"

type timed struct{}

// Run runs a generic job
func (t timed) Run(job rt.Job, ctx *rt.Ctx) (interface{}, error) {
	time.Sleep(time.Second * 5)

	count := 0

	countBytes, err := ctx.Cache.Get(workerKey)
	if err == nil && countBytes != nil {
		s := string(countBytes)
		count, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println(countBytes)
		}
	}
	count++

	j := job.Data().(*Job)
	j.Result = &Result{
		Message: fmt.Sprintf("Finished job #%d at %s", count, time.Now().String()),
	}

	ctx.Cache.Set(workerKey, []byte(fmt.Sprintf("%d", count)), 0)

	return fmt.Sprintf("finished %s", job.String()), nil
}

// OnChange is called when Reactr starts or stops a worker to handle jobs,
// and allows the Runnable to set up before receiving jobs or tear down if needed.
func (t timed) OnChange(change rt.ChangeEvent) error {
	return nil
}
