$(document).ready(function(){
/*
 * 注册表单验证
 */
    jQuery.validator.addMethod("alnum", function(value, element) {
      return this.optional(element) || /^[a-zA-Z0-9]+$/.test(value);
    }, "只能包括英文字母和数字");

    $("#login-form").validate({
        errorElement: "p",
        rules: {
          name: {
            required: true,
          },
          password: {
            required: true,
            minlength: 6,
            alnum:[],
          },

        },
        messages: {
          name: {
            required: "请输入用户名/邮箱",
          },
          password: {
            required: "请输入密码",
            minlength: "密码长度至少为六位",
            alnum: "密码只能为英文或数字",
          },

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