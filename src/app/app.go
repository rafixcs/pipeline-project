package app

type IApp interface {
	Run() string
}

type App struct{}

func (App) Run() string {
	return "running"
}
