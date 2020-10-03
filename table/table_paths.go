package table

import (
	"path/filepath"
)

func (t *Table) SchemaSQLPath() string {
	return filepath.Join(t.SchemaTablesDirPath(), t.SQLFilename())
}

func (t *Table) QuerySQLPath() string {
	return filepath.Join(t.SchemaQueriesDirPath(), t.SQLFilename())
}

func (_ *Table) SchemaTablesDirPath() string {
	return filepath.Join("schema", "tables")
}

func (_ *Table) SchemaQueriesDirPath() string {
	return filepath.Join("schema", "queries")
}

func (t *Table) GoModelDirPath() string {
	return filepath.Join("gomodels", t.Filename())
}

func (t *Table) QuerierInterfaceGoPath() string {
	return filepath.Join(t.GoModelDirPath(), "querier.go")
}

func (t *Table) QuerierMockGoPath() string {
	return filepath.Join(t.GoModelDirPath(), "mock_"+t.Filename(), "querier_mock.go")
}