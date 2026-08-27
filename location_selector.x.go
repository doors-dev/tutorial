// Managed by GoX v0.3.1

//line location_selector.gox:1
package main

import (
	"context"
	"time"
	
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
	"github.com/doors-dev/tutorial/driver"
)

type location struct {
	country driver.Place
	city driver.Place
}

func LocationSelector() gox.Comp {
	loc := doors.NewSource(location{})
	country := doors.DeriveSource(loc,
		func(l location) driver.Place {
			return l.country
		},
		// Propagate country changes back to location. The previous location
		// is dropped on purpose: a new country invalidates the selected city.
		func(_ location, c driver.Place) location {
			return location{
				country: c,
			}
		},
	)
	city := doors.DeriveSource(loc,
		func(l location) driver.Place {
			return l.city
		},
		func(l location, c driver.Place) location {
			l.city = c
			return l
		},
	)
	return locationSelector{
		city: city,
		country: country,
	}
}

type locationSelector struct {
	country doors.Source[driver.Place]
	city doors.Source[driver.Place]
}

//line location_selector.gox:51
func (l locationSelector) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("article"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:54
				__e = __c.Any(placeSelector{
				title: "Country",
				options: new(doors.Door),
				search: driver.Locations.CountriesSearch,
				selected: l.country,
			}); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line location_selector.gox:61
			__e = __c.Any(l.country.Bind(l.selectCity)); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line location_selector.gox:63
}

//line location_selector.gox:65
func (l locationSelector) selectCity(country driver.Place) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line location_selector.gox:66
		if country.IsValid() {
			__e = __c.Init("section"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:68
				__e = __c.Any(placeSelector{
				title: "City",
				options: new(doors.Door),
				search: func(input string) ([]driver.Place, error) {
					return driver.Locations.CitiesSearch(country.Id, input)
				},
				selected: l.city,
			}); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
	return })
//line location_selector.gox:78
}

type placeSelector struct {
	title string
	options *doors.Door
	search func(input string) ([]driver.Place, error)
	selected doors.Source[driver.Place]
}

//line location_selector.gox:87
func (l placeSelector) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line location_selector.gox:88
		__e = __c.Any(l.selected.Bind(func(p driver.Place) gox.Elem {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
//line location_selector.gox:89
			if p.IsValid() {
//line location_selector.gox:90
				__e = __c.Any(l.place(p)); if __e != nil { return }
			} else  {
//line location_selector.gox:92
				__e = __c.Any(l.input()); if __e != nil { return }
			}
		return })
//line location_selector.gox:94
	})); if __e != nil { return }
	return })
//line location_selector.gox:95
}

//line location_selector.gox:97
func (l placeSelector) place(p driver.Place) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("h3"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:98
			__e = __c.Any(l.title); if __e != nil { return }
			__e = __c.Text(": "); if __e != nil { return }
			__e = __c.Init("b"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:98
				__e = __c.Any(p.Name); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("button"); if __e != nil { return }
		{
//line location_selector.gox:100
			__e = __c.Modify(doors.AClick{
			Indicator: doors.IndicateAttr("aria-busy", "true"),
			Scope: new(doors.ScopeBlocking),
			On: func(ctx context.Context, _ doors.RequestPointer) bool {
				l.selected.Update(ctx, driver.Place{})
				return true
			},
		}); if __e != nil { return }
//line location_selector.gox:108
			__e = __c.Set("class", "secondary"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Change"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line location_selector.gox:111
}

//line location_selector.gox:113
func (l placeSelector) input() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line location_selector.gox:115
		loaderID := "loader-" + doors.IDString(l.title)
	inputID := "input-" + doors.IDString(l.title)

		__e = __c.Init("h3"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Select "); if __e != nil { return }
//line location_selector.gox:118
			__e = __c.Many(l.title, " "); if __e != nil { return }
			__e = __c.Init("span"); if __e != nil { return }
			{
//line location_selector.gox:118
				__e = __c.Set("id", loaderID); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.InitVoid("input"); if __e != nil { return }
		{
//line location_selector.gox:120
			__e = __c.Set("id", inputID); if __e != nil { return }
//line location_selector.gox:121
			__e = __c.Modify(doors.AInput{
			Scope: &doors.ScopeDebounce{
				Duration: 300 * time.Millisecond,
				Limit: 600 * time.Millisecond,
			},
			Indicator: doors.IndicateAttrQuery("#" + loaderID, "aria-busy", "true"),
			On: func(ctx context.Context, r doors.RequestInput) bool {
				l.options.Inner(ctx, l.results(r.Event().Value))
				return false
			},
		}); if __e != nil { return }
//line location_selector.gox:132
			__e = __c.Set("type", "search"); if __e != nil { return }
//line location_selector.gox:133
			__e = __c.Set("placeholder", l.title); if __e != nil { return }
//line location_selector.gox:134
			__e = __c.Set("autocomplete", "off"); if __e != nil { return }
		}
		__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:135
		__e = __c.Any(focus(inputID)); if __e != nil { return }
//line location_selector.gox:137
		l.options.Inner(ctx, nil)

//line location_selector.gox:139
		__e = __c.Any(l.options); if __e != nil { return }
	return })
//line location_selector.gox:140
}

func (l placeSelector) results(input string) gox.Elem {
	if len(input) == 0 {
		return nil
	}
	if len(input) < 2 {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("p"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("mark"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("Type at least two letters to search"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line location_selector.gox:149
		return })
	}
	results, _ := l.search(input)
	if len(results) == 0 {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("p"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("i"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("nothing found"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line location_selector.gox:155
		return })
	}
	scope := new(doors.ScopeBlocking)
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ul"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:159
			for _, place := range results {
				__e = __c.Init("li"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("a"); if __e != nil { return }
					{
//line location_selector.gox:162
						__e = __c.Modify(doors.AClick{
						Scope: scope,
						PreventDefault: true,
						On: func(ctx context.Context, _ doors.RequestPointer) bool {
							l.selected.Update(ctx, place)
							return true
						},
					}); if __e != nil { return }
//line location_selector.gox:170
						__e = __c.Set("href", "#"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:171
						__e = __c.Any(place.Name); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
//line location_selector.gox:175
	return })
}

//line location_selector.gox:178
func focus(id string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("script"); if __e != nil { return }
		{
//line location_selector.gox:179
			__e = __c.Set("data:id", id); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Raw("const id = $data(\"id\")\n\t\tconst el = document.getElementById(id)\n\t\tel.focus()"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line location_selector.gox:184
}
