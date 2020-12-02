var urlq = {};
var userid =""
var way = 1
var form={}
$(document).ready(function(){
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

        if (urlq["way"]=="update"){

            form.action = $.act.Search_detail_Sick()
            form.userid = urlq["userid"]
            $.post_send.ajax(form,function(res){
                console.log(res)
                if (res.status ==0){
                    $.toast.text(res.explain)
                    window.location.href=$.page.search()
                    return
                }
                way=2
                $('#name').val(res.name)
                $.auto.choose(':radio[name="gender"]',res.gender)
                $('#age').val(res.age)
                $('#telphone').val(res.telphone)
                $('#hospital_number').val(res.hospital_number)
                $('#attandance_number').val(res.attandance_number)
                $('#disease').val(res.disease)
                $('#add_submit').val("保存患者信息")

            })
        }

    }

    

})



$('#add_submit').on("click",function(){

    form.action = $.act.Add_Sick()
    form.way=way
    form.name = $('#name').val()
    form.gender = $.auto.array('input[name="gender"]:checked')
    form.age=$('#age').val()
    form.telphone=$('#telphone').val()
    form.hospital_number=$('#hospital_number').val()
    form.attandance_number=$('#attandance_number').val()
    form.disease=$('#disease').val()
    form.writer=localStorage.getItem("writer")
    if (!$.check.isNumber(form.age)){$.toast.text("年龄填写错误");return}
    if (!$.check.isNumber(form.telphone)){$.toast.text("手机号填写错误");return}
    if (!$.check.isNumber(form.hospital_number)){$.toast.text("住院号应为数字");return}
    if (!$.check.isNumber(form.attandance_number)){$.toast.text("就诊号应为数字");return}
    if(form.name=="" && form.name=="" && form.name==""){
        $.toast.text("住院号、就诊号、姓名至少填写一个");return
    }

    $.post_send.ajax(form,function(res){
        console.log(res)
        $.toast.text(res.explain)
        if (res.status==1){
            userid=res.data
            $('#add_mask').css("display","block")
            $('#add_dialog').css("display","block")
        }
    })
})


$('#tostatus').on("click",function(){
    localStorage.setItem("back",$.page.index())
    window.location.href=$.page.status() + "?cycleseq=1&init=0&userid="+userid
    
})
$('#toindex').on("click",function(){
    window.location.href=$.page.index()
})