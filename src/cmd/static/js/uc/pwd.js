var pwdValidator = function() {
    var handleSubmit = function() {
        $('.form-horizontal').validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                orgPwd : {
                    remote : {
                        type: "POST",
                        url: "/usercenter/account/verifyOriginalPwd",
                        async: true,
                        data: {
                            orgPwd: function () {
                                return $('#orgPwd').val();
                            }
                        },
                        dataType: "json"
                    }
                },
                pwd1 : {
                    required : true
                },
                pwd2 : {
                    equalTo : "#pwd1"
                }
            },
            messages : {
                orgPwd : "密码不正确",
                pwd1 : "请输入密码",
                pwd2 : "两次密码输入不一致"
            },

            highlight : function(element) {
                $(element).closest('.form-group').addClass('has-error');
            },

            success : function(label) {
                label.closest('.form-group').removeClass('has-error');
                label.remove();
            },

            errorPlacement : function(error, element) {
                element.parent('div').append(error);
            },

            submitHandler : function(form) {
                form.submit();
            }
        });

        $('.form-horizontal input').keypress(function(e) {
            if (e.which == 13) {
                if ($('.form-horizontal').validate().form()) {
                    $('.form-horizontal').submit();
                }
                return false;
            }
        });
    }
    return {
        init : function() {
            handleSubmit();
        }
    };

}();