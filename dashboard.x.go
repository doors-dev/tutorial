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
	days := doors.DeriveBeam(path, func(p Path) int {
		return p.days()
	})
	units := doors.DeriveBeam(path, func(p Path) driver.Units {
		return p.units()
	})
	return city.Bind(func(cityID int) gox.Elem {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
//line dashboard.gox:20
			__e = __c.Any(dashboard{
			cityID: cityID,
			days: days,
			units: units,
		}); if __e != nil { return }
		return })
//line dashboard.gox:25
	})
}

type dashboard struct {
	cityID int
	days doors.Beam[int]
	units doors.Beam[driver.Units]
}

//line dashboard.gox:34
func (d dashboard) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line dashboard.gox:36
		city, _ := driver.Locations.CitiesGet(d.cityID)

//line dashboard.gox:38
		if !city.IsValid() {
//line dashboard.gox:39
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
//line dashboard.gox:43
				__e = __c.Any(city.Name); if __e != nil { return }
				__e = __c.Text(" Weather"); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Text("Weather in "); if __e != nil { return }
//line dashboard.gox:44
				__e = __c.Many(city.Name, ", ", city.Country.Name); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
//line dashboard.gox:46
		__e = __c.Any(d.change()); if __e != nil { return }
//line dashboard.gox:47
		__e = __c.Any(d.menu()); if __e != nil { return }
	return })
//line dashboard.gox:48
}

//line dashboard.gox:50
func (d dashboard) change() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line dashboard.gox:51
		__e = (new(doors.Door)).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.InitContainer(); if __e != nil { return }
			{
//line dashboard.gox:53
				days, _ := d.days.Effect(ctx)
		units, ok := d.units.Effect(ctx)

//line dashboard.gox:56
				if ok {
					__e = __c.Init("a"); if __e != nil { return }
					{
//line dashboard.gox:58
						__e = __c.Set("class", "secondary"); if __e != nil { return }
//line dashboard.gox:59
						__e = __c.Modify(doors.ALink{
					Model: Path{
						Route: Selector,
						Days: daysQuery(days),
						Units: unitsQuery(units),
					},
				}); if __e != nil { return }
//line dashboard.gox:66
						__e = __c.Set("role", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("Change"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
		return })); if __e != nil { return }
	return })
//line dashboard.gox:71
}

//line dashboard.gox:73
func (d dashboard) menu() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("nav"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:75
			__e = __c.Any(d.units.Bind(d.daysNav)); if __e != nil { return }
//line dashboard.gox:76
			__e = __c.Any(d.days.Bind(d.unitNav)); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:78
}

//line dashboard.gox:80
func (d dashboard) daysNav(units driver.Units) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ul"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:82
			for i := range 7 {
//line dashboard.gox:84
				days := i + 1

				__e = __c.Init("li"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("a"); if __e != nil { return }
					{
//line dashboard.gox:88
						__e = __c.Set("class", "secondary"); if __e != nil { return }
//line dashboard.gox:89
						__e = __c.Modify(doors.ALink{
						Model: Path{
							Route: Dashboard,
							CityID: d.cityID,
							Days: daysQuery(days),
							Units: unitsQuery(units),
						},
					}); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:97
						__e = __c.Any(days); if __e != nil { return }
//line dashboard.gox:98
						if days == 1 {
							__e = __c.Text(" day"); if __e != nil { return }
						} else  {
							__e = __c.Text(" days"); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:107
}

//line dashboard.gox:109
func (d dashboard) unitNav(days int) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ul"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:111
			for _, units := range []driver.Units{driver.Metric, driver.Imperial} {
				__e = __c.Init("li"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("a"); if __e != nil { return }
					{
//line dashboard.gox:114
						__e = __c.Set("class", "secondary"); if __e != nil { return }
//line dashboard.gox:115
						__e = __c.Modify(doors.ALink{
						Model: Path{
							Route: Dashboard,
							CityID: d.cityID,
							Days: daysQuery(days),
							Units: unitsQuery(units),
						},
					}); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:123
						__e = __c.Any(units.String()); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:128
}
