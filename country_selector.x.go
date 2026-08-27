// Managed by GoX v0.3.1

//line country_selector.gox:1
package main

import (
	"context"
	"time"
	
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
	"github.com/doors-dev/tutorial/driver"
)

func CountrySelector() gox.Comp {
	return countrySelector{
		options: new(doors.Door),
		selected: doors.NewSource(driver.Place{}),
	}
}

type countrySelector struct {
	options *doors.Door
	selected doors.Source[driver.Place]
}

//line country_selector.gox:24
func (l countrySelector) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line country_selector.gox:25
		__e = __c.Any(l.selected.Bind(func(p driver.Place) gox.Elem {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
//line country_selector.gox:26
			if p.IsValid() {
//line country_selector.gox:27
				__e = __c.Any(l.place(p)); if __e != nil { return }
			} else  {
//line country_selector.gox:29
				__e = __c.Any(l.input()); if __e != nil { return }
			}
		return })
//line country_selector.gox:31
	})); if __e != nil { return }
	return })
//line country_selector.gox:32
}

//line country_selector.gox:34
func (l countrySelector) input() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("h3"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Select Country "); if __e != nil { return }
			__e = __c.Init("span"); if __e != nil { return }
			{
//line country_selector.gox:35
				__e = __c.Set("id", "search-loader"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.InitVoid("input"); if __e != nil { return }
		{
//line country_selector.gox:37
			__e = __c.Modify(doors.AInput{
			Scope: &doors.ScopeDebounce{
				Duration: 300 * time.Millisecond,
				Limit: 600 * time.Millisecond,
			},
			Indicator: doors.IndicateAttrQuery("#search-loader", "aria-busy", "true"),
			On: func(ctx context.Context, r doors.RequestInput) bool {
				input := r.Event().Value
				l.options.Inner(ctx, l.results(input))
				return false
			},
		}); if __e != nil { return }
//line country_selector.gox:49
			__e = __c.Set("type", "search"); if __e != nil { return }
//line country_selector.gox:50
			__e = __c.Set("placeholder", "Country"); if __e != nil { return }
//line country_selector.gox:51
			__e = __c.Set("autocomplete", "off"); if __e != nil { return }
		}
		__e = __c.Submit(); if __e != nil { return }
//line country_selector.gox:53
		l.options.Inner(ctx, nil)

//line country_selector.gox:55
		__e = __c.Any(l.options); if __e != nil { return }
	return })
//line country_selector.gox:56
}

//line country_selector.gox:58
func (l countrySelector) place(p driver.Place) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("h3"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Country: "); if __e != nil { return }
			__e = __c.Init("b"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line country_selector.gox:59
				__e = __c.Any(p.Name); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("button"); if __e != nil { return }
		{
//line country_selector.gox:61
			__e = __c.Modify(doors.AClick{
			Indicator: doors.IndicateAttr("aria-busy", "true"),
			Scope: new(doors.ScopeBlocking),
			On: func(ctx context.Context, _ doors.RequestPointer) bool {
				l.selected.Update(ctx, driver.Place{})
				return true
			},
		}); if __e != nil { return }
//line country_selector.gox:69
			__e = __c.Set("class", "secondary"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Change"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line country_selector.gox:72
}

func (l countrySelector) results(input string) gox.Elem {
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
//line country_selector.gox:81
		return })
	}
	results, _ := driver.Locations.CountriesSearch(input)
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
//line country_selector.gox:87
		return })
	}
	scope := new(doors.ScopeBlocking)
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ul"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line country_selector.gox:91
			for _, place := range results {
				__e = __c.Init("li"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("a"); if __e != nil { return }
					{
//line country_selector.gox:94
						__e = __c.Modify(doors.AClick{
						Scope: scope,
						PreventDefault: true,
						On: func(ctx context.Context, _ doors.RequestPointer) bool {
							l.selected.Update(ctx, place)
							return true
						},
					}); if __e != nil { return }
//line country_selector.gox:102
						__e = __c.Set("href", "#"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line country_selector.gox:103
						__e = __c.Any(place.Name); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
//line country_selector.gox:107
	return })
}
