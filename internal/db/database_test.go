package db

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatal(err)
	}
}

func TestListDatabases(t *testing.T) {
	con, err := InitDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	dbs, err := ListDatabases(con)
	assert.Nil(t, err)
	assert.Greater(t, len(dbs), 0)
}

func TestCreateDatabase(t *testing.T) {
	con, err := InitDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	rowsAffected, err := CreateDatabase(con, "mysqlmeta")
	assert.Nil(t, err)
	assert.Equal(t, rowsAffected, 1)
}
