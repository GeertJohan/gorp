package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func dockerIPAddress(name string) (string, error) {
	dockerInspect := exec.Command("docker", "inspect", "--format='{{.NetworkSettings.IPAddress}}'", name)
	addrBuf := &bytes.Buffer{}
	dockerInspect.Stdout = addrBuf
	err := dockerInspect.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(addrBuf.String()), nil
}

type waitUntilWriter struct {
	buffer  *bytes.Buffer
	trigger string
	doneCh  chan struct{}
}

func newWaitUntilWriter(trigger string) *waitUntilWriter {
	w := &waitUntilWriter{
		buffer:  &bytes.Buffer{},
		trigger: trigger,
		doneCh:  make(chan struct{}),
	}
	return w
}

func (w *waitUntilWriter) Write(p []byte) (n int, err error) {
	w.buffer.Write(p)
	if strings.Contains(w.buffer.String(), w.trigger) {
		select {
		case <-w.doneCh:
		default:
			close(w.doneCh)
		}
	}
	return len(p), nil
}
func (w *waitUntilWriter) Done() {
	<-w.doneCh
}

func dockerWait(name string, trigger string) {
	dockerLogs := exec.Command("docker", "logs", "--follow", name)
	w := newWaitUntilWriter(trigger)
	dockerLogs.Stdout = w
	dockerLogs.Stderr = w
	dockerLogs.Start()
	w.Done()
	dockerLogs.Process.Kill()
}

func dockerStop(name string) error {
	cmdDockerStop := exec.Command("docker", "stop", name)
	linkStdio(cmdDockerStop)
	err := cmdDockerStop.Run()
	if err != nil {
		return err
	}
	return nil
}

func dockerRemove(name string) error {
	cmdDockerRemove := exec.Command("docker", "rm", name)
	linkStdio(cmdDockerRemove)
	err := cmdDockerRemove.Run()
	if err != nil {
		return err
	}
	return nil
}

func linkStdio(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
