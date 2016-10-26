package main

import (
	"bytes"
	"testing"
)

func TestSqlTypeString(t *testing.T) {
	scanner := Scanner{
		Package:   "test",
		Primative: "string",
		Type:      "TestType",
	}

	buf := &bytes.Buffer{}

	if err := SqlType(buf, scanner); err != nil {
		t.Fatal(err)
	}

	if buf.Len() == 0 {
		t.Fatalf("Nothing generated")
	}
}

func TestSqlTypeInt(t *testing.T) {
	scanner := Scanner{
		Package:   "test",
		Primative: "int",
		Type:      "TestType",
	}

	buf := &bytes.Buffer{}

	if err := SqlType(buf, scanner); err != nil {
		t.Fatal(err)
	}

	if buf.Len() == 0 {
		t.Fatalf("Nothing generated")
	}
}

func TestSqlTypeBool(t *testing.T) {
	scanner := Scanner{
		Package:   "test",
		Primative: "bool",
		Type:      "TestType",
	}

	buf := &bytes.Buffer{}

	if err := SqlType(buf, scanner); err != nil {
		t.Fatal(err)
	}

	if buf.Len() == 0 {
		t.Fatalf("Nothing generated")
	}
}

func TestSqlTypePythonDict(t *testing.T) {
	scanner := Scanner{
		Package:   "test",
		Primative: "pythondict",
		Type:      "TestType",
	}

	buf := &bytes.Buffer{}

	if err := SqlType(buf, scanner); err != nil {
		t.Fatal(err)
	}

	if buf.Len() == 0 {
		t.Fatalf("Nothing generated")
	}
}

func TestSqlTypePythonList(t *testing.T) {
	scanner := Scanner{
		Package:   "test",
		Primative: "pythonlist",
		Type:      "TestType",
	}

	buf := &bytes.Buffer{}

	if err := SqlType(buf, scanner); err != nil {
		t.Fatal(err)
	}

	if buf.Len() == 0 {
		t.Fatalf("Nothing generated")
	}
}
