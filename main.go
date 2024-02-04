package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	dbURL := "postgres://postgres:password@localhost:5432/google"

	//for this demonstration consider "google" as a Database in which "youtube" is a relation[table] which has two fields
	//id     --> int
	//tittle --> string

	ctx := context.Background()

	dbpool, dberr := pgxpool.Connect(ctx, dbURL)

	if dberr != nil {
		panic("Unable to connect to database")
	}

	defer dbpool.Close()

	sql := "select * from youtube ;"

	rows, sqlerr := dbpool.Query(ctx, sql)

	if sqlerr != nil {
		panic(fmt.Sprintf("QueryRow failed: %v", sqlerr))
	}

	for rows.Next() {
		var id int
		var tittle string
		rows.Scan(&id, &tittle)

		fmt.Printf("ID::: %d\t Tittle::: %s\n", id, tittle)
	}

	// fetching a specific file

	var id int
	var tittle string

	sql = "select * from youtube where id = $1 ;" //why "$1" means one variable placeholder

	err := dbpool.QueryRow(ctx, sql, 4).Scan(&id, &tittle) //fetching details of song whose id = 4

	if err != nil {
		panic(fmt.Sprintf("QueryRow failed: %v", sqlerr))
	}

	fmt.Printf("Fetched one row--> id:: %d title:: %s", id, tittle)

	//to excute somethying

	rowsByExc, err := dbpool.Exec(ctx, sql, 4)

	fmt.Print("\n")

	fmt.Println(rowsByExc.RowsAffected())

	// Transactions

	tx, err := dbpool.Begin(context.Background())
	if err != nil {
		panic("can' intitate txn")
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "insert into youtube(tittle) values('pal pal dil ke pass') ;")
	if err != nil {
		panic("can' execute txn")
	}

	err = tx.Commit(context.Background())
	if err != nil {
		panic("can' commit txn")
	}

	fmt.Println("Finnally excuted transaction")

}
