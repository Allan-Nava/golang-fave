package modules

import (
	"golang-fave/engine/utils"
	"golang-fave/engine/wrapper"
)

func (this *Modules) RegisterAction_IndexFirstUser() *Action {
	return this.newAction(AInfo{
		Mount: "index-first-user",
	}, func(wrap *wrapper.Wrapper) {
		pf_first_name := utils.Trim(wrap.R.FormValue("first_name"))
		pf_last_name := utils.Trim(wrap.R.FormValue("last_name"))
		pf_email := utils.Trim(wrap.R.FormValue("email"))
		pf_password := utils.Trim(wrap.R.FormValue("password"))

		if pf_email == "" {
			wrap.MsgError(`Please specify user email`)
			return
		}

		if !utils.IsValidEmail(pf_email) {
			wrap.MsgError(`Please specify correct user email`)
			return
		}

		if pf_password == "" {
			wrap.MsgError(`Please specify user password`)
			return
		}

		// Security, check if still need to run this action
		var count int
		err := wrap.DB.QueryRow(
			wrap.R.Context(),
			`SELECT
				COUNT(*)
			FROM
				fave_users
			;`,
		).Scan(
			&count,
		)
		if *wrap.LogCpError(&err) != nil {
			wrap.MsgError(err.Error())
			return
		}
		if count > 0 {
			wrap.MsgError(`CMS is already configured`)
			return
		}

		_, err = wrap.DB.Exec(
			wrap.R.Context(),
			`INSERT INTO fave_users SET
				id = 1,
				first_name = ?,
				last_name = ?,
				email = ?,
				password = MD5(?),
				admin = 1,
				active = 1
			;`,
			pf_first_name,
			pf_last_name,
			pf_email,
			pf_password,
		)
		if err != nil {
			wrap.MsgError(err.Error())
			return
		}

		wrap.ResetCacheBlocks()

		// Reload current page
		wrap.Write(`window.location.reload(false);`)
	})
}
