package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/kareemmahlees/mysql-meta/docs"
	db_handlers "github.com/kareemmahlees/mysql-meta/internal/db"
	"github.com/kareemmahlees/mysql-meta/lib"
	"github.com/kareemmahlees/mysql-meta/utils"
)

func RegisterTablesRoutes(app *fiber.App, db *sqlx.DB) {
	tableGroup := app.Group("tables")
	tableGroup.Get("", utils.RouteHandler(db, handleListTables))
	tableGroup.Get("/:tableName/describe", utils.RouteHandler(db, handeGetTableInfo))
	tableGroup.Post("/:tableName", utils.RouteHandler(db, handleCreateTable))
	tableGroup.Delete("/:tableName", utils.RouteHandler(db, handleDeleteTable))
	tableGroup.Put("/:tableName", utils.RouteHandler(db, handleUpdateTable))
}

func handeGetTableInfo(c *fiber.Ctx, db *sqlx.DB) error {
	if err := lib.ValidateVar(c.Params("tableName"), "required,alpha"); err != nil {
		return c.JSON(lib.ResponseError400(err.Error()))
	}
	info, err := db_handlers.GetTableInfo(db, c.Params("tableName"))
	if err != nil {
		return c.JSON(lib.ResponseError500(err.Error()))
	}
	return c.JSON(info)
}

type HandleListTablesResp struct {
	Tables []string
}

// Lists all tables in the database
//
//	@tags			Tables
//	@description	list tables
//	@router			/tables [get]
//	@produce		json
//	@success		200	{object}	HandleListTablesResp
func handleListTables(c *fiber.Ctx, db *sqlx.DB) error {
	tables, err := db_handlers.ListTables(db)
	if err != nil {
		return c.JSON(lib.ResponseError500(err.Error()))
	}
	return c.JSON(fiber.Map{"tables": tables})
}

type HandleCreateTableResp struct {
	Created string
}
type HandleCreateTableBody struct {
	ColName lib.CreateTableProps
}

// Creates a Table
//
//	@tags			Tables
//	@description	create table
//	@router			/tables/{tableName} [post]
//	@param			tableName	path	string					true	"table name"
//	@param			tableData	body	HandleCreateTableBody	true	"create table data"
//	@accept			json
//	@produce		json
//	@success		201	{object}	HandleCreateTableResp
func handleCreateTable(c *fiber.Ctx, db *sqlx.DB) error {
	if err := lib.ValidateVar(c.Params("tableName"), "required,alpha"); err != nil {
		return c.JSON(lib.ResponseError400(err.Error()))
	}
	var payload map[string]lib.CreateTableProps
	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(lib.ResponseError500(err.Error()))
	}
	for _, v := range payload {
		errs := lib.ValidateStruct(v)
		if len(errs) > 0 {
			return c.Status(400).JSON(lib.ResponseError400(errs))
		}
	}
	err := db_handlers.CreateTable(db, c.Params("tableName"), payload)
	if err != nil {
		return c.JSON(lib.ResponseError500(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"created": c.Params("tableName")})
}

type HandleUpdateDeleteResp struct {
	Success bool
}

// Updates a Table either by add,modify or delete
//
//	@tags			Tables
//	@description	update table
//	@router			/tables/{tableName} [put]
//	@param			tableName	path	string					true	"table name"
//	@param			tableData	body	lib.UpdateTableProps	true	"update table data"
//	@accept			json
//	@produce		json
//	@accept			json
//	@success		200	{object}	HandleUpdateDeleteResp
func handleUpdateTable(c *fiber.Ctx, db *sqlx.DB) error {
	if err := lib.ValidateVar(c.Params("tableName"), "required,alpha"); err != nil {
		return c.JSON(lib.ResponseError400(err.Error()))
	}
	var payload lib.UpdateTableProps
	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(lib.ResponseError500(err.Error()))
	}
	errs := lib.ValidateStruct(payload)
	if len(errs) > 0 {
		return c.JSON(lib.ResponseError400(errs))
	}
	err := db_handlers.UpdateTable(db, c.Params("tableName"), payload)
	if err != nil {
		return c.JSON(lib.ResponseError500(err.Error()))
	}
	return c.JSON(fiber.Map{"success": true})
}

// Deletes a table
//
//	@tags		Tables
//	@decription	delete table
//	@router		/tables/{tableName} [delete]
//	@param		tableName	path	string	true	"table name"
//	@accept		json
//	@produce	json
//	@success	200	{object}	HandleUpdateDeleteResp
func handleDeleteTable(c *fiber.Ctx, db *sqlx.DB) error {
	if err := lib.ValidateVar(c.Params("tableName"), "required,alpha"); err != nil {
		return c.JSON(lib.ResponseError400(err.Error()))
	}
	err := db_handlers.DeleteTable(db, c.Params("tableName"))
	if err != nil {
		return c.JSON(lib.ResponseError500(err.Error()))
	}
	return c.JSON(fiber.Map{"success": true})
}
