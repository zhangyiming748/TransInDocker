package sql

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/nalgeon/redka"
	"os"
	"strings"
	"zhangyiming748/TransInDocker/constant"
)

var (
	db *redka.DB
)

func GetDatabase() *redka.DB {
	return db
}
func Initial() {
	location := strings.Join([]string{constant.GetRoot(), "trans.db"}, string(os.PathSeparator))
	db, _ = redka.Open(location, nil)
}
