$(document).ready(function(){
	$("#login-form").submit(function(){
		var username = $(this).find("input[name='username']").val();
		var password = $(this).find("input[name='password']").val();
		if(username==''||!/^[a-zA-Z0-9]+$/.test(password)){
			$(this).find(".error").show(500);
		}else{
			return true;
		}
        return false; 
    });
})