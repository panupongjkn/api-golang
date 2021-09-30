package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProjectsTable(ctx *context.Context) table.Table {

	projects := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := projects.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Cover", "cover", db.Varchar)
	info.AddField("Like", "like", db.Int)
	info.AddField("Download", "download", db.Int)
	info.AddField("Program", "program", db.Varchar)
	info.AddField("Owner_id", "owner_id", db.Int)

	info.SetTable("projects").SetTitle("Projects").SetDescription("Projects")

	formList := projects.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Cover", "cover", db.Varchar, form.Text)
	formList.AddField("Like", "like", db.Int, form.Number)
	formList.AddField("Download", "download", db.Int, form.Number)
	formList.AddField("Program", "program", db.Varchar, form.Text)
	formList.AddField("Owner_id", "owner_id", db.Int, form.Number)

	formList.SetTable("projects").SetTitle("Projects").SetDescription("Projects")

	return projects
}
