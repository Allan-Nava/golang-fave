package wrapper

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	Images "golang-fave/engine/wrapper/resources/images"
	Others "golang-fave/engine/wrapper/resources/others"
	Scripts "golang-fave/engine/wrapper/resources/scripts"
	Styles "golang-fave/engine/wrapper/resources/styles"
	Templates "golang-fave/engine/wrapper/resources/templates"
)

type tmplDataErrorMsg struct {
	ErrorMessage string
}

func (e *Wrapper) staticResource() bool {
	if e.R.URL.Path == "/assets/sys/styles.css" {
		(*e.W).Header().Set("Content-Type", "text/css")
		(*e.W).Write(Styles.File_assets_sys_styles_css)
		return true
	} else if e.R.URL.Path == "/assets/cp/styles.css" {
		(*e.W).Header().Set("Content-Type", "text/css")
		(*e.W).Write(Styles.File_assets_cp_styles_css)
		return true
	} else if e.R.URL.Path == "/assets/sys/bootstrap.css" {
		(*e.W).Header().Set("Content-Type", "text/css")
		(*e.W).Write(Styles.File_assets_sys_bootstrap_css)
		return true
	} else if e.R.URL.Path == "/assets/sys/jquery.js" {
		(*e.W).Header().Set("Content-Type", "application/javascript; charset=utf-8")
		(*e.W).Write(Scripts.File_assets_sys_jquery_js)
		return true
	} else if e.R.URL.Path == "/assets/sys/popper.js" {
		(*e.W).Header().Set("Content-Type", "application/javascript; charset=utf-8")
		(*e.W).Write(Scripts.File_assets_sys_popper_js)
		return true
	} else if e.R.URL.Path == "/assets/sys/bootstrap.js" {
		(*e.W).Header().Set("Content-Type", "application/javascript; charset=utf-8")
		(*e.W).Write(Scripts.File_assets_sys_bootstrap_js)
		return true
	} else if e.R.URL.Path == "/assets/sys/logo.svg" {
		(*e.W).Header().Set("Content-Type", "image/svg+xml")
		(*e.W).Write(Others.File_assets_sys_logo_svg)
		return true
	} else if e.R.URL.Path == "/assets/sys/bg.png" {
		(*e.W).Header().Set("Content-Type", "image/png")
		(*e.W).Write(Images.File_assets_sys_bg_png)
		return true
	} else if e.R.URL.Path == "/assets/sys/logo.png" {
		(*e.W).Header().Set("Content-Type", "image/png")
		(*e.W).Write(Images.File_assets_sys_logo_png)
		return true
	} else if e.R.URL.Path == "/assets/sys/fave.ico" {
		(*e.W).Header().Set("Content-Type", "image/x-icon")
		(*e.W).Write(Others.File_assets_sys_fave_ico)
		return true
	}
	return false
}

func (e *Wrapper) staticFile() bool {
	file := e.R.URL.Path
	if file != "/" {
		f, err := os.Open(e.DirVhostHome + "/htdocs" + file)
		if err == nil {
			defer f.Close()
			st, err := os.Stat(e.DirVhostHome + "/htdocs" + file)
			if err != nil {
				return false
			}
			if st.Mode().IsDir() {
				return false
			}
			http.ServeFile(*e.W, e.R, e.DirVhostHome+"/htdocs"+file)
			return true
		}
	}
	return false
}

func (e *Wrapper) printPageDefault() {
	// Custom page
	f, err := os.Open(e.DirVhostHome + "/htdocs" + "/index.html")
	if err == nil {
		defer f.Close()
		http.ServeFile(*e.W, e.R, e.DirVhostHome+"/htdocs"+"/index.html")
		return
	}

	// Default page
	tmpl, err := template.New("template").Parse(string(Templates.PageDefault))
	if err != nil {
		e.printTmplPageError(err)
		return
	}
	(*e.W).Header().Set("Content-Type", "text/html")
	tmpl.Execute(*e.W, tmplDataAll{
		System: e.tmplGetSystemData(),
	})
}

func (e *Wrapper) printPage404() {
	// Custom 404 error page
	f, err := ioutil.ReadFile(e.DirVhostHome + "/htdocs" + "/404.html")
	if err == nil {
		(*e.W).WriteHeader(http.StatusNotFound)
		(*e.W).Header().Set("Content-Type", "text/html")
		(*e.W).Write(f)
		return
	}

	// Default error page
	tmpl, err := template.New("template").Parse(string(Templates.PageError404))
	if err != nil {
		e.printTmplPageError(err)
		return
	}
	(*e.W).WriteHeader(http.StatusNotFound)
	(*e.W).Header().Set("Content-Type", "text/html")
	tmpl.Execute(*e.W, tmplDataAll{
		System: e.tmplGetSystemData(),
	})
}

func (e *Wrapper) printTmplPageError(perr error) {
	tmpl, err := template.New("template").Parse(string(Templates.PageTmplError))
	if err != nil {
		(*e.W).WriteHeader(http.StatusInternalServerError)
		(*e.W).Header().Set("Content-Type", "text/html")
		(*e.W).Write([]byte("<h1>Critical engine error!</h1>"))
		(*e.W).Write([]byte("<h2>" + perr.Error() + "</h2>"))
		return
	}
	(*e.W).WriteHeader(http.StatusInternalServerError)
	(*e.W).Header().Set("Content-Type", "text/html")
	tmpl.Execute(*e.W, tmplDataAll{
		System: e.tmplGetSystemData(),
		Data: tmplDataErrorMsg{
			ErrorMessage: perr.Error(),
		},
	})
}
