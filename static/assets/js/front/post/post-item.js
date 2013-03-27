$(document).ready(function(){
  	window.prettyPrint && prettyPrint();
  	$('pre').attr('style', 'overflow:auto');

  	$(".input-error").hide();
	$("#comment-form").validate({
        errorElement: "p",
        rules: {
          content: {
            required: true,
          }
        },
        messages: {
          content: {
            required: "请填写回复内容",
          }
        },
        errorPlacement: function(error, element) {
          element.parent("p").prev("p").html(error.text());
          element.parent("p").prev("p").show();
        },
        success: function(label) {
          label.parent().hide();
        }
    })
});

