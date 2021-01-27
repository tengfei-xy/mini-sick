package main

import (
	pnt "print"
)

func msgMain(msg []byte) []byte {
	if len(msg) != 0 {
		pnt.Json(string(msg))
	}
	msgtype := string(msg[11:22])

	switch msgtype {

	// 更新 患者基本信息
	case Act_Add_Sick:
		var si sickerInfo
		parseJSON(&msg, &si)
		return si.msgMain()

	// 用户登录
	case Act_User_Login:
		var ui userInfo
		parseJSON(&msg, &ui)
		return ui.msgMain()

	// 更新 患者信息的 风险评估
	case Act_Add_Risk:
		var ri riskInfo
		parseJSON(&msg, &ri)
		return ri.msgMain()

	// 更新 患者信息的 护理
	case Act_Add_Nurse:
		var ni nurseInfo
		parseJSON(&msg, &ni)
		return ni.msgMain()

	// 搜索 患者信息的 随访
	case Act_Add_Follow:
		var fi followInfo
		parseJSON(&msg, &fi)
		return fi.msgMain()

	// 搜索 患者
	case Act_Search_Sisk:
		var ss searchSicker
		parseJSON(&msg, &ss)
		return ss.msgMain()

	// 搜索 患者 详细信息
	case Act_Search_detail_Sick:
		var sds searchDeatilSick
		parseJSON(&msg, &sds)
		return sds.msgMain()

	// 搜索 患者 化疗周期
	case Act_Serch_Cycle:
		var ci cycleInfo
		parseJSON(&msg, &ci)
		return ci.msgMain()

	// 搜索 患者 风险评估
	case Act_Req_Risk:
		var rirc riskInfoRec
		parseJSON(&msg, &rirc)
		return rirc.msgMain()

	// 查询 患者 护理表
	case Act_Search_Nurse_Table:
		var ntrc nurseTableRec
		parseJSON(&msg, &ntrc)
		return ntrc.msgMain()

	// 搜索 患者 护理具体信息
	case Act_Req_Nurese:
		var nirc nurseInfoRec
		parseJSON(&msg, &nirc)
		return nirc.msgMain()

	// 搜索 出院
	case Act_Seach_Out_Hospital:
		var ohpc outHospitalRec
		parseJSON(&msg, &ohpc)
		return ohpc.msgMain()

	// 搜索 随访表格
	case Act_Search_Follow_Table:
		var ftrc followTableRec
		parseJSON(&msg, &ftrc)
		return ftrc.msgMain()

	// 搜索 随访表格具体内容
	case Act_Req_Follow:
		var fcrc followContentRec
		parseJSON(&msg, &fcrc)
		return fcrc.msgMain()

	// 搜索 待办
	case Act_Search_Wait:
		var wgrc waitGoRec
		parseJSON(&msg, &wgrc)
		return wgrc.msgMain()

	// 查询 高止吐风险
	case Act_Search_Height_Risk:
		var hrrc heightRiskRec
		parseJSON(&msg, &hrrc)
		return hrrc.msgMain()

	// 查询上次周期内容
	case Act_Req_Cylce_Last:
		var clrc cycleLastRec
		parseJSON(&msg, &clrc)
		return clrc.msgMain()

	// 今日护理
	case Act_Search_Today_Nurse:
		var tonrc toNurseRec
		parseJSON(&msg, &tonrc)
		return tonrc.msgMain()

	// 查询第一周期非药物因素
	case Act_Req_Last_Not_Medication:
		var lnmrc lastnotmedicationRec
		parseJSON(&msg, &lnmrc)
		return lnmrc.msgMain()

	// ID:5 搜索 患者填写的表
	case Act_Search_Sicker_Write_Info:
		var sswic seaSickerWriteInfoRec
		parseJSON(&msg, &sswic)
		return sswic.msgMain()

	// ID:5 提交 患者填写的表
	case Act_Submit_Sicker_Write_Info:
		var swrirs subSickerWriteInfoRec
		parseJSON(&msg, &swrirs)
		return swrirs.msgMain()

	// ID:5 查询 今日患者情况
	case Act_Search_Today_Sicker:
		var tosrc toSickerRec
		parseJSON(&msg, &tosrc)
		return tosrc.msgMain()

	// 查看明日护理和随访数量
	case Act_Cat_Nurse_Follow_Count:
		var catnfcrc catNFCRec
		parseJSON(&msg, &catnfcrc)
		return catnfcrc.msgMain()

	// 默认
	default:
		return nil

	}
	return nil
}
