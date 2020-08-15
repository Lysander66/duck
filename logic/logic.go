package logic

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"text/tabwriter"
	"unicode"
)

var dict = map[string]string{
	"tinyint":   "int32",
	"smallint":  "int32",
	"int":       "int64",
	"float":     "float32",
	"double":    "float64",
	"varchar":   "string",
	"char":      "string",
	"datetime":  "time.Time",
	"date":      "time.Time",
	"time":      "time.Time",
	"timestamp": "time.Time",
}

func GenStruct(schema, name, mysql string) string {
	db, err := gorm.Open("mysql", mysql)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true)
	// db.LogMode(true)

	var result []column
	sql := `SELECT * FROM INFORMATION_SCHEMA.COLUMNS WHERE table_schema = ? AND table_name = ?`
	db.Raw(sql, schema, name).Scan(&result)

	model := "type " + snake2Pascal(name) + " struct {\n"
	var buff bytes.Buffer
	w := new(tabwriter.Writer)
	w.Init(&buff, 4, 4, 4, ' ', 0)
	for _, v := range result {
		dataType, ok := dict[v.DataType]
		if !ok {
			dataType = "string"
		}
		if v.ColumnKey == "PRI" {
			fmt.Fprintf(w, "\t%s\t%s\t`gorm:\"primary_key;column:%s\"`\n", snake2Pascal(v.ColumnName), dataType, v.ColumnName)
			continue
		}
		fmt.Fprintf(w, "\t%s\t%s\t`gorm:\"column:%s\"`\n", snake2Pascal(v.ColumnName), dataType, v.ColumnName)
	}
	w.Flush()
	model += buff.String()
	model += fmt.Sprintf("}\n\nfunc (t *%s) TableName() string {\n", snake2Pascal(name))
	model += fmt.Sprintf("\treturn \"%s\"\n}", name)
	return model
}

type column struct {
	ColumnName             string `gorm:"column:COLUMN_NAME"`
	ColumnType             string `gorm:"column:COLUMN_TYPE"`
	DataType               string `gorm:"column:DATA_TYPE"`
	CharacterMaximumLength string `gorm:"column:CHARACTER_MAXIMUM_LENGTH"`
	IsNullable             string `gorm:"column:IS_NULLABLE"`
	ColumnDefault          string `gorm:"COLUMN_DEFAULT"`
	ColumnKey              string `gorm:"COLUMN_KEY"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT"`
}

func snake2Pascal(word string) string {
	return snake2Camel("_" + word)
}

func snake2Camel(word string) string {
	buffer := new(bytes.Buffer)
	var flag bool
	for _, c := range word {
		if c == '_' {
			flag = true
			continue
		}
		if flag {
			buffer.WriteRune(unicode.ToUpper(c))
			flag = false
		} else {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}

func camel2Snake(word string) string {
	buffer := new(bytes.Buffer)
	for i, c := range word {
		if unicode.IsUpper(c) {
			if i != 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(c))
		} else {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}
