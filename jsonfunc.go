package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"env"
	pnt "print"
)

// ans
func (a *ans) set(s int, e, d string) {
	a.Status = s
	a.Explain = e
	a.Data = d

}
func getAns(status int, explain, data string) ans {
	var a ans
	a.set(status, explain, data)
	return a
}

// 判断 医生/患者登录
func (ui *userInfo) msgMain() []byte {

	var Name string
	var Password string
	var way, log string
	var ulrs userLoginRes

	if ui.Who == "1" {
		// 医生登录
		way = "医生"
		log = fmt.Sprintf("姓名(%s):%s,账号ID:%s,密码:%s", way, ui.Name, ui.Account, ui.Password)
		// 如果密码是空
		if ui.Password == "" {
			pnt.Errorf("登录失败-%s-info:%s", "密码或认证方式不匹配", log)
			return reParseJSON(getAns(1, "登录失败！", ""))
		}

		// 账号登录方式
		if ui.Account != "" {
			err := DB.QueryRow("SELECT name,password FROM users WHERE account=?", ui.Account).Scan(&Name, &Password)
			if err != nil {
				pnt.Errorf("登录错误-%v-info:%s", err, log)
				return reParseJSON(getAns(1, "登录错误！", ""))
			}
			ulrs.Data = Name

			// 姓名登录方式
		} else if ui.Name != "" {

			err := DB.QueryRow("SELECT password FROM users WHERE name=?", ui.Name).Scan(&Password)
			if err != nil {
				pnt.Errorf("登录错误-%v-info:%s", err, log)
				return reParseJSON(getAns(1, "登录错误！", ""))
			}
			ulrs.Data = ui.Name

		}

		// 输入密码 与 数据库查询密码 不匹配
		if ui.Password != Password {
			pnt.Errorf("登录失败-%s-info:%s", "密码或认证方式不匹配", log)
			return reParseJSON(getAns(1, "登录失败！", ""))
		}
		pnt.Infof("登录成功-log:%s", log)

	} else if ui.Who == "2" {
		// 患者登录
		way = "患者登录"
		log = fmt.Sprintf("姓名(%s):%s", way, ui.Name)
		err := DB.QueryRow("SELECT userid,cycle_seq FROM sicker WHERE name=? ORDER BY cycle_seq DESC LIMIT ?", ui.Name, "1").Scan(&ulrs.Userid, &ulrs.Cycle_seq)
		if err != nil {
			pnt.Errorf("登录错误-%v-info:%s", err, log)
			return reParseJSON(getAns(1, "登录错误！", ""))
		}
		pnt.Infof("登录成功-log:%s", log)
		ulrs.Data = ui.Name

	}

	ulrs.Explain = "登录成功！"
	ulrs.Status = 0
	return reParseJSON(ulrs)
}

// 添加/更新 患者基本信息
func (si *sickerInfo) msgMain() []byte {
	var waytext, userid, log string

	if si.Way == 1 {
		waytext = "添加"
		// 从 public.go 中生成患者ID
		userid = createUserID()

	} else {
		waytext = "更新"
		userid = si.Userid
	}

	log = fmt.Sprintf(waytext+"患者姓名:%s,患者ID:%s,住院号:%s,就诊号:%s", si.Name, userid, si.Hospital_number, si.Attandance_number)

	if si.Way == 1 {
		var has bool = true
		// 判断是否重复添加
		if si.Hospital_number == "" {
			if err := DB.QueryRow("SELECT userid FROM sicker WHERE name=? AND attandance_number=?", si.Name, si.Attandance_number).Scan(&userid); err == sql.ErrNoRows {
				has = false
			}
		} else if si.Attandance_number == "" {
			if err := DB.QueryRow("SELECT userid FROM sicker WHERE name=? AND hospital_number=?", si.Name, si.Hospital_number).Scan(&userid); err == sql.ErrNoRows {
				has = false
			}
		} else {
			if err := DB.QueryRow("SELECT userid FROM sicker WHERE name=? AND hospital_number=? AND attandance_number=?", si.Name, si.Hospital_number, si.Attandance_number).Scan(&userid); err == sql.ErrNoRows {
				has = false
			}
		}
		// 如果没有患者信息
		if has == false {
			_, err := DB.Exec("INSERT INTO sicker (userid,name,age,gender,telphone,hospital_number,attandance_number,disease,know,writer) VALUES (?,?,?,?,?,?,?,?,?,?)",
				userid,
				si.Name,
				si.Age,
				si.Gender,
				si.Telphone,
				si.Hospital_number,
				si.Attandance_number,
				si.Disease,
				si.Know,
				si.Writer)

			if err != nil {
				pnt.Errorf(waytext+"患者失败-%v-info:%s", err, log)
				return reParseJSON(getAns(0, waytext+"失败！", ""))
			}
			// 如果有患者
		} else {
			return reParseJSON(getAns(1, "患者已存在！", userid))
		}
	} else {
		_, err := DB.Exec("UPDATE sicker SET name=?,age=?,gender=?,telphone=?,hospital_number=?,attandance_number=?,disease=?,know=?,writer=? WHERE userid=?",
			si.Name,
			si.Age,
			si.Gender,
			si.Telphone,
			si.Hospital_number,
			si.Attandance_number,
			si.Disease,
			si.Know,
			si.Writer,
			si.Userid)
		if err != nil {
			pnt.Errorf(waytext+"患者失败-%v-info:%s", err, log)
			return reParseJSON(getAns(0, waytext+"失败！", ""))
		}
	}

	pnt.Infof(waytext+"患者成功-info:%s", log)
	return reParseJSON(getAns(1, waytext+"信息成功！", userid))
}

// 更新 患者信息的风险评估
func (ri *riskInfo) msgMain() []byte {
	var log string = fmt.Sprintf("患者ID:%s", ri.Userid)
	var waytext string

	if ri.Updated != 0 {
		waytext = "更新"
	} else {
		waytext = "添加"

	}

	// 更新风险评估
	if ri.Updated != 0 {
		// 更新化疗表
		_, err := DB.Exec("UPDATE risk SET program=?,not_medication=?,medication=?,grand=?,pre_program=?,pre_program_diy=?,comment=?,comment_diy=?,need_nurse=?,writer=?,assessment_date=?,assessment_time=?,assessment_timestamp=?,chemotherapy_date=?,chemotherapy_time=?,chemotherapy_timestamp=? WHERE userid=? AND cycle_seq=? ",
			ri.Program,
			ri.Not_medication,
			ri.Medication,
			ri.Grand,
			ri.Pre_program,
			ri.Pre_program_diy,
			ri.Comment,
			ri.Comment_diy,
			ri.Need_nurse,

			ri.Writer,
			ri.Assessment_date,
			ri.Assessment_time,
			ri.Assessment_timestamp,
			ri.Chemotherapy_date,
			ri.Chemotherapy_time,
			ri.Chemotherapy_timestamp,
			ri.Userid,
			ri.Cycle_seq)

		if err != nil {
			pnt.Errorf(waytext+"风险评估(化疗表)失败-%v-info:%s", err, log)
			return reParseJSON(getAns(0, waytext+"失败！", ""))
		}
		// 更新周期表
		if _, err := DB.Exec("UPDATE cycle SET date=?,time=? WHERE userid=? AND cycle_seq=?", ri.Chemotherapy_date, ri.Chemotherapy_time, ri.Userid, ri.Cycle_seq); err != nil {
			pnt.Errorf(waytext+"风险评估(周期表)失败-%v-info:%s", err, log)
			return reParseJSON(getAns(0, waytext+"失败！", ""))
		}

		// 添加风险评估
	} else {

		_, err := DB.Exec("INSERT INTO risk VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
			ri.Userid,
			ri.Cycle_seq,
			ri.Program,
			ri.Not_medication,
			ri.Medication,
			ri.Grand,
			ri.Pre_program,
			ri.Pre_program_diy,
			ri.Comment,
			ri.Comment_diy,
			ri.Need_nurse,
			ri.Writer,
			ri.Assessment_date,
			ri.Assessment_time,
			ri.Assessment_timestamp,
			ri.Chemotherapy_date,
			ri.Chemotherapy_time,
			ri.Chemotherapy_timestamp)

		if err != nil {
			pnt.Errorf(waytext+"风险评估-风险评估-失败-%v-info:%s", err, log)
			return reParseJSON(getAns(0, waytext+"失败！", ""))
		}

		// 更新 化疗周期表
		_, erra := DB.Exec("INSERT cycle(userid,cycle_seq,name,date,time,timestamp) VALUES (?,?,?,?,?,?)",
			ri.Userid,
			ri.Cycle_seq,
			ri.Name,
			ri.Assessment_date,
			ri.Assessment_time,
			ri.Assessment_timestamp)

		if erra != nil {
			pnt.Errorf(waytext+"风险评估-化疗周期失败-%v-info:%s", erra, log)
			return reParseJSON(getAns(0, waytext+"失败！", ""))
		}
		// 更新 患者表
		if _, errb := DB.Exec("UPDATE sicker SET cycle_seq=? WHERE userid=?", ri.Cycle_seq, ri.Userid); errb != nil {
			pnt.Errorf(waytext+"风险评估-患者表失败-%v-info:%s", errb, log)
			return reParseJSON(getAns(0, waytext+"失败！", ""))
		}

	}

	// 是否住院护理,2为否，表示出院
	if ri.Need_nurse == "2" {
		if _, err := DB.Exec("UPDATE cycle SET out_hospital_time=? WHERE userid=? AND cycle_seq=? AND LENGTH(out_hospital_time)=0",
			ri.Assessment_date+" "+ri.Assessment_time, ri.Userid, ri.Cycle_seq); err != nil {
			pnt.Errorf(waytext+"风险评估-周期表-出院判断失败-%v-info:%s", err, log)
			return reParseJSON(getAns(0, waytext+"失败！", ""))
		}
	}
	pnt.Infof(waytext+"风险评估成功-info:%s", log)
	return reParseJSON(getAns(1, waytext+"成功！", ""))
}

// 更新 患者信息的护理评估
func (ni *nurseInfo) msgMain() []byte {

	var log string = fmt.Sprintf("患者ID:%s,化疗周期:%d,护理次序:%d", ni.Userid, ni.Cycle_seq, ni.Nurse_seq)

	// 插入数据
	_, err := DB.Exec(
		"INSERT INTO nurse VALUES (?,?,?,?,?,?,?,?,?,?,?,?)",
		ni.Userid,
		ni.Cycle_seq,
		ni.Nurse_seq,

		ni.Nausea_assessment,
		ni.Emesis_assessment,
		ni.Measure,
		ni.Comment,
		ni.Out_hospital,

		ni.Writer,
		ni.Assessment_date,
		ni.Assessment_time,
		ni.Assessment_timestamp)

	if err != nil {
		pnt.Errorf("插入护理评估-护理表-失败-%s-info:%s", err, log)
		return reParseJSON(getAns(0, "更新失败！", ""))

	}

	pnt.Infof("插入护理评估-护理表-成功-info:%s", log)
	if ni.Out_hospital == "1" {
		_, err := DB.Exec("UPDATE cycle SET out_hospital_time=? WHERE userid=? AND LENGTH(out_hospital_time)=0",
			ni.Assessment_date+" "+ni.Assessment_time, ni.Userid)
		if err != nil {
			pnt.Errorf("插入护理评估-更新化疗周期-失败-%s-info:%s\n", err, log)
			return reParseJSON(getAns(0, "更新失败！", ""))
		}
		return reParseJSON(getAns(1, "更新成功！该患者将进行随访！", ""))

	}
	return reParseJSON(getAns(1, "更新成功！", ""))
}

// 添加 患者信息 随访
func (fi *followInfo) msgMain() []byte {

	var log string = fmt.Sprintf("患者ID:%s,化疗周期:%d,随访周期:%d", fi.Userid, fi.Cycle_seq, fi.Follow_seq)

	// 插入数据
	_, err := DB.Exec(
		"INSERT INTO follow VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		fi.Userid,
		fi.Cycle_seq,
		fi.Follow_seq,
		fi.Hight_risk,
		fi.Emesis_grade,
		fi.Nausea_grade,
		fi.Out_content,
		fi.Out_content_diy,
		fi.Follow_over,
		fi.Satisfaction_1,
		fi.Satisfaction_2,
		fi.Satisfaction_3,
		fi.Satisfaction_4,
		fi.Satisfaction_5,
		fi.Satisfaction_total,
		fi.Writer,
		fi.Follow_follow_date,
		fi.Follow_follow_time,
		fi.Follow_follow_timestamp)

	if err != nil {
		pnt.Errorf("添加随访-随访表失败-%v-log:%s", err, log)
		return reParseJSON(getAns(0, "添加失败！", ""))
	}
	if fi.Follow_over == "1" {
		if _, err := DB.Exec("UPDATE cycle SET follow_over=? WHERE userid=? AND cycle_seq=?", "1", fi.Userid, fi.Cycle_seq); err != nil {
			pnt.Errorf("添加随访-周期表-失败-%v-log:%s", err, log)
			return reParseJSON(getAns(0, "添加失败！", ""))
		}
	}
	pnt.Errorf("随访添加成功-log:%s", log)
	return reParseJSON(getAns(1, "添加成功！", ""))
}

// 搜索 患者
func (ss *searchSicker) msgMain() []byte {
	n := ss.Name
	h := ss.Hospital_number
	a := ss.Attandance_number
	var res searchSickerRes
	var log string = fmt.Sprintf("搜索患者")
	pnt.Info(log)

	t := 0
	var c int = 0
	var gloerr error = nil
	if n != "" {
		if h != "" {
			// 姓名、住院号、就诊号
			if a != "" {
				t = 1
				// 姓名、住院号
			} else {
				t = 2
			}
		} else {
			// 姓名、就诊号
			if a != "" {
				t = 3
				// 姓名
			} else {
				t = 4
			}
		}
	} else {
		if h != "" {
			// 住院号、就诊号
			if a != "" {
				t = 5
				// 住院号
			} else {
				t = 6
			}
		} else {
			// 就诊号
			if a != "" {
				t = 7
			}
		}
	}
	switch t {
	case 1:
		pnt.Search("姓名:%s,住院号:%s,就诊号:%s", n, h, a)
		rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE name=? AND hospital_number=? AND attandance_number=?", n, h, a)
		if err != nil {
			gloerr = err
			break
		}
		defer rows.Close()
		for rows.Next() {
			if serr := rows.Scan(&res.S[c].Name, &res.S[c].Hospital_number, &res.S[c].Attandance_number, &res.S[c].Sicker_id); serr != nil {
				gloerr = serr
			}
			res.S[c].Has = 1
			c++
			if c == 50 {
				pnt.Errorf("%s-扫描将超过50个", log)
				break
			}
		}
	case 2:
		pnt.Search("姓名:%s,住院号:%s", n, h)
		rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE name=? AND hospital_number=?", n, h)
		if err != nil {
			gloerr = err
			break
		}
		defer rows.Close()
		for rows.Next() {
			if serr := rows.Scan(&res.S[c].Name, &res.S[c].Hospital_number, &res.S[c].Attandance_number, &res.S[c].Sicker_id); serr != nil {
				gloerr = serr
			}
			res.S[c].Has = 1
			c++
			if c == 50 {
				pnt.Errorf("%s-扫描将超过50个", log)
				break
			}
		}

	case 3:
		pnt.Search("姓名:%s,就诊号:%s", n, a)
		rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE name=? AND attandance_number=?", n, a)
		if err != nil {
			gloerr = err
			break
		}
		defer rows.Close()
		for rows.Next() {
			if serr := rows.Scan(&res.S[c].Name, &res.S[c].Hospital_number, &res.S[c].Attandance_number, &res.S[c].Sicker_id); serr != nil {
				gloerr = serr
			}
			res.S[c].Has = 1
			c++
			if c == 50 {
				break
			}
		}
	case 4:
		pnt.Search("姓名:%s", n)
		rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE name=?", n)
		if err != nil {
			gloerr = err
			break
		}
		defer rows.Close()
		for rows.Next() {
			if serr := rows.Scan(&res.S[c].Name, &res.S[c].Hospital_number, &res.S[c].Attandance_number, &res.S[c].Sicker_id); serr != nil {
				gloerr = serr
			}
			res.S[c].Has = 1
			c++
			if c == 50 {
				pnt.Errorf("%s-扫描将超过50个", log)
				break
			}
		}

	case 5:
		pnt.Search("住院号:%s,就诊号:%s", h, a)
		rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE hospital_number=? AND attandance_number=?", h, a)
		if err != nil {
			gloerr = err
			break
		}
		defer rows.Close()
		for rows.Next() {
			if serr := rows.Scan(&res.S[c].Name, &res.S[c].Hospital_number, &res.S[c].Attandance_number, &res.S[c].Sicker_id); serr != nil {
				gloerr = serr
			}
			res.S[c].Has = 1
			c++
			if c == 50 {
				pnt.Errorf("%s-扫描将超过50个", log)
				break
			}
		}
	case 6:
		pnt.Search("住院号:%s", h)
		rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE hospital_number=?", h)
		if err != nil {
			gloerr = err
			break
		}
		defer rows.Close()
		for rows.Next() {
			if serr := rows.Scan(&res.S[c].Name, &res.S[c].Hospital_number, &res.S[c].Attandance_number, &res.S[c].Sicker_id); serr != nil {
				gloerr = serr
			}
			res.S[c].Has = 1
			c++
			if c == 50 {
				pnt.Errorf("%s-扫描将超过50个", log)
				break
			}
		}
	case 7:
		pnt.Search("就诊号:%s", a)
		rows, err := DB.Query("SELECT name,hospital_number,attandance_number,userid FROM sicker WHERE attandance_number=?", a)
		if err != nil {
			gloerr = err
			break
		}
		defer rows.Close()
		for rows.Next() {
			if serr := rows.Scan(&res.S[c].Name, &res.S[c].Hospital_number, &res.S[c].Attandance_number, &res.S[c].Sicker_id); serr != nil {
				gloerr = serr
			}
			res.S[c].Has = 1
			c++
			if c == 50 {
				pnt.Errorf("%s-扫描将超过50个", log)
				break
			}
		}
	default:
		return reParseJSON(getAns(1, "暂不支持该类型的搜索方式！", ""))

	}

	if gloerr != nil {
		pnt.Error(gloerr)
		return reParseJSON(getAns(1, "搜索失败！", ""))
	}

	res.Status = 1
	return reParseJSON(res)
}

// 搜索 患者 详细信息:
func (sds *searchDeatilSick) msgMain() []byte {
	var log = fmt.Sprintf("搜索患者详细信息-患者ID:%s", sds.Userid)

	var sdsr searchDeatilSickRes
	err := DB.QueryRow("SELECT name,age,gender,telphone,hospital_number,attandance_number,disease,know FROM sicker WHERE userid=?",
		sds.Userid).Scan(&sdsr.Name, &sdsr.Age, &sdsr.Gender, &sdsr.Telphone, &sdsr.Hospital_number, &sdsr.Attandance_number, &sdsr.Disease, &sdsr.Know)
	if err != nil {
		pnt.Infof("%s-搜索失败,%v", log, err)

		return reParseJSON(getAns(1, "搜索失败！", ""))
	}
	sdsr.Status = 1
	return reParseJSON(sdsr)
}

// 搜索 患者 化疗周期
func (ci *cycleInfo) msgMain() []byte {
	pnt.Search("搜索患者化疗周期-患者ID:%s", ci.Userid)

	var cir cycleInfoRes
	var d, t string
	var c int = 0
	rows, err := DB.Query("SELECT userid,cycle_seq,date,time FROM cycle WHERE userid=?", ci.Userid)

	if err == sql.ErrNoRows {
		cir.Status = 1
		return reParseJSON(cir)

	} else if err != nil {
		pnt.Infof("搜索患者化疗周期-患者ID:%s,%v", ci.Userid, err)
		cir.Status = 0
		return reParseJSON(cir)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&cir.S[c].Userid, &cir.S[c].Cycle_seq, &d, &t)
		cir.S[c].Anstime = d + " " + t
		cir.S[c].Has = 1
		c++
	}
	cir.Status = 1
	return reParseJSON(cir)
}

// 查询 风险评估
func (rirc *riskInfoRec) msgMain() []byte {
	var rirs riskInfoRes
	var last_seq int
	var log string = fmt.Sprintf("查询风险评估-患者ID:%s,化疗周期:%d", rirc.Userid, rirc.Cycle_seq)

	// 至少是第二周期开始查询上次记录
	if rirc.Cycle_seq > 1 {

		// 查询上次风险等级
		if err := DB.QueryRow("SELECT grand FROM risk WHERE userid=? AND cycle_seq=?",
			rirc.Userid, rirc.Cycle_seq-1).Scan(
			&rirs.Last_risk_grand); err != nil && err != sql.ErrNoRows {
			pnt.Errorf("%s(查询上次风险等级)-%v", log, err)
			return reParseJSON(getAns(0, "查询上一周期内容失败！", ""))
		}

		// 查询上次护理呕吐和恶心
		if err := DB.QueryRow("SELECT Emesis_assessment,Nausea_assessment,nurse_seq FROM nurse WHERE userid=? AND cycle_seq=? ORDER BY nurse_seq DESC LIMIT ?",
			rirc.Userid, rirc.Cycle_seq-1, 1).Scan(
			&rirs.Last_nurse_emesis,
			&rirs.Last_nurse_nausea,
			&last_seq); err != nil && err != sql.ErrNoRows {
			pnt.Errorf("%s(查询上次护理记录)-%v", log, err)
			return reParseJSON(getAns(0, "查询上一周期内容失败！", ""))
		}

		// 查询上次随访呕吐和恶心
		if err := DB.QueryRow("SELECT emesis_grade,nausea_grade,follow_seq FROM follow WHERE userid=? AND cycle_seq=? ORDER BY follow_seq DESC LIMIT ?",
			rirc.Userid, rirc.Cycle_seq-1, 1).Scan(
			&rirs.Last_follow_emesis,
			&rirs.Last_follow_nausea,
			&last_seq); err != nil && err != sql.ErrNoRows {
			pnt.Errorf("%s(查询上次随访记录)-%v", log, err)
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
	if err != nil {
		pnt.Error(err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	rirs.Status = 1
	return reParseJSON(rirs)
}

// 查询 患者 护理具体信息
func (nirc *nurseInfoRec) msgMain() []byte {
	var nirs nurseInfoRes
	pnt.Search("护理具体信息-患者ID:%s,化疗周期:%d,护理次序:%d", nirc.Userid, nirc.Cycle_seq, nirc.Nurse_seq)
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
	if err != nil {
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	nirs.Status = 1
	return reParseJSON(nirs)
}

// 查询 患者 护理表
func (ntrc *nurseTableRec) msgMain() []byte {
	var ntrs nurseTableReS
	var i int = 0
	var d, t string
	var log string = fmt.Sprintf("查询护理表-患者ID:%s,化疗周期:%d", ntrc.Userid, ntrc.Cycle_seq)

	pnt.Info(log)
	rows, err := DB.Query("SELECT nurse_seq,assessment_date,assessment_time FROM nurse WHERE userid=? AND cycle_seq=?", ntrc.Userid, ntrc.Cycle_seq)
	if err != nil {
		pnt.Error(err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ntrs.N[i].Nurse_seq, &d, &t)
		if err != nil {
			pnt.Error(err)
			return reParseJSON(getAns(0, "查询失败！", ""))
		}
		ntrs.N[i].Time = d + " " + t
		ntrs.N[i].Has = 1
		i++
		if i == 50 {
			pnt.Errorf("%s-扫描将超过50个", log)
			break
		}
	}
	ntrs.Status = 1
	return reParseJSON(ntrs)
}

// 查询 患者 随访表
func (ftrc *followTableRec) msgMain() []byte {
	var ftrs followTableRes
	var log string = fmt.Sprintf("查询随访表-患者ID:%s", ftrc.Userid)

	var i int = 0
	var d, t string
	pnt.Info(log)
	rows, err := DB.Query("SELECT follow_seq,follow_follow_date,follow_follow_time FROM follow WHERE userid=? AND cycle_seq=?", ftrc.Userid, ftrc.Cycle_seq)
	if err != sql.ErrNoRows && err != nil {
		pnt.Error(err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&ftrs.N[i].Follow_seq, &d, &t)
		if err != nil {
			pnt.Error(err)
			return reParseJSON(getAns(0, "查询失败！", ""))
		}
		ftrs.N[i].Has = 1
		ftrs.N[i].Time = d + " " + t
		i++
		if i == 50 {
			pnt.Errorf("%s-扫描将超过50个", log)
			break
		}

	}

	ftrs.Status = 1
	return reParseJSON(ftrs)
}

// 查询 患者 随访具体内容
func (fcrc *followContentRec) msgMain() []byte {
	pnt.Search("查询随访具体信息-患者ID:%s,化疗周期:%d,随访周期:%d", fcrc.Userid, fcrc.Cycle_seq, fcrc.Follow_seq)
	var fcrs followContentRes
	err := DB.QueryRow("SELECT * from follow WHERE userid=? AND follow_seq=? AND cycle_seq=?", fcrc.Userid, fcrc.Follow_seq, fcrc.Cycle_seq).Scan(
		&fcrs.Userid,
		&fcrs.Cycle_seq,
		&fcrs.Follow_seq,
		&fcrs.Hight_risk,
		&fcrs.Emesis_grade,
		&fcrs.Nausea_grade,
		&fcrs.Out_content,
		&fcrs.Out_content_diy,
		&fcrs.Follow_over,
		&fcrs.Satisfaction_1,
		&fcrs.Satisfaction_2,
		&fcrs.Satisfaction_3,
		&fcrs.Satisfaction_4,
		&fcrs.Satisfaction_5,
		&fcrs.Satisfaction_total,
		&fcrs.Writer,
		&fcrs.Follow_follow_date,
		&fcrs.Follow_follow_time,
		&fcrs.Follow_follow_timestamp)

	if err != nil {
		pnt.Error(err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	fcrs.Status = 1
	return reParseJSON(fcrs)
}

// 查询 患者 是否出院
func (ohpc *outHospitalRec) msgMain() []byte {
	var ohps outHospitalRes
	pnt.Search("查询是否出院-患者ID:%s", ohpc.Userid)
	err := DB.QueryRow(`SELECT out_hospital_time FROM cycle WHERE cycle_seq=? AND userid=?`, ohpc.Cycle_seq, ohpc.Userid).Scan(&ohps.Time)
	if err == sql.ErrNoRows || ohps.Time == "" {
		return reParseJSON(getAns(1, "未出院！", ""))
	} else if err != sql.ErrNoRows && err != nil {
		pnt.Error(err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}

	ohps.Status = 1
	return reParseJSON(ohps)
}

// 查询 今日随访
func (wgrc *waitGoRec) msgMain() []byte {
	var wgrs waitGoRes
	var log = fmt.Sprintf("查询今日随访")

	wgrs.Status = 2
	wgrs.Explain = "今日无随访患者！"
	var i int = 0

	// 随访条件：随访未结束、已出院
	rows, err := DB.Query("SELECT userid,name,cycle_seq FROM cycle WHERE follow_over!=? AND TO_DAYS( NOW( ) ) - TO_DAYS( out_hospital_time) > ? ORDER BY cycle_seq", "1", 2)
	if err == sql.ErrNoRows {
		wgrs.Status = 2
		wgrs.Explain = "今日无随访患者！"
		return reParseJSON(wgrs)
	}
	if err != nil {
		pnt.Errorf("%s-查询化疗表失败-%v", log, err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	defer rows.Close()
	for rows.Next() {
		var t, userid, name string
		var cycle_seq int
		err := rows.Scan(&userid, &name, &cycle_seq)
		if err != nil {
			pnt.Errorf("%s-扫描化疗表失败-%v", log, err)
			return reParseJSON(getAns(0, "查询失败！", ""))
		}
		//pnt.Search("%s 已出院但随访未结束 姓名:%s ID:%s", log, name, userid)

		// 根据 化疗周期、ID查找今天是否有随访
		if err := DB.QueryRow("SELECT follow_follow_date FROM follow WHERE userid=? AND cycle_seq=? AND to_days(follow_follow_date) = to_days(now())", userid, cycle_seq).Scan(&t); err != nil {

			if err == sql.ErrNoRows {
				//pnt.Search("%s 姓名:%s", log, name)
				wgrs.N[i].Has = 1
				wgrs.N[i].Userid = userid
				wgrs.N[i].Name = name
				wgrs.N[i].Cycle_seq = cycle_seq
				wgrs.Status = 1
				i++

				if i == 50 {
					pnt.Errorf("%s-扫描将超过50个", log)
					break
				}
				continue
			}
			pnt.Errorf("%s-查询随访表-%v", log, err)
			return reParseJSON(getAns(0, "查询失败！", ""))
		}
		continue
	}
	return reParseJSON(wgrs)
}

// 查询 随访中高致吐风险
func (hrrc *heightRiskRec) msgMain() []byte {
	var hrrs heightRiskReS
	var log = fmt.Sprintf("查询高致吐风险-患者ID:%s", hrrc.Userid)

	err := DB.QueryRow("SELECT grand FROM risk WHERE userid=? AND cycle_seq=?", hrrc.Userid, hrrc.Cycle_seq).Scan(&hrrs.Height)
	if err == sql.ErrNoRows {
		return reParseJSON(getAns(2, "风险评估的风险等级未填写", ""))
	} else if err != sql.ErrNoRows && err != nil {
		pnt.Errorf("%s-%v", log, err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}

	hrrs.Status = 1
	return reParseJSON(hrrs)
}

// 查询 随访周期 上一次
func (clrc *cycleLastRec) msgMain() []byte {
	var clrs cycleLastRes
	var last_seq int
	var log string = fmt.Sprintf("查询上次随访周期-患者ID:%s,化疗周期:%d", clrc.Userid, clrc.Cycle_seq-1)

	// 至少是第二周期开始查询上次记录
	if clrc.Cycle_seq < 1 {
		reParseJSON(getAns(0, "查询上一周期内容失败！", ""))
	}

	// 查询上次风险等级
	if err := DB.QueryRow("SELECT grand FROM risk WHERE userid=? AND cycle_seq=?",
		clrc.Userid, clrc.Cycle_seq-1).Scan(
		&clrs.Last_risk_grand); err != nil && err != sql.ErrNoRows {
		pnt.Errorf("%s-查询风险表-%v", log, err)
		return reParseJSON(getAns(0, "查询上一周期内容失败！", ""))
	}

	// 查询上次护理呕吐和恶心
	if err := DB.QueryRow("SELECT Emesis_assessment,Nausea_assessment,nurse_seq FROM nurse WHERE userid=? AND cycle_seq=? ORDER BY nurse_seq ASC LIMIT ?",
		clrc.Userid, clrc.Cycle_seq-1, 1).Scan(
		&clrs.Last_nurse_emesis,
		&clrs.Last_nurse_nausea,
		&last_seq); err != nil {
		if err == sql.ErrNoRows {
			clrs.Last_nurse_emesis = ""
			clrs.Last_nurse_nausea = ""

		} else {
			pnt.Errorf("%s-查询护理表-%v", log, err)
			return reParseJSON(getAns(0, "查询上一周期内容失败！", ""))
		}
	}

	// 查询上次随访呕吐和恶心
	if err := DB.QueryRow("SELECT emesis_grade,nausea_grade,follow_seq FROM follow WHERE userid=? AND cycle_seq=? ORDER BY follow_seq ASC LIMIT ?",
		clrc.Userid, clrc.Cycle_seq-1, 1).Scan(
		&clrs.Last_follow_emesis,
		&clrs.Last_follow_nausea, &last_seq); err != nil {
		if err == sql.ErrNoRows {
			clrs.Last_follow_emesis = ""
			clrs.Last_follow_nausea = ""
		} else {
			pnt.Errorf("%s-查询随访表-%v", log, err)
			return reParseJSON(getAns(0, "查询上一周期内容失败！", ""))
		}
	}

	clrs.Status = 1
	return reParseJSON(clrs)
}

// 查询 今日护理
func (tonrc *toNurseRec) msgMain() []byte {

	var wgrs toNurseRes
	var i int = 0
	var log = fmt.Sprintf("查询今日护理")

	// 查询周期表中 没有出院（护理未结束）和随访未结束的病人
	rows, err := DB.Query("SELECT userid,cycle_seq,name FROM cycle WHERE LENGTH(out_hospital_time)=0 AND follow_over=?", 2)
	if err != nil {
		pnt.Errorf("%s(周期表)-%v", log, err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&wgrs.N[i].Userid, &wgrs.N[i].Cycle_seq, &wgrs.N[i].Name)
		if err != nil {
			pnt.Errorf("%s(周期表扫描)-%v", log, err)
			return reParseJSON(getAns(0, "查询失败！", ""))
		}
		var t string
		pnt.Infof("%s 护理、随访未结束 姓名:%s ID:%s", log, wgrs.N[i].Name, wgrs.N[i].Userid)
		// 查询这个化疗周期中 化疗时间是否大于1天

		if derr := DB.QueryRow("SELECT chemotherapy_date FROM risk WHERE TO_DAYS(NOW())-TO_DAYS(chemotherapy_date) <? AND need_nurse=? AND userid=? AND cycle_seq=?", 1, 1, wgrs.N[i].Userid, wgrs.N[i].Cycle_seq).Scan(&t); derr == sql.ErrNoRows {

			pnt.Infof("%s 化疗时间大于1天 姓名:%s ID:%s", log, wgrs.N[i].Name, wgrs.N[i].Userid)

			// 查询这个化疗周期中 护理时间是不是今天
			terr := DB.QueryRow("SELECT assessment_date FROM nurse WHERE userid=? AND to_days(assessment_date)=to_days(now()) AND cycle_seq=? ORDER BY nurse_seq DESC LIMIT ?", wgrs.N[i].Userid, wgrs.N[i].Cycle_seq, 1).Scan(&t)

			if terr == sql.ErrNoRows {
				pnt.Infof("%s 今日护理 姓名:%s ID:%s", log, wgrs.N[i].Name, wgrs.N[i].Userid)

				wgrs.N[i].Has = 1
				i++
			} else if terr != nil {
				pnt.Errorf("%s(护理表扫描)-%v", log, err)
				return reParseJSON(getAns(0, "查询失败！", ""))
			}
		} else if derr != sql.ErrNoRows && derr != nil {
			pnt.Errorf("%s(风险表扫描)-%v", log, err)
			return reParseJSON(getAns(0, "查询失败！", ""))
		}
		// 超过十五个自动结束
		if i == 50 {
			pnt.Errorf("%s-扫描将超过50个", log)
			break
		}

	}

	if i == 0 {
		return reParseJSON(getAns(2, "今日无护理患者", ""))
	}

	wgrs.Status = 1
	return reParseJSON(wgrs)
}

// 查询 第一周期非药物因素
func (lnmrc *lastnotmedicationRec) msgMain() []byte {
	var lnmrs lastnotmedicationRes
	var log string = fmt.Sprintf("查询第一周期非药物因素")

	err := DB.QueryRow("SELECT not_medication FROM risk WHERE userid=? AND cycle_seq=?", lnmrc.Userid, 1).Scan(&lnmrs.Not_medication)
	if err != nil {
		pnt.Errorf("%s(搜索化疗表失败)-%v", log, err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	lnmrs.Status = 1
	return reParseJSON(lnmrs)
}

// 搜索 患者提交的表
func (sswic *seaSickerWriteInfoRec) msgMain() []byte {
	var sswis seaSickerWriteInfoRes
	var log string = fmt.Sprintf("搜索患者填写表-患者ID:%s-化疗周期:%d", sswic.Userid, sswic.Cycle_seq)
	err := DB.QueryRow("SELECT nausea_assessment,emesis_assessment,measure FROM nurse WHERE userid=? AND cycle_seq=?", sswic.Userid, sswic.Cycle_seq).Scan(&sswis.Nausea_assessment, &sswis.Emesis_assessment, &sswis.Measure)
	if err == sql.ErrNoRows {
		sswis.Status = 2
		pnt.Errorf("%s(查询护理表)-%v-为空", log, err)

		return reParseJSON(sswis)

	} else if err != nil && err != sql.ErrNoRows {
		pnt.Errorf("%s(查询护理表失败)-%v", log, err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	} else {
		err := DB.QueryRow("SELECT satisfaction_1,satisfaction_2,satisfaction_3,satisfaction_4,satisfaction_5 FROM follow WHERE userid=? AND cycle_seq=?", sswic.Userid, 1).Scan(&sswis.Satisfaction_1, &sswis.Satisfaction_2, &sswis.Satisfaction_3, &sswis.Satisfaction_4, &sswis.Satisfaction_5)
		if err != nil && err != sql.ErrNoRows {
			pnt.Errorf("%s(查询随访表失败)-%v", log, err)
			return reParseJSON(getAns(0, "查询失败！", ""))
		}
		sswis.Status = 1
		return reParseJSON(sswis)
	}
}

// 提交 患者填写的表
func (swrirc *subSickerWriteInfoRec) msgMain() []byte {
	var swrirs subSickerWriteInfoRes
	var log string = fmt.Sprintf("提交患者填写表-患者ID:%s-化疗周期:%d", swrirc.Userid, swrirc.Cycle_seq)
	// 患者-今日第一次填写
	if swrirc.Updated == 0 {
		_, err := DB.Exec("INSERT INTO nuese (nausea_assessment,emesis_assessment,measure) WHERE userid=? AND cycle_seq=?", swrirc.Userid, swrirc.Cycle_seq)
		if err != nil {
			pnt.Errorf("%s(插入护理表失败)-%v", log, err)
			return reParseJSON(getAns(0, "插入失败！", ""))
		}
		if swrirc.Cycle_seq == 1 {
			_, err := DB.Exec("INSERT INTO follow (satisfaction_1,satisfaction_2,satisfaction_3,satisfaction_4,satisfaction_5,satisfaction_total) WHERE userid=? AND cycle_seq=?", swrirc.Userid, swrirc.Cycle_seq)
			if err != nil {
				pnt.Errorf("%s(插入随访表失败)-%v", log, err)
				return reParseJSON(getAns(0, "查询失败！", ""))
			}
		} else {
			var a, b, c, d, e, t string
			err := DB.QueryRow("SELECT satisfaction_1,satisfaction_2,satisfaction_3,satisfaction_4,satisfaction_5,satisfaction_total FROM follow WHERE userid=? AND cycle_seq=?", swrirc.Userid, swrirc.Cycle_seq).Scan(&a, &b, &c, &d, &e, &t)
			if err != nil {
				pnt.Errorf("%s(查找随访表失败)-%v", log, err)
				return reParseJSON(getAns(0, "查找失败！", ""))
			}
		}
	} else if swrirc.Updated == 1 {
		if _, err := DB.Exec("UPDATE nurse SET nausea_assessment=?,emesis_assessment=?,measure=? WHERE userid=? AND cycle_seq=?",
			swrirc.Nausea_assessment, swrirc.Emesis_assessment, swrirc.Measure, swrirc.Userid, swrirc.Cycle_seq); err != nil {
			pnt.Errorf("%s(更新护理表失败)-%v", log, err)
			return reParseJSON(getAns(0, "更新失败！", ""))
		}
		if err2 := DB.QueryRow("UPDATE follow SET satisfaction_1=?,satisfaction_2=?,satisfaction_3=?,satisfaction_4=?,satisfaction_5=?,satisfaction_total=? WHERE userid=?",
			swrirc.Satisfaction_1, swrirc.Satisfaction_2, swrirc.Satisfaction_3, swrirc.Satisfaction_4, swrirc.Satisfaction_5, swrirc.Satisfaction_total, swrirc.Userid); err2 != nil {
			pnt.Errorf("%s(更新随访表失败)-%v", log, err2)
			return reParseJSON(getAns(0, "更新失败！", ""))
		}

	}
	swrirs.Status = 1
	return nil
}

// 搜索今日患者填写情况
func (tosrc *toSickerRec) msgMain() []byte {
	var tosrs toSickerRes

	var i int = 0
	var log = fmt.Sprintf("查询今日患者填写情况")

	// 查询周期表中 没有出院（护理未结束）和随访未结束的病人
	rows, err := DB.Query("SELECT userid,cycle_seq,name FROM cycle WHERE LENGTH(out_hospital_time)=0 AND follow_over=?", 2)
	if err != nil {
		pnt.Errorf("%s(周期表)-%v", log, err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&tosrs.N[i].Userid, &tosrs.N[i].Cycle_seq, &tosrs.N[i].Name); err != nil {
			pnt.Errorf("%s(周期表扫描)-%v", log, err)
			return reParseJSON(getAns(0, "查询失败！", ""))
		}

		if err := DB.QueryRow("SELECT nausea_assessment,emesis_assessment,nurse_seq FROM nurse WHERE userid=? AND cycle_seq=? ORDER BY nurse_seq DESC LIMIT 1",
			tosrs.N[i].Userid, tosrs.N[i].Cycle_seq).Scan(&tosrs.N[i].Nausea_assessment, &tosrs.N[i].Emesis_assessment, &tosrs.N[i].Nurse_seq); err != nil && err != sql.ErrNoRows {
			pnt.Errorf("%s(扫描护理表)-%v", log, err)
			return reParseJSON(getAns(0, "查询失败！", ""))
		}
		tosrs.N[i].Has = 1
		i++
		if i == 50 {
			pnt.Errorf("%s-扫描将超过50个", log)
			break
		}
	}
	tosrs.Status = 1
	return reParseJSON(tosrs)
}

// 查询明日护理随访量
func (catnfcrc *catNFCRec) msgMain() []byte {
	var catnfcrs catNFCRes
	var log = fmt.Sprintf("查询明日护理随访量")
	// 统计护理
	nurse := DB.QueryRow("SELECT count(cycle_seq) FROM cycle WHERE LENGTH(out_hospital_time)=0 AND follow_over=?", 2)
	if err := nurse.Scan(&catnfcrs.NurseCount); err != nil && err != sql.ErrNoRows {
		pnt.Errorf("%s(统计护理)-%v", log, err)
		return reParseJSON(getAns(1, "查询失败！", ""))
	}

	// 统计随访
	follow := DB.QueryRow("SELECT count(name) FROM cycle WHERE follow_over!=? AND TO_DAYS( NOW( ) )+1 - TO_DAYS( out_hospital_time)+1 > ? ORDER BY cycle_seq", "1", 2)
	if err := follow.Scan(&catnfcrs.FollowCount); err != nil && err != sql.ErrNoRows {
		pnt.Errorf("%s(统计随访)-%v", log, err)
		return reParseJSON(getAns(1, "查询失败！", ""))
	}

	catnfcrs.Status = 0
	return reParseJSON(catnfcrs)
}

// 数据下载-提交
func (dsubrc *downloadSubmitRec) msgMain() []byte {
	var dsubrs downloadSubmitRes
	var log string = fmt.Sprintf("数据下载-提交-提交人:%s-提交时间-:%s-开始时间:%s-结束时间:%s", dsubrc.Writer, dsubrc.Submit, dsubrc.Start, dsubrc.End)
	row, err := DB.Exec("INSERT INTO download (writer,submit,start,end,status) VALUES(?,?,?,?,?)", dsubrc.Writer, dsubrc.Submit, dsubrc.Start, dsubrc.End, 0)
	if err != nil {
		pnt.Errorf("%s(插入下载表)-%v", log, err)
		return reParseJSON(getAns(0, "提交失败！", ""))
	}
	dsubrs.Status = 1
	dsubrs.Explain = "提交成功！请稍后刷新查看!"

	id, errid := row.LastInsertId()
	if errid != nil {
		pnt.Errorf("%s(LastInsertId)-%v", log, err)
	}
	pnt.Infof("%s-ID:%d-状态:0", log, id)
	go downloadMain(log, dsubrc.Writer, dsubrc.Submit, dsubrc.Start, dsubrc.End, id)

	return reParseJSON(dsubrs)
}

// 数据下载-搜索
func (dsearc *downloadSearchRec) msgMain() []byte {
	var dsears downloadSearchRes
	var i int
	var log string = fmt.Sprintf("数据下载-搜索")
	rows, err := DB.Query("SELECT id,writer,start,end,submit,status FROM download ORDER BY id DESC LIMIT ?", 50)
	if err == sql.ErrNoRows {
		pnt.Errorf("%s(扫描下载表-空)-%v", log, err)
		return reParseJSON(getAns(2, "", ""))
	}
	if err != nil && err != sql.ErrNoRows {
		pnt.Errorf("%s(扫描下载表)-%v", log, err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	for rows.Next() {
		if err := rows.Scan(&dsears.N[i].ID, &dsears.N[i].Writer, &dsears.N[i].Start, &dsears.N[i].End, &dsears.N[i].Submit, &dsears.N[i].Status); err != nil {
			return reParseJSON(getAns(0, "查询失败！", ""))
		}
		i++
		if i == 50 {
			pnt.Errorf("%s-扫描将超过50个", log)
			break
		}
	}
	dsears.Status = 1
	return reParseJSON(dsears)
}

// 数据下载-重试
func (dtryrc *downloadTryRec) msgMain() []byte {
	var dtryrs downloadTryRes
	var di downloadInfo
	var log string = fmt.Sprintf("数据下载-重试-ID:%d", dtryrc.ID)
	if err := DB.QueryRow("SELECT writer,start,end,submit,status FROM download WHERE id=?",
		dtryrc.ID).Scan(&di.Writer, &di.Start, &di.End, &di.Submit, &di.Status); err != nil {

		pnt.Errorf("%s(扫描下载表)-%v", log, err)
		return reParseJSON(getAns(0, "查询失败！", ""))
	}
	fn := pugeFileName(di.Submit, di.Start, di.End)
	pnt.Infof("%s-File:%s-Status:%d", log, fn, di.Status)

	switch di.Status {
	// 数据插入,未扫描
	case 0:
		go downloadMain(log, di.Writer, di.Submit, di.Start, di.End, dtryrc.ID)
		dtryrs.Status = 1
		dtryrs.Explain = "正在生成！请稍后刷新！"

	// 扫描数据失败
	case 10:
		if notExistFile(fn) {
			dtryrs.Status = 0
			dtryrs.Explain = "系统错误！请稍后重试！"
			pnt.Errorf("%s(文件不存在)", log)
		}

	// 文件传输失败
	case 11:
		fn = env.PathFileSave + "/" + fn
		if existFile(fn) {
			if err := scpMain(fn); err != nil {
				dtryrs.Status = 0
				dtryrs.Explain = "系统错误！请稍后重试！"
				pnt.Errorf("%s(scp错误)-%v", log, err)
			} else {
				dtryrs.Status = 1
				dtryrs.Explain = "已完成！请刷新界面！"
				DB.Exec("UPDATE download SET status=? WHERE id=?", 9, dtryrc.ID)
			}
		} else {
			dtryrs.Status = 0
			dtryrs.Explain = "系统错误！请稍后重试！"
			pnt.Errorf("%s(文件不存在或其他错误)", log)
		}
	default:
		dtryrs.Explain = "请稍后重试！"
	}
	return reParseJSON(dtryrs)
}
