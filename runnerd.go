package runnerd

import (
	"os/exec"
)

const (
	STOPPED  = iota
	STARTING = iota
	RUNNING  = iota
	BACKOFF  = iota
	STOPPING = iota
	EXITED   = iota
	FATAL    = iota
	UNKNOWN  = iota
)

type Process struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

func (p *Process) State() {
}

type Runner struct {
	Ps []*Process
}
