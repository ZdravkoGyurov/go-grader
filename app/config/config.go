package config

import "time"

// Config for application properties
type Config struct {
	Host                      string
	Port                      int
	MaxExecutorWorkers        int
	MaxExecutorConcurrentJobs int
	DBConnectTimeout          time.Duration
	DBDisconnectTimeout       time.Duration
	GithubTestsRepo           string
}
