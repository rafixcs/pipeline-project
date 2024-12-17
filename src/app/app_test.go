package app

import "testing"

func TestRunApp(t *testing.T) {
	t.Run("Testin running app", func(t *testing.T) {
		app := App{}
		result := app.Run()
		if result != "running" {
			t.Fatalf(`app is not "running" it is "%s"`, result)
		}
	})
}
