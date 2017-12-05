package cmd

func EventAnalyse(logdir string) error {
	obj := NewLocalFabCmd("event_analyse.py")
	return obj.RunShow("analyse", ConfigDir(),logdir)
}
