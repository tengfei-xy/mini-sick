var userid = 0
var cylce_seq_add = 1
var urlq={}
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

        let form={}
        form.action=$.act.Serch_Cycle()
        form.userid=urlq['userid']

        // 查询患者化疗周期
        $.post_send.ajax(form,function(res){
            console.log(res)
            for (let c=0;c<15;c++){
                if (res.data[c].has==1){
                    let a = $.append.table_cell(res.data[c].cycle_seq)
                    let b = $.append.table_cell(res.data[c].time)
                    let div = $('<div id=tostatus_cat data-cycleseq="'+res.data[c].cycle_seq+'" data-sicker="'+userid  +'" class="weui-flex table_box"></div>')
                    div.append(a,b)

                    $("#cycle").append(div)
                    cylce_seq_add += 1
                }

            }
            let but = $('<input style="margin-top:3rem" class="weui-btn weui-btn_primary"></input>')
            but.attr("id", "tostatus_add")
            but.attr("value","添加化疗周期")
            but.attr("data-cycleseq",cylce_seq_add)
            but.attr("data-userid",userid)
            $('#cycle').append(but)
        })
   
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
    window.location.href = $.page.index()
})

$('#cycle').on("click","div#tostatus_cat",function(){
    let seq = $(this).data("cycleseq")
    window.location.href = $.page.status() + "?cycleseq=" + seq + "&init=1&userid="+urlq['userid'];
})
$('#cycle').on("click","input#tostatus_add",function(){
    let seq = $(this).data("cycleseq")
    window.location.href = $.page.status() + "?cycleseq=" + seq + "&init=0&userid="+urlq['userid'];

})