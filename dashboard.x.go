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
		__e = __c.Init("section"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:47
			__e = __c.Any(d.change()); if __e != nil { return }
//line dashboard.gox:48
			__e = __c.Any(d.menu()); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line dashboard.gox:50
		if city.IsValid() {
			__e = __c.Init("section"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:52
				__e = __c.Any(d.charts(city)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
	return })
//line dashboard.gox:55
}

//line dashboard.gox:57
func (d dashboard) change() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line dashboard.gox:58
		__e = (new(doors.Door)).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.InitContainer(); if __e != nil { return }
			{
//line dashboard.gox:60
				days, _ := d.days.Effect(ctx)
		units, ok := d.units.Effect(ctx)

//line dashboard.gox:63
				if ok {
					__e = __c.Init("a"); if __e != nil { return }
					{
//line dashboard.gox:65
						__e = __c.Set("class", "secondary"); if __e != nil { return }
//line dashboard.gox:66
						__e = __c.Modify(doors.ALink{
					Model: Path{
						Route: Selector,
						Days: daysQuery(days),
						Units: unitsQuery(units),
					},
				}); if __e != nil { return }
//line dashboard.gox:73
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
//line dashboard.gox:78
}

//line dashboard.gox:80
func (d dashboard) menu() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("nav"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:82
			__e = __c.Any(d.units.Bind(d.daysNav)); if __e != nil { return }
//line dashboard.gox:83
			__e = __c.Any(d.days.Bind(d.unitNav)); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:85
}

//line dashboard.gox:87
func (d dashboard) daysNav(units driver.Units) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ul"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:89
			for i := range 7 {
//line dashboard.gox:91
				days := i + 1

				__e = __c.Init("li"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:94
					__e = __c.Any(navLink{
					city: d.cityID,
					days: days,
					units: units,
					text: gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.InitContainer(); if __e != nil { return }
						{
//line dashboard.gox:99
							__e = __c.Any(days); if __e != nil { return }
//line dashboard.gox:100
							if days == 1 {
								__e = __c.Text(" day"); if __e != nil { return }
							} else  {
								__e = __c.Text(" days"); if __e != nil { return }
							}
						}
						__e = __c.Close(); if __e != nil { return }
//line dashboard.gox:105
					return }),
				}); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:110
}

//line dashboard.gox:112
func (d dashboard) unitNav(days int) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ul"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:114
			for _, units := range []driver.Units{driver.Metric, driver.Imperial} {
				__e = __c.Init("li"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:116
					__e = __c.Any(navLink{
					city: d.cityID,
					days: days,
					units: units,
					text: units,
				}); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:125
}

type navLink struct {
	city int
	days int
	units driver.Units
	text any
}

//line dashboard.gox:134
func (l navLink) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("a"); if __e != nil { return }
		{
//line dashboard.gox:136
			__e = __c.Set("class", "secondary"); if __e != nil { return }
//line dashboard.gox:137
			__e = __c.Modify(doors.ALink{
			Active: doors.Active{
				Indicator: doors.IndicateAttr("aria-current", "true"),
			},
			Model: Path{
				Route: Dashboard,
				CityID: l.city,
				Days: daysQuery(l.days),
				Units: unitsQuery(l.units),
			},
		}); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:148
			__e = __c.Any(l.text); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:150
}

//line dashboard.gox:152
func (d dashboard) charts(city driver.City) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("div"); if __e != nil { return }
		{
//line dashboard.gox:153
			__e = __c.Set("class", "grid"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:155
				__e = (new(doors.Door)).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
//line dashboard.gox:155
					__e = __c.Any(func() any {
				days, _ := d.days.Effect(ctx)
				units, ok := d.units.Effect(ctx)
				if !ok {
					return nil
				}
				return chart{
					title: "Temperature",
					svg: func() []byte {
						values, _ := driver.Weather.Temperature(ctx, city, units, days)
						svg, _ := driver.ChartLine(values.Values, values.Labels, units.Temperature())
						return svg
					},
				}
			}()); if __e != nil { return }
				return })); if __e != nil { return }
//line dashboard.gox:170
				__e = (new(doors.Door)).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
//line dashboard.gox:170
					__e = __c.Any(func() any {
				days, ok := d.days.Effect(ctx)
				if !ok {
					return nil
				}
				return chart{
					title: "Humidity",
					svg: func() []byte {
						values, _ := driver.Weather.Humidity(ctx, city, days)
						svg, _ := driver.ChartLine(values.Values, values.Labels, "%")
						return svg
					},
				}
			}()); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:186
				__e = (new(doors.Door)).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
//line dashboard.gox:186
					__e = __c.Any(func() any {
				days, ok := d.days.Effect(ctx)
				if !ok {
					return nil
				}
				return chart{
					title: "Weather",
					svg: func() []byte {
						values, _ := driver.Weather.Code(ctx, city, days)
						svg, _ := driver.ChartPie(values.Values)
						return svg
					},
				}
			}()); if __e != nil { return }
				return })); if __e != nil { return }
//line dashboard.gox:200
				__e = (new(doors.Door)).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
//line dashboard.gox:200
					__e = __c.Any(func() any {
				days, _ := d.days.Effect(ctx)
				units, ok := d.units.Effect(ctx)
				if !ok {
					return nil
				}
				return chart{
					title: "Wind Speed",
					svg: func() []byte {
						values, _ := driver.Weather.WindSpeed(ctx, city, units, days)
						svg, _ := driver.ChartLine(values.Values, values.Labels, units.WindSpeed())
						return svg
					},
				}
			}()); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:217
}

type chart struct {
	title string
	svg func() []byte
}

//line dashboard.gox:224
func (c chart) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("article"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("header"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line dashboard.gox:227
				__e = __c.Any(c.title); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.InitVoid("img"); if __e != nil { return }
			{
//line dashboard.gox:229
				__e = __c.Set("height", "auto"); if __e != nil { return }
//line dashboard.gox:229
				__e = __c.Set("width", "100%"); if __e != nil { return }
//line dashboard.gox:229
				__e = __c.Set("src", c.svg()); if __e != nil { return }
//line dashboard.gox:229
				__e = __c.Set("type", "image/svg+xml"); if __e != nil { return }
			}
			__e = __c.Submit(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line dashboard.gox:231
}
