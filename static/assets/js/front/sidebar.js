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
			    		$(".account").prev("h2").hide();
			    		$(".account").html('\
							<div class="avatar">\
								<a href="#"><img width="70" src="/assets/image/default.png" alt="lingjue"></a>\
							</div>\
							<div class="userinfo">\
								<span class="name"><a href="#">'+msg.people.Name+'</a></span>\
								<span class="buttons"><a href="/post/create/" >发文章</a> <a href="/diet/create" >提问题</a></span>\
								<span class="logout"><a href="/logout/">退出</a></span>\
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