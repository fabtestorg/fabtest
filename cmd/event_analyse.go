package cmd

import "fmt"

func EventAnalyse(logdir string) error {
	if logdir == "" {
		return fmt.Errorf("logdir is nil")
	}
	obj := NewLocalFabCmd("event_analyse.py")
	return obj.RunShow("analyse", ConfigDir(),logdir)
}
