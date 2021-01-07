package main

// golang lib
import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	// my lib
	pnt "print"
)

var port string
var podID string
var DB *sql.DB

func mainInitEnv() (string, string) {
	podid := flag.String("id", "000", "请指定ID,默认:000")
	db := flag.String("db", "", "请指定数据库,默认为ID值")
	flag.Parse()

	// 如果数据库未指定，则使用podid指定
	// 用于分开指定，比如测试实例为8000端口，而数据库连接到001
	if *db == "" {
		db = podid
	}
	pnt.Init(fmt.Sprintf("PodID:%s,Start!", *podid))

	return *podid, *db
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
	pnt.Infof("MySQL:%s connection successful!", dbname)

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
