// Managed by GoX v0.3.1

//line dashboard.gox:1
package main

import (
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
	"github.com/doors-dev/tutorial/driver"
)

func WeatherDashboard(path doors.Source[Path]) gox.Comp {
	city := doors.DeriveBeam(path, func(p Path) int {
		return p.CityID
	})
	return city.Bind(func(cityID int) gox.Elem {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
//line dashboard.gox:14
			__e = __c.Any(dashboard{
			cityID: cityID,
		}); if __e != nil { return }
		return })
//line dashboard.gox:17
	})
}

type dashboard struct {
	cityID int
}

//line dashboard.gox:24
func (d dashboard) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line dashboard.gox:26
		city, _ := driver.Locations.CitiesGet(d.cityID)

//line dashboard.gox:28
		if !city.IsValid() {
//line dashboard.gox:29
			__e = __c.Any(doors.Status(404)); if __e != nil { return }
			__e = __c.Init("title"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Text("Not Found"); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Text("City Not Found"); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		} else  {
			__e = __c.Init("title"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:33
				__e = __c.Any(city.Name); if __e != nil { return }
				__e = __c.Text(" Weather"); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Text("Weather in "); if __e != nil { return }
//line dashboard.gox:34
				__e = __c.Many(city.Name, ", ", city.Country.Name); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Init("a"); if __e != nil { return }
		{
//line dashboard.gox:37
			__e = __c.Set("class", "secondary"); if __e != nil { return }
//line dashboard.gox:38
			__e = __c.Modify(doors.ALink{
			Model: Path{Route: Selector},
		}); if __e != nil { return }
//line dashboard.gox:41
			__e = __c.Set("role", "button"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Change"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:44
}
