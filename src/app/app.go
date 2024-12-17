package app

type IApp interface {
	Run() string
	Stop() string
}

type App struct{}

func (App) Run() string {
	return "running"
}

func (App) Stop() string {
	return "stopped"
}
