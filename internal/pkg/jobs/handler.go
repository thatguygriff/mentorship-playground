package jobs

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/suborbital/vektor/vk"
)

var jobs = make(map[string]*Job)

type Job struct {
	ID     string  `json:"id"`
	Href   string  `json:"href,omitempty"`
	Result *Result `json:"result,omitempty"`
}

type Result struct {
	Message string `json:"message,omitempty"`
}

func Register(r *http.Request, ctx *vk.Ctx) (interface{}, error) {
	uuid := uuid.New().String()
	job := &Job{
		ID:   uuid,
		Href: fmt.Sprintf("%s/%s", r.URL.Path, uuid),
	}

	jobs[uuid] = job
	return vk.Respond(http.StatusCreated, job), nil
}

func Check(r *http.Request, ctx *vk.Ctx) (interface{}, error) {
	jobID := ctx.Params.ByName("jobID")

	if jobs[jobID] == nil {
		return nil, vk.Err(http.StatusNotFound, fmt.Sprintf("No Job with ID %s was found", jobID))
	}

	return jobs[jobID], nil
}
