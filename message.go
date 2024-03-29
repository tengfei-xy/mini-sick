package main

import (
	"env"
	pnt "print"
)

func msgMain(msg []byte) []byte {
	if len(msg) != 0 {
		pnt.Json(string(msg))
	}
	msgtype := string(msg[11:22])

	switch msgtype {

	// 更新 患者基本信息
	case env.ActAddSick:
		var si sickerInfo
		parseJSON(&msg, &si)
		return si.msgMain()

	// 用户登录
	case env.ActUserLogin:
		var ui userInfo
		parseJSON(&msg, &ui)
		return ui.msgMain()

	// 更新 患者信息的 风险评估
	case env.ActAddRisk:
		var ri riskInfo
		parseJSON(&msg, &ri)
		return ri.msgMain()

	// 更新 患者信息的 护理
	case env.ActAddNurse:
		var ni nurseInfo
		parseJSON(&msg, &ni)
		return ni.msgMain()

	// 搜索 患者信息的 随访
	case env.ActAddFollow:
		var fi followInfo
		parseJSON(&msg, &fi)
		return fi.msgMain()

	// 搜索 患者
	case env.ActSearchSisk:
		var ss searchSicker
		parseJSON(&msg, &ss)
		return ss.msgMain()

	// 搜索 患者 详细信息
	case env.ActSearchdetailSick:
		var sds searchDeatilSick
		parseJSON(&msg, &sds)
		return sds.msgMain()

	// 搜索 患者 化疗周期
	case env.ActSerchCycle:
		var ci cycleInfo
		parseJSON(&msg, &ci)
		return ci.msgMain()

	// 搜索 患者 风险评估
	case env.ActReqRisk:
		var rirc riskInfoRec
		parseJSON(&msg, &rirc)
		return rirc.msgMain()

	// 查询 患者 护理表
	case env.ActSearchNurseTable:
		var ntrc nurseTableRec
		parseJSON(&msg, &ntrc)
		return ntrc.msgMain()

	// 搜索 患者 护理具体信息
	case env.ActReqNurese:
		var nirc nurseInfoRec
		parseJSON(&msg, &nirc)
		return nirc.msgMain()

	// 搜索 出院
	case env.ActSeachOutHospital:
		var ohpc outHospitalRec
		parseJSON(&msg, &ohpc)
		return ohpc.msgMain()

	// 搜索 随访表格
	case env.ActSearchFollowTable:
		var ftrc followTableRec
		parseJSON(&msg, &ftrc)
		return ftrc.msgMain()

	// 搜索 随访表格具体内容
	case env.ActReqFollow:
		var fcrc followContentRec
		parseJSON(&msg, &fcrc)
		return fcrc.msgMain()

	// 搜索 待办
	case env.ActSearchWait:
		var wgrc waitGoRec
		parseJSON(&msg, &wgrc)
		return wgrc.msgMain()

	// 查询 高止吐风险
	case env.ActSearchHeightRisk:
		var hrrc heightRiskRec
		parseJSON(&msg, &hrrc)
		return hrrc.msgMain()

	// 查询上次周期内容
	case env.ActReqCylceLast:
		var clrc cycleLastRec
		parseJSON(&msg, &clrc)
		return clrc.msgMain()

	// 今日护理
	case env.ActSearchTodayNurse:
		var tonrc toNurseRec
		parseJSON(&msg, &tonrc)
		return tonrc.msgMain()

	// 查询第一周期非药物因素
	case env.ActReqLastNotMedication:
		var lnmrc lastnotmedicationRec
		parseJSON(&msg, &lnmrc)
		return lnmrc.msgMain()

	// ID:5 搜索 患者填写的表
	case env.ActSearchSickerWriteInfo:
		var sswic seaSickerWriteInfoRec
		parseJSON(&msg, &sswic)
		return sswic.msgMain()

	// ID:5 提交 患者填写的表
	case env.ActSubmitSickerWriteInfo:
		var swrirs subSickerWriteInfoRec
		parseJSON(&msg, &swrirs)
		return swrirs.msgMain()

	// ID:5 查询 今日患者情况
	case env.ActSearchTodaySicker:
		var tosrc toSickerRec
		parseJSON(&msg, &tosrc)
		return tosrc.msgMain()

	// 查看明日护理和随访数量
	case env.ActCatNurseFollowCount:
		var catnfcrc catNFCRec
		parseJSON(&msg, &catnfcrc)
		return catnfcrc.msgMain()

	// 数据下载提交
	case env.ActDownloadSubmit:
		var dsubrc downloadSubmitRec
		parseJSON(&msg, &dsubrc)
		return dsubrc.msgMain()

	// 数据下载搜索
	case env.ActDownloadSearch:
		var dsearc downloadSearchRec
		parseJSON(&msg, &dsearc)
		return dsearc.msgMain()

	// 数据下载-重试
	case env.ActDownloadTry:
		var dtryrc downloadTryRec
		parseJSON(&msg, &dtryrc)
		return dtryrc.msgMain()

	// 默认
	default:
		return nil
	}
	return nil
}
