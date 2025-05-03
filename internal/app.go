package internal

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) Start() {}
