package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	// my lib
	"env"
	pnt "print"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/go-sql-driver/mysql"
)

type data struct {
	xls        *excelize.File
	sTime      string
	eTime      string
	nTime      string
	writer     string
	log        string
	id         int64
	saveFile   string
	sickerName map[string]string
}

func pugeFileName(a, b, c string) string {
	return fmt.Sprintf("%s__%s__%s.xlsx", a, b, c)
}
func (d *data) getFileName() {
	d.saveFile = fmt.Sprintf("%s/%s__%s__%s.xlsx", env.PathFileSave, d.nTime, d.sTime, d.eTime)
}
func (d *data) setTitle(sheet, title string) int {
	list := strings.Split(title, ",")
	for seq, value := range list {
		xy, _ := excelize.CoordinatesToCellName(seq+1, 1)
		d.xls.SetCellValue(sheet, xy, value)
	}
	return len(list)
}
func (d *data) writeRows(sheet string, list []string, rowsSeq int) {
	for seq, value := range list {
		xy, _ := excelize.CoordinatesToCellName(seq+1, rowsSeq)
		d.xls.SetCellValue(sheet, xy, value)
	}
}

func downloadMain(log, writer, nTime, sTime, eTime string, id int64) {

	var d data

	runStartTime := time.Now()
	d.sTime = sTime
	d.eTime = eTime
	d.nTime = nTime
	d.writer = writer
	d.xls = excelize.NewFile()
	d.id = id
	d.log = log

	if len(d.sTime) == 0 || len(d.eTime) == 0 || len(d.nTime) == 0 || len(d.writer) == 0 {
		if _, err := DB.Exec("UPDATE download SET status=? WHERE id=?", 12, d.id); err != nil {
			pnt.Errorf("%s(数据不完整)-%v", d.log, err)
			return
		}
	}

	if err := d.queryMain(); err != nil {
		pnt.Errorf("%s(查询主函数失败)-%v", d.log, err)
		if _, err := DB.Exec("UPDATE download SET status=? WHERE id=?", 10, d.id); err != nil {
			pnt.Errorf("%s(扫描数据失败)-%v", d.log, err)
			return
		}
		return
	}

	if _, err := DB.Exec("UPDATE download SET status=? WHERE id=?", 1, d.id); err != nil {
		pnt.Errorf("%s(更新表失败)-%v", d.log, err)
		return
	}

	pnt.Infof("Collection over! Deal Time:%s", time.Since(runStartTime))

	if err := scpMain(d.saveFile); err != nil {
		pnt.Errorf("%s(传输失败)-%v", d.log, err)
		DB.Exec("UPDATE download SET status=? WHERE id=?", 11, d.id)
		return
	}
	DB.Exec("UPDATE download SET status=? WHERE id=?", 9, d.id)

	pnt.Infof("scp over! Deal Time:%s", time.Since(runStartTime))
	return

}
func scpMain(filename string) error {
	var path string
	if PodID == 0 {
		path = env.PathNginxMiniSickBase + env.PathNginxTestMiniSick
	} else {
		path = env.PathNginxMiniSickBase + env.PathNginxMiniSick
	}

	c := exec.Command("scp", "-P", "17655", filename, fmt.Sprintf("%spod%d/file", path, PodID))
	if err := c.Start(); err != nil {
		pnt.Errorf("scp start down:%v", err)
		return err
	}
	pnt.Infof("%v", c.Args)

	if err := c.Wait(); err != nil {
		pnt.Errorf("scp wait down:%v", err)
		return err
	}

	return nil
}

// 查询主要过程
func (d *data) queryMain() error {

	if err := d.queryName(); err != nil {
		return err
	}
	if PodID == 3 || PodID == 4 {
		if err := d.querySicker34(); err != nil {
			return err
		}
	} else {
		if err := d.querySicker(); err != nil {
			return err
		}
	}
	if err := d.queryRisk(); err != nil {
		return err
	}
	if err := d.queryNurse(); err != nil {
		return err
	}
	if err := d.queryFollow(); err != nil {
		return err
	}
	if err := d.Total(); err != nil {
		return err
	}
	return nil
}

// 查询 患者姓名
func (d *data) queryName() error {
	var userid, name string
	rows, err := DB.Query("SELECT userid,name FROM sicker WHERE userid IN (SELECT distinct userid FROM risk WHERE chemotherapy_date >=? AND chemotherapy_date<= ?)", d.sTime, d.eTime)
	if err != nil {
		pnt.Errorf("%s(查询患者姓名个失败)-%v", d.log, err)
		return err
	}
	columns, err := rows.Columns()
	if err != nil {
		pnt.Errorf("%s(获取患者表列名出错)-%v", d.log, err)
		return err
	}
	list := make(map[string]string, len(columns))
	for rows.Next() {
		err := rows.Scan(&userid, &name)
		if err != nil {
			pnt.Errorf("%s(扫描患者id和姓名失败)-%v", d.log, err)
			return err
		}
		list[userid] = name
	}

	d.sickerName = list
	return nil
}

// 查询 患者信息 适用podID:3、4
func (d *data) querySicker34() error {
	var i int = 2
	var name, age, gender, telphone, hospital_number, disease, know string
	const sheet = "患者信息表"
	d.xls.NewSheet(sheet)
	colCount := d.setTitle(sheet, "姓名,年龄,性别,电话,出生年月日,诊断,患者是否知情")
	list := make([]string, colCount)

	rows, err := DB.Query("SELECT name,age,gender,telphone,hospital_number,disease,know FROM sicker WHERE userid IN (SELECT distinct userid FROM risk WHERE chemotherapy_date >=? AND chemotherapy_date<= ?)", d.sTime, d.eTime)
	if err != nil {
		pnt.Errorf("%s(查询患者-查询患者表失败)-%v", d.log, err)
		return err
	}

	for rows.Next() {
		err := rows.Scan(&name, &age, &gender, &telphone, &hospital_number, &disease, &know)
		if err != nil {
			pnt.Errorf("%s(查询患者-扫描失败)-%v", d.log, err)
			return err
		}

		list[0] = name
		list[1] = age
		list[2] = gender
		list[3] = telphone
		list[4] = hospital_number
		list[5] = disease
		list[6] = convSickerKnow(know)

		d.writeRows(sheet, list, i)
		i++
	}
	return nil
}

// 查询 患者信息 适用一般
func (d *data) querySicker() error {
	var i int = 2
	var name, age, gender, telphone, hospital_number, attandance_number, disease, know string
	const sheet = "患者信息表"
	d.xls.NewSheet(sheet)
	colCount := d.setTitle(sheet, "姓名,年龄,性别,电话,住院号,就诊号,诊断,患者是否知情")
	list := make([]string, colCount)

	rows, err := DB.Query("SELECT name,age,gender,telphone,hospital_number,attandance_number,disease,know FROM sicker WHERE userid IN (SELECT distinct userid FROM risk WHERE chemotherapy_date >=? AND chemotherapy_date<= ?)", d.sTime, d.eTime)
	if err != nil {
		pnt.Errorf("%s(查询患者-查询患者表失败)-%v", d.log, err)
		return err
	}

	for rows.Next() {
		err := rows.Scan(&name, &age, &gender, &telphone, &hospital_number, &attandance_number, &disease, &know)
		if err != nil {
			pnt.Errorf("%s(查询患者-扫描失败)-%v", d.log, err)
			return err
		}

		list[0] = name
		list[1] = age
		list[2] = gender
		list[3] = telphone
		list[4] = hospital_number
		list[5] = attandance_number
		list[6] = disease
		list[7] = convSickerKnow(know)

		d.writeRows(sheet, list, i)
		i++
	}
	return nil
}

// 查询 风险评估
func (d *data) queryRisk() error {
	var i int = 2
	const sheet = "CINV风险评估表"
	var userid, cycle_seq, program, not_medication, medication, grand, pre_program, pre_program_diy, comment, comment_diy, need_nurse, writer, assessment_date, assessment_time, assessment_timestamp, chemotherapy_date, chemotherapy_time, chemotherapy_timestamp string
	d.xls.NewSheet(sheet)
	colCount := d.setTitle(sheet, "姓名,化疗周期,评估日期,评估时间,化疗日期,化疗时间,化疗方案,药物因素,非药物因素性别,CINV风险等级,预处理方案,预处理方案（其他）,备注,备注(其他),是否需要住院,填写人")

	rows, err := DB.Query("SELECT * FROM risk WHERE userid IN (SELECT distinct userid FROM risk WHERE chemotherapy_date >= ? AND chemotherapy_date <= ?)", d.sTime, d.eTime)
	if err != nil {
		pnt.Errorf("%s(风险评估-查询风险评估错误)-%v", d.log, err)
		return err
	}

	list := make([]string, colCount)
	for rows.Next() {

		err := rows.Scan(&userid, &cycle_seq, &program, &not_medication, &medication, &grand, &pre_program, &pre_program_diy, &comment, &comment_diy, &need_nurse, &writer, &assessment_date, &assessment_time, &assessment_timestamp, &chemotherapy_date, &chemotherapy_time, &chemotherapy_timestamp)
		if err != nil {
			pnt.Errorf("%s(风险评估-扫描失败)-%v", d.log, err)
			return err
		}
		list[0] = d.sickerName[userid]
		list[1] = cycle_seq
		list[2] = assessment_date
		list[3] = assessment_time
		list[4] = chemotherapy_date
		list[5] = chemotherapy_time
		list[6] = program
		list[7] = convRiskMedication(medication)
		list[8] = convRiskNotMedication(not_medication)
		list[9] = convRiskGrand(grand)
		list[10] = convRiskPreProgram(pre_program)
		list[11] = pre_program_diy
		list[12] = convRiskComment(comment)
		list[13] = comment_diy
		list[14] = convRiskNeedNurse(need_nurse)
		list[15] = writer

		d.writeRows(sheet, list, i)
		i++
	}
	return nil
}

// 查询 护理表
func (d *data) queryNurse() error {
	var i int = 2
	const sheet = "CINV护理表"
	var userid, cycle_seq, nurse_seq, nausea_assessment, emesis_assessment, measure, comment, out_hospital, writer, assessment_date, assessment_time, assessment_timestamp string
	d.xls.NewSheet(sheet)
	colCount := d.setTitle(sheet, "姓名,化疗周期,护理次序,恶心分级评估,呕吐分级评估,护理措施,备注,是否出院,护理评估日期,护理评估时间,填写人")
	list := make([]string, colCount)

	rows, err := DB.Query("SELECT * FROM nurse WHERE userid IN (SELECT distinct userid FROM risk WHERE chemotherapy_date >= ? AND chemotherapy_date <= ?)", d.sTime, d.eTime)
	if err != nil {
		pnt.Errorf("%s(护理表-查询护理表错误)-%v", d.log, err)
		return err
	}
	for rows.Next() {
		err := rows.Scan(&userid, &cycle_seq, &nurse_seq, &nausea_assessment, &emesis_assessment, &measure, &comment, &out_hospital, &writer, &assessment_date, &assessment_time, &assessment_timestamp)
		if err != nil {
			pnt.Errorf("%s(护理表-扫描失败)-%v", d.log, err)
			return err
		}
		list[0] = d.sickerName[userid]
		list[1] = cycle_seq
		list[2] = nurse_seq
		list[3] = convNurseNauseaAssessment(nausea_assessment)
		list[4] = convNurseEmesisAssessment(emesis_assessment)
		list[5] = convNurseMeasure(measure)
		list[6] = comment
		list[7] = convNurseOutHospital(out_hospital)
		list[8] = assessment_date
		list[9] = assessment_time
		list[10] = writer

		d.writeRows(sheet, list, i)
		i++
	}
	return nil
}

// 查询 随访表
func (d *data) queryFollow() error {
	var i int = 2
	const sheet = "CINV随访表"
	var userid, cycle_seq, follow_seq, hight_risk, emesis_grade, nausea_grade, out_content, out_content_diy, follow_over, satisfaction_1, satisfaction_2, satisfaction_3, satisfaction_4, satisfaction_5, satisfaction_total, writer, follow_follow_date, follow_follow_time, follow_follow_timestamp string
	d.xls.NewSheet(sheet)
	colCount := d.setTitle(sheet, "姓名,化疗周期,随访次序,CINV高风险等级,呕吐分级,恶心分级,出院后专科指导内容,出院后专科指导内容(其他内容),随访是否结束,随访时间,随访日期,填表人,满意度调查：1.住院期间，您是否知晓无呕病房？,满意度调查：2.医护人员为您介绍了CINV相关知识吗？,满意度调查：3.您是否拥有《患者关爱手册》及在护士指导下使用？,满意度调查：4.住院期间，您有恶心、呕吐时，医护人员能及时处理并耐心解决疑问吗？,满意度调查：5.在您悲伤、焦虑时，医护人员会不会安慰、帮助您吗？,满意度调查：总分")
	list := make([]string, colCount)

	rows, err := DB.Query("SELECT * FROM follow WHERE userid IN (SELECT distinct userid FROM risk WHERE chemotherapy_date >= ? AND chemotherapy_date <= ?)", d.sTime, d.eTime)
	if err != nil {
		pnt.Errorf("%s(随访表-查询随访表错误)-%v", d.log, err)
		return err
	}
	for rows.Next() {
		err := rows.Scan(&userid, &cycle_seq, &follow_seq, &hight_risk, &emesis_grade, &nausea_grade, &out_content, &out_content_diy, &follow_over, &satisfaction_1, &satisfaction_2, &satisfaction_3, &satisfaction_4, &satisfaction_5, &satisfaction_total, &writer, &follow_follow_date, &follow_follow_time, &follow_follow_timestamp)
		if err != nil {
			pnt.Errorf("%s(随访表-扫描失败)-%v", d.log, err)
			return err
		}
		list[0] = d.sickerName[userid]
		list[1] = cycle_seq
		list[2] = follow_seq
		list[3] = convFollowHightRisk(hight_risk)
		list[4] = convFollowEmesisGrade(emesis_grade)
		list[5] = convFollowNauseaGrade(nausea_grade)
		list[6] = convFollowOutContent(out_content)
		list[7] = out_content_diy
		list[8] = convFollowFollowOver(follow_over)
		list[9] = follow_follow_date
		list[10] = follow_follow_time
		list[11] = writer
		list[12] = convFollowSatisfactionTable(satisfaction_1)
		list[13] = convFollowSatisfactionTable(satisfaction_2)
		list[14] = convFollowSatisfactionTable(satisfaction_3)
		list[15] = convFollowSatisfactionTable(satisfaction_4)
		list[16] = convFollowSatisfactionTable(satisfaction_5)
		list[17] = convFollowSatisfactionTotal(satisfaction_total)

		d.writeRows(sheet, list, i)
		i++
	}
	return nil
}

// 汇总，保存文件
func (d *data) Total() error {

	d.xls.DeleteSheet("Sheet1")
	d.xls.SetActiveSheet(0)
	d.getFileName()
	pnt.Infof("%s-save as %s", d.log, d.saveFile)
	if err := d.xls.SaveAs(d.saveFile); err != nil {
		pnt.Errorf("%s(保存文件失败)-%v", d.log, err)
		return err
	}
	return nil
}
