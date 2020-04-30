package main

import (
	"encoding/csv"
	"fmt"
	"github.com/atotto/clipboard"
	"io"
	"strconv"
	"strings"
)

/**
クリップボードのタブ区切り文字列をテーブル定義 DDL に変換する。
カラムは、次の内容になっていること。
「列名	データ型	サイズ	PK	-	-	-	-	NOT NULL	Default」
*/
func main() {
	// クリップボードから読み込み
	data, _ := clipboard.ReadAll()
	rd := csv.NewReader(strings.NewReader(data))
	// TAB 区切り（TSV）
	rd.Comma = '\t'

	var pk []string
	firstLine := true

	fmt.Println("CREATE TABLE IF NOT EXISTS TTTTT (")
	for {
		record, err := rd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("読み込みエラー: ", err)
			break
		}
		if !firstLine {
			fmt.Print(",\n")
		}

		// PK 項目
		if len(record[3]) != 0 {
			pk = append(pk, "`"+record[0]+"`")
		}

		// カラム型を分割（bigint unsigned など）
		colTypes := strings.Split(record[1], " ")
		// 型名
		colType := strings.ToUpper(colTypes[0])
		if len(record[2]) != 0 {
			// サイズが指定されている場合
			colType += "(" + record[2] + ")"
		}
		if len(colTypes) == 2 {
			// 型が bigint unsigned のような場合
			// unsigned を大文字化して追加
			colType += " " + strings.ToUpper(colTypes[1])
		}

		notNull := ""
		if len(record[8]) != 0 {
			// NotNull の場合
			notNull = "NOT NULL"
		}

		defaultValue := ""
		if len(record[9]) != 0 {
			// デフォルト値が指定されている場合
			if len(notNull) != 0 && record[9] == "NULL" {
				// NotNull カラムなのに、デフォルト値が Null の場合
				println(record[0] + " NOT NULL Column but Default NULL")
			} else if _, err = strconv.Atoi(record[9]); err == nil || record[9] == "NULL" {
				// デフォルト値が数値の場合
				defaultValue = "DEFAULT " + record[9]
			} else {
				// デフォルト値が文字列の場合
				defaultValue = "DEFAULT '" + record[9] + "'"
			}
		}

		// 標準出力にカラムの定義を出力
		fmt.Printf("  `%s` %s %s %s", record[0], colType, notNull, defaultValue)
		firstLine = false
	}
	if len(pk) != 0 {
		// PK がある場合
		fmt.Printf(",\n  PRIMARY KEY(%s)", strings.Join(pk, ", "))
	}
	fmt.Print("\n)\n")
}
