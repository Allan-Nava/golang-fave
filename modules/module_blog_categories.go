package modules

import (
	"html"
	"strings"

	"golang-fave/engine/wrapper"
	"golang-fave/utils"
)

func (this *Modules) blog_GetCategorySelectOptions(wrap *wrapper.Wrapper, id int, parentId int) string {
	result := ``
	rows, err := wrap.DB.Query(
		`SELECT
			node.id,
			node.user,
			node.name,
			node.alias,
			(COUNT(parent.id) - 1) AS depth
		FROM
			blog_cats AS node,
			blog_cats AS parent
		WHERE
			node.lft BETWEEN parent.lft AND parent.rgt
		GROUP BY
			node.id
		ORDER BY
			node.lft ASC
		;`,
	)
	if err == nil {
		values := make([]string, 5)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		idStr := utils.IntToStr(id)
		parentIdStr := utils.IntToStr(parentId)
		for rows.Next() {
			err = rows.Scan(scan...)
			if err == nil {
				disabled := ""
				if string(values[0]) == idStr {
					disabled = " disabled"
				}
				selected := ""
				if string(values[0]) == parentIdStr {
					selected = " selected"
				}
				sub := strings.Repeat("&mdash; ", utils.StrToInt(string(values[4])))
				result += `<option value="` + html.EscapeString(string(values[0])) + `"` + disabled + selected + `>` + sub + html.EscapeString(string(values[2])) + `</option>`
			}
		}
	}
	return result
}

func (this *Modules) blog_GetCategoryParentId(wrap *wrapper.Wrapper, id int) int {
	var parentId int
	_ = wrap.DB.QueryRow(`
		SELECT
			parent.id
		FROM
			blog_cats AS node,
			blog_cats AS parent
		WHERE
			node.lft BETWEEN parent.lft AND parent.rgt AND
			node.id = ? AND
			parent.id <> ?
		ORDER BY
			parent.lft DESC
		LIMIT 1;`,
		id,
		id,
	).Scan(
		&parentId,
	)
	return parentId
}
