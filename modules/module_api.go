package modules

import (
	"net/http"
	"os"

	"golang-fave/assets"
	"golang-fave/engine/fetdata"
	"golang-fave/engine/wrapper"
	"golang-fave/utils"
)

func (this *Modules) RegisterModule_Api() *Module {
	return this.newModule(MInfo{
		WantDB: true,
		Mount:  "api",
		Name:   "Api",
		Order:  803,
		System: true,
		Icon:   assets.SysSvgIconPage,
		Sub:    &[]MISub{},
	}, func(wrap *wrapper.Wrapper) {
		if len(wrap.UrlArgs) == 2 && wrap.UrlArgs[0] == "api" && wrap.UrlArgs[1] == "products" {
			if (*wrap.Config).API.XML.Enabled == 1 {
				// Fix url
				if wrap.R.URL.Path[len(wrap.R.URL.Path)-1] != '/' {
					http.Redirect(wrap.W, wrap.R, wrap.R.URL.Path+"/"+utils.ExtractGetParams(wrap.R.RequestURI), 301)
					return
				}

				target_file := wrap.DHtdocs + string(os.PathSeparator) + "products.xml"
				if !utils.IsFileExists(target_file) {
					data := []byte(this.api_GenerateEmptyXml(wrap))

					// Make empty file
					if file, err := os.Create(target_file); err == nil {
						file.Write(data)
						file.Close()
					}

					// Make regular XML
					data = []byte(this.api_GenerateXml(wrap))

					// Save file
					wrap.RemoveProductXmlCacheFile()
					if file, err := os.Create(target_file); err == nil {
						file.Write(data)
						file.Close()
					}

					wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
					wrap.W.Header().Set("Content-Type", "text/xml; charset=utf-8")
					wrap.W.WriteHeader(http.StatusOK)
					wrap.W.Write(data)
				} else {
					http.ServeFile(wrap.W, wrap.R, target_file)
				}
			} else {
				wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				wrap.W.WriteHeader(http.StatusNotFound)
				wrap.W.Write([]byte("Disabled!"))
			}
		} else if len(wrap.UrlArgs) == 1 {
			// Fix url
			if wrap.R.URL.Path[len(wrap.R.URL.Path)-1] != '/' {
				http.Redirect(wrap.W, wrap.R, wrap.R.URL.Path+"/"+utils.ExtractGetParams(wrap.R.RequestURI), 301)
				return
			}

			// Some info
			wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			wrap.W.WriteHeader(http.StatusOK)
			wrap.W.Write([]byte("Fave engine API mount point!"))
		} else {
			// User error 404 page
			wrap.RenderFrontEnd("404", fetdata.New(wrap, true, nil, nil), http.StatusNotFound)
			return
		}
	}, func(wrap *wrapper.Wrapper) (string, string, string) {
		// No any page for back-end
		return "", "", ""
	})
}
