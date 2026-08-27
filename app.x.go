// Managed by GoX v0.3.1

//line app.gox:1
package main

import (
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

type App struct{}

//line app.gox:10
func (a App) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Raw("<!doctype html>"); if __e != nil { return }
		__e = __c.Init("html"); if __e != nil { return }
		{
//line app.gox:12
			__e = __c.Set("lang", "en"); if __e != nil { return }
//line app.gox:12
			__e = __c.Set("data-theme", "dark"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("head"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:14
					__e = __c.Set("charset", "utf-8"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:15
					__e = __c.Set("name", "viewport"); if __e != nil { return }
//line app.gox:15
					__e = __c.Set("content", "width=device-width, initial-scale=1"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("title"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("Hello Doors!"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:18
					__e = __c.Modify(doors.ResourceExternal("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css")); if __e != nil { return }
//line app.gox:19
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
//line app.gox:22
					__e = __c.Set("class", "container"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line app.gox:23
					__e = __c.Any(LocationSelector()); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line app.gox:27
}
