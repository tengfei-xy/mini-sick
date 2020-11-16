package main
// golang lib
import(
	"net/http"
	"io/ioutil"
	"fmt"
    "database/sql"
)
import (
	_ "github.com/go-sql-driver/mysql"
)
// my lib
import(
	pnt "print"

)
var DB * sql.DB

func main (){


	DB = initMySQL()
	pnt.Init("mini-sick Start!")

	http.HandleFunc("/", root)
	http.HandleFunc("/mini-sick", index)
	pnt.Info(http.ListenAndServeTLS("0.0.0.0:8080", "../ssl/mini.xunyang.site.pem", "../ssl/mini.xunyang.site.key", nil))

}
func index(w http.ResponseWriter, r *http.Request){

	r.ParseForm()
	msg,err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		pnt.Error(err)
	}

	pnt.Json(string(msg))

	w.Write(msgMain(msg))
}
func root(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("ok"))

}

func initMySQL() * sql.DB {

	Username := `root`
	Password := `if(hdc==MYSQL)`
	UnixSocket := `/tmp/mysql.sock`
	Database := `mini_sick_poda`

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