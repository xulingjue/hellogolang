$(document).ready(function(){
	prettyPrint();
	$("#post_form").submit(function(){
		if($("input[name='title']").val()==""){
			alert("请输入标题内容！");
			return false;
		}

	/*	if($("input[name='contetnum']").val()==0){
			alert("请输入正文内容！");
			return false;
		}*/
		return true;
	})

	CKEDITOR.replace('post-content');
});

