package database

import ("database/sql"
	      "github.com/go-sql-driver/mysql"
)

db, err := sql.Open("")

if err != nil{
  panic(err)
}
