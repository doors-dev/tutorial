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

func LocationSelector(apply func(ctx context.Context, city int)) gox.Comp {
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
		apply: apply,
		scope: new(doors.ScopeConcurrent),
	}
}

type locationSelector struct {
	country doors.Source[driver.Place]
	city doors.Source[driver.Place]
	apply func(ctx context.Context, city int)
	scope *doors.ScopeConcurrent
}

//line location_selector.gox:55
func (l locationSelector) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("article"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:58
				__e = __c.Any(placeSelector{
				title: "Country",
				options: new(doors.Door),
				search: driver.Locations.CountriesSearch,
				selected: l.country,
				scope: l.scope.Scope(1),
			}); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line location_selector.gox:66
			__e = __c.Any(l.country.Bind(l.selectCity)); if __e != nil { return }
//line location_selector.gox:67
			__e = __c.Any(l.city.Bind(l.submit)); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line location_selector.gox:69
}

//line location_selector.gox:71
func (l locationSelector) submit(city driver.Place) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line location_selector.gox:72
		if city.IsValid() {
			__e = __c.InitVoid("hr"); if __e != nil { return }
			{
			}
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("button"); if __e != nil { return }
			{
//line location_selector.gox:75
				__e = __c.Modify(doors.AClick{
				Indicator: doors.IndicateAttr("aria-busy", "true"),
				Scope: l.scope.Scope(0),
				On: func(ctx context.Context, _ doors.RequestPointer) bool {
					l.apply(ctx, city.Id)
					return true
				},
			}); if __e != nil { return }
//line location_selector.gox:83
				__e = __c.Set("id", "confirm"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Text("Confirm"); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line location_selector.gox:86
			__e = __c.Any(focus("confirm")); if __e != nil { return }
		}
	return })
//line location_selector.gox:88
}

//line location_selector.gox:90
func (l locationSelector) selectCity(country driver.Place) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line location_selector.gox:91
		if country.IsValid() {
			__e = __c.Init("section"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:93
				__e = __c.Any(placeSelector{
				title: "City",
				options: new(doors.Door),
				search: func(input string) ([]driver.Place, error) {
					return driver.Locations.CitiesSearch(country.Id, input)
				},
				scope: l.scope.Scope(1),
				selected: l.city,
			}); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
	return })
//line location_selector.gox:104
}

type placeSelector struct {
	title string
	options *doors.Door
	search func(input string) ([]driver.Place, error)
	selected doors.Source[driver.Place]
	scope doors.Scopes
}

//line location_selector.gox:114
func (l placeSelector) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line location_selector.gox:115
		__e = __c.Any(l.selected.Bind(func(p driver.Place) gox.Elem {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
//line location_selector.gox:116
			if p.IsValid() {
//line location_selector.gox:117
				__e = __c.Any(l.place(p)); if __e != nil { return }
			} else  {
//line location_selector.gox:119
				__e = __c.Any(l.input()); if __e != nil { return }
			}
		return })
//line location_selector.gox:121
	})); if __e != nil { return }
	return })
//line location_selector.gox:122
}

//line location_selector.gox:124
func (l placeSelector) place(p driver.Place) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("h3"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:125
			__e = __c.Any(l.title); if __e != nil { return }
			__e = __c.Text(": "); if __e != nil { return }
			__e = __c.Init("b"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:125
				__e = __c.Any(p.Name); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("button"); if __e != nil { return }
		{
//line location_selector.gox:127
			__e = __c.Modify(doors.AClick{
			Indicator: doors.IndicateAttr("aria-busy", "true"),
			Scope: new(doors.ScopeBlocking).And(l.scope),
			On: func(ctx context.Context, _ doors.RequestPointer) bool {
				l.selected.Update(ctx, driver.Place{})
				return true
			},
		}); if __e != nil { return }
//line location_selector.gox:135
			__e = __c.Set("class", "secondary"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Change"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line location_selector.gox:138
}

//line location_selector.gox:140
func (l placeSelector) input() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line location_selector.gox:142
		loaderID := "loader-" + doors.IDString(l.title)
	inputID := "input-" + doors.IDString(l.title)

		__e = __c.Init("h3"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Select "); if __e != nil { return }
//line location_selector.gox:145
			__e = __c.Many(l.title, " "); if __e != nil { return }
			__e = __c.Init("span"); if __e != nil { return }
			{
//line location_selector.gox:145
				__e = __c.Set("id", loaderID); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.InitVoid("input"); if __e != nil { return }
		{
//line location_selector.gox:147
			__e = __c.Set("id", inputID); if __e != nil { return }
//line location_selector.gox:148
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
//line location_selector.gox:159
			__e = __c.Set("type", "search"); if __e != nil { return }
//line location_selector.gox:160
			__e = __c.Set("placeholder", l.title); if __e != nil { return }
//line location_selector.gox:161
			__e = __c.Set("autocomplete", "off"); if __e != nil { return }
		}
		__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:162
		__e = __c.Any(focus(inputID)); if __e != nil { return }
//line location_selector.gox:164
		l.options.Inner(ctx, nil)

//line location_selector.gox:166
		__e = __c.Any(l.options); if __e != nil { return }
	return })
//line location_selector.gox:167
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
//line location_selector.gox:176
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
//line location_selector.gox:182
		return })
	}
	scope := new(doors.ScopeBlocking)
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ul"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:186
			for _, place := range results {
				__e = __c.Init("li"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("a"); if __e != nil { return }
					{
//line location_selector.gox:189
						__e = __c.Modify(doors.AClick{
						Scope: scope,
						PreventDefault: true,
						On: func(ctx context.Context, _ doors.RequestPointer) bool {
							l.selected.Update(ctx, place)
							return true
						},
					}); if __e != nil { return }
//line location_selector.gox:197
						__e = __c.Set("href", "#"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line location_selector.gox:198
						__e = __c.Any(place.Name); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
//line location_selector.gox:202
	return })
}

//line location_selector.gox:205
func focus(id string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("script"); if __e != nil { return }
		{
//line location_selector.gox:206
			__e = __c.Set("data:id", id); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Raw("const id = $data(\"id\")\n\t\tconst el = document.getElementById(id)\n\t\tel.focus()"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line location_selector.gox:211
}
