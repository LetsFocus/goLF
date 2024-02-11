package cronjobs

import "sync"

// Job represents a cron job.
type Job struct {
	Spec string
	Task func()
}

// CronManager manages cron jobs.
type CronManager struct {
	jobs []*Job
	stop chan struct{}
	mu   sync.Mutex
}
