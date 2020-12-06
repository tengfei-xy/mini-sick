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
// struct -> json
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
// json -> struct
func reParseJson(v interface{}) []byte{ 
    textbyte,err := json.Marshal(v)
    if err !=nil {
        pnt.Errorwd(v,err)
    }
    return textbyte
}
// ans
func (a * ans) set(s int,e,d string){
    a.Status = s
    a.Explain = e
    a.Data = d

}
func getAns(status int,explain ,data string) ans{
    var a ans
    a.set(status,explain,data)
    return a
}

// 判断 医生登录
func (ui * userInfo) msgMain() []byte{

    var    Name string
    var    Password string
    var    log = fmt.Sprintf("姓名:%s,账号ID:%s,密码:%s",ui.Name,ui.Account,ui.Password)
    pnt.Info(log)
    // 如果密码是空
    if ui.Password == ""{ return reParseJson(getAns(1,"登录失败！","")) }

    // 账号登录方式
    if ui.Account != "" {
        err := DB.QueryRow("SELECT name,password FROM users WHERE account=?",ui.Account).Scan(&Name,&Password)
        if err != nil{
            pnt.Errorf("用户登录错误-%v-info:%s",err,log)
            return reParseJson(getAns(1,"登录错误！",""))
        }

    // 姓名登录方式
    }else if ui.Name!="" {

        err := DB.QueryRow("SELECT password FROM users WHERE name=?",ui.Name).Scan(&Password)
        if err != nil{
            pnt.Errorf("用户登录错误-%v-info:%s",err,log)
            return reParseJson(getAns(1,"登录错误！",""))
        }
    }
    
    // 输入密码 与 数据库查询密码 不匹配
    if ui.Password != Password{

        pnt.Errorf("用户登录失败-%s-info:%s","密码或认证方式不匹配",log)
        return reParseJson(getAns(1,"登录失败！",""))
    }

    pnt.Infof("用户登录成功-log:%s",log)
    return reParseJson(getAns(0,"登录成功！",Name))
    
}

// 添加 患者基本信息
func (si * sickerInfo) msgMain() []byte {
    var waytext string
    userid := createUserID()

    log := fmt.Sprintf(waytext + "患者姓名:%s,患者ID:%s,住院号:%s,就诊号:%s",si.Name,userid,si.Hospital_number,si.Attandance_number)

    if  si.Way == 1{
        waytext = "添加"
    }else{
        waytext = "更新"
    }

    // 从 public.go 中生成患者ID
    if si.Way == 1 {

        _,err := DB.Exec("INSERT INTO sicker (userid,name,age,gender,telphone,hospital_number,attandance_number,disease,writer) VALUES (?,?,?,?,?,?,?,?,?)",
            userid,
            si.Name,
            si.Age,
            si.Gender,
            si.Telphone,
            si.Hospital_number,
            si.Attandance_number,
            si.Disease,
            si.Writer)

        if err != nil{
            pnt.Errorf(waytext + "患者失败-%v-info:%s",err,log)
            return reParseJson(getAns(0,waytext + "失败！",""))
        }
    }else{
        _,errf := DB.Exec("UPDATE sicker SET name=?,age=?,gender=?,telphone=?,hospital_number=?,attandance_number=?,disease=?,writer=? WHERE userid=?",
            si.Name,
            si.Age,
            si.Gender,
            si.Telphone,
            si.Hospital_number,
            si.Attandance_number,
            si.Disease,
            si.Writer,
            si.Userid)
        if errf != nil{
            pnt.Errorf(waytext + "患者失败-%v-info:%s",errf,log)
            return reParseJson(getAns(0,waytext + "失败！",""))
        }
    }

    
    pnt.Infof(waytext + "患者成功-info:%s",log)
    return reParseJson(getAns(1,waytext +"信息成功！",userid))
    
}

// 更新 患者信息的风险评估
func (ri * riskInfo) msgMain() []byte {
    var log string = fmt.Sprintf("患者ID:%s",ri.Userid)
    var waytext string


    if ri.Updated != 0{
        waytext = "更新"
    }else{
        waytext = "添加"

    }

    // 更新风险评估
    if ri.Updated != 0{
        _,err := DB.Exec("UPDATE risk SET program=?,not_medication=?,medication=?,grand=?,pre_program=?,pre_program_diy=?,comment=?,comment_diy=?,need_nurse=?,writer=?,assessment_date=?,assessment_time=?,assessment_timestamp=?,chemotherapy_date=?,chemotherapy_time=?,chemotherapy_timestamp=? WHERE userid=? AND cycle_seq=? ",
        ri.Program        	,
        ri.Not_medication	,
        ri.Medication     	,
        ri.Grand          	,
        ri.Pre_program 		,
        ri.Pre_program_diy 	,
        ri.Comment    		,
        ri.Comment_diy     	,
        ri.Need_nurse       ,

        ri.Writer       	,
        ri.Assessment_date 	,
        ri.Assessment_time 	,
        ri.Assessment_timestamp,
        ri.Chemotherapy_date ,     
        ri.Chemotherapy_time ,     
        ri.Chemotherapy_timestamp,
        ri.Userid 	     	,
        ri.Cycle_seq)
        
        if err != nil{
            pnt.Errorf(waytext + "风险评估失败-%v-info:%s",err,log)
            return reParseJson(getAns(0,waytext+"失败！",""))
        }


    // 添加风险评估
    } else{

        _,err := DB.Exec("INSERT INTO risk VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
            ri.Userid 	     	,
            ri.Cycle_seq    	,
            ri.Program        	,
            ri.Not_medication	,
            ri.Medication     	,
            ri.Grand          	,
            ri.Pre_program 		,
            ri.Pre_program_diy 	,
            ri.Comment    		,
            ri.Comment_diy     	,
            ri.Need_nurse       ,
            ri.Writer       	,
            ri.Assessment_date 	,
            ri.Assessment_time 	,
            ri.Assessment_timestamp,
            ri.Chemotherapy_date ,     
            ri.Chemotherapy_time ,     
            ri.Chemotherapy_timestamp)
            
        if err != nil{
            pnt.Errorf(waytext + "风险评估-风险评估-失败-%v-info:%s",err,log)
            return reParseJson(getAns(0,waytext+"失败！",""))
        }

        // 更新 化疗周期表
        _,erra := DB.Exec("INSERT cycle(userid,cycle_seq,name,date,time,timestamp) VALUES (?,?,?,?,?,?)",
            ri.Userid 	     	,
            ri.Cycle_seq    	,
            ri.Name             ,
            ri.Assessment_date 	,
            ri.Assessment_time 	,
            ri.Assessment_timestamp)

        if erra != nil{
            pnt.Errorf(waytext + "风险评估-化疗周期失败-%v-info:%s",erra,log)
            return reParseJson(getAns(0,waytext + "失败！",""))
        }
        // 更新 患者表
        if _,errb := DB.Exec("UPDATE sicker SET cycle_seq=? WHERE userid=?",ri.Cycle_seq,ri.Userid);errb != nil{
            pnt.Errorf(waytext + "风险评估-患者表失败-%v-info:%s",errb,log)
            return reParseJson(getAns(0,waytext + "失败！",""))
        }


    }

    // 是否住院护理,2为否，表示出院
    if ri.Need_nurse=="2"{
        if _,err := DB.Exec("UPDATE cycle SET out_hospital_time=? WHERE userid=? AND cycle_seq=? AND LENGTH(out_hospital_time)=0",
        ri.Assessment_date + " " + ri.Assessment_time,ri.Userid,ri.Cycle_seq);err!=nil{
            pnt.Errorf(waytext + "风险评估-周期表-出院判断失败-%v-info:%s",err,log)
            return reParseJson(getAns(0,waytext+"失败！",""))
        }
    }
    pnt.Infof(waytext + "风险评估成功-info:%s",log)
    return reParseJson(getAns(1,waytext+"成功！",""))
    
     
}

// 更新 患者信息的护理评估
func (ni * nurseInfo) msgMain() []byte {

    var log string = fmt.Sprintf("患者ID:%s,化疗周期:%d,护理次序:%d",ni.Userid,ni.Cycle_seq,ni.Nurse_seq)

    // 插入数据
    _,err := DB.Exec(
        "INSERT INTO nurse VALUES (?,?,?,?,?,?,?,?,?,?,?,?)",
        ni.Userid               ,
        ni.Cycle_seq            ,
        ni.Nurse_seq            ,

        ni.Nausea_assessment    ,
        ni.Emesis_assessment    ,
        ni.Measure              ,
        ni.Comment              ,
        ni.Out_hospital         ,

        ni.Writer               ,
        ni.Assessment_date      ,
        ni.Assessment_time      ,
        ni.Assessment_timestamp)

    if err != nil{
        pnt.Errorf("插入护理评估-护理表-失败-%s-info:%s",err,log)
        return reParseJson(getAns(0,"更新失败！",""))

    }
    

    pnt.Infof("插入护理评估-护理表-成功-info:%s",log)
    if ni.Out_hospital == "1"{
         _,err := DB.Exec("UPDATE cycle SET out_hospital_time=? WHERE userid=? AND LENGTH(out_hospital_time)=0",
         ni.Assessment_date + " " + ni.Assessment_time,ni.Userid)
        if err !=nil{
            pnt.Errorf("插入护理评估-更新化疗周期-失败-%s-info:%s\n",err,log)
            return reParseJson(getAns(0,"更新失败！",""))
        }
        return reParseJson(getAns(1,"更新成功！该患者将进行随访！",""))

    }
    return reParseJson(getAns(1,"更新成功！",""))
    

}

// 添加 患者信息 随访
func (fi * followInfo) msgMain() []byte {

    var log string = fmt.Sprintf("患者ID:%s,化疗周期:%d,随访周期:%d",fi.Userid,fi.Cycle_seq,fi.Follow_seq)

    // 插入数据
    _,err := DB.Exec(
        "INSERT INTO follow VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
        fi.Userid                  ,
        fi.Cycle_seq               ,
        fi.Follow_seq              ,
        fi.Hight_risk              ,
        fi.Emesis_grade            ,
        fi.Nausea_grade            ,
        fi.Out_content             ,
        fi.Out_content_diy         ,
        fi.Follow_over             ,
        fi.Satisfaction_1          ,
        fi.Satisfaction_2          ,
        fi.Satisfaction_3          ,
        fi.Satisfaction_4          ,
        fi.Satisfaction_5          ,
        fi.Satisfaction_total      ,
        fi.Writer 				   ,
        fi.Follow_follow_date      ,
        fi.Follow_follow_time      ,
        fi.Follow_follow_timestamp)

        if err != nil{
            pnt.Errorf("添加随访-随访表失败-%v-log:%s",err,log)
            return reParseJson(getAns(0,"添加失败！",""))
        }
        if fi.Follow_over == "1"{
            if _,err := DB.Exec("UPDATE cycle SET follow_over=? WHERE userid=? AND cycle_seq=?","1",fi.Userid,fi.Cycle_seq);err != nil{
                pnt.Errorf("添加随访-周期表-失败-%v-log:%s",err,log)
                return reParseJson(getAns(0,"添加失败！",""))
            }
        }
        pnt.Errorf("随访添加成功-log:%s",log)
        return reParseJson(getAns(1,"添加成功！",""))
        
}

// 搜索 患者
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
        pnt.Search("姓名:%s,住院号:%s,就诊号:%s",n,h,a)
        rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE name=? AND hospital_number=? AND attandance_number=?",n,h,a)
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
        pnt.Search("姓名:%s,住院号:%s",n,h)
        rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE name=? AND hospital_number=?",n,h)
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
        
    case 3:
        pnt.Search("姓名:%s,就诊号:%s",n,a)
        rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE name=? AND attandance_number=?",n,a)
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
    case 4:
        pnt.Search("姓名:%s",n)
        rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE name=?",n)
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
        
    case 5:
        pnt.Search("住院号:%s,就诊号:%s",h,a)
        rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE hospital_number=? AND attandance_number=?",h,a)
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
    case 6:
        pnt.Search("住院号:%s",h)
        rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE hospital_number=?",h)
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
    case 7:
        pnt.Search("就诊号:%s",a)
        rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE attandance_number=?",a)
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
    default:
        return reParseJson(getAns(1,"暂不支持该类型的搜索方式！",""))

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
    pnt.Search("搜索患者详细信息-患者ID:%s",sds.Userid)

    var sdsr searchDeatilSickRes
    err := DB.QueryRow("SELECT name,age,gender,telphone,hospital_number,attandance_number,disease FROM sicker WHERE userid=?",
    sds.Userid).Scan(&sdsr.Name,&sdsr.Age,&sdsr.Gender,&sdsr.Telphone,&sdsr.Hospital_number,&sdsr.Attandance_number,&sdsr.Disease)
    if err != nil{
        pnt.Infof("搜索患者详细信息-患者ID:%s-搜索失败,%v",sds.Userid,err)

        return reParseJson(getAns(1,"搜索失败！",""))
    }
    sdsr.Status=1
    return reParseJson(sdsr)
}

// 搜索 患者 化疗周期
func (ci * cycleInfo) msgMain() []byte {
    pnt.Search("搜索患者化疗周期-患者ID:%s",ci.Userid)

    var cir cycleInfoRes
    var d,t string
    var c int = 0
    rows,err := DB.Query("SELECT userid,cycle_seq,date,time FROM cycle WHERE userid=?",ci.Userid)

    if err == sql.ErrNoRows{
        cir.Status=1
        return reParseJson(cir)

    }else if err != nil{
        pnt.Infof("搜索患者化疗周期-患者ID:%s,%v",ci.Userid,err)
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
    var last_seq int
    pnt.Search("查询风险评估-患者ID:%s,化疗周期:%d",rirc.Userid,rirc.Cycle_seq)

    // 至少是第二周期开始查询上次记录
    if (rirc.Cycle_seq >1){

    // 查询上次风险等级
    if err := DB.QueryRow("SELECT grand FROM risk WHERE userid=? AND cycle_seq=?",
    rirc.Userid,rirc.Cycle_seq-1).Scan(
        &rirs.Last_risk_grand);
        err !=nil && err!=sql.ErrNoRows{
            pnt.Errorf("查询上次风险等级记录失败-%v",err)
            return reParseJson(getAns(0,"查询上一周期内容失败！",""))
        }

    // 查询上次护理呕吐和恶心
    if err := DB.QueryRow("SELECT Emesis_assessment,Nausea_assessment,nurse_seq FROM nurse WHERE userid=? AND cycle_seq=? ORDER BY nurse_seq DESC LIMIT ?",
    rirc.Userid,rirc.Cycle_seq-1,1).Scan(
        &rirs.Last_nurse_emesis,
        &rirs.Last_nurse_nausea,
        &last_seq);err !=nil && err!=sql.ErrNoRows{
            pnt.Errorf("查询上次护理记录失败-%v",err)
            return reParseJson(getAns(0,"查询上一周期内容失败！",""))
        }
            
    // 查询上次随访呕吐和恶心
    if err := DB.QueryRow("SELECT emesis_grade,nausea_grade,follow_seq FROM follow WHERE userid=? AND cycle_seq=? ORDER BY follow_seq DESC LIMIT ?",
    rirc.Userid,rirc.Cycle_seq-1,1).Scan(
        &rirs.Last_follow_emesis,
        &rirs.Last_follow_nausea,
        &last_seq);err !=nil && err!=sql.ErrNoRows{
            pnt.Errorf("查询上次随访记录失败-%v",err)
        }
    }

    err := DB.QueryRow("SELECT * FROM risk WHERE userid=? AND cycle_seq=?",
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
            &rirs.Need_nurse,

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
    pnt.Search("护理具体信息-患者ID:%s,化疗周期:%d,护理次序:%d",nirc.Userid,nirc.Cycle_seq,nirc.Nurse_seq)
    err := DB.QueryRow("SELECT * FROM nurse WHERE userid=? AND cycle_seq=? AND nurse_seq=?",
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
    pnt.Search("查询护理表-患者ID:%s,化疗周期:%d",ntrc.Userid,ntrc.Cycle_seq)
    rows,err := DB.Query("SELECT nurse_seq,assessment_date,assessment_time FROM nurse WHERE userid=? AND cycle_seq=?",ntrc.Userid,ntrc.Cycle_seq)
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
    pnt.Search("查询随访表-患者ID:%s",ftrc.Userid)
    rows,err := DB.Query("SELECT follow_seq,follow_follow_date,follow_follow_time FROM follow WHERE userid=? AND cycle_seq=?",ftrc.Userid,ftrc.Cycle_seq)
    if err != sql.ErrNoRows && err !=nil{
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

    ftrs.Status = 1
    return reParseJson(ftrs)

}

// 查询 患者 随访具体内容
func (fcrc * followContentRec) msgMain() []byte{
    pnt.Infof("查询随访具体信息-患者ID:%s,化疗周期:%d,随访周期:%d",fcrc.Userid,fcrc.Cycle_seq,fcrc.Follow_seq)
    var fcrs followContentRes
    err := DB.QueryRow("SELECT * from follow WHERE userid=? AND follow_seq=? AND cycle_seq=?",fcrc.Userid,fcrc.Follow_seq,fcrc.Cycle_seq).Scan(
        &fcrs.Userid                 ,
        &fcrs.Cycle_seq             ,
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
    var ohps outHospitalRes
    pnt.Search("查询是否出院-患者ID:%s",ohpc.Userid)
    err := DB.QueryRow(`SELECT out_hospital_time FROM cycle WHERE cycle_seq=? AND userid=?`,ohpc.Cycle_seq,ohpc.Userid).Scan(&ohps.Time)
    if err ==sql.ErrNoRows || ohps.Time== ""{
        return reParseJson(getAns(1,"未出院！",""))
    } else if err !=sql.ErrNoRows && err != nil{
        pnt.Error(err)
        return reParseJson(getAns(0,"查询失败！",""))
    }

    ohps.Status=1
    return reParseJson(ohps)
}

// 搜索 待办
func (wgrc * waitGoRec) msgMain() []byte{
    var wgrs waitGoRes
    var    log = fmt.Sprintf("查询待办")

    var i = 0
    nt := time.Now()
	neight:= time.Date(nt.Year(),nt.Month(),nt.Day(),0,0,0,0,time.Local)
    rows,err := DB.Query("SELECT userid,name,out_hospital_time,cycle_seq FROM cycle WHERE follow_over!=? AND LENGTH(out_hospital_time)!=? ORDER BY cycle_seq DESC LIMIT ?","1",0,1)
    if err == sql.ErrNoRows{
        wgrs.Status=1
        wgrs.Explain="今日无随访患者！"
        return reParseJson(wgrs)
    }
    if err !=nil {
        pnt.Errorf("%s-%v",log,err)
        return reParseJson(getAns(0,"查询失败！",""))
    }
    defer rows.Close()
    for rows.Next(){
        var t string
        err := rows.Scan(&wgrs.N[i].Userid,&wgrs.N[i].Name,&t,&wgrs.N[i].Cycle_seq)
        if err != nil{
            pnt.Errorf("%s-%v",err,log)
            return reParseJson(getAns(0,"查询失败！",""))
        }

        t = strings.Split(t," ")[0]
        s := strings.Split(t,"-")
        y,_ := strconv.Atoi(s[0])
        m,_ := strconv.Atoi(strings.TrimPrefix(s[1],"0"))
        d,_ := strconv.Atoi(strings.TrimPrefix(s[2],"0"))
    
        n:= time.Date(y,time.Month(m),d,0,0,0,0,time.Local)
        sub :=int(neight.Sub(n).Hours()) 

        if sub > 24 {
            pnt.Search("%s,姓名:%s,%d月%d日",log,wgrs.N[i].Name,m,d)

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

// 查询 随访中高致吐风险
func (hrrc * heightRiskRec) msgMain() []byte{
    var hrrs heightRiskReS
    var    log = fmt.Sprintf("查询高致吐风险-患者ID:%s",hrrc.Userid)

    err := DB.QueryRow("SELECT grand FROM risk WHERE userid=? AND cycle_seq=?",hrrc.Userid,hrrc.Cycle_seq).Scan(&hrrs.Height)
    if err == sql.ErrNoRows{
        return reParseJson(getAns(2,"风险评估的风险等级未填写",""))
    }else if err !=sql.ErrNoRows && err != nil{
        pnt.Errorf("%s-%v",err,log)
        return reParseJson(getAns(0,"查询失败！",""))
    }

    hrrs.Status = 1
    return reParseJson(hrrs)

}

// 查询 随访周期 上一次
func (clrc * cycleLastRec) msgMain() []byte{
    var clrs cycleLastRes
    var last_seq int
    pnt.Search("查询上次随访周期-患者ID:%s,化疗周期:%d",clrc.Userid,clrc.Cycle_seq-1)

    // 至少是第二周期开始查询上次记录
    if (clrc.Cycle_seq <1){reParseJson(getAns(0,"查询上一周期内容失败！",""))}
    
    // 查询上次风险等级
    if err := DB.QueryRow("SELECT grand FROM risk WHERE userid=? AND cycle_seq=?",
    clrc.Userid,clrc.Cycle_seq-1).Scan(
        &clrs.Last_risk_grand);
        err !=nil && err!=sql.ErrNoRows{
            pnt.Errorf("查询上次风险等级记录失败-%v",err)
            return reParseJson(getAns(0,"查询上一周期内容失败！",""))
        }

    // 查询上次护理呕吐和恶心
    if err := DB.QueryRow("SELECT Emesis_assessment,Nausea_assessment,nurse_seq FROM nurse WHERE userid=? AND cycle_seq=? ORDER BY nurse_seq DESC LIMIT ?",
    clrc.Userid,clrc.Cycle_seq-1,1).Scan(
        &clrs.Last_nurse_emesis,
        &clrs.Last_nurse_nausea,
        &last_seq);err !=nil && err!=sql.ErrNoRows{
            pnt.Errorf("查询上次护理记录失败-%v",err)
            return reParseJson(getAns(0,"查询上一周期内容失败！",""))
        }
        
    // 查询上次随访呕吐和恶心
    if err := DB.QueryRow("SELECT emesis_grade,nausea_grade,follow_seq FROM follow WHERE userid=? AND cycle_seq=? ORDER BY follow_seq DESC LIMIT ?",
        clrc.Userid,clrc.Cycle_seq-1,1).Scan(
        &clrs.Last_follow_emesis,
        &clrs.Last_follow_nausea,&last_seq);
    err !=nil && err!=sql.ErrNoRows{
        pnt.Errorf("查询上次随访记录失败-%v",err)
        return reParseJson(getAns(0,"查询上一周期内容失败！",""))

    }
    clrs.Status =1
    return reParseJson(clrs)
    
}