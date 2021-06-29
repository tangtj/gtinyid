package dao

import (
	"database/sql"
	"errors"
	"log"
)

const (
	querySegmentInfo = "select id, biz_type, step, incr, max_id from id_info where biz_type = ?"

	updateSegmentInfo = "update id_info set max_id = max_id + step,version = version + 1 where biz_type = ? and version = ?"
)

func getSegmentInfo(bizType string) {
	rdb := GetDb()

	tx, err := rdb.Begin()
	if err != nil {
		log.Printf("开启事物异常 %s", err)
		return
	}
	row := tx.QueryRow(querySegmentInfo, bizType)
	if e := row.Err(); e != nil {
		if errors.Is(e, sql.ErrNoRows) {

		}
	}

}
