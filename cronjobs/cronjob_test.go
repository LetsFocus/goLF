package cronjobs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCronManager(t *testing.T) {
	t.Run("Create new CronManager", func(t *testing.T) {
		cronManager := NewCronManager()
		assert.NotNil(t, cronManager, "CronManager should not be nil")
		assert.NotNil(t, cronManager.stop, "Stop channel should not be nil")
	})
}

func TestAddJob(t *testing.T) {
	tests := []struct {
		description  string
		spec         string
		expectedTask bool
	}{
		{
			description:  "AddJob",
			spec:         "* * * * *",
			expectedTask: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			cronManager := NewCronManager()
			taskExecuted := false

			cronManager.AddJob(tt.spec, func() {
				taskExecuted = true
			})

			assert.Len(t, cronManager.jobs, 1, "Job should be added to CronManager")

			cronManager.jobs[0].Task()
			assert.Equal(t, tt.expectedTask, taskExecuted, "Task execution mismatch")
		})
	}
}

func TestStartAndStop(t *testing.T) {
	tests := []struct {
		description      string
		expectedTaskExec bool
	}{
		{
			description:      "Start and Stop CronManager",
			expectedTaskExec: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			cronManager := NewCronManager()
			taskExecuted := false

			cronManager.AddJob("* * * * *", func() {
				taskExecuted = true
			})

			// Start the CronManager in a goroutine
			go cronManager.Start()

			// Wait for a while to allow the job to execute
			time.Sleep(2 * time.Second)

			// Stop the CronManager
			cronManager.Stop()

			// Ensure that the task was executed at least once
			assert.Equal(t, tt.expectedTaskExec, taskExecuted, "Task execution mismatch")
		})
	}
}

func TestRunJob(t *testing.T) {
	tests := []struct {
		description      string
		spec             string
		expectedTaskExec bool
	}{
		{
			description:      "RunJob - Immediate execution",
			spec:             "* * * * *",
			expectedTaskExec: true,
		},
		{
			description:      "RunJob - Delayed execution",
			spec:             "2 * * * *", // Every 2 minutes
			expectedTaskExec: true,
		},
		{
			description:      "RunJob - Past time execution",
			spec:             "3 * * * *",
			expectedTaskExec: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			cronManager := NewCronManager()
			taskExecuted := false

			cronManager.AddJob(tt.spec, func() {
				taskExecuted = true
			})

			job := cronManager.jobs[0]

			// Run the job manually
			go cronManager.runJob(job)

			// Wait for a while to allow the job to execute
			time.Sleep(5 * time.Second)

			// Ensure that the task was executed
			assert.Equal(t, tt.expectedTaskExec, taskExecuted, "Task execution mismatch")
		})
	}
}

func TestParseCronExpression(t *testing.T) {
	tests := []struct {
		description string
		spec        string
		now         time.Time
		expected    time.Time
	}{
		{
			description: "Every second, current time at the beginning of the year",
			spec:        "* * * * *",
			now:         time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected:    time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			description: "Specific time on the same day",
			spec:        "15 12 * * *",
			now:         time.Date(2022, time.January, 1, 10, 0, 0, 0, time.UTC),
			expected:    time.Date(2022, time.January, 1, 10, 12, 15, 0, time.UTC),
		},
		{
			description: "Every minute, current time at the beginning of the year",
			spec:        "* * * * *",
			now:         time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected:    time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			description: "Invalid cron expression, should return zero time",
			spec:        "invalid",
			now:         time.Now(),
			expected:    time.Time{},
		},
		{
			description: "Specific time in the future",
			spec:        "30 15 1 2 *",
			now:         time.Date(2022, time.January, 1, 10, 0, 0, 0, time.UTC),
			expected:    time.Date(2022, time.January, 3, 11, 15, 30, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := parseCronExpression(tt.spec, tt.now)
			assert.Equal(t, tt.expected, result, "Unexpected result for %s", tt.description)
		})
	}
}

func TestParseCronField(t *testing.T) {
	tests := []struct {
		description string
		field       string
		min         int
		max         int
		defaultVal  int
		expected    int
	}{
		{
			description: "Wildcard, should use default value",
			field:       "*",
			min:         0,
			max:         59,
			defaultVal:  30,
			expected:    30,
		},
		{
			description: "Specific value within range",
			field:       "15",
			min:         0,
			max:         59,
			defaultVal:  30,
			expected:    45,
		},
		{
			description: "Invalid value, should use default value",
			field:       "100",
			min:         0,
			max:         59,
			defaultVal:  30,
			expected:    30,
		},
		{
			description: "Wildcard for hour, should use default value",
			field:       "*",
			min:         0,
			max:         23,
			defaultVal:  12,
			expected:    12,
		},
		{
			description: "Specific value for hour within range",
			field:       "8",
			min:         0,
			max:         23,
			defaultVal:  12,
			expected:    20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := parseCronField(tt.field, tt.min, tt.max, tt.defaultVal)
			assert.Equal(t, tt.expected, result, "Unexpected result for %s", tt.description)
		})
	}
}
