

$(document).ready(function(){
    

 	$("#myform").validate({
		  errorPlacement: function(error, element) {
			 alert("hello world!");
		  },
		  rules: {
		     fname: "required",
		     lname: {
		       required: true,
		       email: true
		     }
		  },
		  messages: {
		    fname: "Please specify your name",
		    lname: {
		       required: "We need your email address to contact you",
		       email: "Your email address must be in the format of name@domain.com"
		    }
		 }
 	})
});
