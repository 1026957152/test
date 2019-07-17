var loginValidator = function () {
    var handleSubmit = function () {
        $('#login_form').validate({
            errorElement: 'span',
            errorClass: 'help-block',
            focusInvalid: false,
            rules: {
                email: {
                    required: true,
                    email: true
                },
                password: {
                    required: true
                },
                checkCode: {
                    required:true,
                    remote: {// verify the checkcode is correct or not
                        type: "POST",
                        url: "/homepage/user/verifyCheckCode",
                        async: true,
                        data: {
                            checkCode: function () {
                                return $('#checkCode').val();
                            }
                        },
                        dataType: "json"
                    }
                }
            },
            messages: {
                email: "请输入正确的邮箱地址",
                password: "请输入密码",
                checkCode: {
                    required: "请输入验证码",
                    remote: "验证码错误，请重新输入"
                }
            },

            success: function (label) {
                label.remove();
            },

            errorPlacement: function (error, element) {
                $('#serverMsg').html('');
                $('#msg').append(error);
            },

            submitHandler: function (form) {
                form.submit();
            }
        });

        $('#login_form input').keypress(function (e) {
            if (e.which == 13) {
                var loginForm = $('#login_form');
                if (loginForm.validate().form()) {
                    loginForm.submit();
                }
                return false;
            }
        });
    }
    return {
        init: function () {
            handleSubmit();
        }
    };

}();