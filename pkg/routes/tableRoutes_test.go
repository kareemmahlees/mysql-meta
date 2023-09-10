package routes

import (
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kareemmahlees/mysql-meta/pkg/db"
	"github.com/kareemmahlees/mysql-meta/utils"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func TestHandleListTables(t *testing.T) {
	app := fiber.New()

	con, err := db.InitDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	RegisterTablesRoutes(app, con)
	req := httptest.NewRequest("GET", "http://localhost:4000/tables", nil)

	resp, _ := app.Test(req)
	payload := utils.ReadBody(resp.Body)

	tables, ok := payload["tables"]
	assert.True(t, ok)
	assert.IsType(t, reflect.Slice, reflect.TypeOf(tables).Kind())
}

func TestHandleCreateTable(t *testing.T) {
	app := fiber.New()

	con, err := db.InitDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	RegisterTablesRoutes(app, con)
	req := httptest.NewRequest("POST", "http://localhost:4000/tables/testHandleCreateTable", strings.NewReader(`
	{
    "name": {
        "type": "text",
        "nullable": true,
        "default": "kareem",
        "unique": true
    	}
	}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.Nil(t, err)

	payload := utils.ReadBody(resp.Body)
	assert.Equal(t, payload["created"], "testHandleCreateTable")

	req = httptest.NewRequest("GET", "http://localhost:4000/tables", nil)

	resp, _ = app.Test(req)
	payload = utils.ReadBody(resp.Body)

	assert.Contains(t, payload["tables"], "testHandleCreateTable")
}

func TestHandleUpdateTable(t *testing.T) {
	app := fiber.New()

	con, err := db.InitDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	RegisterTablesRoutes(app, con)

	req := httptest.NewRequest("POST", "http://localhost:4000/tables/testHandleUpdateTable", strings.NewReader(`{
		"name":{
			"type":"text",
			"default":"defaultText",
			"unique":true,
			"nullable":false
		}
	}`))
	req.Header.Set("Content-Type", "application/json")
	_, err = app.Test(req)
	assert.Nil(t, err)

	req = httptest.NewRequest("PUT", "http://localhost:4000/tables/testHandleUpdateTable", strings.NewReader(`{
		"operation":{
			"type":"add",
			"data":{
				"age":"int"
			}
		}
	}`))
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	resp, err = app.Test(req)
	assert.Nil(t, err)
	payload := utils.ReadBody(resp.Body)
	assert.True(t, payload["success"].(bool), true)

	req = httptest.NewRequest("PUT", "http://localhost:4000/tables/testHandleUpdateTable", strings.NewReader(`{
		"operation":{
			"type":"modify",
			"data":{
				"name":"varchar(255)"
			}
		}
	}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(req)
	assert.Nil(t, err)
	payload = utils.ReadBody(resp.Body)
	assert.True(t, payload["success"].(bool), true)

	req = httptest.NewRequest("PUT", "http://localhost:4000/tables/testHandleUpdateTable", strings.NewReader(`{
		"operation":{
			"type":"delete",
			"data":["name","age"]
		}
	}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(req)
	assert.Nil(t, err)
	payload = utils.ReadBody(resp.Body)
	assert.True(t, payload["success"].(bool), true)
}

func TestHandleDeleteTalbe(t *testing.T) {
	app := fiber.New()

	con, err := db.InitDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	RegisterTablesRoutes(app, con)

	req := httptest.NewRequest("POST", "http://localhost:4000/tables/testHandleDeleteTable", strings.NewReader(`{
		"name":{
			"type":"text",
			"default":"defaultText",
			"unique":true,
			"nullable":false
		}
	}`))
	req.Header.Set("Content-Type", "application/json")
	_, err = app.Test(req)
	assert.Nil(t, err)

	req = httptest.NewRequest("DELETE", "http://localhost:4000/tables/testHandleDeleteTable", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.Nil(t, err)

	payload := utils.ReadBody(resp.Body)
	assert.True(t, payload["success"].(bool), true)
}