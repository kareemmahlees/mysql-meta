// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CreateDatabaseResponse struct {
	// number of created records
	Created int `json:"created"`
}

type CreateTableData struct {
	ColName *string           `json:"colName,omitempty"`
	Props   *CreateTableProps `json:"props,omitempty"`
}

type CreateTableProps struct {
	Type     *string     `json:"type,omitempty"`
	Nullable *bool       `json:"nullable,omitempty"`
	Default  interface{} `json:"default,omitempty"`
	Unique   *bool       `json:"unique,omitempty"`
}

type CreateTableResponse struct {
	// name of created table
	Created string `json:"created"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

// Table info like field name returned from table query
type TableInfo struct {
	Field   *string     `json:"field,omitempty"`
	Type    *string     `json:"type,omitempty"`
	Null    *string     `json:"null,omitempty"`
	Key     *string     `json:"key,omitempty"`
	Default interface{} `json:"default,omitempty"`
	Extra   interface{} `json:"extra,omitempty"`
}

type UpdateTableData struct {
	Operation *UpdateTableProps `json:"operation"`
}

type UpdateTableProps struct {
	Type UpdateTableOperationTypes `json:"type"`
	// an array of string that must be specified if operation type is delete
	ColumnsToDelete []*string `json:"columnsToDelete,omitempty"`
	// a map of column name and type that must be specified if operation type is add
	ColumnsToAdd map[string]interface{} `json:"columnsToAdd,omitempty"`
	// a map of column name and type that must be specified if operation type is modify
	ColumnsToModify map[string]interface{} `json:"columnsToModify,omitempty"`
}

type UpdateTableOperationTypes string

const (
	UpdateTableOperationTypesAdd    UpdateTableOperationTypes = "add"
	UpdateTableOperationTypesModify UpdateTableOperationTypes = "modify"
	UpdateTableOperationTypesDelete UpdateTableOperationTypes = "delete"
)

var AllUpdateTableOperationTypes = []UpdateTableOperationTypes{
	UpdateTableOperationTypesAdd,
	UpdateTableOperationTypesModify,
	UpdateTableOperationTypesDelete,
}

func (e UpdateTableOperationTypes) IsValid() bool {
	switch e {
	case UpdateTableOperationTypesAdd, UpdateTableOperationTypesModify, UpdateTableOperationTypesDelete:
		return true
	}
	return false
}

func (e UpdateTableOperationTypes) String() string {
	return string(e)
}

func (e *UpdateTableOperationTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UpdateTableOperationTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UpdateTableOperationTypes", str)
	}
	return nil
}

func (e UpdateTableOperationTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
