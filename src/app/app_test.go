package app

import "testing"

func TestRunApp(t *testing.T) {
	t.Run("Testing running app", func(t *testing.T) {
		app := App{}
		result := app.Run()
		if result != "running" {
			t.Fatalf(`app is not "running" it is "%s"`, result)
		}
	})
}

func TestStopApp(t *testing.T) {
	t.Run("Testing stopping app", func(t *testing.T) {
		app := App{}
		result := app.Stop()
		if result != "stopped" {
			t.Fatalf(`app is not "stopped" it is "%s"`, result)
		}
	})
}
