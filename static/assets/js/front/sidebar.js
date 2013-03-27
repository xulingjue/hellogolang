$(document).ready(function(){
	$("#login-form").submit(function(){
		var username = $(this).find("input[name='username']").val();
		var password = $(this).find("input[name='password']").val();
		var loginErr = false;

		if(username==''||!/^[a-zA-Z0-9]+$/.test(password)){
			$(this).find(".error").show(300);
		}else{
			//ajax登录
			$.ajax({
			   	type: "POST",
			   	url: "/ajaxlogin/",
			   	dataType:"json",
			   	data: {
			   		"name":username,
			   		"password":password
			   	},
			   	success: function(msg){
			    	if(msg.result=="success"){
			    		$("account").html('\
							<div class="avatar">\
								<a href="#"><img width="70" src="/assets/image/default3.png" alt="lingjue"></a>\
							</div>\
							<div class="userinfo">\
								<span class="name"><a href="#">'+msg.people.Name+'</a></span>\
								<span class="score">发帖：50</span>\
								<span class="score">提问：50</span>\
							</div>\
							<div class="clear"></div>\
							<div class="user-btn-panel">\
								<span><a href="/post/create/" rel="nofollow" class="btncjcp">发文章</a></span>\
								<span><a href="/diet/create" rel="nofollow" class="btnjlrj">提问题</a></span>\
								<div class="clear"></div>\
							</div>\
							<div class="clear"></div>\
			    			')
			    		//更新界面
			    	}else{
			    		$("#login-form").find(".error").show(300);
			    	}
			   	}
			})
		}

        return false; 
    });
})