// Managed by GoX v0.3.1

//line login.gox:1
package main

import (
	"context"
	"net/http"
	"time"
	
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
	"github.com/doors-dev/tutorial/driver"
)

const userLogin = "admin"
const userPassword = "password123"
const sessionDuration = time.Hour * 24

func Login(session doors.Source[driver.Session]) gox.Comp {
	return login{
		session: session,
		message: new(doors.Door),
	}
}

type login struct {
	session doors.Source[driver.Session]
	message *doors.Door
}

type loginData struct {
	Login string `form:"login"`
	Password string `form:"password"`
}

func (l login) submit(ctx context.Context, r doors.RequestForm[loginData]) bool {
	if r.Data().Login != userLogin || r.Data().Password != userPassword {
		l.message.Inner(ctx, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("p"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("mark"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("wrong password or login"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line login.gox:36
		return }))
		return false
	}
	l.message.Inner(ctx, nil)
	session := driver.Sessions.Add(r.Data().Login, sessionDuration)
	r.SetCookie(&http.Cookie{
		Name: "session",
		Value: session.Token,
		Expires: time.Now().Add(sessionDuration),
		Path: "/",
		HttpOnly: true,
	})
	doors.SessionExpire(ctx, sessionDuration)
	l.session.Update(ctx, session)
	return true
}

//line login.gox:53
func (l login) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Log In"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("h1"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Log In"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("form"); if __e != nil { return }
		{
//line login.gox:57
			__e = __c.Modify(doors.ASubmit[loginData]{
			Scope: new(doors.ScopeBlocking),
			Indicator: doors.IndicateAttrQuery("#login-submit", "aria-busy", "true"),
			On: l.submit,
		}); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("fieldset"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("label"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("Login"); if __e != nil { return }
					__e = __c.InitVoid("input"); if __e != nil { return }
					{
//line login.gox:66
						__e = __c.Set("name", "login"); if __e != nil { return }
//line login.gox:67
						__e = __c.Set("required", "true"); if __e != nil { return }
					}
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("label"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("Password"); if __e != nil { return }
					__e = __c.InitVoid("input"); if __e != nil { return }
					{
//line login.gox:72
						__e = __c.Set("type", "password"); if __e != nil { return }
//line login.gox:73
						__e = __c.Set("name", "password"); if __e != nil { return }
//line login.gox:74
						__e = __c.Set("required", "true"); if __e != nil { return }
					}
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line login.gox:76
				__e = __c.Any(l.message); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("button"); if __e != nil { return }
			{
//line login.gox:78
				__e = __c.Set("id", "login-submit"); if __e != nil { return }
//line login.gox:78
				__e = __c.Set("type", "submit"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Text("Log In"); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line login.gox:80
}
