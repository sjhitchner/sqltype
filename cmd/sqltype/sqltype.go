package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"go/build"
	"go/format"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

// value is one of these
//    int64
//    float64
//    bool
//    []byte
//    string
//    time.Time
//    nil - for NULL values

const ScannerTemplate = `
package {{ .Package }}

import (
	"fmt"
	"database/sql/driver"
)

func (t *{{ .Type }}) Scan(value interface{}) error {
	s, ok := value.({{ .Primative }})
	if !ok {
		return fmt.Errorf("Can't convert %v to {{ .Primative }}", value)
	}

	*t = {{ .Type }}(s)
	return nil
}

func (t {{ .Type }}) Value(v interface{}) (driver.Value, error) {
	return {{ .Primative }}(t), nil
}
`

type Scanner struct {
	Package   string
	Primative string
	Type      string
}

var (
	primativeType string
	typeName      string
)

func init() {
	flag.StringVar(&primativeType, "primative", "", "Corresponding primative type")
	flag.StringVar(&typeName, "type", "", "Name of type")
}

func main() {
	flag.Parse()

	if typeName == "" {
		log.Fatal("Need to provide --type")
	}

	if primativeType == "" {
		log.Fatal("Need to provide --primative")
	}

	packageName, err := getPackageName()
	if err != nil {
		log.Fatal(err)
	}

	scanner := Scanner{
		Package:   packageName,
		Primative: primativeType,
		Type:      typeName,
	}

	filename := fmt.Sprintf("%s_sql.go", strings.ToLower(typeName))
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := SqlType(file, scanner); err != nil {
		log.Fatal(err)
	}
}

func SqlType(writer io.Writer, scanner Scanner) error {
	t := template.Must(template.New("scanner").Parse(ScannerTemplate))

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, scanner); err != nil {
		return errors.Wrap(err, "error excuting template")
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "error sourcing go struct")
	}

	if _, err := writer.Write(b); err != nil {
		return errors.Wrap(err, "error writing to file")
	}

	return nil
}

func getPackageName() (string, error) {
	directory := "."

	pkg, err := build.Default.ImportDir(directory, 0)
	if err != nil {
		return "", err
	}

	return pkg.Name, err
}
