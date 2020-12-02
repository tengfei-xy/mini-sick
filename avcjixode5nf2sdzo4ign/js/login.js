$('#login_submit').on("click",function(){
    let form={}
    form.action = $.act.User_Login()
    form.name = $('#name').val()
    form.account = $('#account').val()
    form.password = $('#password').val()
    $.post_send.ajax(form,function(res){
        $.toast.text(res.explain)
        if (res.status == 0){
            localStorage.setItem("writer",res.data)
            window.location.href=$.page.index()
        }
    })
})
