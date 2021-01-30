package main

// 患者表
/*
CREATE TABLE IF NOT EXISTS `sicker`(
   `userid` CHAR(15) NOT NULL,
   `name` VARCHAR(6),
   `age` CHAR(3) ,
   `gender` CHAR(2) ,
   `telphone` CHAR(11) ,
   `hospital_number` VARCHAR(12),
   `attandance_number` VARCHAR(12),
   `disease` VARCHAR(20) ,
   `know` char(11) DEFAULT '0',
   `cycle_seq` INT DEFAULT '0',
   `writer` VARCHAR(10) NOT NULL,
   `write_data`DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
   PRIMARY KEY ( `userid` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type sickerInfo struct {
	Action string `json:"action"`

	Userid            string `json:"userid"`
	Name              string `json:"name"`
	Age               string `json:"age"`
	Gender            string `json:"gender"`
	Telphone          string `json:"telphone"`
	Hospital_number   string `json:"hospital_number"`
	Attandance_number string `json:"attandance_number"`
	Disease           string `json:"disease"`
	Know              string `json:"know"`
	Writer            string `json:"writer"`
	Way               int    `json:"way"`
}

// 用户表
/*
CREATE TABLE IF NOT EXISTS `users`(
   `account` VARCHAR(15) NOT NULL,
   `name` VARCHAR(10) DEFAULT '',
   `password` VARCHAR(20) NOT NULL,
   `type` tinyint NOT NULL,
   PRIMARY KEY ( `account` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type userInfo struct {
	Account  string `json:"account"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Who      string `json:"who"`
}

// 风险评估表
/*
CREATE TABLE IF NOT EXISTS `risk`(
   `userid` CHAR(15) NOT NULL,
   `cycle_seq` INT NOT NULL,
   `program` VARCHAR(30) DEFAULT '',
   `not_medication` VARCHAR(20) DEFAULT '',
   `medication` VARCHAR(2) DEFAULT '',
   `grand` VARCHAR(2) DEFAULT '',

   `pre_program` VARCHAR(20) DEFAULT '',
   `pre_program_diy` VARCHAR(50),

   `comment` VARCHAR(20),
   `comment_diy` VARCHAR(50),
   `need_nurse` CHAR(1) DEFAULT '1',

   `writer` VARCHAR(10) NOT NULL,
   `assessment_date` CHAR(10) NOT NULL,
   `assessment_time` CHAR(5) NOT NULL,
   `assessment_timestamp` CHAR(13) NOT NULL,
   `chemotherapy_date` CHAR(10) NOT NULL,
   `chemotherapy_time` CHAR(5) NOT NULL,
   `chemotherapy_timestamp` CHAR(13) NOT NULL,
   PRIMARY KEY ( `userid`,`cycle_seq` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type riskInfo struct {
	Action    string `json:"Action"`
	Userid    string `json:"userid"`
	Cycle_seq int    `json:"cycle_seq"`

	Program         string `json:"program"`
	Not_medication  string `json:"not_medication"`
	Medication      string `json:"medication"`
	Grand           string `json:"grand"`
	Pre_program     string `json:"pre_program"`
	Pre_program_diy string `json:"pre_program_diy"`
	Comment         string `json:"comment"`
	Comment_diy     string `json:"comment_diy"`
	Need_nurse      string `json:"need_nurse"`

	Writer                 string `json:"writer"`
	Assessment_date        string `json:"assessment_date"`
	Assessment_time        string `json:"assessment_time"`
	Assessment_timestamp   string `json:"assessment_timestamp"`
	Chemotherapy_date      string `json:"chemotherapy_date"`
	Chemotherapy_time      string `json:"chemotherapy_time"`
	Chemotherapy_timestamp string `json:"chemotherapy_timestamp"`

	Name    string `json:"name"`
	Updated int    `json:"updated"`
}

/* 护理评估表
CREATE TABLE IF NOT EXISTS `nurse`(
   `userid` CHAR(15) NOT NULL,
   `cycle_seq` INT NOT NULL,
   `nurse_seq` INT NOT NULL,

   `nausea_assessment` CHAR(1) NOT NULL,
   `emesis_assessment` CHAR(1) NOT NULL,
   `measure` VARCHAR(30) NOT NULL,
   `comment` VARCHAR(50),
   `out_hospital` VARCHAR(3) NOT NULL,

   `writer` VARCHAR(10) NOT NULL,
   `assessment_date` CHAR(10) NOT NULL,
   `assessment_time` CHAR(5) NOT NULL,
   `assessment_timestamp` CHAR(13) NOT NULL,
   PRIMARY KEY ( `userid`,`cycle_seq`,`nurse_seq` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type nurseInfo struct {
	Userid    string `json:"userid"`
	Cycle_seq int    `json:"cycle_seq"`
	Nurse_seq int    `json:"nurse_seq"`

	Nausea_assessment string `json:"nausea_assessment"`
	Emesis_assessment string `json:"emesis_assessment"`
	Measure           string `json:"measure"`
	Comment           string `json:"comment"`
	Out_hospital      string `json:"out_hospital"`

	Writer               string `json:"writer"`
	Assessment_date      string `json:"nurse_assessment_date"`
	Assessment_time      string `json:"nurse_assessment_time"`
	Assessment_timestamp string `json:"nurse_assessment_timestamp"`
}

// 随访记录
/*
 CREATE TABLE IF NOT EXISTS `follow`(
   `userid` CHAR(15) NOT NULL,
   `cycle_seq` INT NOT NULL,
   `follow_seq` INT NOT NULL,
   `hight_risk` CHAR(3) NOT NULL,
   `emesis_grade` CHAR(3) NOT NULL,
   `nausea_grade` CHAR(3) NOT NULL,
   `out_content` VARCHAR(20) NOT NULL,
   `out_content_diy` VARCHAR(100),
   `follow_over` CHAR(3) NOT NULL,
   `satisfaction_1` CHAR(3),
   `satisfaction_2` CHAR(3),
   `satisfaction_3` CHAR(3),
   `satisfaction_4` CHAR(3),
   `satisfaction_5` CHAR(3),
   `satisfaction_total` CHAR(3),
   `writer` VARCHAR(10) NOT NULL,
   `follow_follow_date` CHAR(10) NOT NULL,
   `follow_follow_time` CHAR(5) NOT NULL,
   `follow_follow_timestamp` CHAR(13) NOT NULL,
   PRIMARY KEY ( `userid`,`cycle_seq`,`follow_seq` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type followInfo struct {
	Userid             string `json:"userid"`
	Cycle_seq          int    `json:"cycle_seq"`
	Follow_seq         int    `json:"follow_seq"`
	Hight_risk         string `json:"hight_risk"`
	Emesis_grade       string `json:"emesis_grade"`
	Nausea_grade       string `json:"nausea_grade"`
	Out_content        string `json:"out_content"`
	Out_content_diy    string `json:"out_content_diy"`
	Follow_over        string `json:"follow_over"`
	Satisfaction_1     string `json:"satisfaction_1"`
	Satisfaction_2     string `json:"satisfaction_2"`
	Satisfaction_3     string `json:"satisfaction_3"`
	Satisfaction_4     string `json:"satisfaction_4"`
	Satisfaction_5     string `json:"satisfaction_5"`
	Satisfaction_total string `json:"satisfaction_total"`

	Writer                  string `json:"writer"`
	Follow_follow_date      string `json:"follow_follow_date"`
	Follow_follow_time      string `json:"follow_follow_time"`
	Follow_follow_timestamp string `json:"follow_follow_timestamp"`
}

// 化疗周期
/*
 CREATE TABLE IF NOT EXISTS `cycle`(
   `userid` CHAR(15) NOT NULL,
   `cycle_seq` INT NOT NULL,
   `name` VARCHAR(10) NOT NULL,
   `out_hospital_time` CHAR(17) DEFAULT '',
   `follow_over` CHAR(1) DEFAULT '2',
   `date` CHAR(10) NOT NULL,
   `time` CHAR(5) NOT NULL,
   `timestamp` CHAR(13) NOT NULL,
PRIMARY KEY ( `userid`,`cycle_seq` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
// 数据下载
/*
 CREATE TABLE IF NOT EXISTS `download`(
   `id` INT NOT NULL AUTO_INCREMENT,
   `writer` VARCHAR(10) NOT NULL,
   `start` CHAR(10) NOT NULL,
   `end` CHAR(10) NOT NULL,
   `submit` CHAR(10) NOT NULL,
   `status` tinyint NOT NULL,
 PRIMARY KEY ( `id` )
 )ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;
*/
type downloadInfo struct {
	ID     int64  `json:"id"`
	Writer string `json:"writer"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Submit string `json:"submit"`
	Status int    `json:"status"`
}

// 小程序应答结构体
type ans struct {
	Status  int    `json:"status"`
	Explain string `json:"explain"`
	Data    string `json:"data"`
}
type baseinfo struct {
	Userid    string `json:"userid"`
	Cycle_seq int    `json:"cycle_seq"`
}
type cycleInfo struct {
	Action string `json:"Action"`
	Userid string `json:"userid"`
}
type cycleInfoRes struct {
	ans
	S [50]cycleInfoResS `json:"data"`
}
type cycleInfoResS struct {
	baseinfo
	Anstime string `json:"time"`
	Has     int    `json:"has"`
}
type searchSicker struct {
	Name              string `json:"name"`
	Hospital_number   string `json:"hospital_number"`
	Attandance_number string `json:"attandance_number"`
	Sicker_id         string `json:"userid"`
	Has               int    `json:"has"`
}
type searchSickerRes struct {
	ans
	S [50]searchSicker `json:"data"`
}
type searchDeatilSick struct {
	Action string `json:"action"`
	Userid string `json:"userid"`
}
type searchDeatilSickRes struct {
	ans
	Name              string `json:"name"`
	Age               string `json:"age"`
	Gender            string `json:"gender"`
	Telphone          string `json:"telphone"`
	Hospital_number   string `json:"hospital_number"`
	Attandance_number string `json:"attandance_number"`
	Disease           string `json:"disease"`
	Know              string `json:"know"`
}

type riskInfoRec struct {
	Action string `json:"action"`
	baseinfo
}
type riskInfoRes struct {
	ans

	baseinfo

	Program         string `json:"program"`
	Not_medication  string `json:"not_medication"`
	Medication      string `json:"medication"`
	Grand           string `json:"grand"`
	Pre_program     string `json:"pre_program"`
	Pre_program_diy string `json:"pre_program_diy"`
	Comment         string `json:"comment"`
	Comment_diy     string `json:"comment_diy"`
	Need_nurse      string `json:"need_nurse"`

	Last_risk_grand    string `json:"last_risk_grand"`
	Last_nurse_emesis  string `json:"last_nurse_emesis"`
	Last_nurse_nausea  string `json:"last_nurse_nausea"`
	Last_follow_emesis string `json:"last_follow_emesis"`
	Last_follow_nausea string `json:"last_follow_nausea"`

	Writer                 string `json:"writer"`
	Assessment_date        string `json:"assessment_date"`
	Assessment_time        string `json:"assessment_time"`
	Assessment_timestamp   string `json:"assessment_timestamp"`
	Chemotherapy_date      string `json:"chemotherapy_date"`
	Chemotherapy_time      string `json:"chemotherapy_time"`
	Chemotherapy_timestamp string `json:"chemotherapy_timestamp"`
}
type nurseTableRec struct {
	Action string `json:"action"`
	baseinfo
}
type nurseTable struct {
	Nurse_seq int    `json:"nurse_seq"`
	Time      string `json:"time"`
	Has       int    `json:"has"`
}
type nurseTableReS struct {
	ans
	N [50]nurseTable
}
type nurseInfoRec struct {
	Action string `json:"action"`
	baseinfo
	Nurse_seq int `json:"nurse_seq"`
}
type nurseInfoRes struct {
	ans
	baseinfo

	Nurse_seq         string `json:"nurse_seq"`
	Nausea_assessment string `json:"nausea_assessment"`
	Emesis_assessment string `json:"emesis_assessment"`
	Measure           string `json:"measure"`
	Comment           string `json:"comment"`
	Out_hospital      string `json:"out_hospital"`

	Writer               string `json:"writer"`
	Assessment_date      string `json:"nurse_assessment_date"`
	Assessment_time      string `json:"nurse_assessment_time"`
	Assessment_timestamp string `json:"nurse_assessment_timestamp"`
}

type outHospitalRec struct {
	Action string `json:"action"`
	baseinfo
}
type outHospitalRes struct {
	ans
	Explain string `json:"explain"`
	Time    string `json:"time"`
}
type followTable struct {
	Follow_seq int    `json:"follow_seq"`
	Time       string `json:"time"`
	Has        int    `json:"has"`
}
type followTableRec struct {
	Action string `json:"action"`
	baseinfo
}
type followTableRes struct {
	ans
	N [50]followTable
}
type followContentRec struct {
	Action string `json:"action"`
	baseinfo
	Follow_seq int `json:"follow_seq"`
}
type followContentRes struct {
	ans
	baseinfo

	Follow_seq              int    `json:"follow_seq"`
	Hight_risk              string `json:"hight_risk"`
	Emesis_grade            string `json:"emesis_grade"`
	Nausea_grade            string `json:"nausea_grade"`
	Out_content             string `json:"out_content"`
	Out_content_diy         string `json:"out_content_diy"`
	Follow_over             string `json:"follow_over"`
	Satisfaction_1          string `json:"satisfaction_1"`
	Satisfaction_2          string `json:"satisfaction_2"`
	Satisfaction_3          string `json:"satisfaction_3"`
	Satisfaction_4          string `json:"satisfaction_4"`
	Satisfaction_5          string `json:"satisfaction_5"`
	Satisfaction_total      string `json:"satisfaction_total"`
	Writer                  string `json:"writer"`
	Follow_follow_date      string `json:"follow_follow_date"`
	Follow_follow_time      string `json:"follow_follow_time"`
	Follow_follow_timestamp string `json:"follow_follow_timestamp"`
}
type waitGo struct {
	baseinfo
	Name string `json:"name"`
	Has  int    `json:"has"`
}
type waitGoRec struct {
	Action string `json:"action"`
}
type waitGoRes struct {
	ans
	N [50]waitGo
}
type heightRiskRec struct {
	Action string `json:"action"`
	baseinfo
}
type heightRiskReS struct {
	ans
	Height int `json:"height"`
}
type cycleLastRec struct {
	Action string `json:"action"`
	baseinfo
}
type cycleLastRes struct {
	ans
	Last_risk_grand    string `json:"last_risk_grand"`
	Last_nurse_emesis  string `json:"last_nurse_emesis"`
	Last_nurse_nausea  string `json:"last_nurse_nausea"`
	Last_follow_emesis string `json:"last_follow_emesis"`
	Last_follow_nausea string `json:"last_follow_nausea"`
}

type toNurse struct {
	baseinfo
	Name string `json:"name"`
	Has  int    `json:"has"`
}
type toNurseRec struct {
	Action string `json:"action"`
}
type toNurseRes struct {
	ans
	N [50]toNurse
}
type lastnotmedicationRec struct {
	Action string `json:"action"`
	Userid string `json:"userid"`
}
type lastnotmedicationRes struct {
	ans
	Not_medication string `json:"not_medication"`
}

// ID:5 start
type userLoginRes struct {
	ans
	baseinfo
}
type seaSickerWriteInfoRec struct {
	Action string `json:"action"`
	baseinfo
}
type seaSickerWriteInfoRes struct {
	ans
	Nausea_assessment string `json:"nausea_assessment"`
	Emesis_assessment string `json:"emesis_assessment"`
	Measure           string `json:"measure"`
	Satisfaction_1    string `json:"satisfaction_1"`
	Satisfaction_2    string `json:"satisfaction_2"`
	Satisfaction_3    string `json:"satisfaction_3"`
	Satisfaction_4    string `json:"satisfaction_4"`
	Satisfaction_5    string `json:"satisfaction_5"`
}
type subSickerWriteInfoRec struct {
	Action string `json:"action"`
	baseinfo

	Updated            int    `json:"updated"`
	Nausea_assessment  string `json:"nausea_assessment"`
	Emesis_assessment  string `json:"emesis_assessment"`
	Measure            string `json:"measure"`
	Satisfaction_1     string `json:"satisfaction_1"`
	Satisfaction_2     string `json:"satisfaction_2"`
	Satisfaction_3     string `json:"satisfaction_3"`
	Satisfaction_4     string `json:"satisfaction_4"`
	Satisfaction_5     string `json:"satisfaction_5"`
	Satisfaction_total string `json:"satisfaction_total"`
}
type subSickerWriteInfoRes struct {
	ans
}
type toSickerRec struct {
	Action string `json:"action"`
}
type toSickerRes struct {
	ans
	N [50]toSickerInfo
}
type toSickerInfo struct {
	baseinfo
	Name              string `json:"name"`
	Nurse_seq         int    `json:"nurse_seq"`
	Nausea_assessment string `json:"nausea_assessment"`
	Emesis_assessment string `json:"emesis_assessment"`

	Has int `json:"has"`
}

// ID:5 start over

type catNFCRec struct {
	Action string `json:"action"`
}
type catNFCRes struct {
	ans
	FollowCount int `json:"followcount"`
	NurseCount  int `json:"nursecount"`
}
type downloadSubmitRec struct {
	Action string `json:"action"`
	downloadInfo
}
type downloadSubmitRes struct {
	ans
}
type downloadSearchRec struct {
	Action string `json:"action"`
}
type downloadSearchRes struct {
	ans
	N [50]downloadInfo
}
type downloadTryRec struct {
	Action string `json:"action"`
	ID     int64  `json:"id"`
}
type downloadTryRes struct {
	ans
}
