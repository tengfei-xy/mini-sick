package main

import "strings"

// 护理表 是否出院
func convBaseIf(i string) string {
	switch i {
	case "1":
		i = `是`
	case "2":
		i = `否`
	default:
		i = `否`
	}
	return i
}

// 患者是否知情
func convSickerKnow(i string) string {

	switch i {
	case "1":
		i = `知情`
	case "2":
		i = `不知情`
	default:
		i = ""
	}
	return i
}

// 风险评估 药物因素
func convRiskMedication(i string) string {
	switch i {
	case "1":
		i = `高致吐风险方案（呕吐发生率>90%）`
	case "2":
		i = `中致吐风险方案（呕吐发生率30%-90%）`
	case "3":
		i = `低致吐风险方案（呕吐发生率10%-30%）`
	case "4":
		i = `轻微致吐风险方案（呕吐发生率<10%）`
	default:
		i = ""
	}
	return i
}

// 风险评估 非药物因素
func convRiskNotMedication(i string) string {
	var res string
	for _, j := range strings.Split(i, ",") {
		switch j {
		case "1":
			res += "1. 女性\n"
		case "2":
			res += "2. 小于55岁年轻患者\n"
		case "3":
			res += "3. 低酒精摄入（每周<5<100g/次）\n"
		case "4":
			res += "4. 既往有晕动证\n"
		case "5":
			res += "5. 既往化疗出现呕吐\n"
		case "6":
			res += "6. 怀孕期间有妊娠反应\n"
		case "7":
			res += "7. 有焦虑等情绪因素\n"
		default:
			res += ""
		}
	}
	return res
}

// 风险评估 风险等级
func convRiskGrand(i string) string {
	switch i {
	case "1":
		i = `高`
	case "2":
		i = `中`
	case "3":
		i = `低`
	default:
		i = ""
	}
	return i
}

// 风险评估 预处理方案
func convRiskPreProgram(i string) string {
	var res string
	for _, j := range strings.Split(i, ",") {
		switch j {
		case "1":
			res += "1. 地塞米松 + NK-1拮抗剂（阿瑞匹坦） + 5-HT3拮抗剂（司琼类）\n"
		case "2":
			res += "2. NK-1拮抗剂（阿瑞匹坦） + 5-HT3拮抗剂（司琼类）\n"
		case "3":
			res += "3. 地塞米松 + 5-HT3拮抗剂（司琼类）\n"
		case "4":
			res += "4. 地塞米松\n"
		case "5":
			res += "5. 5-HT3拮抗剂（司琼类）\n"
		case "6":
			res += "6. 苯二氮卓类\n"
		case "7":
			res += "7. 奥氮平\n"
		case "8":
			res += "8. 其他\n"
		default:
			res += ". "
		}
	}
	return res
}

// 风险评估 备注
func convRiskComment(i string) string {
	var res string
	for _, j := range strings.Split(i, ",") {
		switch j {
		case "1":
			res += "1. 患者放弃使用NK-1"
		case "2":
			res += "2. 患者上一周期化疗未产生明显CINV"
		case "3":
			res += "3. 其他"
		default:
			res += ""
		}
	}
	return res
}

// 风险评估 需要住院护理
func convRiskNeedNurse(i string) string {
	switch i {
	case "1":
		i = `需要`
	case "2":
		i = `不需要`
	default:
		i = ""
	}
	return i
}

// 护理表 恶心分级评估
func convNurseNauseaAssessment(i string) string {
	switch i {
	case "1":
		i = `0级：无症状`
	case "2":
		i = `1级：食欲降低，不伴进食习惯改变`
	case "3":
		i = `2级：经口进食减少，无明显体重下降、无脱水或营养不良`
	case "4":
		i = `3级：经口摄入能量和水份不足，需鼻饲、静脉营养或住院治疗`
	default:
		i = ""
	}
	return i
}

// 护理表 呕吐分级评估
func convNurseEmesisAssessment(i string) string {
	switch i {
	case "1":
		i = `0级：无症状`
	case "2":
		i = `1级：24小时内发作1~2次（间隔5分钟)`
	case "3":
		i = `2级：24小时内发作3~5次（间隔5分钟)`
	case "4":
		i = `3级：24小时内发作6次或以上（间隔5分钟）；鼻饲、全肠外营养或住院治疗`
	case "5":
		i = `4级：危及生命，需要紧急治疗`
	case "6":
		i = `5级：死亡`
	default:
		i = ""
	}
	return i
}

// 护理表 护理措施
func convNurseMeasure(i string) string {
	var res string
	for _, j := range strings.Split(i, ",") {
		switch j {
		case "1":
			res += "1. 安慰患者、解释病情\n"
		case "2":
			res += "2. 创造良好的治疗环境\n"
		case "3":
			res += "3. 饮食指导，口服高热量饮料\n"
		case "4":
			res += "4. 指导有氧运动（散步、慢跑、快走等）\n"
		case "5":
			res += "5. 分散注意力\n"
		case "6":
			res += "6. 音乐疗法\n"
		case "7":
			res += "7. 穴位按摩\n"
		case "8":
			res += "8. 遵医嘱静脉营养\n"
		}
	}
	return res
}

// 护理表 是否出院
func convNurseOutHospital(i string) string {
	switch i {
	case "1":
		i = `已出院`
	case "2":
		i = `未出院`
	}
	return i
}

// 随访表
func convFollowHightRisk(i string) string {
	return convBaseIf(i)
}

// 随访表 恶心分级评估
func convFollowEmesisGrade(i string) string {
	switch i {
	case "1":
		i = `0级：无症状`
	case "2":
		i = `1级：食欲降低，不伴进食习惯改变`
	case "3":
		i = `2级：经口进食减少，无明显体重下降、无脱水或营养不良`
	case "4":
		i = `3级：经口摄入能量和水份不足，需鼻饲、静脉营养或住院治疗`
	default:
		i = ""
	}
	return i
}

// 随访表 呕吐分级评估
func convFollowNauseaGrade(i string) string {
	switch i {
	case "1":
		i = `0级：无症状`
	case "2":
		i = `1级：24小时内发作1~2次（间隔5分钟)`
	case "3":
		i = `2级：24小时内发作3~5次（间隔5分钟)`
	case "4":
		i = `3级：24小时内发作6次或以上（间隔5分钟）；鼻饲、全肠外营养或住院治疗`
	case "5":
		i = `4级：危及生命，需要紧急治疗`
	case "6":
		i = `5级：死亡`
	default:
		i = ""
	}
	return i
}

// 随访表 出院后专科指导内容
func convFollowOutContent(i string) string {
	var res string
	for _, j := range strings.Split(i, ",") {
		switch j {
		case "1":
			res += "1. 服药\n"
		case "2":
			res += "2. 饮食\n"
		case "3":
			res += "3. 心理\n"
		case "4":
			res += "4. 康复锻炼\n"
		case "5":
			res += "5. 填写《患者关爱手册》\n"
		case "6":
			res += "6. 其他\n"
		}
	}
	return res
}

// 随访表 随访结束
func convFollowFollowOver(i string) string {
	switch i {
	case "1":
		i = `随访结束`
	case "2":
		i = `随访未结束`
	}
	return i
}

// 随访表 满意调查表
func convFollowSatisfactionTable(i string) string {
	switch i {
	case "1":
		i = `是`
	case "2":
		i = `否`
	default:
		i = ``
	}
	return i
}

// 随访表 满意调查表 总分
func convFollowSatisfactionTotal(i string) string {
	if i == "0" {
		return ""
	}
	return i
}
