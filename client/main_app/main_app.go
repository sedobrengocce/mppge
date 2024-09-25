package mainapp

import (
    "github.com/veandco/go-sdl2/sdl"
)

type App struct {
    Window *sdl.Window
    Renderer *sdl.Renderer
    Running bool
    spf int
    RenderFunc func()
    UpdateFunc func()
    EventFunc func()
}

func NewApp(title string, width, height int32) (*App, error) {
    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        return nil, err
    }

    window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
    if err != nil {
        return nil, err
    }

    renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
        return nil, err
    }

    spf := 1000 / 60
    
    running := false

    return &App{window, renderer, running, spf, nil, nil, nil}, nil
}

func (app *App) Stop() {
    app.Running = false
}

func (app *App) SetFPS(fps int) {
    app.spf = 1000 / fps
}

func (app *App) SetRenderFunc(f func()) {
    app.RenderFunc = f
}

func (app *App) SetUpdateFunc(f func()) {
    app.UpdateFunc = f
}

func (app *App) SetEventFunc(f func()) {
    app.EventFunc = f
}

func (app *App) Run() {
    app.Running = true
    for app.Running {
        start := sdl.GetTicks64()
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            if app.EventFunc != nil {
                app.EventFunc()
            }
            switch event.(type) {
            case *sdl.QuitEvent:
                app.Running = false
            }
        }

        if app.UpdateFunc != nil {
            app.UpdateFunc()
        }

        if app.RenderFunc != nil {
            app.RenderFunc()
        }

        end := sdl.GetTicks64()
        if end - start < uint64(app.spf) {
            sdl.Delay(uint32(app.spf) - uint32(end - start))
        }
    }
}
