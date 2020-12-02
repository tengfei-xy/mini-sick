$('#toindex').on("click",function(){
    window.location.href = $.page.index()
})
$(document).ready(function(){
    let form = {}
    form.action=$.act.Search_Wait()
    $.post_send.ajax(form,function(res){
        console.log(res)
        if (res.status == 0){
            $.toast.text(res.explain)
            return
        }
        for (let i=0;i<res.N.length;i++){
            if (res.N[i].has==1){
                console.log(res.N[i].name)
                let div=$('<div class="weui-flex"></div>').append($('<div id="tofollow" class="weui-flex__item follow-title" data-cycleseq="'+res.N[i].cycle_seq+'" data-userid="'+res.N[i].userid+'"></div>').text(res.N[i].name))
           
                $("#wait").append(div)
            }else{
                break
            }
        }
    })
})

$('#wait').on('click',"div#tofollow",function(){
    window.location.href=$.page.status() + "?way=wait&userid="+$(this).data("userid") + "&cycleseq="+$(this).data("cycleseq")
})