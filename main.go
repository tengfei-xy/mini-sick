package main
// golang lib
import(
	"net/http"
	"io/ioutil"
	"fmt"
	"database/sql"
	"strconv"
	"strings"
	"path/filepath"
	"os"
)
import(
	_ "github.com/go-sql-driver/mysql"
)
// my lib
import(
	pnt "print"
)
var port		string
var podID		string 
var DB 		*	sql.DB

func mainInitEnv(){
	// 将所在路径作为端口
	ex, err := os.Executable()
    if err != nil {
        panic(err)
	}
	path := filepath.Dir(ex)
	pathindex := strings.LastIndex(path,"/") +1
	podID = path[pathindex:len(path)]
	port = "80" + podID

	// 将所在路径作为数据库名称
	x,_ := strconv.Atoi(podID)
	podID = fmt.Sprintf("%x",x)
	DB = mainInitMySQL()  

}

func mainInitMySQL() * sql.DB {

	Username := `root`
	Password := `if(hdc==MYSQL)`
	UnixSocket := `/tmp/mysql.sock`
	Database := `mini_sick_pod`  + podID

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
	pnt.Info("MySQL connection successful")

	return db
}

func main (){
	mainInitEnv()
	pnt.Infof("PodID:%s,Port:%s,Start!",podID,port)
	http.HandleFunc("/", index)
	go pnt.Info(http.ListenAndServe("0.0.0.0:" + port, nil))

}
func index(w http.ResponseWriter, r *http.Request){
	msg,err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		pnt.Error(err)
	}
	pnt.Json(string(msg))
	w.Write(msgMain(msg))
}