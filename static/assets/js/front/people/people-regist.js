$(document).ready(function(){
/*
 * 注册表单验证
 */ 
    jQuery.validator.addMethod("alnum", function(value, element) {
      return this.optional(element) || /^[a-zA-Z0-9\u4E00-\u9FA5]+$/.test(value);
    }, "只能包括英文字母和数字");

    $("#regist_form").validate({
        onkeyup:true,
        errorElement: "p",
        rules: {
          name: {
            required: true,
            minlength: 2,
            maxlength: 10,
            remote:"/people/isexist/",
            alnum:[],
          },
          email: {
            required: true,
            email: true,
            remote:"/people/isexist/",
          },
          password: {
            required: true,
            minlength: 6,
            alnum:[],
          },
          repassword:{
            equalTo:"#password"
          }
        },
        messages: {
          name: {
            required: "&nbsp;请输入用户名",
            minlength: "&nbsp;用户名长度至少为两位",
            maxlength: "&nbsp;用户名长度最长为十位",
            remote: "&nbsp;该用户名已被注册",
            alnum: "&nbsp;用户名含有特殊字符",
          },
          email: {
            required: "&nbsp;请输入邮箱",
            email: "&nbsp;请输入正确的邮箱地址",
            remote: "&nbsp;该邮箱已被注册",
          },
          password: {
            required: "&nbsp;请输入密码",
            minlength: "&nbsp;密码长度至少为六位",
            alnum: "&nbsp;密码只能为英文或数字",
          },
          repassword:{
            equalTo:"&nbsp;两次输入的密码不一致",
          } 
        },
        errorPlacement: function(error, element) {
          element.next(".input-error").html(error.text());
        },
        success: function(label) {
          label.parent("div").removeClass("ym-error");
          label.remove();
        }
    })
});