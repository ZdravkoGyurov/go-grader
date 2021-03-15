package config

import "time"

// Config for application properties
type Config struct {
	Host                      string
	Port                      int
	MaxExecutorWorkers        int
	MaxExecutorConcurrentJobs int
	DatabaseURI               string
	DBConnectTimeout          time.Duration
	DBDisconnectTimeout       time.Duration
	DatabaseName              string
	ServerShutdownTimeout     time.Duration
	SessionCookieName         string
	TestsGitUser              string
	TestsGitRepo              string
}
