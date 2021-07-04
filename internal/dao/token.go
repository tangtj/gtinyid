package dao

import (
	"github.com/tangtj/gtinyid/base"
	"github.com/tangtj/gtinyid/internal/model"
	"log"
)

const (
	queryToken string = "select id, biz_type,token from id_token"
)

func GetAllIdToken() ([]*model.IdToken, error) {
	db := GetDb()
	rows, err := db.Query(queryToken)
	if err != nil {
		log.Print("查询失败")
		return []*model.IdToken{}, base.DBQueryError
	}
	ret := make([]*model.IdToken, 0)
	for rows.Next() {
		token := &model.IdToken{}
		if ierr := rows.Scan(&token.Id, &token.BizType, &token.Token); ierr != nil {
			log.Printf("查询token异常 %s", ierr)
		} else {
			ret = append(ret, token)
		}
	}
	return ret, nil
}
