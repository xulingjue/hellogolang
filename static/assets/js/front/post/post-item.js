$(document).ready(function(){
  	window.prettyPrint && prettyPrint();
  	$('pre').attr('style', 'overflow:auto');

  	$(".input-error").hide();
  	$("#comment-form").validate({
          onsubmit:true,
          onfocusout:false,
          onkeyup:false,
          onclick:false,
          rules: {
            content: {
              required: true
            }
          },
          messages: {
            content: {
              required: "请输入回复内容"
            }
          },
          errorPlacement: function(error, element) {
            if(error.text()!=''){
              element.parent("p").prev("p").html(error.text());
              element.parent("p").prev("p").show();
            }
          },
          success: function(label) {
            label.parent().hide();
            return false;
          }
      }
    )
});

