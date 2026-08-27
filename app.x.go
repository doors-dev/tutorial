// Managed by GoX v0.3.1

//line app.gox:1
package main

import (
	"context"
	_ "embed"
	
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
	"github.com/doors-dev/tutorial/driver"
)

//go:embed style.css
var styles []byte

type Path struct {
	Route Route `/:" | :CityID"`
	CityID int
	Units *driver.Units `query:"units"`
	Days *int `query:"days"`
}

func (p Path) days() int {
	if p.Days == nil {
		return 7
	}
	return min(max(*p.Days, 1), 7)
}

func daysQuery(days int) *int {
	if days == 7 {
		return nil
	}
	return &days
}

func (p Path) units() driver.Units {
	if p.Units == nil || *p.Units != driver.Imperial {
		return driver.Metric
	}
	return driver.Imperial
}

func unitsQuery(units driver.Units) *driver.Units {
	if units == driver.Metric {
		return nil
	}
	return &units
}

type Route int

const (
	Selector Route = iota
	Dashboard
)

type App struct{}

//line app.gox:59
func (a App) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Raw("<!doctype html>"); if __e != nil { return }
		__e = __c.Init("html"); if __e != nil { return }
		{
//line app.gox:61
			__e = __c.Set("lang", "en"); if __e != nil { return }
//line app.gox:61
			__e = __c.Set("data-theme", "dark"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("head"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:63
					__e = __c.Set("charset", "utf-8"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:64
					__e = __c.Set("name", "viewport"); if __e != nil { return }
//line app.gox:64
					__e = __c.Set("content", "width=device-width, initial-scale=1"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:66
					__e = __c.Modify(doors.ResourceExternal("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css")); if __e != nil { return }
//line app.gox:67
					__e = __c.Set("rel", "stylesheet"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:68
					__e = __c.Set("href", styles); if __e != nil { return }
//line app.gox:68
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
//line app.gox:71
					__e = __c.Set("class", "container"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line app.gox:72
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
//line app.gox:76
							__e = __c.Any(doors.Status(404)); if __e != nil { return }
							__e = __c.Init("h1"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Text("Location Not Found"); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line app.gox:78
					return })),
				)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line app.gox:83
}

//line app.gox:85
func (a App) content(path doors.Source[Path]) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line app.gox:86
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
//line app.gox:91
				__e = __c.Any(LocationSelector(func(ctx context.Context, city int) {
				path.Mutate(ctx, func(p Path) Path {
					p.Route = Dashboard
					p.CityID = city
					return p
				})
			})); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line app.gox:98
		return })),
		doors.RouteDefault(WeatherDashboard),
	)); if __e != nil { return }
	return })
//line app.gox:101
}
