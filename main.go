package main

import (
	"context"
	"net/http"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

func main() {
	// Create a Doors app with a factory for our App.
	app := doors.NewApp(func(ctx context.Context, r doors.Request) gox.Comp {
		return App{}
	})

	// Serve the Doors app on port 8080.
	if err := http.ListenAndServe(":8080", app); err != nil {
		panic(err)
	}
}
