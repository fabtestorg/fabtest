package cmd

func EventAnalyse() error {
	obj := NewLocalFabCmd("event_analyse.py")
	return obj.RunShow("analyse", ConfigDir())
}
