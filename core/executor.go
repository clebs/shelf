package core

import (
	"fmt"
	"os/exec"
)

// AsyncExecutor is a set of commands that run concurrently
type AsyncExecutor []*exec.Cmd

// Run executes all commands concurrently
func (asyncExec AsyncExecutor) Run() {
	n := len(asyncExec)
	ch := make(chan *string, n)

	for _, cmd := range asyncExec {
		worker := NewWorker(cmd, ch)
		go worker.Run()
	}

	for i := 0; i < n; i++ {
		fmt.Println(*<-ch)
	}
}

// SyncExecutor is a set of commands that run in sequence
type SyncExecutor []*exec.Cmd

// Run executes all commands sequentially
func (syncExec SyncExecutor) Run() {
	for _, cmd := range syncExec {
		out, err := cmd.Output()
		if err != nil {
			println(err)
		} else {
			println(out)
		}
	}
}
