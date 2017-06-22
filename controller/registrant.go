package controller

// import (
// 	"html/template"
// 	"net/http"
// 	"strconv"

// 	"github.com/emurmotol/nmsrs/helper"
// 	"github.com/emurmotol/nmsrs/lang"
// 	"github.com/emurmotol/nmsrs/model"
// 	"github.com/unrolled/render"
// )

// func GetRegistrants(w http.ResponseWriter, r *http.Request) {
// 	db := database.Con()
// 	defer db.Close()

// 	query := db.Model(&model.Registrant{})
// 	query.Count(&count)
// 	page, err := strconv.Atoi(r.URL.Query().Get("page"))

// 	if err != nil {
// 		page = 1
// 	}

// 	pagination := &helper.Paginator{
// 		Page:     page,
// 		Limit:    limit,
// 		Count:    count,
// 		Interval: interval,
// 		QueryURL: r.URL.Query(),
// 	}

// 	if page > pagination.PageCount() {
// 		pagination.Page = 1
// 	}
// 	registrants := []model.Registrant{}
// 	query.Offset(pagination.Offset()).Limit(limit).Find(&registrants)

// 	data := make(map[string]interface{})
// 	data["title"] = "Registrants"
// 	data["authUser"] = authUser(r)
// 	data["registrants"] = registrants
// 	data["q"] = r.URL.Query().Get("q")
// 	data["pagination"] = helper.Pager{
// 		Markup:     template.HTML(pagination.String()),
// 		Count:      pagination.Count,
// 		StartIndex: pagination.StartIndex(),
// 		EndIndex:   pagination.EndIndex(),
// 	}
// 	flashAlert := helper.GetFlash(w, r, "alert")

// 	if flashAlert != nil {
// 		alert := flashAlert.(helper.Alert)
// 		data["alert"] = template.HTML(alert.String())
// 	}
// 	rd.HTML(w, http.StatusOK, "registrant/index", data)
// }

// func CreateRegistrant(w http.ResponseWriter, r *http.Request) {
// 	data := make(map[string]interface{})
// 	flashAlert := helper.GetFlash(w, r, "alert")

// 	if flashAlert != nil {
// 		alert := flashAlert.(helper.Alert)
// 		data["alert"] = template.HTML(alert.String())
// 	}
// 	createRegistrantForm := helper.GetFlash(w, r, "createRegistrantForm")

// 	if createRegistrantForm != nil {
// 		data["createRegistrantForm"] = createRegistrantForm.(model.CreateRegistrantForm)
// 	}
// 	data["civilStats"] = model.CivilStats()
// 	data["sexes"] = model.Sexes()
// 	data["empStats"] = model.EmpStats()
// 	data["disabilities"] = model.Disabilities()
// 	data["title"] = "Create Registrant"
// 	data["authUser"] = authUser(r)
// 	rd.HTML(w, http.StatusOK, "registrant/create", data, render.HTMLOptions{Layout: "layouts/wizard"})
// }

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
// 		RegisteredAt: createRegistrantForm.RegisteredAt,
// 		IAccept:      createRegistrantForm.IAccept,
// 	}
// 	newRegistrant := registrant.Create()
// 	hasPhoto := false

// 	if createRegistrantForm.PhotoFile != nil {
// 		hasPhoto = true
// 	}
// 	registInfo := model.RegistInfo{
// 		RegistrantId:   newRegistrant.Id,
// 		FamilyName:     createRegistrantForm.FamilyName,
// 		GivenName:      createRegistrantForm.GivenName,
// 		MiddleName:     createRegistrantForm.MiddleName,
// 		Birthdate:      createRegistrantForm.Birthdate,
// 		Password:       createRegistrantForm.Password,
// 		HasPhoto:       hasPhoto,
// 		StSub:          createRegistrantForm.StSub,
// 		CityMunId:      createRegistrantForm.CityMunId,
// 		ProvId:         createRegistrantForm.ProvId,
// 		BrgyId:         createRegistrantForm.BrgyId,
// 		CivilStatId:    createRegistrantForm.CivilStatId,
// 		CivilStatOther: createRegistrantForm.CivilStatOther,
// 		SexId:          createRegistrantForm.SexId,
// 		Age:            createRegistrantForm.Age,
// 		Height:         createRegistrantForm.Height,
// 		Weight:         createRegistrantForm.Weight,
// 		LandlineNo:     createRegistrantForm.LandlineNo,
// 		MobileNo:       createRegistrantForm.MobileNo,
// 		Email:          createRegistrantForm.Email,
// 	}
// 	registInfo.Create()

// 	registEmp := model.RegistEmp{
// 		RegistrantId: newRegistrant.Id,
// 		EmpStatId:    createRegistrantForm.EmpStatId,
// 		UnEmpStatId:  createRegistrantForm.UnEmpStatId,
// 		TocId:        createRegistrantForm.TocId,
// 		Alfw:         createRegistrantForm.Alfw,
// 		PassportNo:   createRegistrantForm.PassportNo,
// 		Pned:         createRegistrantForm.Pned,
// 	}
// 	registEmp.Create()
// 	http.Redirect(w, r, "/registrants/create", http.StatusFound)
// }

// func RegistrantEmailTaken(w http.ResponseWriter, r *http.Request) {
// 	if taken := model.RegistrantEmailTaken(r.URL.Query().Get("email")); taken {
// 		data := make(map[string]string)
// 		data["error"] = lang.Get("email_taken")
// 		rd.JSON(w, http.StatusNotFound, data)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }
