$(document).ready(function(){
/*
 * 注册表单验证
 */
    $("#regist_form").validate({
        onkeyup: false,
        errorElement: "p",
        errorClass: "ym-message",
        rules: {
          name: {
            required: true,
            minlength: 6,
            remote: "process.php"
          },
          email: {
            required: true,
            email: true,
            remote: "process.php"
          },
          password: {
            required: true,
            minlength: 6
          },
          repassword:{
            equalTo:"#password"
          }
        },
        messages: {
          name: {
            required: "请输入用户名",
            minlength: "用户名长度至少为六位",
            remote: "用户名已被注册"
          },
          email: {
            required: "请输入邮箱",
            email: "请输入正确的邮箱地址",
            remote: "邮箱已被注册"
          },
          password: {
            required: "请输入密码",
            minlength: "密码长度至少为六位"
          },
          repassword: "两次输入的密码不一致"
        },
        errorPlacement: function(error, element) {

          element.parent("div").addClass("ym-error").prepend(error);
        },
        success: function(label) {
          label.parent("div").removeClass("ym-error");
          label.remove();
        }
    })
});