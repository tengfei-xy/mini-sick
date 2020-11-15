package main
import (
    "database/sql"
    "encoding/json"
    "bytes"
    "io"
    "fmt"
	"time"
	"strings"
	"strconv"
)
import (
    _ "github.com/go-sql-driver/mysql"
)
import(
    pnt "print"
)
//
//
// json base function
//
//
func parseJSON(unmsg * []byte, v interface{}) error {

    dec := json.NewDecoder(bytes.NewReader(*unmsg))
    for {
        if err := dec.Decode(&v); err == io.EOF {
            break
        } else if err != nil {
            return err
        }
    }
    return nil
}

func reParseJson(v interface{}) []byte{ 
    textbyte,err := json.Marshal(v)
    if err !=nil {
        pnt.Errorwd(v,err)
    }
    return textbyte
}
//
//
// ans
//
//
func (a * ans) set(s int,e,d string){
    a.Status = s
    a.Explain = e
    a.Data = d

}
func getAns(status int,explain string,data string) ans{
    var a ans
    a.set(status,explain,data)
    return a
}

//
//
// 判断 医生登录
//
//
func (ui * userInfo) msgMain() []byte{
    // var u user
    // type user struct{
    //     Name string
    //     Password string
    // }
    var    Name string
    var    Password string
    pnt.Search("医生登录-姓名:%s,账号ID:%s,密码:%s\n",ui.Name,ui.Account,ui.Password)
    err := DB.QueryRow("SELECT name,password FROM users WHERE account=?",ui.Account).Scan(&Name,&Password)
    if err != nil{
        pnt.Error(err)
        return reParseJson(getAns(1,"登录错误！",""))

    }
    if Name=="" && ui.Name == ""{
        return reParseJson(getAns(1,"第一次登录其先输入姓名信息！",""))
    }

    // if Name == ""{
    //     _,err:= DB.Exec("update users SET name=? where account=?",ui.Name,ui.Account)
    //     if err != nil{
    //         pnt.Error(err)
    //         return reParseJson(getAns(1,"登录失败！",""))
    //     }
    // }
    if ui.Password != Password{
        return reParseJson(getAns(1,"登录失败！",""))
    }
    return reParseJson(getAns(0,"登录成功！",Name))
    
}
//
//
// 更新 患者基本信息
//
//
func (si * sickerInfo) msgMain() []byte {
    userid := createUserID()
    
    _,err := DB.Exec("insert into sicker (userid,name,age,gender,telphone,hospital_number,attandance_number,disease,out_hospital,writer) values (?,?,?,?,?,?,?,?,?,?)",
        userid,
        si.Name,
        si.Age,
        si.Gender,
        si.Telphone,
        si.Hospital_number,
        si.Attandance_number,
        si.Disease,
        "",
        si.Writer)

    if err != nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"添加失败！",""))

    }else{
        pnt.MySQL(si.Name,userid+" 添加患者成功")
        return reParseJson(getAns(1,"添加成功！",userid))
    }
}

//
//
// 更新 患者信息的风险评估
//
//
func (ri * riskInfo) msgMain() []byte {
    var notMedicationWord string
    var preProgramWord string
    var commentWord string
    var diy bool = false
    
    // 多选 非药物因素
    for _,i := range ri.Not_medication{
        notMedicationWord +=  i + ","
    }

    // 多选 删除 预处理止吐方案 选填留白
    for _,i := range ri.Pre_program{
        if i == "diy" {
            diy = true
        }
        preProgramWord += i + ","
        
    }
    
    if !diy {ri.Pre_program_diy = ""}
    diy = false

    // 删除 备注 选填留白
    for _,i := range ri.Comment{
        if i == "diy" {
            diy = true
        }
        commentWord +=  i + ","
    }
    if !diy {ri.Comment_diy = ""}
    diy = false

    // 生成MySQL特定字段
    notMedicationWord = notMedicationWord[0:len(notMedicationWord)-1]
    preProgramWord = preProgramWord[0:len(preProgramWord)-1]
    commentWord = commentWord[0:len(commentWord)-1]

    _,err := DB.Exec("insert into risk values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
        ri.Userid 	     	,
        ri.Cycle_seq    	,
        ri.Program        	,
        notMedicationWord	,
        ri.Medication     	,
        ri.Grand          	,
        preProgramWord		,
        ri.Pre_program_diy 	,
        commentWord 		,
        ri.Comment_diy     	,
        ri.Writer       	,
        ri.Assessment_date 	,
        ri.Assessment_time 	,
        ri.Assessment_timestamp,
        ri.Chemotherapy_date ,     
        ri.Chemotherapy_time ,     
        ri.Chemotherapy_timestamp)
        
    if err != nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"更新失败！",""))
    }

    _,erra := DB.Exec("insert into cycle values (?,?,?,?,?)",
        ri.Userid 	     	,
        ri.Cycle_seq    	,
        ri.Assessment_date 	,
        ri.Assessment_time 	,
        ri.Assessment_timestamp)

    if erra != nil{
        pnt.Error(erra)
        return reParseJson(getAns(0,"更新失败！",""))
    }else{
        pnt.MySQL(ri.Userid," 风险评估 更新成功")
        return reParseJson(getAns(1,"更新成功！",""))
    }
    
     
}
//
//
// 更新 患者信息的护理评估
//
//
func (ni * nurseInfo) msgMain() []byte {

    var MeasureWord string
    // 多选 非药物因素
    for _,i := range ni.Measure{
        MeasureWord +=  i + ","
    }
    MeasureWord = MeasureWord[0:len(MeasureWord)-1]
    
    // 插入数据
    _,err := DB.Exec(
        "insert into nurse values (?,?,?,?,?,?,?,?,?,?,?,?)",
        ni.Userid               ,
        ni.Cycle_seq            ,
        ni.Nurse_seq            ,

        ni.Nausea_assessment    ,
        ni.Emesis_assessment    ,
        MeasureWord             ,
        ni.Comment              ,
        ni.Out_hospital         ,

        ni.Writer               ,
        ni.Assessment_date      ,
        ni.Assessment_time      ,
        ni.Assessment_timestamp)

    if err != nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"更新失败！",""))

    }
    t := ni.Assessment_date + " " + ni.Assessment_time
    if ni.Out_hospital == "1"{
        _,err := DB.Exec("UPDATE sicker SET out_hospital=?,follow_over=? where userid=?",t,"2",ni.Userid)
        if err !=nil{
            pnt.Error(err)
            return reParseJson(getAns(0,"更新失败！",""))
        }
    }
    pnt.MySQL(fmt.Sprintf("插入护理评估 患者ID:%s,化疗周期:%d,护理次序:%d 成功",ni.Userid,ni.Cycle_seq,ni.Nurse_seq))
    return reParseJson(getAns(1,"更新成功！",""))
    

}

//
//
// 更新 患者信息 随访
//
//
func (fi * followInfo) msgMain() []byte {

    var outContentWord string
    var diy bool = false
    var satisfaction_total int
    // 多选 非药物因素
    for _,i := range fi.Out_content{
        if i == "diy" {
            diy = true
        }
        outContentWord +=  i + ","
    }
    if !diy {fi.Out_content_diy = ""}
    outContentWord = outContentWord[0:len(outContentWord)-1]

    // 计算 满意调查分数
    if fi.Satisfaction_if==1 {
        if fi.Satisfaction_1 == "1" { satisfaction_total+=20}
        if fi.Satisfaction_2 == "1" { satisfaction_total+=20}
        if fi.Satisfaction_3 == "1" { satisfaction_total+=20}
        if fi.Satisfaction_4 == "1" { satisfaction_total+=20}
        if fi.Satisfaction_5 == "1" { satisfaction_total+=20}
    }else{
        fi.Satisfaction_1 = ""
        fi.Satisfaction_2 = ""
        fi.Satisfaction_3 = ""
        fi.Satisfaction_4 = ""
        fi.Satisfaction_5 = ""
    }

    // 插入数据
    _,err := DB.Exec(
        "insert into follow values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
        fi.Userid                  ,
        fi.Follow_seq              ,
        fi.Hight_risk              ,
        fi.Emesis_grade            ,
        fi.Nausea_grade            ,
        outContentWord             ,
        fi.Out_content_diy         ,
        fi.Follow_over             ,
        fi.Satisfaction_1          ,
        fi.Satisfaction_2          ,
        fi.Satisfaction_3          ,
        fi.Satisfaction_4          ,
        fi.Satisfaction_5          ,
        satisfaction_total         ,
        fi.Writer 				   ,
        fi.Follow_follow_date      ,
        fi.Follow_follow_time      ,
        fi.Follow_follow_timestamp)

        if err != nil{
            pnt.Error(err)
            return reParseJson(getAns(0,"更新失败！",""))
        }
        
        // 随访结束
        if fi.Follow_over == "1"{
            _,err := DB.Exec("UPDATE sicker set follow_over=? where userid=?",1,fi.Userid)
            if err != nil{
                pnt.Error(err)
                return reParseJson(getAns(0,"更新失败！",""))
            }
        }
        pnt.MySQL(fmt.Sprintf("%s %s 护理评估 更新成功",fi.Userid,fi.Follow_seq))
        return reParseJson(getAns(1,"更新成功！",""))
        
}

//
//
// 搜索 患者
//
//
func (ss * searchSicker) msgMain() []byte {
    n := ss.Name 
    h := ss.Hospital_number
    a := ss.Attandance_number
    var res searchSickerRes
    t := 0 
    var c int = 0
    var gloerr error = nil
    if n != ""{
        if h != ""{
            // 姓名、住院号、就诊号
            if a != ""{
                t = 1
            // 姓名、住院号
            }else{
                t = 2
            }
        } else{
            // 姓名、就诊号
            if a != ""{
                t = 3
            // 姓名
            }else{
                t = 4
            }
        }
    }else{
        if h != ""{
            // 住院号、就诊号
            if a != ""{
                t = 5
            // 住院号
            }else{
                t = 6
            }
        }else{
            // 就诊号
            if a != ""{
                t = 7
            }
        }
    }
    switch t {
    case 1:
        pnt.Search("姓名:%s,住院号:%s,就诊号:%s\n",n,h,a)
        rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker where name=? and hospital_number=? and attandance_number=?",n,h,a)
        if err !=nil{
            gloerr = err
            break
        }
        defer rows.Close()
        for rows.Next(){
            if serr := rows.Scan(&res.S[c].Name,&res.S[c].Hospital_number,&res.S[c].Attandance_number,&res.S[c].Sicker_id);serr != nil{
                gloerr = serr
            }
            res.S[c].Has = 1
            c++
            if c== 15{
                break
            }
        }
    case 2:
        pnt.Search("姓名:%s,住院号:%s\n",n,h)
        //DB.Query("SELECT name,hospital_number,attandance_number FROM where name=? and hospital_number=? and attandance_number=?",n,h,z)
        
    case 3:
        pnt.Search("姓名:%s,就诊号:%s\n",n,a)
        
    case 4:
        pnt.Search("姓名:%s\n",n)
        
    case 5:
        pnt.Search("住院号:%s,就诊号:%s\n",h,a)
    case 6:
        pnt.Search("住院号:%s\n",h)
    case 7:
        pnt.Search("就诊号:%s\n",a)
    default:
        pnt.Search("空搜索")
    }

    if gloerr != nil{
        pnt.Error(gloerr)
        return reParseJson(getAns(1,"搜索失败！",""))
    }

    res.Status=1
    return reParseJson(res)

}
// 搜索 患者 详细信息:
func (sds * searchDeatilSick) msgMain() []byte{
    var sdsr searchDeatilSickRes
    err := DB.QueryRow("SELECT name,age,gender,telphone,hospital_number,attandance_number,disease FROM sicker where userid=?",
    sds.Userid).Scan(&sdsr.Name,&sdsr.Age,&sdsr.Gender,&sdsr.Telphone,&sdsr.Hospital_number,&sdsr.Attandance_number,&sdsr.Disease)
    if err != nil{
        pnt.Error(err)
        sdsr.Status=0
        return reParseJson(sdsr)
    }
    sdsr.Status=1
    return reParseJson(sdsr)
}

// 搜索 患者 化疗周期
func (ci * cycleInfo) msgMain() []byte {
    var cir cycleInfoRes
    var d,t string
    var c int = 0
    rows,err := DB.Query("SELECT userid,cycle_seq,date,time FROM cycle WHERE userid=?",ci.Userid)

    if err != sql.ErrNoRows{
        cir.Status=1
    }else if err != nil{
        pnt.Error(err)
        cir.Status=0
        return reParseJson(cir)
    }
    defer rows.Close()

    for rows.Next(){
        rows.Scan(&cir.S[c].Userid,&cir.S[c].Cycle_seq,&d,&t)
        cir.S[c].Anstime = d + " " + t
        cir.S[c].Has = 1
        c++
    }
    cir.Status=1
    return reParseJson(cir)
}
// 查询 风险评估
func (rirc * riskInfoRec) msgMain() []byte{
    var rirs riskInfoRes
    pnt.Search("查询风险评估-患者ID:%s,化疗周期:%d\n",rirc.Userid,rirc.Cycle_seq)
    err := DB.QueryRow("SELECT * FROM risk WHERE userid=? and cycle_seq=?",
        rirc.Userid,
        rirc.Cycle_seq).Scan(
            &rirs.Userid,
            &rirs.Cycle_seq,
            &rirs.Program,
            &rirs.Not_medication,
            &rirs.Medication,
            &rirs.Grand,
            &rirs.Pre_program,
            &rirs.Pre_program_diy,
            &rirs.Comment,
            &rirs.Comment_diy,
            &rirs.Writer,
            &rirs.Assessment_date,
            &rirs.Assessment_time,
            &rirs.Assessment_timestamp,
            &rirs.Chemotherapy_date,
            &rirs.Chemotherapy_time,
            &rirs.Chemotherapy_timestamp)
    if err !=nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"查询失败！",""))
    }
    rirs.Status=1
    return reParseJson(rirs)
}
// 查询 患者 护理具体信息
func (nirc * nurseInfoRec) msgMain() []byte{
    var nirs nurseInfoRes
    pnt.Search("护理具体信息-患者ID:%s,化疗周期:%d,护理次序:%d\n",nirc.Userid,nirc.Cycle_seq,nirc.Nurse_seq)
    err := DB.QueryRow("SELECT * FROM nurse WHERE userid=? and cycle_seq=? and nurse_seq=?",
    nirc.Userid,
    nirc.Cycle_seq,
    nirc.Nurse_seq).Scan(
        &nirs.Userid,
        &nirs.Cycle_seq,
        &nirs.Nurse_seq,
        &nirs.Nausea_assessment,
        &nirs.Emesis_assessment,
        &nirs.Measure,
        &nirs.Comment,
        &nirs.Out_hospital,
        &nirs.Writer,
        &nirs.Assessment_date,
        &nirs.Assessment_time,
        &nirs.Assessment_timestamp)
    if err != nil{
        return reParseJson(getAns(0,"查询失败！",""))
    }
    nirs.Status=1
    return reParseJson(nirs)

}
// 查询 患者 护理表
func (ntrc * nurseTableRec) msgMain() []byte{
    var ntrs nurseTableReS
    var i int = 0
    var d,t string
    pnt.Search("查询护理表-患者ID:%s,化疗周期:%d\n",ntrc.Userid,ntrc.Cycle_seq)
    rows,err := DB.Query("SELECT nurse_seq,assessment_date,assessment_time FROM nurse WHERE userid=? and cycle_seq=?",ntrc.Userid,ntrc.Cycle_seq)
    if err !=nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"查询失败！",""))
    }
    defer rows.Close()
    for rows.Next(){
        err := rows.Scan(&ntrs.N[i].Nurse_seq,&d,&t)
        if err != nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"查询失败！",""))
        }
        ntrs.N[i].Time = d + " " + t
        ntrs.N[i].Has = 1
        i++
        if i ==15 {
            break
        }
    }
    ntrs.Status = 1
    return reParseJson(ntrs)


}
// 查询 患者 随访表
func (ftrc * followTableRec) msgMain() []byte{
    var ftrs followTableRes
    var i int = 0
    var d,t string 
    pnt.Search("查询随访表-患者ID:%s\n",ftrc.Userid)
    rows,err := DB.Query("SELECT follow_seq,follow_follow_date,follow_follow_time FROM follow where userid=?",ftrc.Userid)
    if err !=nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"查询失败！",""))
    }
    defer rows.Close()
    for rows.Next(){

        err := rows.Scan(&ftrs.N[i].Follow_seq,&d,&t)
        if err != nil{
            pnt.Error(err)
            return reParseJson(getAns(0,"查询失败！",""))
        }
        ftrs.N[i].Has = 1
        ftrs.N[i].Time = d + " " +t
        i++
        if i==15 {
            break
        }

    }

    ftrs.Status = 0
    return reParseJson(ftrs)

}
func (fcrc * followContentRec) msgMain() []byte{
    var fcrs followContentRes
    err := DB.QueryRow("SELECT * from follow where userid=? and follow_seq=?",fcrc.Userid,fcrc.Follow_seq).Scan(
        &fcrs.Userid                 ,
        &fcrs.Follow_seq             ,
        &fcrs.Hight_risk             ,
        &fcrs.Emesis_grade           ,
        &fcrs.Nausea_grade           ,
        &fcrs.Out_content            ,
        &fcrs.Out_content_diy        ,
        &fcrs.Follow_over            ,
        &fcrs.Satisfaction_1         ,
        &fcrs.Satisfaction_2         ,
        &fcrs.Satisfaction_3         ,
        &fcrs.Satisfaction_4         ,
        &fcrs.Satisfaction_5         ,
        &fcrs.Satisfaction_total     ,
        &fcrs.Writer                 ,
        &fcrs.Follow_follow_date     ,
        &fcrs.Follow_follow_time     ,
        &fcrs.Follow_follow_timestamp)
    
    if err != nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"查询失败！",""))
    }
    fcrs.Status=1
    return reParseJson(fcrs)
}
// 查询 患者 是否出院
func (ohpc * outHospitalRec) msgMain() []byte{
    var d,t string
    var ohps outHospitalRes
    pnt.Search("查询是否出院-患者ID:%s\n",ohpc.Userid)
    err := DB.QueryRow("SELECT assessment_date,assessment_time FROM nurse where userid=? and out_hospital=? and follow_over=?",ohpc.Userid,"1",2).Scan(&d,&t)
    if err !=nil{
        pnt.Error(err)
        ohps.Status=0
    }
    // if err == ErrNoRows{
    //     ohps.Explain="未出院"
    //     ohps.Status=2

    // }else if err != sql.ErrNoRows{
    //     ohps.Explain= "出院信息查询失败！"
    //     ohps.Status=1
    // }else{
    //     ohps.Status=0
    //     ohps.Time = d + " " + t
    // }
    ohps.Status=1
    ohps.Time = d + " " + t
    return reParseJson(ohps)
}

// 搜索 待办
func (wgrc * waitGoRec) msgMain() []byte{
    var wgrs waitGoRes
    var i = 0
    nt := time.Now()

    pnt.Search("查询待办\n")
    rows,err := DB.Query("SELECT out_hospital,name,userid FROM sicker WHERE follow_over=2")
    if err !=nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"查询失败！",""))
    }
    defer rows.Close()
    for rows.Next(){
        var t string
        err := rows.Scan(&t,&wgrs.N[i].Name,&wgrs.N[i].Userid)
        if err != nil{
            pnt.Error(err)
            return reParseJson(getAns(0,"查询失败！",""))
        }

        t = strings.Split(t," ")[0]
        s := strings.Split(t,"-")
        y,_ := strconv.Atoi(s[0])
        m,_ := strconv.Atoi(strings.TrimPrefix(s[1],"0"))
        d,_ := strconv.Atoi(strings.TrimPrefix(s[2],"0"))
    
        n:= time.Date(y,time.Month(m),d,0,0,0,0,time.Local)
        if int(nt.Sub(n).Hours()) > 48 {
            pnt.Search("查询待办-姓名:%s,%d年%d月%d日\n",wgrs.N[i].Name,y,m,d)
            wgrs.N[i].Has=1
            i++
            if i ==15 {
                break
            }
        }


    }
    wgrs.Status = 1
    return reParseJson(wgrs)

}

