var registerValidator = function () {
    var handleSubmit = function () {
        $('#register_form').validate({
            errorElement: 'span',
            errorClass: 'help-block',
            focusInvalid: false,
            rules: {
                userName: {
                    required: true,
                    remote: {// verify the username is existed or not
                        type: "POST",
                        url: "/homepage/user/verifyUserNameAllowed",
                        async: true,
                        data: {
                            userName: function () {
                                return $('#userName').val();
                            }
                        },
                        dataType: "json"
                    }
                },
                password: {
                    required: true,
                    minlength: 6
                },
                password_confirm: {
                    equalTo: '#password'
                }
            },
            messages: {
                userName: {
                    required: "请填写用户名",
                    remote: "该用户名已经存在"
                },
                password: {
                    required: "请输入密码",
                    minlength: "密码至少6位以上"
                },
                password_confirm: "密码不一致"
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
                var form = $('#register_form');
                if (form.validate().form()) {
                    form.submit();
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