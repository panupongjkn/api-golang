package tables

import (
	"log"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProjectImagesTable(ctx *context.Context) table.Table {
	log.Println(ctx.Request.URL.Query().Get("test"))
	projectImages := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))
	log.Println(projectImages)

	info := projectImages.GetInfo().Where("id", "=", "1").HideFilterArea()
	log.Println(info)

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	// info.AddField("Created_at", "created_at", db.Datetime)
	// info.AddField("Updated_at", "updated_at", db.Datetime)
	// info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Image_url", "image_url", db.Varchar)
	info.AddField("Project_id", "project_id", db.Int)

	info.SetTable("project_images").SetTitle("ProjectImages").SetDescription("ProjectImages")

	formList := projectImages.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Image_url", "image_url", db.Varchar, form.Text)
	formList.AddField("Project_id", "project_id", db.Int, form.Number)

	formList.SetTable("project_images").SetTitle("ProjectImages").SetDescription("ProjectImages")

	return projectImages
}
