package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewCurso struct {
	Name    string
	IsEdit  bool
	Data    models.Curso
	Widgets []models.Curso
}

var tmplc = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/curso/index.html", "web/curso/form.html"))

func CursoList(w http.ResponseWriter, req *http.Request) {

	curso := models.Curso{}

	cursos, _ := curso.FindAll(cfig.DB)

	for _, lis := range cursos {
		fmt.Println(lis.ToString())
		fmt.Println("Matriculas: ", len(lis.Matriculas))
		if len(lis.Matriculas) > 0 {
			for _, d := range lis.Matriculas {
				fmt.Println(d.ToString())
				fmt.Println("=============================")
			}
		}
		fmt.Println("--------------------")
	}

	// Create
	//cfig.DB.Create(&models.Alumno{Name: "Juan", City: "Juliaca"})
	lis := []models.Curso{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewCurso{
		Name:    "Curso",
		Widgets: lis,
	}

	err := tmpla.ExecuteTemplate(w, "curso/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CursoForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Curso
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Nombres = r.FormValue("nombres")
		d.Codigo = r.FormValue("codigo")
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
		http.Redirect(w, r, "/curso/index", 301)
	}

	data := ViewCurso{
		Name:   "Curso",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmpla.ExecuteTemplate(w, "curso/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CursoDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.Curso
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}
	http.Redirect(w, r, "/curso/index", 301)
}
