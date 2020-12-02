var urlq = {};
var glo_nurse_seq = 1
var glo_follow_seq = 1
var sicker={}

$('docuemnt').ready(function(){

    let url = decodeURI(window.location.search); 
    if (url.indexOf("?") != 1) {
        url = url.substr(1)
        var queryArr = url.split("&");

        queryArr.forEach(function (item) {
            var key = item.split("=")[0];
            var value = item.split("=")[1];
            if (value==""){widows.location.href($.page.login())}
            urlq[key] = decodeURIComponent(value);
        });

    }

    // 新页面，不用从服务器参数数据
    if (parseInt(urlq['init']) == 0){
        let date = $.time.current_date
        let time = $.time.current_time
        $('#risk_assessment_date').text(date)
        $('#risk_assessment_time').text(time)
        $('#risk_chemotherapy_date').text(date)
        $('#risk_chemotherapy_time').text(time)
    }

    // 填充患者信息
    let form ={}
    form.action = $.act.Search_detail_Sick()
    form.userid = urlq['userid']
    $.post_send.ajax(form,function(res){
        console.log("患者信息",res)

    localStorage.setItem("sicker",JSON.stringify(res))
    sicker = res

    // 填充患者信息 到主菜单
    let sickerinfo = $(
        "<p>姓名：" + sicker.name + "</p>" +
        "<p>性别：" + sicker.gender + "</p>" +
        "<p>年龄：" + sicker.age + "</p>" +
        "<p>电话：" + sicker.telphone + "</p>" +
        "<p>住院号：" + sicker.hospital_number + "</p>" +
        "<p>就诊号：" + sicker.attandance_number + "</p>" +
        "<p>诊断：" + sicker.disease + "</p>")
        $('#sicker_info').html(sickerinfo)
            
    })

    //sicker=JSON.parse(localStorage.getItem("sicker"))


    if (urlq['way']=="wait"){
        $('#risk_title').css("display","none")
        $('#nurse_title').css("display","none")

        return
    }
})
$('#cat_sicker').on("click",function(){
    $('#show_sicker').fadeIn(200);
})
$('#close_sicker').on("click",function(){
    $('#show_sicker').fadeOut(200);
})
$('#toindex').on("click",function(){
    window.location.href = $.page.index()
})
$('#toback').on("click",function(){
    window.location.href = $.page.cycle() + "?userid="+ urlq["userid"]
    //window.location.href = localStorage.setItem("back")
})

// 风险评估 标题点击
$('#risk_title').on("click",function(){
    if ($('#risk_content').css("display") == "none"){
        $('#risk_content').css("display","block") 
    } else{
        $('#risk_content').css("display","none")
        return
    }

    // 获取上次记录
    if (parseInt(urlq['init']) == 0 && parseInt(urlq['cycleseq'])>1){
        console.log("查询上一周期记录")
        let form={}
        form.action = $.act.Req_Cylce_Last()
        form.updated = parseInt(urlq['init'])
        form.cycle_seq = parseInt(urlq['cycleseq'])
        form.userid = urlq['userid']
        $.post_send.ajax(form,function(res){
            console.log(res)
            if (res.status == 1){
                $('#risk_last').css("display","block")
            }else{
                $.toast.text(res.explain)
            }
            res.last_nurse_emesis == "" ? last_nurse_emesis="无" : last_nurse_emesis = parseInt(res.last_nurse_emesis)-1 + "级"
            res.last_nurse_nausea == "" ? last_nurse_nausea="无" : last_nurse_nausea = parseInt(res.last_nurse_nausea)-1 + "级"
            res.last_follow_emesis == "" ? last_follow_emesis="无" : last_follow_emesis = parseInt(res.last_follow_emesis)-1 + "级"
            res.last_follow_nausea == "" ? last_follow_nausea="无" : last_follow_nausea = parseInt(res.last_follow_nausea)-1 + "级"
            switch (res.last_risk_grand){
            case "1":
                gui_rand="高风险"
                break;
            case "2":
                gui_rand="中风险"
                break;
            case "3":
                gui_rand="低风险"
                break;
            default:
                gui_rand= "无"
                break;
            }

            $('#last_risk_grand').text(gui_rand)
            $('#last_nurse_emesis').text(last_nurse_emesis)
            $('#last_nurse_nausea').text(last_nurse_nausea)
            $('#last_follow_emesis').text(last_follow_emesis)
            $('#last_follow_nausea').text(last_follow_nausea)
        })
    }
    // 如果是新界面
    if (parseInt(urlq['init']) == 0){return}
    
    let form={}
    form.action = $.act.Req_Risk()
    form.updated = parseInt(urlq['init'])
    form.cycle_seq = parseInt(urlq['cycleseq'])
    form.userid = urlq['userid']
    $.post_send.ajax(form,function(res){
        console.log(res)
        if (res.status != 1) {$.toast.text(res.explain)}

        // 风险 重设
        $(':checkbox[name="risk_not_medication"]').removeAttr('checked');
        $(':radio[name="risk_medication"]').removeAttr('checked');
        $(':radio[name="risk_grand"]').removeAttr('checked');
        $(':checkbox[name="risk_pre_program"]').removeAttr('checked');
        $(':checkbox[name="risk_comment"]').removeAttr('checked');
        $(':radio[name="risk_need_nurse"]').removeAttr('checked');

        // 上一周期更新
        if (parseInt(urlq['cycleseq'])>1){
            $('#risk_last').css("display","block")
            res.last_nurse_emesis == "" ? last_nurse_emesis="无" : last_nurse_emesis = parseInt(res.last_nurse_emesis)-1 + "级"
            res.last_nurse_nausea == "" ? last_nurse_nausea="无" : last_nurse_nausea = parseInt(res.last_nurse_nausea)-1 + "级"
            res.last_follow_emesis == "" ? last_follow_emesis="无" : last_follow_emesis = parseInt(res.last_follow_emesis)-1 + "级"
            res.last_follow_nausea == "" ? last_follow_nausea="无" : last_follow_nausea = parseInt(res.last_follow_nausea)-1 + "级"
            switch (res.last_risk_grand){
            case "1":
                gui_rand="高风险"
                break;
            case "2":
                gui_rand="中风险"
                break;
            case "3":
                gui_rand="低风险"
                break;
            default:
                gui_rand= "无"
                break;
            }

            $('#last_risk_grand').text(gui_rand)
            $('#last_nurse_emesis').text(last_nurse_emesis)
            $('#last_nurse_nausea').text(last_nurse_nausea)
            $('#last_follow_emesis').text(last_follow_emesis)
            $('#last_follow_nausea').text(last_follow_nausea)
        }
        
        // 风险赋值
        $('#risk_assessment_date').text(res.assessment_date)
        $('#risk_assessment_time').text(res.assessment_time)
        $('#risk_chemotherapy_date').text(res.chemotherapy_date)
        $('#risk_chemotherapy_time').text(res.chemotherapy_time)

        $('#risk_program').val(res.program)
        $.auto.choose(':checkbox[name="risk_not_medication"]',res.not_medication)
        $.auto.choose(':radio[name="risk_medication"]',res.medication)
        $.auto.choose(':radio[name="risk_grand"]',res.grand)

        $.auto.choose(':checkbox[name="risk_pre_program"]',res.pre_program)
        $.auto.choose_diy('#risk_pre_program_diy',res.pre_program,res.pre_program_diy)

        $.auto.choose(':checkbox[name="risk_comment"]',res.comment)
        $.auto.choose_diy('#risk_comment_diy',res.comment,res.comment_diy)
        $.auto.choose(':radio[name="risk_need_nurse"]',res.need_nurse)

        // 只有随访最后一条是出院时,显示 添加护理记录 按钮
        if (res.need_nurse !="1"){
            $('#nurse_add').css("display","block")
        }
        if (res.need_nurse =="2"){
            $('#nurse_title').css("display","none")
        }
    })
    
    
})
// 护理 标题点击
$('#nurse_title').on("click",function(){
    ($('#nurse_content').css("display") == "none") ? $('#nurse_content').css("display","block") : $('#nurse_content').css("display","none")

    // 如果是新界面
    if (parseInt(urlq['init']) == 0){return }
    let form={}
    form.action = $.act.Search_Nurse_Table()
    form.cycle_seq = parseInt(urlq['cycleseq'])
    form.userid = urlq['userid']

    $.post_send.ajax(form,function(res){
        console.log(res)

        if (res.status == 0){
            $.toast.text(res.status)
            return
        }
        $("#nurse_table_content").empty()
        for (let i =0;i<res.N.length;i++){
            if(res.N[i].has==1){
                let a=$.append.table_small_cell(res.N[i].nurse_seq)
                let b=$.append.table_small_cell(res.N[i].time)

                let div = $('<div id="seq_nurse_det" data-nurseseq="'+res.N[i].nurse_seq  +'" class="weui-flex"></div>')
                div.append(a,b)
                $("#nurse_table_content").append(div)
                glo_nurse_seq = i+2
            }
        }
    })
})
// 随访登记 标题点击
$('#follow_title').on("click",function(){
    // 输出出院时间
    let form={}
    form.action = $.act.Seach_Out_Hospital()
    form.cycle_seq=parseInt(urlq['cycleseq'])
    form.userid=urlq['userid']
    $.post_send.ajax(form,function(res){
        console.log("post_res",res)

        if (res.status !=1 || res.time==""){
            $.toast.text(res.explain)
            return
        }
        console.log("出院时间:",res.time)
        var out_hospital_time = res.time
        let show = $('#follow_content').css("display")
        if (show == "none"){
            $('#follow_content').css("display","block")
        }else{
            $('#follow_content').css("display","none")
        }
    
        $('#out_hospital_time').text(out_hospital_time)
    
        // 查询随访信息
        form.action = $.act.Search_Follow_Table()
        $.post_send.ajax(form,function(res){
            console.log(res)
            if (res.status == 0){
                return
            }
            $("#follow_table_content").empty()
            for (let i =0;i<res.N.length;i++){
                if(res.N[0].has==0){break}
                if(res.N[i].has==1){
                    let a=$.append.table_small_cell(res.N[i].follow_seq)
                    let b=$.append.table_small_cell(res.N[i].time)
    
                    let div = $('<div id="seq_follow_det" data-followseq="'+res.N[i].follow_seq  +'" class="weui-flex"></div>')
                    div.append(a,b)
                    $("#follow_table_content").append(div)
                    glo_follow_seq = i+2
                }
            }
            
        })
        // 随访结束关闭添加护理记录按钮
        if (sicker.follow_over == "1"){
            $("#follow_add").css("display","none")
        }

    })



})


// 风险评估 提交
$('#risk_submit').on("click",function(){

    let form={}
    form.action=$.act.Add_Risk()
    form.updated = parseInt(urlq['init'])

    form.assessment_date=$('#risk_assessment_date').text()
    form.assessment_time=$('#risk_assessment_time').text()
    form.assessment_timestamp = $.time.transTimestamp(form.assessment_date,form.assessment_time)
    form.chemotherapy_date=$('#risk_chemotherapy_date').text()
    form.chemotherapy_time=$('#risk_chemotherapy_time').text()
    form.chemotherapy_timestamp = $.time.transTimestamp(form.chemotherapy_date,form.chemotherapy_time)
    form.comment=$.auto.array('input[name="risk_comment"]:checked')
    form.comment_diy=$('#risk_comment_diy').val()
    form.grand=$('input[name="risk_grand"]:checked').val()
    form.medication=$.auto.array('input[name="risk_medication"]:checked')
    form.need_nurse=$.auto.array('input[name="risk_need_nurse"]:checked')
    form.not_medication=$.auto.array('input[name="risk_not_medication"]:checked')
    form.pre_program=$.auto.array('input[name="risk_pre_program"]:checked')
    form.pre_program_diy=$('#risk_pre_program_diy').val()
    form.program=$('#risk_program').val()
    form.name=sicker.name


    form.cycle_seq=parseInt(seq=urlq['cycleseq'])
    form.userid=urlq['userid']
    form.writer=localStorage.getItem("writer")
    $.post_send.ajax(form,function(res){
        $.toast.text(res.explain)
        if (res.status=="1"){
            if (form.need_nurse=="2"){
                $('#nurse_title').css("display","none")
            }else if (form.need_nurse=="1"){
                $('#nurse_title').css("display","block")
            }
            urlq['init']="1"
        }
    })
})
// 护理 提交
$('#nurse_submit').on("click",function(){
    console.log(glo_nurse_seq)
    let form={}
    form.action=$.act.Add_Nurse()
    form.userid = urlq['userid']
    form.writer = localStorage.getItem("writer")
    form.cycle_seq = parseInt(urlq['cycleseq'])
    form.nurse_seq=parseInt($('#nurse_seq').text())

    form.nurse_assessment_date = $('#nurse_assessment_date').text()
    form.nurse_assessment_time = $('#nurse_assessment_time').text()
    form.nurse_assessment_timestamp = $.time.transTimestamp(form.nurse_assessment_date, form.nurse_assessment_time)
    form.nausea_assessment = $.auto.array('input[name="nurse_nausea_assessment"]:checked')
    form.emesis_assessment = $.auto.array('input[name="nurse_emesis_assessment"]:checked')
    form.measure = $.auto.array('input[name="nurse_measure"]:checked')
    form.comment = $('#nurse_comment').val()
    form.out_hospital = $.auto.array('input[name="nurse_out_hospital"]:checked')

    $.post_send.ajax(form,function(res){
        console.log(res)
        $.toast.text(res.explain)
        if (res.status ==1){
            $('#nurse_submit').css("display","none")
            $('#nurse_add').css("display","block")
            //出院
            if (res.out_hospital=="1"){
                $('#nurse_submit').css("display","none")
            }
        }
        console.log("添加提交的护理表单")
        let a=$.append.table_small_cell(form.nurse_seq)
        let b=$.append.table_small_cell(form.nurse_assessment_date + " " +form.nurse_assessment_time)

        let div = $('<div id="seq_nurse_det" data-nurseseq="'+$('#nurse_seq').text()  +'" class="weui-flex"></div>')
        div.append(a,b)
        $("#nurse_table_content").append(div)
        glo_nurse_seq +=1
    })

})
// 随访 提交
$('#follow_submit').on("click",function(){
    let form = {}
    if ($('#add_satisfaction_table').css("display")=="none"){
        $('input[name="satisfaction_1"]').eq[0].attr("checked",false)
        $('input[name="satisfaction_1"]').eq[1].attr("checked",false)
        $('input[name="satisfaction_2"]').eq[0].attr("checked",false)
        $('input[name="satisfaction_2"]').eq[1].attr("checked",false)
        $('input[name="satisfaction_3"]').eq[0].attr("checked",false)
        $('input[name="satisfaction_3"]').eq[1].attr("checked",false)
        $('input[name="satisfaction_4"]').eq[0].attr("checked",false)
        $('input[name="satisfaction_4"]').eq[1].attr("checked",false)
        $('input[name="satisfaction_5"]').eq[0].attr("checked",false)
        $('input[name="satisfaction_5"]').eq[1].attr("checked",false)
        form.satisfaction_total=""
    }


    form.action = $.act.Add_Follow()
    form.userid=urlq['userid']
    form.cycle_seq=parseInt(urlq['cycleseq'])
    form.follow_seq = parseInt($('#follow_seq').text())
    form.hight_risk = $.auto.array('input[name="follow_hight_risk"]:checked')
    form.emesis_grade = $.auto.array('input[name="follow_emesis_grade"]:checked')
    form.nausea_grade = $.auto.array('input[name="follow_nausea_grade"]:checked')
    form.out_content = $.auto.array('input[name="follow_out_content"]:checked')
    form.out_content_diy = $('#follow_out_content_diy').val()
    form.follow_over = $.auto.array('input[name="follow_over"]:checked')

    form.satisfaction_1=$.auto.array('input[name="satisfaction_1"]:checked')
    form.satisfaction_2=$.auto.array('input[name="satisfaction_2"]:checked')
    form.satisfaction_3=$.auto.array('input[name="satisfaction_3"]:checked')
    form.satisfaction_4=$.auto.array('input[name="satisfaction_4"]:checked')
    form.satisfaction_5=$.auto.array('input[name="satisfaction_5"]:checked')
    let total = (form.satisfaction_1+form.satisfaction_2+form.satisfaction_3+form.satisfaction_4+form.satisfaction_5)
    form.satisfaction_total = String((total.split("1").length-1)*20)
    form.follow_follow_date = $('#follow_follow_date').text()
    form.follow_follow_time = $('#follow_follow_time').text()
    form.follow_follow_timestamp = $.time.transTimestamp(form.follow_follow_date,form.follow_follow_time)
    form.writer=localStorage.getItem("writer")

    $.post_send.ajax(form,function(res){
        console.log(res)
        $.toast.text(res.explain)
        if (res.status == 1){
            console.log("添加提交的随访表单")
            let a=$.append.table_small_cell(form.follow_seq)
            let b=$.append.table_small_cell(form.follow_follow_date + " " + form.follow_follow_time)
        
            let div = $('<div id="seq_follow_det" data-followseq="'+$('#follow_seq').text()  +'" class="weui-flex"></div>')
            div.append(a,b)
            $("#follow_table_content").append(div)
            $('#follow_submit').css("display","none")

            // 如果随访结束
            if (form.follow_over=="1"){
                sicker.follow_over="1"
                localStorage.setItem("sicker",sicker)
                $('#follow_add').css("display","none")
            }else{
                $('#follow_add').css("display","block")

            }
        
            if ($('#satisfaction_table')=="none"){
            $('#add_satisfaction_table').css("display","block")
            }
            glo_follow_seq +=1
        }
    })
})


// 护理 添加
$('#nurse_add').on("click",function(){
    console.log(glo_nurse_seq)

    // 隐藏 添加护理记录 按钮
    $('#nurse_add').css("display","none")

    // 显示 具体护理单
    $('#nurese_table_body').css("display","block")

    $('#nurse_seq').text(glo_nurse_seq)

    let date = $.time.current_date
    let time = $.time.current_time
    $('#nurse_assessment_date').text(date)
    $('#nurse_assessment_time').text(time)
    $('#nurse_submit').css("display","block")
})
// 随访 添加
$('#follow_add').on("click",function(){

    // 隐藏 添加随访记录 按钮
    $('#follow_add').css("display","none")

    // 查询是否是高风险
    let form ={}
    form.action=$.act.Search_Height_Risk()
    form.userid=urlq['userid']
    form.cycle_seq = parseInt(urlq['cycleseq'])
    $.post_send.ajax(form,function(res){
        console.log(res)
        if (res.status == 1){
            console.log("风险等级",res.height)
            $(':radio[name="follow_hight_risk"]').removeAttr('checked');
            
            if (res.height==1){
                choose=0
            }else{
                choose=1

            }
            $(':radio[name="follow_hight_risk"]').eq(choose).attr("checked",true)
        } else{
            $.toast.text(res.explain)
        }
        
    })

    let date = $.time.current_date
    let time = $.time.current_time
    $('#follow_follow_date').text(date)
    $('#follow_follow_time').text(time)

    $('#follow_seq').text(glo_follow_seq)
    $('#add_satisfaction_table').text("添加满意度调查表（当前状态：不上传数据）")
    // 显示 具体护理单
    $('#follow_table_body').css("display","block")
    $('#add_satisfaction_table').css("display","block")
    $('#satisfaction_table').css("display","none")
    $('#follow_submit').css("display","block")
    $('#satisfaction_total_score').css("display","none")
})


// 护理 已更新 查看历史
$('#nurse_table_content').on("click",'div#seq_nurse_det',function(){
    

    let form={}
    form.action = $.act.Req_Nurese()
    form.userid = urlq['userid']
    form.writer = localStorage.getItem("writer")
    form.cycle_seq = parseInt(urlq['cycleseq'])
    form.nurse_seq = $(this).data("nurseseq")

    $.post_send.ajax(form,function(res){

        if (res.status ==0){
            $.toast.text(res.explain)
            return
        }
        // 护理 重设
        $(':radio[name="nurse_nausea_assessment"]').removeAttr('checked');
        $(':radio[name="nurse_emesis_assessment"]').removeAttr('checked');
        $(':checkbox[name="nurse_measure"]').removeAttr('checked');
        $(':radio[name="nurse_out_hospital"]').removeAttr('checked');


        // 护理 赋值
        $('#nurese_table_body').css("display","block")
        console.log(res)
        $('#nurse_seq').text(res.nurse_seq)
        $('#nurse_assessment_date').text(res.nurse_assessment_date)
        $('#nurse_assessment_time').text(res.nurse_assessment_time)
        $.auto.choose(':radio[name="nurse_nausea_assessment"]',res.nausea_assessment)
        $.auto.choose(':radio[name="nurse_emesis_assessment"]',res.emesis_assessment)
        $.auto.choose(':checkbox[name="nurse_measure"]',res.measure)
        $('#nurse_comment').val(res.comment)
        $.auto.choose(':radio[name="nurse_out_hospital"]',res.out_hospital)
        $('#nurse_submit').css("display","none")
    })
})
// 随访 已更新 查看历史
$('#follow_table_content').on("click",'div#seq_follow_det',function(){
    let follow_seq=$(this).data("followseq")
    let form={}
    form.action = $.act.Req_Follow()
    form.userid=urlq['userid']
    form.cycle_seq=parseInt(urlq['cycleseq'])

    form.follow_seq = parseInt(follow_seq)
    
    $.post_send.ajax(form,function(res){
        console.log(res)
        if (res.status == 0 ){
            $.toast.text(res.explain)
            return
        }
        // 随访 重设
        $(':radio[name="follow_hight_risk"]').removeAttr('checked');
        $(':radio[name="follow_emesis_grade"]').removeAttr('checked');
        $(':radio[name="follow_nausea_grade"]').removeAttr('checked');
        $(':checkbox[name="follow_out_content"]').removeAttr('checked');
        $(':radio[name="follow_over"]').removeAttr('checked');
        $(':radio[name="satisfaction_1"]').removeAttr('checked');
        $(':radio[name="satisfaction_2"]').removeAttr('checked');
        $(':radio[name="satisfaction_3"]').removeAttr('checked');
        $(':radio[name="satisfaction_4"]').removeAttr('checked');
        $(':radio[name="satisfaction_5"]').removeAttr('checked');
        // 随访 赋值
        $("#follow_table_body").css("display","block")
        $("#follow_seq").text(res.follow_seq)
        $('#follow_follow_date').text(res.follow_follow_date)
        $('#follow_follow_time').text(res.follow_follow_time)
        $.auto.choose(':radio[name="follow_hight_risk"]',res.hight_risk)
        $.auto.choose(':radio[name="follow_emesis_grade"]',res.emesis_grade)
        $.auto.choose(':radio[name="follow_nausea_grade"]',res.nausea_grade)
        $.auto.choose(':checkbox[name="follow_out_content"]',res.out_content)
        $('#follow_out_content_diy').val(res.out_content_diy)

        $.auto.choose(':radio[name="follow_over"]',res.follow_over)
        $.auto.choose(':radio[name="satisfaction_1"]',res.satisfaction_1)
        $.auto.choose(':radio[name="satisfaction_2"]',res.satisfaction_2)
        $.auto.choose(':radio[name="satisfaction_3"]',res.satisfaction_3)
        $.auto.choose(':radio[name="satisfaction_4"]',res.satisfaction_4)
        $.auto.choose(':radio[name="satisfaction_5"]',res.satisfaction_5)
        if (res.satisfaction_total!="0" || res.satisfaction_total!=""){
            $('#satisfaction_table').css("display","block")
            $('#satisfaction_total_score').css("display","block")
            $('#add_satisfaction_table').text("关闭满意度调查表（当前状态：上传数据）")
            $('#satisfaction_score').text(res.satisfaction_total)
        }
        $('#follow_submit').css("display","none")
        
    })
})







// 时间选择器 风险评估 评估日期
$('#risk_assessment_date').on('click', function () {
    $.time.wxpickerdate(this,'风险评估日期')
});
// 时间选择器 风险评估 评估时间
$('#risk_assessment_time').on('click', function () {
    $.time.wxpickertime(this,'风险评估时间')
});
// 时间选择器 风险评估 化疗日期
$('#risk_chemotherapy_date').on('click', function () {                        
    $.time.wxpickerdate(this,'风险化疗日期')
});
// 时间选择器 风险评估 化疗时间
$('#risk_chemotherapy_time').on('click', function () {
    $.time.wxpickertime(this,'风险化疗时间')
});       
// 时间选择器 护理评估 日期
$('#nurse_chemotherapy_date').on('click', function () {                        
    $.time.wxpickerdate(this,'护理日期')
});
// 时间选择器 护理评估 时间
$('#nurse_chemotherapy_time').on('click', function () {
    $.time.wxpickertime(this,'护理时间')
});       
// 时间选择器 随访 日期
$('#follow_follow_date').on('click', function () {                        
    $.time.wxpickerdate(this,'随访日期')
});
// 时间选择器 随访 时间
$('#follow_follow_time').on('click', function () {
    $.time.wxpickertime(this,'随访时间')
});       

// 查看风险等级参照图
$('#risk_grand_show').on("click",function(){                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                
    $('#risk_grand_pic').css("display","block")
})
// 查看风险等级参照图
$('#risk_grand_pic').on("click",function(){
    $('#risk_grand_pic').css("display","none")
})
$('#add_satisfaction_table').on("click",function(){
    let status = $('#satisfaction_table').css("display")
    if (status == "none"){
        $('#satisfaction_table').css("display","block")
        $('#add_satisfaction_table').text("关闭满意度调查表（当前状态：上传数据）")
    }else{
        $('#satisfaction_table').css("display","none")
        $('#add_satisfaction_table').text("添加满意度调查表（当前状态：不上传数据）")

    }
})

$('input[type=radio][name="risk_medication"]').change(function () {
    let c =$('input[name="risk_not_medication"]:checked').length

    let medication=$('input[type=radio][name="risk_medication"]:checked').val()
    var $choose=$('input[name=risk_grand]')
    $choose.removeAttr("checked")
    switch (medication){
    case "1":
        $choose.eq(0).attr("checked",true)
        break
    case "2":
        if(c>=2){
            $choose.eq(0).attr("checked",true)
        } else{
            $choose.eq(1).attr("checked",true)
        }
        break
    case "3":
        if(c>=2) {
            $choose.eq(1).attr("checked",true)
        }else {
            $choose.eq(2).attr("checked",true)
        }
        break
    case "4":
        $choose.eq(2).attr("checked",true)
        break
    }
})
$('input[type=checkbox][name="risk_not_medication"]').change(function () {
    let c =$('input[name="risk_not_medication"]:checked').length

    let medication=$('input[type=radio][name="risk_medication"]:checked').val()
    var $choose=$('input[name=risk_grand]')
    $choose.removeAttr("checked")

    switch (medication){
        case "1":
            $choose.eq(0).attr("checked",true)
            break
        case "2":
            if(c>=2){
                $choose.eq(0).attr("checked",true)
            } else{
                $choose.eq(1).attr("checked",true)
            }
            break
        case "3":
            if(c>=2) {
                $choose.eq(1).attr("checked",true)
            }else {
                $choose.eq(2).attr("checked",true)
            }
            break
        case "4":
            $choose.eq(2).attr("checked",true)
            break
        }
   
    

})

