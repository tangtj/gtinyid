package dao

import (
	"database/sql"
	"errors"
	"github.com/tangtj/gtinyid/base"
	"log"
)

const (
	querySegmentInfo = "select max_id, step, incr,version from id_info where biz_type = ?"

	updateSegmentInfo = "update id_info set max_id = max_id + step,version = version + 1 where biz_type = ? and version = ?"
)

func GetNextSegment(bizType string) (*base.Segment, error) {
	rdb := GetDb()

	tx, err := rdb.Begin()
	defer func() {
		if e := recover(); e != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err != nil {
		log.Printf("开启事物异常 %s", err)
		return nil, base.DBQueryError
	}
	row := tx.QueryRow(querySegmentInfo, bizType)
	if e := row.Err(); e != nil {
		if errors.Is(e, sql.ErrNoRows) {
			return nil, base.DBQueryNotFound
		}
	}
	var maxId, step, incr, ver int64
	if err := row.Scan(&maxId, &step, &incr, &ver); err != nil {
		log.Printf("查询号段信息异常 : %s", err)
	}
	segment := base.NewSegment(bizType, maxId, step, incr)

	result, err := tx.Exec(updateSegmentInfo, bizType, ver)
	if err != nil {
		log.Printf("更新号段信息异常 %s", err)
		return nil, base.DBExecError
	}
	affect, _ := result.RowsAffected()
	if affect > 0 {
		return segment, nil
	} else {
		return nil, base.DBUpdateNoAffect
	}
}
