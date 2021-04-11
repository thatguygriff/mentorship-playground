package jobs

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/suborbital/reactr/rt"
	"github.com/suborbital/vektor/vk"
)

const defaultDelay = 60

type Manager struct {
	jobs   map[string]*Job
	reactr *rt.Reactr
}

func NewManager() *Manager {
	r := rt.New()
	r.Handle("timed", timed{})

	return &Manager{
		jobs:   make(map[string]*Job),
		reactr: r,
	}
}

type Job struct {
	ID     string  `json:"id"`
	Href   string  `json:"href,omitempty"`
	Result *Result `json:"result,omitempty"`
}

type Result struct {
	Message string `json:"message,omitempty"`
}

func (m *Manager) Register(r *http.Request, ctx *vk.Ctx) (interface{}, error) {
	uuid := uuid.New().String()
	job := &Job{
		ID:   uuid,
		Href: fmt.Sprintf("%s/%s", r.URL.Path, uuid),
	}
	m.jobs[uuid] = job

	m.reactr.Do(m.reactr.Job("timed", job))

	return vk.Respond(http.StatusCreated, job), nil
}

func (m *Manager) RegisterSync(r *http.Request, ctx *vk.Ctx) (interface{}, error) {
	uuid := uuid.New().String()
	job := &Job{
		ID:   uuid,
		Href: fmt.Sprintf("%s/%s", r.URL.Path, uuid),
	}
	m.jobs[uuid] = job

	m.reactr.Do(m.reactr.Job("timed", job)).Then()

	return vk.Respond(http.StatusCreated, job), nil
}

func (m *Manager) Check(r *http.Request, ctx *vk.Ctx) (interface{}, error) {
	jobID := ctx.Params.ByName("jobID")

	if m.jobs[jobID] == nil {
		return nil, vk.Err(http.StatusNotFound, fmt.Sprintf("No Job with ID %s was found", jobID))
	}

	return m.jobs[jobID], nil
}
