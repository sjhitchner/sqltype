package main

import (
	//"database/sql"
	//"database/sql/driver"
	"bytes"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

const CREATE = `
CREATE TABLE temporary (
    field2 timestamp without time zone,
    field3 bytea,
    field1 boolean,
    field6 bytea,
    field7 text,
    field4 bigint,
    field5 double precision,
    id character varying(36) NOT NULL
);


ALTER TABLE public.temporary OWNER TO kiip;

--
-- Name: temporary_pkey; Type: CONSTRAINT; Schema: public; Owner: kiip; Tablespace: 
--

ALTER TABLE ONLY temporary
    ADD CONSTRAINT temporary_pkey PRIMARY KEY (id);

`

type Primatives struct {
	Id     string    `db:"id"`
	Field1 bool      `db:"field1"`
	Field2 time.Time `db:"field2"`
	//Field3 Dictionary `db:"field3"`
	Field3 string  `db:"field3"`
	Field4 int     `db:"field4"`
	Field5 float32 `db:"field5"`
	//Field6 List       `db:"field6"`
	Field6 string `db:"field6"`
	Field7 string `db:"field7"`
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
	return buf.String()
}

func TestDatabasePrimativeSerialization(t *testing.T) {
	db := DBConnection(t)
	temp := Primatives{}
	err := db.Get(&temp, "SELECT id, field1, field2, field3, field4, field5, field6, field7 FROM temporary WHERE id=$1", "test01")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(temp)
}

type Customs struct {
	Id     Id         `db:"id"`
	Field1 Bool       `db:"field1"`
	Field2 Time       `db:"field2"`
	Field3 Dictionary `db:"field3"`
	Field4 Int        `db:"field4"`
	Field5 Float      `db:"field5"`
	Field6 List       `db:"field6"`
	Field7 String     `db:"field7"`
}

func (t Customs) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Id     string %s\n", t.Id)
	fmt.Fprintf(buf, "Field1 bool   %t\n", t.Field1)
	fmt.Fprintf(buf, "Field2 time   %v\n", t.Field2)
	fmt.Fprintf(buf, "Field3 dict   %s\n", t.Field3)
	fmt.Fprintf(buf, "Field4 int    %d\n", t.Field4)
	fmt.Fprintf(buf, "Field5 float  %f\n", t.Field5)
	fmt.Fprintf(buf, "Field6 list   %s\n", t.Field6)
	fmt.Fprintf(buf, "Field7 string %s\n", t.Field7)
	return buf.String()
}

func TestDatabaseCustomSerialization(t *testing.T) {
	db := DBConnection(t)
	temp := Customs{}
	err := db.Get(&temp, "SELECT id, field1, field2, field3, field4, field5, field6, field7 FROM temporary WHERE id=$1", "test01")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(temp)
}

func DBConnection(t *testing.T) *sqlx.DB {
	//dsn := "host=localhost user=steve password='kiip' dbname=kiip port=5432 sslmode=disable"
	dsn := "host=localhost user=steve dbname=steve port=5432 sslmode=disable"
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}

	Setup(t, conn)
	return conn
}

func Setup(t *testing.T, db *sqlx.DB) {
	Teardown(t, db)

	tx := db.MustBegin()
	for i := 0; i < 10; i++ {
		tx.MustExec(
			"INSERT INTO temporary (id, field1, field2, field3, field4, field5, field6, field7) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
			fmt.Sprintf("test%02d", i),
			true,
			time.Now().UTC(),
			"{}",
			34,
			4.5,
			"[]",
			"hello, world",
		)
	}
	tx.Commit()
}

func Teardown(t *testing.T, db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM temporary")
	tx.Commit()
}

type Id string
type Bool bool
type Time time.Time
type Dictionary string
type Int int
type Float float32
type List string
type String string
