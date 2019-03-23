package sqlw

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"golang-fave/consts"
)

func log(query string, s time.Time, transaction bool) {
	color := "0;33"
	if transaction {
		color = "1;33"
	}
	msg := query
	if reg, err := regexp.Compile("[\\s\\t]+"); err == nil {
		msg = strings.Trim(reg.ReplaceAllString(msg, " "), " ")
	}
	if reg, err := regexp.Compile("[\\s\\t]+;$"); err == nil {
		msg = reg.ReplaceAllString(msg, ";")
	}
	if consts.ParamDebug {
		t := time.Now().Sub(s).Seconds()
		if consts.IS_WIN {
			fmt.Fprintln(os.Stdout, "[SQL] "+msg+fmt.Sprintf(" %.3f ms", t))
		} else {
			fmt.Fprintln(os.Stdout, "\033["+color+"m[SQL] "+msg+fmt.Sprintf(" %.3f ms", t)+"\033[0m")
		}
	}
}
