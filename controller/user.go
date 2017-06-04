package controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"fmt"

	"github.com/emurmotol/nmsrs.v4/constant"
	"github.com/emurmotol/nmsrs.v4/database"
	"github.com/emurmotol/nmsrs.v4/helper"
	"github.com/emurmotol/nmsrs.v4/lang"
	"github.com/emurmotol/nmsrs.v4/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := database.Conn()
	defer db.Close()
	query := db.Model(&model.User{}).Not("email", model.SuperuserEmail)

	if val, ok := r.URL.Query()["is_admin"]; ok {
		isAdmin, err := strconv.Atoi(val[0])

		if err == nil {
			query = query.Where("is_admin = ?", isAdmin)
		}
	}

	if val, ok := r.URL.Query()["q"]; ok {
		q := database.WrapLike(val[0])
		query = query.Where("name LIKE ? OR email LIKE ?", q, q)
	}
	query.Count(&count)
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		page = 1
	}

	pagination := &helper.Paginator{
		Page:     page,
		Limit:    limit,
		Count:    count,
		Interval: interval,
		QueryURL: r.URL.Query(),
	}

	if page > pagination.PageCount() {
		pagination.Page = 1
	}
	users := []model.User{}
	query.Offset(pagination.Offset()).Limit(limit).Find(&users)

	data := make(map[string]interface{})
	data["title"] = "Users"
	data["authUser"] = authUser(r)
	data["users"] = users
	data["q"] = r.URL.Query().Get("q")
	data["pagination"] = helper.Pager{
		Markup:     template.HTML(pagination.String()),
		Count:      pagination.Count,
		StartIndex: pagination.StartIndex(),
		EndIndex:   pagination.EndIndex(),
	}
	flashAlert := helper.GetFlash(w, r, "alert")

	if flashAlert != nil {
		alert := flashAlert.(helper.Alert)
		data["alert"] = template.HTML(alert.String())
	}
	rd.HTML(w, http.StatusOK, "user/index", data)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashAlert := helper.GetFlash(w, r, "alert")

	if flashAlert != nil {
		alert := flashAlert.(helper.Alert)
		data["alert"] = template.HTML(alert.String())
	}
	data["title"] = "Show User"
	data["user"] = r.Context().Value(constant.UserCtxKey).(*model.User)
	data["authUser"] = authUser(r)
	rd.HTML(w, http.StatusOK, "user/show", data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashAlert := helper.GetFlash(w, r, "alert")

	if flashAlert != nil {
		alert := flashAlert.(helper.Alert)
		data["alert"] = template.HTML(alert.String())
	}
	createUserForm := helper.GetFlash(w, r, "createUserForm")

	if createUserForm != nil {
		data["createUserForm"] = createUserForm.(model.CreateUserForm)
	}
	data["title"] = "Create User"
	data["authUser"] = authUser(r)
	rd.HTML(w, http.StatusOK, "user/create", data)
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		panic(err)
	}
	photoFile, photoHeader, err := r.FormFile("photo")

	if err != nil {
		if err != http.ErrMissingFile {
			panic(err)
		}
	}
	delete(r.PostForm, "photo")
	createUserForm := model.CreateUserForm{}

	if err := decoder.Decode(&createUserForm, r.PostForm); err != nil {
		panic(err)
	}
	createUserForm.PhotoFile = photoFile
	createUserForm.PhotoHeader = photoHeader

	if !createUserForm.IsValid() {
		helper.SetFlash(w, r, "createUserForm", createUserForm)
		CreateUser(w, r)
		return
	}
	user := model.User{
		Name:     createUserForm.Name,
		Email:    createUserForm.Email,
		Password: createUserForm.Password,
		IsAdmin:  createUserForm.IsAdmin,
	}
	newUser, err := user.Create()

	if err != nil {
		panic(err)
	}

	if createUserForm.PhotoFile != nil {
		if err := newUser.SetPhoto(createUserForm.PhotoFile); err != nil {
			panic(err)
		}
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "success",
		Content: fmt.Sprintf(lang.Get("user_success_create"), newUser.Name),
	})
	http.Redirect(w, r, "/users", http.StatusFound)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashAlert := helper.GetFlash(w, r, "alert")

	if flashAlert != nil {
		alert := flashAlert.(helper.Alert)
		data["alert"] = template.HTML(alert.String())
	}
	editProfileForm := helper.GetFlash(w, r, "editProfileForm")

	if editProfileForm != nil {
		data["editProfileForm"] = editProfileForm.(model.EditProfileForm)
	}
	passwordResetForm := helper.GetFlash(w, r, "passwordResetForm")

	if passwordResetForm != nil {
		data["passwordResetForm"] = passwordResetForm.(model.PasswordResetForm)
	}
	data["title"] = "Edit User"
	data["user"] = r.Context().Value(constant.UserCtxKey).(*model.User)
	data["authUser"] = authUser(r)
	rd.HTML(w, http.StatusOK, "user/edit", data)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		panic(err)
	}
	userCtx := r.Context().Value(constant.UserCtxKey).(*model.User)

	if r.PostFormValue("_method") == "PUT" {
		delete(r.PostForm, "_method")
		photoFile, photoHeader, err := r.FormFile("photo")

		if err != nil {
			if err != http.ErrMissingFile {
				panic(err)
			}
		}
		delete(r.PostForm, "photo")
		editProfileForm := model.EditProfileForm{}

		if err := decoder.Decode(&editProfileForm, r.PostForm); err != nil {
			panic(err)
		}
		editProfileForm.ID = userCtx.ID
		editProfileForm.PhotoFile = photoFile
		editProfileForm.PhotoHeader = photoHeader

		if !editProfileForm.IsValid() {
			helper.SetFlash(w, r, "editProfileForm", editProfileForm)
			EditUser(w, r)
			return
		}
		user := model.User{
			ID:      editProfileForm.ID,
			Name:    editProfileForm.Name,
			Email:   editProfileForm.Email,
			IsAdmin: editProfileForm.IsAdmin,
		}

		if err := user.UpdateProfile(); err != nil {
			panic(err)
		}

		if editProfileForm.PhotoFile != nil {
			if err := user.SetPhoto(editProfileForm.PhotoFile); err != nil {
				panic(err)
			}
		}
		helper.SetFlash(w, r, "alert", helper.Alert{
			Type:    "success",
			Content: fmt.Sprintf(lang.Get("user_success_update"), user.Name),
		})
		http.Redirect(w, r, fmt.Sprintf("/users/%d/edit", user.ID), http.StatusFound)
		return
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "danger",
		Content: lang.Get("method_invalid"),
	})
	http.Redirect(w, r, fmt.Sprintf("/users/%d/edit", userCtx.ID), http.StatusFound)
}

func UserPasswordReset(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	userCtx := r.Context().Value(constant.UserCtxKey).(*model.User)

	if r.PostFormValue("_method") == "PUT" {
		delete(r.PostForm, "_method")
		passwordResetForm := model.PasswordResetForm{}

		if err := decoder.Decode(&passwordResetForm, r.PostForm); err != nil {
			panic(err)
		}

		if !passwordResetForm.IsValid() {
			helper.SetFlash(w, r, "passwordResetForm", passwordResetForm)
			EditUser(w, r)
			return
		}
		user := model.User{
			ID:       userCtx.ID,
			Password: passwordResetForm.NewPassword,
		}

		if err := user.ResetPassword(); err != nil {
			panic(err)
		}
		helper.SetFlash(w, r, "alert", helper.Alert{
			Type:    "success",
			Content: fmt.Sprintf(lang.Get("password_success_update")),
		})
		http.Redirect(w, r, fmt.Sprintf("/users/%d/edit", user.ID), http.StatusFound)
		return
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "danger",
		Content: lang.Get("method_invalid"),
	})
	http.Redirect(w, r, fmt.Sprintf("/users/%d/edit", userCtx.ID), http.StatusFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	if r.PostFormValue("_method") == "DELETE" {
		userCtx := r.Context().Value(constant.UserCtxKey).(*model.User)

		if err := userCtx.Delete(); err != nil {
			panic(err)
		}

		helper.SetFlash(w, r, "alert", helper.Alert{
			Type:    "success",
			Content: fmt.Sprintf(lang.Get("user_success_delete"), userCtx.Name),
		})
		http.Redirect(w, r, "/users", http.StatusFound)
		return
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "danger",
		Content: lang.Get("method_invalid"),
	})
	http.Redirect(w, r, "/users", http.StatusFound)
}

func DeleteManyUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	if r.PostFormValue("_method") == "DELETE" {
		var ids []int64

		if err := json.Unmarshal([]byte(r.PostFormValue("ids")), &ids); err != nil {
			panic(err)
		}

		if err := model.DeleteManyUser(ids); err != nil {
			http.Redirect(w, r, "/users", http.StatusFound)
			return
		}
		helper.SetFlash(w, r, "alert", helper.Alert{
			Type:    "success",
			Content: fmt.Sprintf(lang.Get("user_success_delete"), fmt.Sprintf("%d users ", len(ids))),
		})
		http.Redirect(w, r, "/users", http.StatusFound)
		return
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "danger",
		Content: lang.Get("method_invalid"),
	})
	http.Redirect(w, r, "/users", http.StatusFound)
}

func UserPhoto(w http.ResponseWriter, r *http.Request) {
	userCtx := r.Context().Value(constant.UserCtxKey).(*model.User)
	http.ServeFile(w, r, userCtx.GetPhoto())
}

func UserEmailTaken(w http.ResponseWriter, r *http.Request) {
	if taken, _ := model.UserEmailTaken(r.URL.Query().Get("email")); taken {
		data := make(map[string]string)
		data["error"] = lang.Get("email_taken")
		rd.JSON(w, http.StatusNotFound, data)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UserEmailExists(w http.ResponseWriter, r *http.Request) {
	if taken, _ := model.UserEmailTaken(r.URL.Query().Get("email")); !taken {
		data := make(map[string]string)
		data["error"] = lang.Get("email_not_recognized")
		rd.JSON(w, http.StatusNotFound, data)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UserEmailCheck(w http.ResponseWriter, r *http.Request) {
	userCtx := r.Context().Value(constant.UserCtxKey).(*model.User)

	if same, _ := model.UserEmailSameAsOld(userCtx.ID, r.URL.Query().Get("email")); !same {
		if taken, _ := model.UserEmailTaken(r.URL.Query().Get("email")); taken {
			data := make(map[string]string)
			data["error"] = lang.Get("email_taken")
			rd.JSON(w, http.StatusNotFound, data)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
