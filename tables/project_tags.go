package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProjectTagsTable(ctx *context.Context) table.Table {

	projectTags := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := projectTags.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Project_id", "project_id", db.Int)

	info.SetTable("project_tags").SetTitle("ProjectTags").SetDescription("ProjectTags")

	formList := projectTags.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Project_id", "project_id", db.Int, form.Number)

	formList.SetTable("project_tags").SetTitle("ProjectTags").SetDescription("ProjectTags")

	return projectTags
}
