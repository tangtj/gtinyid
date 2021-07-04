package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tangtj/gtinyid/internal/config"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var dbClient []*sql.DB

var rd *rand.Rand

var dbSize = 0

func init() {
	rd = rand.New(rand.NewSource(time.Now().UnixNano()))
	c := config.GetConfig()

	for i, s := range c.Db {
		s = strings.TrimPrefix(s, "mysql://")
		rdb, err := sql.Open("mysql", s)
		if err != nil {
			panic("数据库连接异常，index ：" + strconv.Itoa(i))
		}
		dbClient = append(dbClient, rdb)
	}
	dbSize = len(dbClient)
}

func GetDb() *sql.DB {
	return dbClient[rd.Intn(dbSize)]
}
