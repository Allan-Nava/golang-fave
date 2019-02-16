package modules

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"net/http"
	"os"
	"strconv"

	"golang-fave/engine/wrapper"
	"golang-fave/utils"
)

func (this *Modules) RegisterModule_Index() *Module {
	return this.newModule(MInfo{
		WantDB: true,
		Mount:  "index",
		Name:   "Pages",
	}, func(wrap *wrapper.Wrapper) {
		// Front-end
		wrap.W.WriteHeader(http.StatusOK)
		wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		wrap.W.Header().Set("Content-Type", "text/html; charset=utf-8")
		wrap.W.Write([]byte(`INDEX FrontEnd func call (` + wrap.CurrModule + `)`))
	}, func(wrap *wrapper.Wrapper) {
		// Back-end
		wrap.W.WriteHeader(http.StatusOK)
		wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		wrap.W.Header().Set("Content-Type", "text/html; charset=utf-8")
		wrap.W.Write([]byte(`INDEX BackEnd func call (` + wrap.CurrModule + `)`))
	})
}

func (this *Modules) RegisterAction_MysqlSetup() *Action {
	return this.newAction(AInfo{
		WantDB: false,
		Mount:  "mysql",
	}, func(wrap *wrapper.Wrapper) {
		pf_host := wrap.R.FormValue("host")
		pf_port := wrap.R.FormValue("port")
		pf_name := wrap.R.FormValue("name")
		pf_user := wrap.R.FormValue("user")
		pf_password := wrap.R.FormValue("password")

		if pf_host == "" {
			wrap.MsgError(`Please specify host for MySQL connection`)
			return
		}

		if pf_port == "" {
			wrap.MsgError(`Please specify host port for MySQL connection`)
			return
		}

		if _, err := strconv.Atoi(pf_port); err != nil {
			wrap.MsgError(`MySQL host port must be integer number`)
			return
		}

		if pf_name == "" {
			wrap.MsgError(`Please specify MySQL database name`)
			return
		}

		if pf_user == "" {
			wrap.MsgError(`Please specify MySQL user`)
			return
		}

		// Try connect to mysql
		db, err := sql.Open("mysql", pf_user+":"+pf_password+"@tcp("+pf_host+":"+pf_port+")/"+pf_name)
		if err != nil {
			wrap.MsgError(err.Error())
			return
		}
		defer db.Close()
		err = db.Ping()
		if err != nil {
			wrap.MsgError(err.Error())
			return
		}

		// Try to install all tables
		_, err = db.Query(fmt.Sprintf(
			"CREATE TABLE `%s`.`users` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'AI', `first_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT 'User first name', `last_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT 'User last name', `email` VARCHAR(64) NOT NULL COMMENT 'User email', `password` VARCHAR(32) NOT NULL COMMENT 'User password (MD5)', PRIMARY KEY (`id`)) ENGINE = InnoDB;",
			pf_name))
		if err != nil {
			wrap.MsgError(err.Error())
			return
		}
		_, err = db.Query(fmt.Sprintf(
			"CREATE TABLE `%s`.`pages` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'AI', `user` int(11) NOT NULL COMMENT 'User id', `name` varchar(255) NOT NULL COMMENT 'Page name', `slug` varchar(255) NOT NULL COMMENT 'Page url part', `content` text NOT NULL COMMENT 'Page content', `meta_title` varchar(255) NOT NULL DEFAULT '' COMMENT 'Page meta title', `meta_keywords` varchar(255) NOT NULL DEFAULT '' COMMENT 'Page meta keywords', `meta_description` varchar(510) NOT NULL DEFAULT '' COMMENT 'Page meta description', `datetime` datetime NOT NULL COMMENT 'Creation date/time', `status` enum('draft','public','trash') NOT NULL COMMENT 'Page status', PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;",
			pf_name))
		if err != nil {
			wrap.MsgError(err.Error())
			return
		}

		// Save mysql config file
		err = utils.MySqlConfigWrite(wrap.DConfig+string(os.PathSeparator)+"mysql.json", pf_host, pf_port, pf_name, pf_user, pf_password)
		if err != nil {
			wrap.MsgError(err.Error())
			return
		}

		// Reload current page
		wrap.Write(`window.location.reload(false);`)
	})
}

func (this *Modules) RegisterAction_CpFirstUser() *Action {
	return this.newAction(AInfo{
		WantDB: true,
		Mount:  "first_user",
	}, func(wrap *wrapper.Wrapper) {
		pf_first_name := wrap.R.FormValue("first_name")
		pf_last_name := wrap.R.FormValue("last_name")
		pf_email := wrap.R.FormValue("email")
		pf_password := wrap.R.FormValue("password")

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

		_, err := wrap.DB.Query(
			"INSERT INTO `users` SET `first_name` = ?, `last_name` = ?, `email` = ?, `password` = MD5(?);",
			pf_first_name, pf_last_name, pf_email, pf_password)
		if err != nil {
			wrap.MsgError(err.Error())
			return
		}

		// Reload current page
		wrap.Write(`window.location.reload(false);`)
	})
}

// All actions here...
// User login
// User logout
