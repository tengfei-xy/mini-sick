package main

// golang lib
import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	// my lib
	pnt "print"
)

var PodID int
var DB *sql.DB

func mainInitEnv() (string, string) {
	pPort := flag.String("port", "000", "请指定外部端口,百位数0为生产环境,1为测试环境,默认:000")
	pDB := flag.String("db", "000", "请指定数据库,默认为ID值")
	flag.Parse()

	pnt.Init(fmt.Sprintf("Port:%s,DB:%s,Start!", *pPort, *pDB))
	PodID, _ = strconv.Atoi(*pDB)
	return *pPort, *pDB
}

func mainInitMySQL(dbname string) *sql.DB {

	Username := `root`
	Password := `if(hdc==MYSQL)`
	UnixSocket := `/tmp/mysql.sock`
	Database := `mini_sick_` + dbname

	linkAddress := fmt.Sprintf("%s:%s@%s(%s)/%s", Username, Password, "unix", UnixSocket, Database)

	// 启动连接
	db, err := sql.Open("mysql", linkAddress)
	if err != nil {
		panic(err)
	}
	// 连接测试
	if err = db.Ping(); err != nil {
		panic(err)
	}
	pnt.Init("MySQL connection successful")

	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		pnt.Error(err)
	}
	w.Write(msgMain(msg))
}

func main() {
	port, db := mainInitEnv()
	DB = mainInitMySQL(db)
	http.HandleFunc("/", index)
	go pnt.Info(http.ListenAndServe("0.0.0.0:8"+port, nil))

}
