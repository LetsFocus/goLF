package cronjobs

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// NewCronManager creates a new CronManager.
func NewCronManager() *CronManager {
	return &CronManager{
		stop: make(chan struct{}),
	}
}

// AddJob adds a new job to the cron manager.
func (c *CronManager) AddJob(spec string, task func()) {
	c.mu.Lock()
	defer c.mu.Unlock()

	job := &Job{
		Spec: spec,
		Task: task,
	}

	c.jobs = append(c.jobs, job)
}

// Start starts the cron manager.
func (c *CronManager) Start() {
	for _, job := range c.jobs {
		go func(job *Job) {
			c.runJob(job)
		}(job)
	}
}

// Stop stops the cron manager.
func (c *CronManager) Stop() {
	close(c.stop)
}

// runJob executes the task associated with a given Job at scheduled intervals.
// It continuously calculates the next run time based on the Job's cron expression,
// sleeps until that time, and then triggers the associated task.
func (c *CronManager) runJob(job *Job) {
	for {
		select {
		case <-c.stop:
			return
		default:
			now := time.Now()
			nextRun := parseCronExpression(job.Spec, now)

			sleepDuration := time.Until(nextRun)

			select {
			case <-c.stop:
				return
			case <-time.After(sleepDuration):
				job.Task()
			}
		}
	}
}

// parseCronExpression parses a cron expression and calculates the next run time.
func parseCronExpression(spec string, now time.Time) time.Time {
	fields := strings.Fields(spec)

	if len(fields) != 5 {
		log.Printf("Invalid cron expression: %s", spec)
		return time.Time{}
	}

	// Parse each field.
	second := parseCronField(fields[0], 0, 59, now.Second())
	minute := parseCronField(fields[1], 0, 59, now.Minute())
	hour := parseCronField(fields[2], 0, 23, now.Hour())
	dayOfMonth := parseCronField(fields[3], 1, 31, now.Day())
	month := parseCronField(fields[4], 1, 12, int(now.Month()))

	// Calculate the next run time based on the cron expression.
	nextRun := time.Date(now.Year(), time.Month(month), dayOfMonth, hour, minute, second, 0, now.Location())

	// If the calculated time is in the past, add one day
	if nextRun.Before(now) {
		nextRun = nextRun.Add(1 * time.Second)

	}

	return nextRun
}

func parseCronField(field string, min, max, defaultValue int) int {
	if field == "*" {
		return defaultValue
	}

	value, err := strconv.Atoi(field)
	if err != nil || value < min || value > max {
		log.Printf("Invalid field in cron expression: %s", field)
		return defaultValue
	}

	return value + defaultValue
}
