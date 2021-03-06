// This file is generated by GoAdmin CLI adm.
package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// Generators is a map of table models.
//
// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "project_images" => http://localhost:9033/admin/info/project_images
// "project_tags" => http://localhost:9033/admin/info/project_tags
// "projects" => http://localhost:9033/admin/info/projects
// "users" => http://localhost:9033/admin/info/users
//
// example end
//
var Generators = map[string]table.Generator{

	"project_images": GetProjectImagesTable,
	"project_tags":   GetProjectTagsTable,
	"projects":       GetProjectsTable,
	"users":          GetUsersTable,

	// generators end
}
