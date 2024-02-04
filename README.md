# pgX 
pgx is a pure Go driver and toolkit for PostgreSQL.

# To use
Intialize a module file like this
```go mod init github.com/shubhammishra-1```

install these libraries

```go get github.com/jackc/pgx/v5```

```go get github.com/jackc/pgx/v5/pgxpool```


# About database URL 
it is a string that is used to connect database. generally it is fetch from ```.env``` enviornment variables doesnot expose into main logic.

dbURL := ```"postgres://username:password@localhost:5432/database_name"```


## Creating connection pool

there are two ways to create connection pool..

Way 1 it is not concurrency safe

```conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))```

Way 2 it is concurrency safe

```dbpool, dberr := pgxpool.Connect(ctx, dbURL)```

## To Close created database pool after program cycle

make sure you closed database after complete excetuion of programs

```defer dbpool.Close()```

# Query Operations

to fetch some details from Psql DB use these operations.

## dbpool.Query(ctx, sqlString,Dval1,Dval2...)

it returns one or more than one rows. here Dval1,Dval2.. are dynamics variables.. which is optional in sql statement.

for example:: ```sqlString= "select * from youtube where id >= $1 AND id <= $2"```

dval1=4 and dval2=8 means fetch all rows from from 4 to  8

it returns ```rows,err:= dbpool.Query(ctx, sqlString,Dval1,Dval2...)```

to map fetched rows to our go data structure. make sure your datatypes matches with specified datatypes and with same variable names

```go
for rows.Next() {
	rows.Scan(&d1, &d2 ,...)
}
```

## dbPool.QueryRow(context.Background(), sqlString, dval,...).Scan(&var1, &var2...)
to fecth a single rows it can be used
it returns error only 

for example::: 

```go 
var name string
var weight int64
err := conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
if err != nil {
    return err
}
```

## dbpool.Exec(ctx, sqlString, dval...)

Exec() used to execute a query that does not return a result set. 
to excute some operations... it majorly used for update,delete,insert,rows affected checking only... 

for example ::

```go 
commandTag, err := conn.Exec(context.Background(), "delete from widgets where id=$1", 42)
if err != nil {
    return err
}
if commandTag.RowsAffected() != 1 {
    return errors.New("No row found to delete")
}
```

# Transactions
Transactions are started by calling Begin. 
The Tx returned from Begin also implements the Begin method. This can be used to implement pseudo nested transactions. These are internally implemented with savepoints.

```go 

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
    
```


## Why $1,$2,$3... used in SQL string

it is used to act as placeholder for some specfied variables.. which will be replaced by some values to get details...

$1 means first placeholder
$2 means second placeholder
so on...


# Reference

Please refer this documentation for some more features associated with pgX

```https://pkg.go.dev/github.com/jackc/pgx```