package main

// 患者表
/* 
CREATE TABLE IF NOT EXISTS `sicker`(
   `userid` char(15) NOT NULL,
   `name` VARCHAR(6) NOT NULL,
   `age` INT NOT NULL,
   `gender` CHAR(2) NOT NULL,
   `telphone` CHAR(11) NOT NULL,
   `hospital_number` INT,
   `attandance_number` INT,
   `disease` VARCHAR(20) NOT NULL,
   `out_hospital` char(16),
   `follow_over` tinyint,
   `writer` VARCHAR(10) NOT NULL,
   `write_data`DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
   PRIMARY KEY ( `userid` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type sickerInfo struct {
    Action                  string  `json:"action"`
    Name                    string  `json:"name"`
    Age                     int     `json:"age"`
    Gender                  string  `json:"gender"`
    Telphone                string  `json:"telphone"`
    Hospital_number         int     `json:"hospital_number"`
    Attandance_number       int     `json:"attandance_number"`
    Disease                 string  `json:"disease"`
    Writer                  string  `json:"writer"`
}


// 用户表
/* 
CREATE TABLE IF NOT EXISTS `users`(
   `account` varchar(15) NOT NULL,
   `name` varchar(10) DEFAULT '',
   `password` varchar(20) NOT NULL,
   `type` tinyint NOT NULL,
   PRIMARY KEY ( `account` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
type:用户类型
1:test
2:nurse
*/
type userInfo struct {
    Account                 string  `json:"account"`
    Name                    string  `json:"name"`
    Password                string  `json:"password"`
}


// 风险评估表
/* 
CREATE TABLE IF NOT EXISTS `risk`(
   `userid` char(15) NOT NULL,
   `cycle_seq` int NOT NULL,
   `program` VARCHAR(30) NOT NULL,
   `not_medication` varchar(20) NOT NULL,
   `medication`  varchar(2) NOT NULL,
   `grand` varchar(2) NOT NULL,

   `pre_program` varchar(20) NOT NULL,
   `pre_program_diy` varchar(50),

   `comment` varchar(20),
   `comment_diy` varchar(50),

   `writer` varchar(10) NOT NULL,
   `assessment_date` char(10) NOT NULL,
   `assessment_time` char(5) NOT NULL,
   `assessment_timestamp` CHAR(13) NOT NULL,
   `chemotherapy_date` char(10) NOT NULL,
   `chemotherapy_time` char(5) NOT NULL,
   `chemotherapy_timestamp` CHAR(13) NOT NULL,
   PRIMARY KEY ( `userid`,`cycle_seq` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type riskInfo struct {
    Action                  string  `json:"Action"`
    Userid                  string  `json:"userid"`
    Cycle_seq               int     `json:"cycle_seq"`

    Program                 string  `json:"program"`
    Not_medication       [] string  `json:"not_medication"`
    Medication              string  `json:"medication"`
    Grand                   string  `json:"grand"`
    Pre_program          [] string  `json:"pre_program"`
    Pre_program_diy         string  `json:"pre_program_diy"`
    Comment              [] string  `json:"comment"`
    Comment_diy             string  `json:"comment_diy"`

    Writer                  string  `json:"writer"`
    Assessment_date         string  `json:"assessment_date"`
    Assessment_time         string  `json:"assessment_time"`
    Assessment_timestamp    string  `json:"assessment_timestamp"`
    Chemotherapy_date       string  `json:"chemotherapy_date"`
    Chemotherapy_time       string  `json:"chemotherapy_time"`
    Chemotherapy_timestamp  string  `json:"chemotherapy_timestamp"`
}


/* 护理评估表
CREATE TABLE IF NOT EXISTS `nurse`(
   `userid` char(15) NOT NULL,
   `cycle_seq` int NOT NULL,
   `nurse_seq` int NOT NULL,

   `nausea_assessment` char(1) NOT NULL,
   `emesis_assessment` char(1) NOT NULL,
   `measure` varchar(30) NOT NULL,
   `comment` varchar(50),
   `out_hospital` varchar(3) NOT NULL,

   `writer` varchar(10) NOT NULL,
   `assessment_date` char(10) NOT NULL,
   `assessment_time` char(5) NOT NULL,
   `assessment_timestamp` CHAR(13) NOT NULL,
   PRIMARY KEY ( `userid`,`cycle_seq`,`nurse_seq` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type nurseInfo struct{
    Userid                  string  `json:"userid"`
    Cycle_seq               int     `json:"cycle_seq"`
    Nurse_seq               int     `json:"nurse_seq"`

    Nausea_assessment       string  `json:"nausea_assessment"` 
    Emesis_assessment       string  `json:"emesis_assessment"`
    Measure              [] string  `json:"measure"`
    Comment                 string  `json:"comment"`
    Out_hospital            string  `json:"out_hospital"`

    Writer                  string  `json:"writer"`
    Assessment_date         string  `json:"nurse_assessment_date"`
    Assessment_time         string  `json:"nurse_assessment_time"`
    Assessment_timestamp    string  `json:"nurse_assessment_timestamp"`

}

// 随访记录
/* 
 CREATE TABLE IF NOT EXISTS `follow`(
   `userid` char(15) NOT NULL,
   `follow_seq` int NOT NULL,
   `hight_risk` char(3) NOT NULL,
   `emesis_grade` char(3) NOT NULL,
   `nausea_grade` char(3) NOT NULL,
   `out_content` varchar(20) NOT NULL,
   `out_content_diy` varchar(100),
   `follow_over` char(3) NOT NULL,
   `satisfaction_1` char(3),
   `satisfaction_2` char(3),
   `satisfaction_3` char(3),
   `satisfaction_4` char(3),
   `satisfaction_5` char(3),
   `satisfaction_total` int,
   `writer` varchar(10) NOT NULL,
   `follow_follow_date` char(10) NOT NULL,
   `follow_follow_time` char(5) NOT NULL,
   `follow_follow_timestamp` CHAR(13) NOT NULL,
   PRIMARY KEY ( `userid`,`follow_seq` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type followInfo struct{
    Userid                  string  `json:"userid"`
    Follow_seq              int     `json:"follow_seq"`
    Hight_risk              string  `json:"hight_risk"`
    Emesis_grade            string  `json:"emesis_grade"`
    Nausea_grade            string  `json:"nausea_grade"`
    Out_content           []string  `json:"out_content"`
    Out_content_diy         string  `json:"out_content_diy"`
    Follow_over             string  `json:"follow_over"`
    Satisfaction_if         int     `json:"satisfaction_if"`
    Satisfaction_1          string  `json:"satisfaction_1"`
    Satisfaction_2          string  `json:"satisfaction_2"`
    Satisfaction_3          string  `json:"satisfaction_3"`
    Satisfaction_4          string  `json:"satisfaction_4"`
    Satisfaction_5          string  `json:"satisfaction_5"`

    Writer                  string  `json:"writer"`
    Follow_follow_date      string  `json:"follow_follow_date"`
    Follow_follow_time      string  `json:"follow_follow_time"`
    Follow_follow_timestamp string  `json:"follow_follow_timestamp"`
}

// 化疗周期
/*
 CREATE TABLE IF NOT EXISTS `cycle`(
   `userid` char(15) NOT NULL,
   `cycle_seq` int NOT NULL,
   `date` char(10) NOT NULL,
   `time` char(5) NOT NULL,
   `timestamp` CHAR(13) NOT NULL,
PRIMARY KEY ( `userid`,`cycle_seq` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type cycleInfo struct{
    Action                  string  `json:"Action"`
    Userid                  string  `json:"userid"`
}
type cycleInfoRes struct{
    Status                  int     `json:"status"`
    S              [15] cycleInfoResS `json:"data"`
}
type cycleInfoResS struct{
    Userid                  string  `json:"userid"`
    Cycle_seq               int     `json:"cycle_seq"`
    Anstime                 string  `json:"time"`
    Has                     int     `json:"has"`
}
// 小程序应答结构体
type ans struct{
    Status                  int     `json:"status"`
    Explain                 string  `json:"explain"`
    Data                    string  `json:"data"`
}
type searchSicker struct{
    Name                    string  `json:"name"`
    Hospital_number         string  `json:"hospital_number"`
    Attandance_number       string  `json:"attandance_number"`
    Sicker_id               string  `josn:"sicker_id"`
    Has                     int     `json:"has"`
}
type searchSickerRes struct{

    Status                  int     `json:"status"`
    S            [15] searchSicker `json:"data"`
}
type searchDeatilSick struct{
    Action                  string  `json:"action"`
    Userid                  string  `json:"userid"`
}
type searchDeatilSickRes struct{
    Status                  int     `json:"status"`
    Name                    string  `json:"name"`
    Age                     int     `json:"age"`
    Gender                  string  `json:"gender"`
    Telphone                string  `json:"telphone"`
    Hospital_number         int     `json:"hospital_number"`
    Attandance_number       int     `json:"attandance_number"`
    Disease                 string  `json:"disease"`
}

type riskInfoRec struct{
    Action                  string  `json:"action"`
    Userid                  string  `json:"userid"`
    Cycle_seq               int  `json:"cycle_seq"`

}
type riskInfoRes struct{
    Status                  int     `json:"status"`

    Userid                  string  `json:"userid"`
    Cycle_seq               int     `json:"cycle_seq"`

    Program                 string  `json:"program"`
    Not_medication          string  `json:"not_medication"`
    Medication              string  `json:"medication"`
    Grand                   string  `json:"grand"`
    Pre_program             string  `json:"pre_program"`
    Pre_program_diy         string  `json:"pre_program_diy"`
    Comment                 string  `json:"comment"`
    Comment_diy             string  `json:"comment_diy"`

    Writer                  string  `json:"writer"`
    Assessment_date         string  `json:"assessment_date"`
    Assessment_time         string  `json:"assessment_time"`
    Assessment_timestamp    string  `json:"assessment_timestamp"`
    Chemotherapy_date       string  `json:"chemotherapy_date"`
    Chemotherapy_time       string  `json:"chemotherapy_time"`
    Chemotherapy_timestamp  string  `json:"chemotherapy_timestamp"`
}
type nurseTableRec struct {
    Action                  string  `json:"action"`
    Userid                  string  `json:"userid"`
    Cycle_seq               int     `json:"cycle_seq"`
}
type nurseTable struct{
    Nurse_seq               int     `json:"nurse_seq"`
    Time                    string  `json:"time"`
    Has                     int     `json:"has"`
}
type nurseTableReS struct {
    Status                  int     `json:"status"`
    N                   [15]nurseTable
}
type nurseInfoRec struct{
    Action                  string  `json:"action"`
    Userid                  string  `json:"userid"`
    Cycle_seq               int     `json:"cycle_seq"`
    Nurse_seq               int     `json:"nurse_seq"`
}
type nurseInfoRes struct{
    Status                  int     `json:"status"`
    Userid                  string  `json:"userid"`
    Cycle_seq               int     `json:"cycle_seq"`
    Nurse_seq               int     `json:"nurse_seq"`

    Nausea_assessment       string  `json:"nausea_assessment"` 
    Emesis_assessment       string  `json:"emesis_assessment"`
    Measure                 string  `json:"measure"`
    Comment                 string  `json:"comment"`
    Out_hospital            string  `json:"out_hospital"`

    Writer                  string  `json:"writer"`
    Assessment_date         string  `json:"nurse_assessment_date"`
    Assessment_time         string  `json:"nurse_assessment_time"`
    Assessment_timestamp    string  `json:"nurse_assessment_timestamp"`
}

type outHospitalRec struct{
    Action                  string  `json:"action"`
    Userid                  string  `json:"userid"`
}
type outHospitalRes struct{
    Status                  int     `json:"status"`
    Explain                 string  `json:"explain"`
    Time                    string  `json:"time"`
}
type followTable struct{
    Follow_seq              string `json:"follow_seq"`
    Time                    string  `json:"time"`
    Has                     int     `json:"has"`

}
type followTableRec struct{
    Action                  string  `json:"action"`
    Userid                  string  `json:"userid"`
}
type followTableRes struct{
    Status                  int     `json:"status"`
    N                     [15]followTable
}
type followContentRec struct{
    Action                  string  `json:"action"`
    Userid                  string  `json:"userid"`
    Follow_seq              string  `json:"follow_seq"`
}
type followContentRes struct{
    Status                  int     `json:"status"`
    
    Userid                  string  `json:"userid"`
    Follow_seq              int     `json:"follow_seq"`
    Hight_risk              string  `json:"hight_risk"`
    Emesis_grade            string  `json:"emesis_grade"`
    Nausea_grade            string  `json:"nausea_grade"`
    Out_content             string  `json:"out_content"`
    Out_content_diy         string  `json:"out_content_diy"`
    Follow_over             string  `json:"follow_over"`
    Satisfaction_1          string  `json:"satisfaction_1"`
    Satisfaction_2          string  `json:"satisfaction_2"`
    Satisfaction_3          string  `json:"satisfaction_3"`
    Satisfaction_4          string  `json:"satisfaction_4"`
    Satisfaction_5          string  `json:"satisfaction_5"`
    Satisfaction_total      string  `json:"satisfaction_total"`
    Writer                  string  `json:"writer"`
    Follow_follow_date      string  `json:"follow_follow_date"`
    Follow_follow_time      string  `json:"follow_follow_time"`
    Follow_follow_timestamp string  `json:"follow_follow_timestamp"`
}
type waitGo struct{
    Name                    string  `json:"name"`
    Userid                  string  `json:"userid"`
    Has                     int     `json:"has"`
}
type waitGoRec struct{
    Action                  string  `json:"action"`
}
type waitGoRes struct{
    Status                  int     `json:"status"`
    N                   [ 15] waitGo
}