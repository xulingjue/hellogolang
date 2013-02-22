$(document).ready(function(){
/*
 * 注册表单验证
 */
    jQuery.validator.addMethod("alnum", function(value, element) {
      return this.optional(element) || /^[a-zA-Z0-9]+$/.test(value);
    }, "只能包括英文字母和数字");

    $("#regist_form").validate({
        onkeyup: false,
        errorElement: "p",
        errorClass: "ym-message",
        rules: {
          name: {
            required: true,
            minlength: 3,
            maxlength: 10,
            //remote:"check.php",
          },
          email: {
            required: true,
            email: true,
            //remote:"check.php",
          },
          password: {
            required: true,
            minlength: 6,
            alnum:[]
          },
          repassword:{
            equalTo:"#password"
          }
        },
        messages: {
          name: {
            required: "请输入用户名",
            minlength: "用户名长度至少为四位",
            maxlength: "用户名长度最长为十位"
          },
          email: {
            required: "请输入邮箱",
            email: "请输入正确的邮箱地址",
          
          },
          password: {
            required: "请输入密码",
            minlength: "密码长度至少为六位"
          },
          repassword: "两次输入的密码不一致"
        },
        errorPlacement: function(error, element) {
          element.prev("label").find("span").html(error.text());
        },
        success: function(label) {
          label.parent("div").removeClass("ym-error");
          label.remove();
        }
    })
});