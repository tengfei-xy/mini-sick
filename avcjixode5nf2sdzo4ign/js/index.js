var urlq = {};

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

    }
})
$(function(){
    var $search = $('#search'),
        $add = $('#add'),
        $quit = $('#quit'),
        $download=$('#download'),
        $wait=$('#wait')
    $search.on('click', function(){
        window.location.href = $.page.search() 
    });
    $add.on('click', function(){
        window.location.href = $.page.add()
    });
   
    $download.on('click', function(){
        $.toast.text("功能优化中")
    });
    $quit.on('click', function(){
        window.location.href = $.page.login();
    });
    $wait.on('click',function(){
        let today = (new Date()).getDay()
        console.log(today)
        if (today ==1 || today ==4){
            window.location.href = $.page.wait()
        }else{
            $.toast.text("今日不进行随访！")
            return
        }
    })
});