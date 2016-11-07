package main

import (
	//"database/sql"
	//"database/sql/driver"
	"bytes"
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const CreateSQL = `
CREATE TABLE IF NOT EXISTS temporary (
    id character varying(36) NOT NULL,
    field1 boolean,
    field2 timestamp without time zone,
    field3 bytea,
    field4 bigint,
    field5 double precision,
    field6 bytea,
    field7 text,
	field8 text
);
`

//const SelectSQL = "SELECT id, field1, field2, field3, field4, field5, field6, field7 FROM temporary WHERE id=$1"
const SelectSQL = "SELECT id, field1, field3, field4, field5, field7, field8 FROM temporary WHERE id=$1"

//go:generate sqltype -primative string -type Id
type Id string

//go:generate sqltype -primative bool -type Bool
type Bool bool

type Int int

type Time time.Time

//go:generate sqltype -primative pythondict -type Dictionary
type Dictionary map[string]interface{}
type Float float32
type List string

//go:generate sqltype -primative string -type String
type String string

type Primatives struct {
	Id     string    `db:"id"`
	Field1 bool      `db:"field1"`
	Field2 time.Time `db:"field2"`
	Field3 string    `db:"field3"`
	Field4 int       `db:"field4"`
	Field5 float32   `db:"field5"`
	Field6 string    `db:"field6"`
	Field7 string    `db:"field7"`
	Field8 *string   `db:"field8"`
}

func (t Primatives) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Id     string %s\n", t.Id)
	fmt.Fprintf(buf, "Field1 bool   %t\n", t.Field1)
	fmt.Fprintf(buf, "Field2 time   %v\n", t.Field2)
	fmt.Fprintf(buf, "Field3 dict   %s\n", t.Field3)
	fmt.Fprintf(buf, "Field4 int    %d\n", t.Field4)
	fmt.Fprintf(buf, "Field5 float  %f\n", t.Field5)
	fmt.Fprintf(buf, "Field6 list   %s\n", t.Field6)
	fmt.Fprintf(buf, "Field7 string %s\n", t.Field7)
	fmt.Fprintf(buf, "Field8 string %s\n", t.Field8)
	return buf.String()
}

type Customs struct {
	Id     Id   `db:"id"`
	Field1 Bool `db:"field1"`
	//Field2 Time       `db:"field2"`
	Field3 Dictionary `db:"field3"`
	Field4 Int        `db:"field4"`
	Field5 Float      `db:"field5"`
	//Field6 List       `db:"field6"`
	Field7 String `db:"field7"`
	Field8 String `db:"field8"`
}

func (t Customs) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Id     String %s\n", t.Id)
	fmt.Fprintf(buf, "Field1 Bool   %t\n", t.Field1)
	//fmt.Fprintf(buf, "Field2 Time   %v\n", t.Field2)
	fmt.Fprintf(buf, "Field3 Dict   %s\n", t.Field3)
	fmt.Fprintf(buf, "Field4 Int    %d\n", t.Field4)
	fmt.Fprintf(buf, "Field5 Float  %f\n", t.Field5)
	//fmt.Fprintf(buf, "Field6 List   %s\n", t.Field6)
	fmt.Fprintf(buf, "Field7 String %s\n", t.Field7)
	fmt.Fprintf(buf, "Field8 String %s\n", t.Field8)
	return buf.String()
}

func main() {
	flag.Parse()

	db, err := DBConnection()
	CheckError(err)

	SetupDB(db)

	primative := Primatives{}
	CheckError(
		db.Get(
			&primative,
			SelectSQL,
			"test01",
		))
	fmt.Println(primative)

	custom := Customs{}
	CheckError(
		db.Get(
			&custom,
			SelectSQL,
			"test01",
		))
	fmt.Println(custom)
}

func SetupDB(db *sqlx.DB) {
	db.MustExec(CreateSQL)

	tx := db.MustBegin()
	tx.MustExec("DELETE FROM temporary")
	for i := 0; i < 10; i++ {
		tx.MustExec(
			`INSERT INTO temporary (id, field1, field2, field3, field4, field5, field6, field7, field8)
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			fmt.Sprintf("test%02d", i),
			true,
			time.Now().UTC(),
			Dictionary{
				"foo": "bar",
			},
			34,
			4.5,
			"[]",
			"hello, world",
			nil,
		)
	}
	tx.Commit()
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func DBConnection() (*sqlx.DB, error) {
	dsn := "host=localhost user=steve dbname=steve port=5432 sslmode=disable"
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
