package main

import (
	"context"
	_ "context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx context.Context
)

func main() {
	fmt.Println("DB Connect Example")
	db, err := sql.Open("mysql", "ysh8361:ysh8361@tcp(192.168.0.35:3456)/tutorial")
	if err != nil {
		fmt.Println("DB conn fail")
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("DB conn succ")

	err = db.Ping()
	if err != nil {
		fmt.Println("db ping fail")
	} else {
		fmt.Println("db ping Succ")
	}

	var (
		MSG_SEQ     int
		PROC_RESULT sql.NullInt64
		IN_GROUP    sql.NullString
		OUT_GROUP   sql.NullInt64
		SERIAL_NUM  sql.NullString
		CALLER_NO   sql.NullString
		CALLEE_NO   sql.NullString
		CALLBACK_NO sql.NullString
		IN_TIME     sql.NullString
		REPORT_TIME sql.NullString
	)

	var querys = "SELECT " +
		"MSG_SEQ, PROC_RESULT, IN_GROUP, " +
		"OUT_GROUP, SERIAL_NUM, CALLER_NO, " +
		"CALLEE_NO, CALLBACK_NO, IN_TIME, " +
		"REPORT_TIME " +
		"FROM TSMS_SND_MSG;"
	rows, err := db.Query(querys)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&MSG_SEQ,
			&PROC_RESULT,
			&IN_GROUP,
			&OUT_GROUP,
			&SERIAL_NUM,
			&CALLER_NO,
			&CALLEE_NO,
			&CALLBACK_NO,
			&IN_TIME,
			&REPORT_TIME); err != nil {

			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		fmt.Print("RESULT : ")
		fmt.Print(" MSG_SEQ : ", MSG_SEQ)
		if PROC_RESULT.Valid {
			fmt.Print(" PROC_RESULT : ", PROC_RESULT.Int64)
		} else {
			fmt.Print(" PROC_RESULT : null")
		}
		if IN_GROUP.Valid {
			fmt.Print(" IN_GROUP : ", IN_GROUP.String)
		} else {
			fmt.Print(" IN_GROUP : null")
		}
		if OUT_GROUP.Valid {
			fmt.Print(" OUT_GROUP : ", OUT_GROUP.Int64)
		} else {
			fmt.Print(" OUT_GROUP : null")
		}
		if SERIAL_NUM.Valid {
			fmt.Print(" SERIAL_NUM : ", SERIAL_NUM.String)
		} else {
			fmt.Print(" SERIAL_NUM : null")
		}
		if CALLER_NO.Valid {
			fmt.Print(" CALLER_NO : ", CALLER_NO.String)
		} else {
			fmt.Print(" CALLER_NO : null")
		}
		if CALLEE_NO.Valid {
			fmt.Print(" CALLEE_NO : ", CALLEE_NO.String)
		} else {
			fmt.Print(" CALLEE_NO : null")
		}
		if CALLBACK_NO.Valid {
			fmt.Print(" CALLBACK_NO : ", CALLBACK_NO.String)
		} else {
			fmt.Print(" CALLBACK_NO : null")
		}
		if IN_TIME.Valid {
			fmt.Print(" IN_TIME : ", IN_TIME.String)
		} else {
			fmt.Print(" IN_TIME : null")
		}
		if REPORT_TIME.Valid {
			fmt.Print(" REPORT_TIME : ", REPORT_TIME.String)
		} else {
			fmt.Print(" REPORT_TIME : null")
		}
		fmt.Println(" ]")

	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
