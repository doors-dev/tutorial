package main

import (
	"context"
	"net/http"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
	"github.com/doors-dev/tutorial/driver"
)

func main() {
	app := doors.NewApp(func(ctx context.Context, r doors.Request) gox.Comp {
		session := doors.SessionStore(ctx).Init("session", func() any {
			var s driver.Session
			c, err := r.GetCookie("session")
			if err == nil {
				s = driver.Sessions.Get(c.Value)
			}
			return doors.NewSource(s)
		}).(doors.Source[driver.Session])
		return App{
			session: session,
		}
	})

	if err := http.ListenAndServe(":8080", app); err != nil {
		panic(err)
	}
}
