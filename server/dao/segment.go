package dao

import (
	"database/sql"
	"errors"
	"github.com/tangtj/gtinyid/server/errorx"
	"github.com/tangtj/gtinyid/server/model"
	"log"
)

const (
	querySegmentInfo = "select biz_type,max_id, step, incr,version from id_info where biz_type = ?"

	updateSegmentInfo = "update id_info set max_id = max_id + step,version = version + 1 where biz_type = ? and version = ?"
)

func GetNextSegment(bizType string) (*model.SegmentId, error) {
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
		return nil, errorx.DBQueryError
	}
	row := tx.QueryRow(querySegmentInfo, bizType)
	if e := row.Err(); e != nil {
		if errors.Is(e, sql.ErrNoRows) {
			return nil, errorx.DBQueryNotFound
		}
	}
	s := &model.SegmentId{}
	ver := 0
	if err := row.Scan(&s.BizType, &s.MaxId, &s.Step, &s.Incr, &ver); err != nil {
		log.Printf("查询号段信息异常 : %s", err)
	}
	s.CurrentId, s.StartId = s.MaxId, s.MaxId
	s.MaxId = s.StartId + s.Step

	result, err := tx.Exec(updateSegmentInfo, bizType, ver)
	if err != nil {
		log.Printf("更新号段信息异常 %s", err)
		return nil, errorx.DBExecError
	}
	affect, _ := result.RowsAffected()
	if affect > 0 {
		return s, nil
	} else {
		return nil, errorx.DBUpdateNoAffect
	}
}
