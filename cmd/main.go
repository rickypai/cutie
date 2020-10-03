package main

import (
	"log"
	"os"

	"github.com/jackc/pgconn"
	"github.com/rickypai/cutie/modelgen"
	"github.com/rickypai/cutie/table"
)

func main() {
	tables, err := table.TablesFromConfig(".cutie.yaml")
	if err != nil {
		log.Fatalf("error loading tables: %s", err)
	}

	pgConfig, err := pgconn.ParseConfig(os.Getenv("DATABASE_URL"))

	err = table.DumpAll(pgConfig, tables)
	if err != nil {
		log.Fatalf("error dumping tables: %s", err)
	}

	err = modelgen.GenerateSQLCModels(tables)
	if err != nil {
		log.Fatalf("error generating SQLC models: %s", err)
	}

	err = modelgen.GenerateSQLCModelMocks(tables)
	if err != nil {
		log.Fatalf("error generating SQLC mocks: %s", err)
	}

	log.Println("All done!")
}
