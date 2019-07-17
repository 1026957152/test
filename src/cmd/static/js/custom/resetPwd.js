var resetPwdValidator = function () {
    var handleSubmit = function () {
        $('#resetPwd_form').validate({
            errorElement: 'span',
            errorClass: 'help-block',
            focusInvalid: false,
            rules: {
                password: {
                    required: true,
                    minlength: 6
                },
                password_confirm: {
                    equalTo: '#password'
                }
            },
            messages: {
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
                var form = $('#resetPwd_form');
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