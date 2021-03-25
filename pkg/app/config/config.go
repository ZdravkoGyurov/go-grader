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
	DBRequestTimeout          time.Duration
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
		DatabaseURI:               "mongodb://localhost:27017,localhost:27018,localhost:27019/grader?replicaSet=gograder-mongo-set",
		DBConnectTimeout:          30 * time.Second,
		DBDisconnectTimeout:       30 * time.Second,
		DBRequestTimeout:          30 * time.Second,
		DatabaseName:              "grader",
		ServerShutdownTimeout:     5 * time.Second,
		SessionCookieName:         "Grader",
		TestsGitUser:              "ZdravkoGyurov",
		TestsGitRepo:              "grader-docker-tests",
	}
}
