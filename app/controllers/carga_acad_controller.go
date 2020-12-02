package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewCarga_acad struct {
	Name    string
	IsEdit  bool
	Data    models.Carga_acad
	Widgets []models.Carga_acad
	Cursos []models.Curso
}

var tmplca = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/carga_acad/index.html", "web/carga_acad/form.html"))

func Carga_acadList(w http.ResponseWriter, req *http.Request) {

	lis := []models.Carga_acad{}
	if err := cfig.DB.Preload("Curso").Find(&lis).Error; err != nil { // Preload("Curso") carga los objetos Curso relacionado
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := ViewCarga_acad{
		Name:    "Carga_acad",
		Widgets: lis,
	}
	err := tmplca.ExecuteTemplate(w, "carga_acad/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Carga_acadForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Carga_acad
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	curso := models.Curso{}
	cursos, _ := curso.GetAll(cfig.DB) // para mostrar los cursos en un combobox

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Semestre = r.FormValue("semestre")
		//n, err := strconv.Atoi(r.FormValue("alumno_id"))
		//if err != nil {
		//	log.Printf("Invalid ID: %v - %v\n", n, err)
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		d.CursoId = r.FormValue("alumno_id") //n
		if id != "" {
			if err := cfig.DB.Save(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}

		} else {
			if err := cfig.DB.Create(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}
		}
		http.Redirect(w, r, "/carga_acad/index", 301)
	}

	data := ViewCarga_acad{
		Name:    "Carga_acad",
		Data:    d,
		IsEdit:  IsEdit,
		Cursos: cursos,
	}

	err := tmplca.ExecuteTemplate(w, "carga_acad/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Carga_acadDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var d models.Carga_acad
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		//log.Printf("No save  %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}

	http.Redirect(w, r, "/carga_acad/index", 301)
}