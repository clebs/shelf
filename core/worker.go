package core

import "os/exec"

// Worker is a Command that runs concurrently, it forwards the execution output to a channel
type Worker struct {
	Command *exec.Cmd
	Output  chan *string
}

// Run executes the command and sends output (stdout and stderr) to the channel
func (w Worker) Run() {
	out, err := w.Command.CombinedOutput()
	if err != nil {
		s := string(err.Error())
		w.Output <- &s
		return
	}
	s := string(out)
	w.Output <- &s
}

// NewWorker creates a new Worker instance
func NewWorker(cmd *exec.Cmd, ch chan *string) Worker {
	return Worker{
		Command: cmd,
		Output:  ch}
}
