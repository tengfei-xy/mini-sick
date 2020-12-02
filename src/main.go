package main
// golang lib
import(
	"net/http"
	"io/ioutil"
	"fmt"
	"database/sql"
	"path/filepath"
)
import (
	_ "github.com/go-sql-driver/mysql"
)
// my lib
import(
	pnt "print"
)
//https://mini.xunyang.site:8080/avcjixode5nf2sdzo4ign/
const podid string = "avcjixode5nf2sdzo4ign"
var DB * sql.DB
func main (){

	DB = initMySQL()
	pnt.Init("mini-sick-poda Start!")

	go http.HandleFunc("/"+podid+"/", root)
	go http.HandleFunc("/"+podid +"/mini-sick", index)
	pnt.Info(http.ListenAndServeTLS("0.0.0.0:8080", "../ssl/mini.xunyang.site.pem", "../ssl/mini.xunyang.site.key", nil))

}
func index(w http.ResponseWriter, r *http.Request){
	pnt.IP(r.RemoteAddr)
	msg,err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		pnt.Error(err)
	}
	pnt.Json(string(msg))
	w.Write(msgMain(msg))

}
func root(w http.ResponseWriter, r *http.Request){
        var reqURLExt string = filepath.Ext(r.URL.Path)
        var reqURL string = r.URL.Path
        var res []byte
        //var err error

        switch reqURLExt {
        case ".css":
                w.Header().Set("Content-Type", "text/css")
        case ".png":
                w.Header().Set("Content-Type", "image/png")
        case ".ico":
                w.Header().Set("Content-Type", "image/x-ico")
        case ".js":
                w.Header().Set("Content-Type", "application/javascript")
        case ".jpg":
                w.Header().Set("Content-Type", "image/jpeg")
        default:
                w.Header().Set("Content-Type", "text/html")
		}
        switch reqURL {
        case "/avcjixode5nf2sdzo4ign/":
                res,_ = ioutil.ReadFile(".." + reqURL +"login.html")
        default:
                res,_ = ioutil.ReadFile(".."+reqURL)
		}

        w.Write(res)
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