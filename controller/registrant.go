package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/emurmotol/nmsrs/database"
	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/model"
	"github.com/unrolled/render"
)

func GetRegistrants(w http.ResponseWriter, r *http.Request) {
	db := database.Conn()
	defer db.Close()

	query := db.Model(&model.Registrant{})
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
	registrants := []model.Registrant{}
	query.Offset(pagination.Offset()).Limit(limit).Find(&registrants)

	data := make(map[string]interface{})
	data["title"] = "Registrants"
	data["authUser"] = authUser(r)
	data["registrants"] = registrants
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
	rd.HTML(w, http.StatusOK, "registrant/index", data)
}

// func ShowRegistrant(w http.ResponseWriter, r *http.Request) {
// 	data := make(map[string]interface{})
// 	flashAlert := helper.GetFlash(w, r, "alert")

// 	if flashAlert != nil {
// 		alert := flashAlert.(helper.Alert)
// 		data["alert"] = template.HTML(alert.String())
// 	}
// 	data["title"] = "Show Registrant"
// 	data["registrant"] = r.Context().Value(constant.RegistrantCtxKey).(*model.Registrant)
// 	data["authUser"] = authUser(r)
// 	rd.HTML(w, http.StatusOK, "registrant/show", data)
// }

func CreateRegistrant(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashAlert := helper.GetFlash(w, r, "alert")

	if flashAlert != nil {
		alert := flashAlert.(helper.Alert)
		data["alert"] = template.HTML(alert.String())
	}
	// createRegistrantForm := helper.GetFlash(w, r, "createRegistrantForm")

	// if createRegistrantForm != nil {
	// 	data["createRegistrantForm"] = createRegistrantForm.(model.CreateRegistrantForm)
	// }
	data["title"] = "Create Registrant"
	data["authUser"] = authUser(r)
	rd.HTML(w, http.StatusOK, "registrant/create", data, render.HTMLOptions{Layout: "layouts/wizard"})
}

// func StoreRegistrant(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseMultipartForm(0); err != nil {
// 		panic(err)
// 	}
// 	photoFile, photoHeader, err := r.FormFile("photo")

// 	if err != nil {
// 		if err != http.ErrMissingFile {
// 			panic(err)
// 		}
// 	}
// 	delete(r.PostForm, "photo")
// 	createRegistrantForm := model.CreateRegistrantForm{}

// 	if err := decoder.Decode(&createRegistrantForm, r.PostForm); err != nil {
// 		panic(err)
// 	}
// 	createRegistrantForm.PhotoFile = photoFile
// 	createRegistrantForm.PhotoHeader = photoHeader

// 	if !createRegistrantForm.IsValid() {
// 		helper.SetFlash(w, r, "createRegistrantForm", createRegistrantForm)
// 		CreateRegistrant(w, r)
// 		return
// 	}
// 	registrant := model.Registrant{
// 		Name:     createRegistrantForm.Name,
// 		Email:    createRegistrantForm.Email,
// 		Password: createRegistrantForm.Password,
// 		IsAdmin:  createRegistrantForm.IsAdmin,
// 	}
// 	newRegistrant, err := registrant.Create()

// 	if err != nil {
// 		panic(err)
// 	}

// 	if createRegistrantForm.PhotoFile != nil {
// 		if err := newRegistrant.SetPhoto(createRegistrantForm.PhotoFile); err != nil {
// 			panic(err)
// 		}
// 	}
// 	helper.SetFlash(w, r, "alert", helper.Alert{
// 		Type:    "success",
// 		Content: fmt.Sprintf(lang.Get("registrant_success_create"), newRegistrant.Name),
// 	})
// 	http.Redirect(w, r, "/registrants", http.StatusFound)
// }

// func EditRegistrant(w http.ResponseWriter, r *http.Request) {
// 	data := make(map[string]interface{})
// 	flashAlert := helper.GetFlash(w, r, "alert")

// 	if flashAlert != nil {
// 		alert := flashAlert.(helper.Alert)
// 		data["alert"] = template.HTML(alert.String())
// 	}
// 	editProfileForm := helper.GetFlash(w, r, "editProfileForm")

// 	if editProfileForm != nil {
// 		data["editProfileForm"] = editProfileForm.(model.EditProfileForm)
// 	}
// 	passwordResetForm := helper.GetFlash(w, r, "passwordResetForm")

// 	if passwordResetForm != nil {
// 		data["passwordResetForm"] = passwordResetForm.(model.PasswordResetForm)
// 	}
// 	data["title"] = "Edit Registrant"
// 	data["registrant"] = r.Context().Value(constant.RegistrantCtxKey).(*model.Registrant)
// 	data["authUser"] = authUser(r)
// 	rd.HTML(w, http.StatusOK, "registrant/edit", data)
// }

// func UpdateRegistrant(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseMultipartForm(0); err != nil {
// 		panic(err)
// 	}
// 	registrantCtx := r.Context().Value(constant.RegistrantCtxKey).(*model.Registrant)

// 	if r.PostFormValue("_method") == "PUT" {
// 		delete(r.PostForm, "_method")
// 		photoFile, photoHeader, err := r.FormFile("photo")

// 		if err != nil {
// 			if err != http.ErrMissingFile {
// 				panic(err)
// 			}
// 		}
// 		delete(r.PostForm, "photo")
// 		editProfileForm := model.EditProfileForm{}

// 		if err := decoder.Decode(&editProfileForm, r.PostForm); err != nil {
// 			panic(err)
// 		}
// 		editProfileForm.ID = registrantCtx.ID
// 		editProfileForm.PhotoFile = photoFile
// 		editProfileForm.PhotoHeader = photoHeader

// 		if !editProfileForm.IsValid() {
// 			helper.SetFlash(w, r, "editProfileForm", editProfileForm)
// 			EditRegistrant(w, r)
// 			return
// 		}
// 		registrant := model.Registrant{
// 			ID:      editProfileForm.ID,
// 			Name:    editProfileForm.Name,
// 			Email:   editProfileForm.Email,
// 			IsAdmin: editProfileForm.IsAdmin,
// 		}

// 		if err := registrant.UpdateRegistrant(); err != nil {
// 			panic(err)
// 		}

// 		if editProfileForm.PhotoFile != nil {
// 			if err := registrant.SetPhoto(editProfileForm.PhotoFile); err != nil {
// 				panic(err)
// 			}
// 		}
// 		helper.SetFlash(w, r, "alert", helper.Alert{
// 			Type:    "success",
// 			Content: fmt.Sprintf(lang.Get("registrant_success_update"), registrant.Name),
// 		})
// 		http.Redirect(w, r, fmt.Sprintf("/registrants/%d/edit", registrant.ID), http.StatusFound)
// 		return
// 	}
// 	helper.SetFlash(w, r, "alert", helper.Alert{
// 		Type:    "danger",
// 		Content: lang.Get("method_invalid"),
// 	})
// 	http.Redirect(w, r, fmt.Sprintf("/registrants/%d/edit", registrantCtx.ID), http.StatusFound)
// }

// func DeleteRegistrant(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		panic(err)
// 	}

// 	if r.PostFormValue("_method") == "DELETE" {
// 		registrantCtx := r.Context().Value(constant.RegistrantCtxKey).(*model.Registrant)

// 		if err := registrantCtx.Delete(); err != nil {
// 			panic(err)
// 		}

// 		helper.SetFlash(w, r, "alert", helper.Alert{
// 			Type:    "success",
// 			Content: fmt.Sprintf(lang.Get("registrant_success_delete"), registrantCtx.Name),
// 		})
// 		http.Redirect(w, r, "/registrants", http.StatusFound)
// 		return
// 	}
// 	helper.SetFlash(w, r, "alert", helper.Alert{
// 		Type:    "danger",
// 		Content: lang.Get("method_invalid"),
// 	})
// 	http.Redirect(w, r, "/registrants", http.StatusFound)
// }

// func DeleteManyRegistrant(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		panic(err)
// 	}

// 	if r.PostFormValue("_method") == "DELETE" {
// 		var ids []int64

// 		if err := json.Unmarshal([]byte(r.PostFormValue("ids")), &ids); err != nil {
// 			panic(err)
// 		}

// 		if err := model.DeleteManyRegistrant(ids); err != nil {
// 			http.Redirect(w, r, "/registrants", http.StatusFound)
// 			return
// 		}
// 		helper.SetFlash(w, r, "alert", helper.Alert{
// 			Type:    "success",
// 			Content: fmt.Sprintf(lang.Get("registrant_success_delete"), fmt.Sprintf("%d registrants ", len(ids))),
// 		})
// 		http.Redirect(w, r, "/registrants", http.StatusFound)
// 		return
// 	}
// 	helper.SetFlash(w, r, "alert", helper.Alert{
// 		Type:    "danger",
// 		Content: lang.Get("method_invalid"),
// 	})
// 	http.Redirect(w, r, "/registrants", http.StatusFound)
// }

// func RegistrantPhoto(w http.ResponseWriter, r *http.Request) {
// 	registrantCtx := r.Context().Value(constant.RegistrantCtxKey).(*model.Registrant)
// 	http.ServeFile(w, r, registrantCtx.GetPhoto())
// }

// func RegistrantEmailTaken(w http.ResponseWriter, r *http.Request) {
// 	if taken, _ := model.RegistrantEmailTaken(r.URL.Query().Get("email")); taken {
// 		data := make(map[string]string)
// 		data["error"] = lang.Get("email_taken")
// 		rd.JSON(w, http.StatusNotFound, data)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

// func RegistrantEmailCheck(w http.ResponseWriter, r *http.Request) {
// 	registrantCtx := r.Context().Value(constant.RegistrantCtxKey).(*model.Registrant)

// 	if same, _ := model.RegistrantEmailSameAsOld(registrantCtx.ID, r.URL.Query().Get("email")); !same {
// 		if taken, _ := model.RegistrantEmailTaken(r.URL.Query().Get("email")); taken {
// 			data := make(map[string]string)
// 			data["error"] = lang.Get("email_taken")
// 			rd.JSON(w, http.StatusNotFound, data)
// 			return
// 		}
// 	}
// 	w.WriteHeader(http.StatusOK)
// }
