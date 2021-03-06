package modules

import (
	"html"

	"golang-fave/engine/assets"
	"golang-fave/engine/builder"
	"golang-fave/engine/consts"
	"golang-fave/engine/sqlw"
	"golang-fave/engine/utils"
	"golang-fave/engine/wrapper"
)

func (this *Modules) RegisterModule_NotifyMail() *Module {
	return this.newModule(MInfo{
		Mount:  "notify-mail",
		Name:   "Mail notifier",
		Order:  803,
		System: true,
		Icon:   assets.SysSvgIconEmail,
		Sub: &[]MISub{
			{Mount: "default", Name: "All", Show: true, Icon: assets.SysSvgIconList},
			{Mount: "success", Name: "Success", Show: true, Icon: assets.SysSvgIconList},
			{Mount: "in-progress", Name: "In progress", Show: true, Icon: assets.SysSvgIconList},
			{Mount: "error", Name: "Error", Show: true, Icon: assets.SysSvgIconList},
		},
	}, nil, func(wrap *wrapper.Wrapper) (string, string, string) {
		content := ""
		sidebar := ""
		if wrap.CurrSubModule == "" || wrap.CurrSubModule == "default" || wrap.CurrSubModule == "success" || wrap.CurrSubModule == "in-progress" || wrap.CurrSubModule == "error" {
			ModuleName := "All"
			ModulePagination := "/cp/" + wrap.CurrModule + "/"
			ModuleSqlWhere := ""

			if wrap.CurrSubModule == "success" {
				ModuleName = "Success"
				ModulePagination = "/cp/" + wrap.CurrModule + "/success/"
				ModuleSqlWhere = " WHERE fave_notify_mail.status = 1"
			} else if wrap.CurrSubModule == "in-progress" {
				ModuleName = "In progress"
				ModulePagination = "/cp/" + wrap.CurrModule + "/in-progress/"
				ModuleSqlWhere = " WHERE fave_notify_mail.status = 2 OR fave_notify_mail.status = 3"
			} else if wrap.CurrSubModule == "error" {
				ModuleName = "Error"
				ModulePagination = "/cp/" + wrap.CurrModule + "/error/"
				ModuleSqlWhere = " WHERE fave_notify_mail.status = 0"
			}

			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: ModuleName},
			})
			content += builder.DataTable(
				wrap,
				"fave_notify_mail",
				"id",
				"DESC",
				&[]builder.DataTableRow{
					{
						DBField: "id",
					},
					{
						DBField:     "email",
						NameInTable: "Email / Subject",
						CallBack: func(values *[]string) string {
							subject := html.EscapeString((*values)[2])
							if subject != "" {
								subject = `<div><small>` + subject + `</small></div>`
							}
							error_message := html.EscapeString((*values)[5])
							if error_message != "" {
								error_message = `<div><small><b>` + error_message + `</b></small></div>`
							}
							return `<div>` + html.EscapeString((*values)[1]) + `</div>` + subject + error_message
						},
					},
					{
						DBField: "subject",
					},
					{
						DBField:     "datetime",
						DBExp:       "UNIX_TIMESTAMP(`datetime`)",
						NameInTable: "Date / Time",
						Classes:     "d-none d-md-table-cell",
						CallBack: func(values *[]string) string {
							t := int64(utils.StrToInt((*values)[3]))
							return `<div>` + utils.UnixTimestampToFormat(t, "02.01.2006") + `</div>` +
								`<div><small>` + utils.UnixTimestampToFormat(t, "15:04:05") + `</small></div>`
						},
					},
					{
						DBField:     "status",
						NameInTable: "Status",
						Classes:     "d-none d-sm-table-cell",
						CallBack: func(values *[]string) string {
							return builder.CheckBox(utils.StrToInt((*values)[4]))
						},
					},
					{
						DBField: "error",
					},
				},
				nil,
				ModulePagination,
				func() (int, error) {
					var count int
					return count, wrap.DB.QueryRow(
						wrap.R.Context(),
						"SELECT COUNT(*) FROM `fave_notify_mail`"+ModuleSqlWhere+";",
					).Scan(&count)
				},
				func(limit_offset int, pear_page int) (*sqlw.Rows, error) {
					return wrap.DB.Query(
						wrap.R.Context(),
						`SELECT
							fave_notify_mail.id,
							fave_notify_mail.email,
							fave_notify_mail.subject,
							UNIX_TIMESTAMP(`+"`fave_notify_mail`.`datetime`"+`) AS datetime,
							fave_notify_mail.status,
							fave_notify_mail.error
						FROM
							fave_notify_mail
						`+ModuleSqlWhere+`
						ORDER BY
							fave_notify_mail.id DESC
						LIMIT ?, ?;`,
						limit_offset,
						pear_page,
					)
				},
				true,
			)
		}
		return this.getSidebarModules(wrap), content, sidebar
	})
}
