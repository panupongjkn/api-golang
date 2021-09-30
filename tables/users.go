package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersTable(ctx *context.Context) table.Table {

	users := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := users.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Password", "password", db.Varchar)
	info.AddField("First_name", "first_name", db.Varchar)
	info.AddField("Last_name", "last_name", db.Varchar)
	info.AddField("Profile_image", "profile_image", db.Varchar).FieldImage("50", "auto", "http://localhost:8080/uploads/")
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList := users.GetForm()
	formList.AddField("Profile_image", "profile_image", db.Varchar, form.File)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("First_name", "first_name", db.Varchar, form.Text)
	formList.AddField("Last_name", "last_name", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).FieldNowWhenInsert().FieldHideWhenCreate()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).FieldNowWhenInsert().FieldHideWhenCreate()

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	return users
}
