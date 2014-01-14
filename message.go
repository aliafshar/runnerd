package runnerd

import (
  "os"
)

const(
  MSG_OUT = iota
  MSG_STDERR
  MSG_EXIT
  PIPE_STDERR
  PIPE_STDOUT
)

type Message interface {
  Type() int
}

type PipeMessage struct {
  Body []byte
  Stream int
}

func (m *PipeMessage) Type() int {
  return MSG_OUT
}

type ExitMessage struct {
  ProcessState *os.ProcessState
}

func (m *ExitMessage) Type() int {
  return MSG_EXIT
}
