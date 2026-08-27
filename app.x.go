// Managed by GoX v0.3.1

//line app.gox:1
package main

import (
	"context"
	
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

type Path struct {
	Route Route `/:" | :CityID"`
	CityID int
}

type Route int

const (
	Selector Route = iota
	Dashboard
)

type App struct{}

//line app.gox:24
func (a App) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Raw("<!doctype html>"); if __e != nil { return }
		__e = __c.Init("html"); if __e != nil { return }
		{
//line app.gox:26
			__e = __c.Set("lang", "en"); if __e != nil { return }
//line app.gox:26
			__e = __c.Set("data-theme", "dark"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("head"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:28
					__e = __c.Set("charset", "utf-8"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:29
					__e = __c.Set("name", "viewport"); if __e != nil { return }
//line app.gox:29
					__e = __c.Set("content", "width=device-width, initial-scale=1"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:31
					__e = __c.Modify(doors.ResourceExternal("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css")); if __e != nil { return }
//line app.gox:32
					__e = __c.Set("rel", "stylesheet"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("body"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("main"); if __e != nil { return }
				{
//line app.gox:35
					__e = __c.Set("class", "container"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line app.gox:36
					__e = __c.Any(doors.Route(
					doors.RouteModel(a.content),
					doors.RouteDefaultComp[doors.Location](gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.InitContainer(); if __e != nil { return }
						{
							__e = __c.Init("title"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Text("Not Found"); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line app.gox:40
							__e = __c.Any(doors.Status(404)); if __e != nil { return }
							__e = __c.Init("h1"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Text("Location Not Found"); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line app.gox:42
					return })),
				)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line app.gox:47
}

//line app.gox:49
func (a App) content(path doors.Source[Path]) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line app.gox:50
		__e = __c.Any(path.Route(
		doors.RouteMatch(func(p Path) bool {
			return p.Route == Selector
		}).Comp(gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.InitContainer(); if __e != nil { return }
			{
				__e = __c.Init("title"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("Select Location"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line app.gox:55
				__e = __c.Any(LocationSelector(func(ctx context.Context, city int) {
				path.Update(ctx, Path{
					Route: Dashboard,
					CityID: city,
				})
			})); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line app.gox:61
		return })),
		doors.RouteDefault(WeatherDashboard),
	)); if __e != nil { return }
	return })
//line app.gox:64
}
