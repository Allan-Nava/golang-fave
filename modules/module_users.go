package modules

import (
	"html"

	"golang-fave/assets"
	"golang-fave/consts"
	"golang-fave/engine/builder"
	"golang-fave/engine/wrapper"
)

func (this *Modules) RegisterModule_Users() *Module {
	return this.newModule(MInfo{
		WantDB: true,
		Mount:  "users",
		Name:   "Users",
		Order:  800,
		System: true,
		Icon:   assets.SysSvgIconUser,
		Sub: &[]MISub{
			{Mount: "default", Name: "List of Users", Show: true, Icon: assets.SysSvgIconList},
			{Mount: "add", Name: "Add New User", Show: true, Icon: assets.SysSvgIconPlus},
			{Mount: "modify", Name: "Modify user", Show: false},
		},
	}, nil, func(wrap *wrapper.Wrapper) (string, string, string) {
		content := ""
		sidebar := ""
		if wrap.CurrSubModule == "" || wrap.CurrSubModule == "default" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "List of Users"},
			})
			content += builder.DataTable(wrap, "users", "id", "DESC", []builder.DataTableRow{
				{
					DBField: "id",
				},
				{
					DBField:     "email",
					NameInTable: "Email / Name",
					CallBack: func(values *[]string) string {
						email := `<a href="/cp/` + wrap.CurrModule + `/modify/` + (*values)[0] + `/">` + html.EscapeString((*values)[1]) + `</a>`
						name := html.EscapeString((*values)[2])
						if name != "" && (*values)[3] != "" {
							name += ` ` + (*values)[3]
						}
						if name != "" {
							name = `<div><small>` + name + `</small></div>`
						}
						return `<div>` + email + `</div>` + name
					},
				},
				{
					DBField: "first_name",
				},
				{
					DBField: "last_name",
				},
			}, func(values *[]string) string {
				return `<a class="ico" href="/cp/` + wrap.CurrModule + `/modify/` + (*values)[0] + `/">` +
					assets.SysSvgIconEdit + `</a>` +
					`<a class="ico" href="#">` + assets.SysSvgIconRemove + `</a>`
			}, "/cp/"+wrap.CurrModule+"/")
		} else if wrap.CurrSubModule == "add" || wrap.CurrSubModule == "modify" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Add New User"},
			})
			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "users-modify",
				},
				{
					Kind:  builder.DFKHidden,
					Name:  "id",
					Value: "0",
				},
				{
					Kind:    builder.DFKText,
					Caption: "First Name",
					Name:    "first_name",
					Value:   "",
				},
				{
					Kind:    builder.DFKText,
					Caption: "Last Name",
					Name:    "last_name",
					Value:   "",
				},
				{
					Kind:     builder.DFKEmail,
					Caption:  "Email",
					Name:     "email",
					Value:    "",
					Required: true,
				},
				{
					Kind:    builder.DFKPassword,
					Caption: "Password",
					Name:    "password",
					Hint:    "Leave the field blank to not change the password",
				},
				{
					Kind: builder.DFKMessage,
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  "Add",
					Target: "add-edit-button",
				},
			})
			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Add</button>`
		}
		return this.getSidebarModules(wrap), content, sidebar
	})
}
