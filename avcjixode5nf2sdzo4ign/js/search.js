var urlq = {};
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
})

$('#toindex').on("click",function(){
    window.location.href = $.page.index()
})


$('#search_submit').on('click',function(){
    console.log("开始搜索")
    $("#search_out").empty()
    let form ={}
    form.action=$.act.Search_Sisk()
    form.name=$('#name').val()
    form.hospital_number=$('#hospital_number').val()
    form.attandance_number=$('#attandance_number').val()

    $.post_send.ajax(form,function(res){
        console.log(res)
        if (res.status != 1){
            $.toast.text(res.explain)
        }
        if (res.data[0].has == 0) {
            $.toast.text("搜索结果为空")
        }

        let a = $.append.table_cell("住院号")
        let b = $.append.table_cell("就诊号")
        let d = $.append.table_cell("姓名")
        $.append.table("#search_out",a,b,d)

        for (let c=0;c<15;c++){
            if (res.data[c].has==1){
                let a = $.append.table_cell(res.data[c].hospital_number)
                let b = $.append.table_cell(res.data[c].attandance_number)
                let d = $.append.table_cell(res.data[c].name)

                let div = $('<div id=tocycle data-userid="'+res.data[c].userid+'" class="weui-flex"></div>')
                div.append(a,b,d)
                $("#search_out").append(div)
            }
        }
    })
})


$('#search_out').on("click","div#tocycle",function(e){
    let id = $(this).data("userid")
    window.location.href = $.page.cycle() + "?userid=" + id;
})

$('#search_out').on("taphold","div#tocycle",function(e){
    let id = $(this).data("userid")
    console.log(233)
    window.location.href = $.page.add() + "?way=update&userid=" + id;
})
