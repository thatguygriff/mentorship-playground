package server

import (
	"github.com/suborbital/vektor/vk"
	"github.com/suborbital/vektor/vlog"
	"github.com/thatguygriff/mentorship-playground/internal/pkg/jobs"
)

func Start() {
	// Create Job Manager
	jobsManager := jobs.NewManager()

	// Create HTTP server on port 8080
	server := vk.New(vk.UseAppName("Playground"), vk.UseHTTPPort(8080), vk.UseLogger(vlog.Default()))
	server.GET("/jobs/:jobID", jobsManager.Check)
	server.POST("/jobs", jobsManager.Register)
	server.POST("/jobs/sync", jobsManager.RegisterSync)

	server.Start()
}
