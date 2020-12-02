package print
import(
	"fmt"
	"time"
)
func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func Init(s interface{}){
	fmt.Println(now()," Init      ",s)
}
func IP(ip string){
	fmt.Println(now()," IP        ",ip)
}
func Info(s interface{}){
	fmt.Println(now()," Info      ",s)
}
func Infof(f string,a ...interface{}){
	fmt.Println(now()," Info      ",fmt.Sprintf(f,a...))
}
func MySQL(s ...interface{}){
	fmt.Print(now(),"  MySQL     ")
	fmt.Println(s...)
}
func Json(s interface{}){
	fmt.Println(now()," Json      ",s)
}
func Request(s string){
	fmt.Println(now()," Request   ",s)
}
func Search(f string,s ...interface{}){
	fmt.Println(now()," Search    ",fmt.Sprintf(f,s...))
}
func Error(s error){
	fmt.Println(now()," Error     ",s)
}
func Errorf(f string,a ...interface{}){
	fmt.Println(now()," Error     ",fmt.Sprintf(f,a...))
}
func Errorwd(v interface {},s error){
	fmt.Println(now()," Error     ",v,s)
}