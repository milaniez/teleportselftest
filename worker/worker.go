package worker

import (
	"os/exec"
	"time"
)

type JobStatus int

const (
	JobStatusNotStarted JobStatus = iota
	JobStatusInProgress
	JobStatusSuccess
	JobStatusFailUserAborted
	JobStatusFailExitCode
	JobStatusFailTimeout
	JobStatusFailCPUExceed
	JobStatusFailMemExceed
)

type Job struct {
	cmd      exec.Cmd
	status   JobStatus
	maxTime  time.Duration
	maxMemKB int64
}

func JobNew(maxTime time.Duration, maxMemKB int64) Job {
	return Job{maxTime: maxTime, maxMemKB: maxMemKB}
}