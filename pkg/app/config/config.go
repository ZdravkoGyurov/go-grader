package config

import "time"

// Config for application properties
type Config struct {
	Host                      string
	Port                      int
	ServerReadTimeout         time.Duration
	ServerWriteTimeout        time.Duration
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

func DefaultConfig() Config {
	return Config{
		Host:                      "0.0.0.0",
		Port:                      8080,
		ServerReadTimeout:         30 * time.Second,
		ServerWriteTimeout:        30 * time.Second,
		MaxExecutorWorkers:        5,
		MaxExecutorConcurrentJobs: 100,
		DatabaseURI:               "mongodb://host.docker.internal:27017",
		DBConnectTimeout:          30 * time.Second,
		DBDisconnectTimeout:       30 * time.Second,
		DatabaseName:              "grader",
		ServerShutdownTimeout:     5 * time.Second,
		SessionCookieName:         "Grader",
		TestsGitUser:              "ZdravkoGyurov",
		TestsGitRepo:              "grader-docker-tests",
	}
}
