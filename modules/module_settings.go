package modules

import (
	"html"
	"io/ioutil"
	"os"

	"golang-fave/assets"
	"golang-fave/consts"
	"golang-fave/engine/builder"
	"golang-fave/engine/wrapper"
	"golang-fave/utils"
)

func (this *Modules) RegisterModule_Settings() *Module {
	return this.newModule(MInfo{
		WantDB: false,
		Mount:  "settings",
		Name:   "Settings",
		Order:  801,
		System: true,
		Icon:   assets.SysSvgIconGear,
		Sub: &[]MISub{
			{Mount: "default", Name: "Robots.txt", Show: true, Icon: assets.SysSvgIconBug},
			{Mount: "pagination", Name: "Pagination", Show: true, Icon: assets.SysSvgIconList},
		},
	}, nil, func(wrap *wrapper.Wrapper) (string, string, string) {
		content := ""
		sidebar := ""

		if wrap.CurrSubModule == "" || wrap.CurrSubModule == "default" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Robots.txt"},
			})

			fcont := []byte(``)
			fcont, _ = ioutil.ReadFile(wrap.DTemplate + string(os.PathSeparator) + "robots.txt")

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-robots-txt",
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group last"><div class="row"><div class="col-12"><textarea class="form-control autosize" id="lbl_content" name="content" placeholder="" autocomplete="off">` + html.EscapeString(string(fcont)) + `</textarea></div></div></div>`
					},
				},
				{
					Kind: builder.DFKMessage,
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="row"><div class="col-12"><div class="sys-messages"></div></div></div>`
					},
				},
				{
					Kind: builder.DFKSubmit,
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="row d-lg-none"><div class="col-12"><button type="submit" class="btn btn-primary" data-target="add-edit-button">Save</button></div></div>`
					},
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		} else if wrap.CurrSubModule == "pagination" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Pagination"},
			})

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-pagination",
				},
				{
					Kind:     builder.DFKNumber,
					Caption:  "Blog main page",
					Name:     "blog-index",
					Min:      "1",
					Max:      "100",
					Required: true,
					Value:    utils.IntToStr((*wrap.Config).Blog.Pagination.Index),
				},
				{
					Kind:     builder.DFKNumber,
					Caption:  "Blog category page",
					Name:     "blog-category",
					Min:      "1",
					Max:      "100",
					Required: true,
					Value:    utils.IntToStr((*wrap.Config).Blog.Pagination.Category),
				},
				{
					Kind:    builder.DFKText,
					Caption: "",
					Name:    "",
					Value:   "",
					CallBack: func(field *builder.DataFormField) string {
						return `<hr>`
					},
				},
				{
					Kind:     builder.DFKNumber,
					Caption:  "Shop main page",
					Name:     "shop-index",
					Min:      "1",
					Max:      "100",
					Required: true,
					Value:    utils.IntToStr((*wrap.Config).Shop.Pagination.Index),
				},
				{
					Kind:     builder.DFKNumber,
					Caption:  "Shop category page",
					Name:     "shop-category",
					Min:      "1",
					Max:      "100",
					Required: true,
					Value:    utils.IntToStr((*wrap.Config).Shop.Pagination.Category),
				},
				{
					Kind: builder.DFKMessage,
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  "Save",
					Target: "add-edit-button",
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		}
		return this.getSidebarModules(wrap), content, sidebar
	})
}
