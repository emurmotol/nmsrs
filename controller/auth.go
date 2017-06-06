package controller

import (
	"html/template"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/emurmotol/nmsrs/constant"
	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/lang"
	"github.com/emurmotol/nmsrs/model"
	"github.com/goware/jwtauth"
	"github.com/unrolled/render"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashAlert := helper.GetFlash(w, r, "alert")

	if flashAlert != nil {
		alert := flashAlert.(helper.Alert)
		data["alert"] = template.HTML(alert.String())
	}
	loginForm := helper.GetFlash(w, r, "loginForm")

	if loginForm != nil {
		data["loginForm"] = loginForm.(model.LoginForm)
	}
	data["title"] = "Login"
	rd.HTML(w, http.StatusOK, "auth/login", data, render.HTMLOptions{Layout: "layouts/auth"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	loginForm := model.LoginForm{}

	if err := decoder.Decode(&loginForm, r.PostForm); err != nil {
		panic(err)
	}

	if !loginForm.IsValid() {
		helper.SetFlash(w, r, "loginForm", loginForm)
		GetLogin(w, r)
		return
	}
	user := authenticate(loginForm.Email, loginForm.Password)

	if user == nil {
		helper.SetFlash(w, r, "loginForm", loginForm)
		helper.SetFlash(w, r, "alert", helper.Alert{
			Type:    "danger",
			Content: lang.Get("wrong_credentials"),
		})
		GetLogin(w, r)
		return
	}
	claims := jwtauth.Claims{}
	claims.SetIssuedNow()
	setExpiryDuration, _ := env.Conf.Int("pkg.jwtauth.setExpiryDuration")
	claims.SetExpiry(time.Now().Add(time.Hour * time.Duration(setExpiryDuration))) // 2 weeks
	claims["userID"] = user.ID
	tokenAuth := r.Context().Value(constant.TokenAuthCtxKey).(*jwtauth.JwtAuth)
	_, tokenString, _ := tokenAuth.Encode(claims)
	tokenName, _ := env.Conf.String("pkg.jwtauth.tokenName")

	http.SetCookie(w, &http.Cookie{
		Name:       tokenName,
		Value:      tokenString,
		Path:       "/",
		RawExpires: "0",
	})

	if val, ok := r.URL.Query()["redirect"]; ok && val[0] != "" {
		http.Redirect(w, r, val[0], http.StatusFound)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenName, _ := env.Conf.String("pkg.jwtauth.tokenName")

	http.SetCookie(w, &http.Cookie{
		Name:   tokenName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	// todo: delete sessions
	http.Redirect(w, r, "/login", http.StatusFound)
}

func authenticate(email, password string) *model.User {
	user := model.UserByEmail(email)

	if user == nil {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil
	}
	return user
}

func authUser(r *http.Request) *model.User {
	authUserCtxVal := r.Context().Value(constant.AuthUserCtxKey)

	if &authUserCtxVal != nil {
		return authUserCtxVal.(*model.User)
	}
	return nil
}
