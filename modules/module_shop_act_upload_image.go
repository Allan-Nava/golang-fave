package modules

import (
	"bytes"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"golang-fave/engine/wrapper"
	"golang-fave/utils"
)

func (this *Modules) RegisterAction_ShopUploadImage() *Action {
	return this.newAction(AInfo{
		WantDB:    true,
		Mount:     "shop-upload-image",
		WantAdmin: true,
	}, func(wrap *wrapper.Wrapper) {
		pf_id := wrap.R.FormValue("id")
		pf_count := wrap.R.FormValue("count")

		if !utils.IsNumeric(pf_id) {
			wrap.MsgError(`Inner system error`)
			return
		}

		if !utils.IsNumeric(pf_count) {
			wrap.MsgError(`Inner system error`)
			return
		}

		pf_count_int := utils.StrToInt(pf_count)
		if pf_count_int <= 0 {
			wrap.MsgError(`Inner system error`)
			return
		}

		// Proccess all files
		for i := 1; i <= pf_count_int; i++ {
			post_field_name := "file_" + utils.IntToStr(i-1)
			if file, handler, err := wrap.R.FormFile(post_field_name); err == nil {
				if handler.Filename != "" {
					if fileBytes, err := ioutil.ReadAll(file); err == nil {
						if _, _, err := image.Decode(bytes.NewReader(fileBytes)); err == nil {
							if err := os.MkdirAll(wrap.DHtdocs+string(os.PathSeparator)+"products"+string(os.PathSeparator)+"images"+string(os.PathSeparator)+pf_id, os.ModePerm); err == nil {
								target_file_name := utils.Int64ToStr(time.Now().Unix()+int64(i-1)) + filepath.Ext(handler.Filename)
								target_file_full := wrap.DHtdocs + string(os.PathSeparator) + "products" + string(os.PathSeparator) + "images" + string(os.PathSeparator) + pf_id + string(os.PathSeparator) + target_file_name
								if err := wrap.DB.Transaction(func(tx *wrapper.Tx) error {
									// Block rows
									if _, err := tx.Exec("SELECT id FROM shop_products WHERE id = ? FOR UPDATE;", utils.StrToInt(pf_id)); err != nil {
										return err
									}

									// Insert row
									if _, err := tx.Exec(
										`INSERT INTO shop_product_images SET
											product_id = ?,
											filename = ?
										;`,
										utils.StrToInt(pf_id),
										target_file_name,
									); err != nil {
										return err
									}

									// Write file to disk
									if err := ioutil.WriteFile(target_file_full, fileBytes, 0664); err != nil {
										return err
									}
									return nil
								}); err == nil {
									wrap.Write(`$('#list-images').append('<div class="attached-img"><a href="/products/images/` + pf_id + `/` + target_file_name + `" title="` + target_file_name + `" target="_blank"><img src="/products/images/` + pf_id + `/thumb-0-` + target_file_name + `" onerror="fave.ShopProductsRetryImage(this);" /></a><a href="javascript:fave.ShopProductsDeleteImage(this, ` + pf_id + `, \'` + target_file_name + `\');">Delete</a></div>');`)
								}
							}
						}
					}
				}
				file.Close()
			}
		}

		wrap.RecreateProductXmlFile()

		wrap.ResetCacheBlocks()
	})
}
