package controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"fmt"

	"github.com/emurmotol/nmsrs/constant"
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/lang"
	"github.com/emurmotol/nmsrs/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	query := []bson.M{}
	query = append(query, bson.M{
		"email": bson.M{"$ne": model.SuperUserEmail},
	})

	if val, ok := r.URL.Query()["isAdmin"]; ok {
		isAdmin, err := strconv.ParseBool(val[0])

		if err != nil {
			panic(err)
		}
		query = append(query, bson.M{"isAdmin": isAdmin})
	}

	if val, ok := r.URL.Query()["q"]; ok {
		regex := bson.M{"$regex": bson.RegEx{Pattern: val[0], Options: "i"}}
		query = append(query, bson.M{
			"$or": []bson.M{
				bson.M{"name": regex},
				bson.M{"email": regex},
			},
		})
	}
	count, _ := db.C("users").Find(bson.M{"$and": query}).Count()
	defer db.Close()
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		page = 1
	}

	pagination := &helper.Paginator{
		Page:     page,
		Limit:    limit,
		Count:    count,
		Interval: interval,
		QueryUrl: r.URL.Query(),
	}

	if page > pagination.PageCount() {
		pagination.Page = 1
	}
	users := []model.User{}
	db.C("users").Find(bson.M{"$and": query}).Sort("-createdAt").Skip(pagination.Offset()).Limit(limit).All(&users)

	data := make(map[string]interface{})
	data["title"] = "Users"
	data["auth"] = model.Auth(r)
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
	data["auth"] = model.Auth(r)
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
	data["auth"] = model.Auth(r)
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
	createUserForm := new(model.CreateUserForm)

	if err := decoder.Decode(createUserForm, r.PostForm); err != nil {
		panic(err)
	}
	createUserForm.PhotoFile = photoFile
	createUserForm.PhotoHeader = photoHeader

	if !createUserForm.IsValid() {
		helper.SetFlash(w, r, "createUserForm", createUserForm)
		CreateUser(w, r)
		return
	}
	user := &model.User{
		Name:     createUserForm.Name,
		Email:    createUserForm.Email,
		Password: createUserForm.Password,
		IsAdmin:  createUserForm.IsAdmin,
	}
	newUser := user.Create()

	if createUserForm.PhotoFile != nil {
		if err := newUser.SetPhoto(createUserForm.PhotoFile); err != nil {
			panic(err)
		}
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "success",
		Content: fmt.Sprintf(lang.Get("userSuccessCreate"), newUser.Name),
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
	data["auth"] = model.Auth(r)
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
		editProfileForm := new(model.EditProfileForm)

		if err := decoder.Decode(editProfileForm, r.PostForm); err != nil {
			panic(err)
		}
		editProfileForm.HexId = userCtx.Id.Hex()
		editProfileForm.PhotoFile = photoFile
		editProfileForm.PhotoHeader = photoHeader

		if !editProfileForm.IsValid() {
			helper.SetFlash(w, r, "editProfileForm", editProfileForm)
			EditUser(w, r)
			return
		}
		user := &model.User{
			Id:      bson.ObjectIdHex(editProfileForm.HexId),
			Name:    editProfileForm.Name,
			Email:   editProfileForm.Email,
			IsAdmin: editProfileForm.IsAdmin,
		}
		user.UpdateProfile()

		if editProfileForm.PhotoFile != nil {
			if err := user.SetPhoto(editProfileForm.PhotoFile); err != nil {
				panic(err)
			}
		}
		helper.SetFlash(w, r, "alert", helper.Alert{
			Type:    "success",
			Content: fmt.Sprintf(lang.Get("userSuccessUpdate"), user.Name),
		})
		http.Redirect(w, r, fmt.Sprintf("/users/%s/edit", user.Id.Hex()), http.StatusFound)
		return
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "danger",
		Content: lang.Get("methodInvalid"),
	})
	http.Redirect(w, r, fmt.Sprintf("/users/%s/edit", userCtx.Id.Hex()), http.StatusFound)
}

func UserPasswordReset(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	userCtx := r.Context().Value(constant.UserCtxKey).(*model.User)

	if r.PostFormValue("_method") == "PUT" {
		delete(r.PostForm, "_method")
		passwordResetForm := new(model.PasswordResetForm)

		if err := decoder.Decode(passwordResetForm, r.PostForm); err != nil {
			panic(err)
		}

		if !passwordResetForm.IsValid() {
			helper.SetFlash(w, r, "passwordResetForm", passwordResetForm)
			EditUser(w, r)
			return
		}
		user := &model.User{
			Id:       userCtx.Id,
			Password: passwordResetForm.NewPassword,
		}
		user.ResetPassword()
		helper.SetFlash(w, r, "alert", helper.Alert{
			Type:    "success",
			Content: fmt.Sprintf(lang.Get("passwordSuccessUpdate")),
		})
		http.Redirect(w, r, fmt.Sprintf("/users/%s/edit", user.Id.Hex()), http.StatusFound)
		return
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "danger",
		Content: lang.Get("methodInvalid"),
	})
	http.Redirect(w, r, fmt.Sprintf("/users/%s/edit", userCtx.Id.Hex()), http.StatusFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	if r.PostFormValue("_method") == "DELETE" {
		userCtx := r.Context().Value(constant.UserCtxKey).(*model.User)
		userCtx.Delete()
		helper.SetFlash(w, r, "alert", helper.Alert{
			Type:    "success",
			Content: fmt.Sprintf(lang.Get("userSuccessDelete"), userCtx.Name),
		})
		http.Redirect(w, r, "/users", http.StatusFound)
		return
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "danger",
		Content: lang.Get("methodInvalid"),
	})
	http.Redirect(w, r, "/users", http.StatusFound)
}

func DeleteManyUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	if r.PostFormValue("_method") == "DELETE" {
		var hexIds []string

		if err := json.Unmarshal([]byte(r.PostFormValue("userHexIds")), &hexIds); err != nil {
			panic(err)
		}
		model.DeleteManyUser(hexIds)
		helper.SetFlash(w, r, "alert", helper.Alert{
			Type:    "success",
			Content: fmt.Sprintf(lang.Get("userSuccessDelete"), fmt.Sprintf("%d users ", len(hexIds))),
		})
		http.Redirect(w, r, "/users", http.StatusFound)
		return
	}
	helper.SetFlash(w, r, "alert", helper.Alert{
		Type:    "danger",
		Content: lang.Get("methodInvalid"),
	})
	http.Redirect(w, r, "/users", http.StatusFound)
}

func UserPhoto(w http.ResponseWriter, r *http.Request) {
	userCtx := r.Context().Value(constant.UserCtxKey).(*model.User)
	http.ServeFile(w, r, userCtx.GetPhoto())
}

func UserEmailTaken(w http.ResponseWriter, r *http.Request) {
	if taken := model.UserEmailTaken(r.URL.Query().Get("email")); taken {
		data := make(map[string]string)
		data["error"] = lang.Get("emailTaken")
		rd.JSON(w, http.StatusNotFound, data)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UserEmailExists(w http.ResponseWriter, r *http.Request) {
	if taken := model.UserEmailTaken(r.URL.Query().Get("email")); !taken {
		data := make(map[string]string)
		data["error"] = lang.Get("emailNotRecognized")
		rd.JSON(w, http.StatusNotFound, data)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UserEmailCheck(w http.ResponseWriter, r *http.Request) {
	userCtx := r.Context().Value(constant.UserCtxKey).(*model.User)

	if same := model.UserEmailSameAsOld(userCtx.Id, r.URL.Query().Get("email")); !same {
		if taken := model.UserEmailTaken(r.URL.Query().Get("email")); taken {
			data := make(map[string]string)
			data["error"] = lang.Get("emailTaken")
			rd.JSON(w, http.StatusNotFound, data)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
