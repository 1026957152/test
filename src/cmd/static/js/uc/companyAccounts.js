var companyAccountValidator = function () {
    var handleSubmit = function () {
        $('#userForm').validate({
            errorElement: 'span',
            errorClass: 'help-block',
            focusInvalid: false,
            rules: {
                realName: {
                    required: true,
                    maxlength: 50
                },
                emailAddress: {
                    email: true
                },
                pwd: {
                    required: true,
                    minlength: 6
                }
            },

            highlight: function (element) {
                $(element).closest('.form-group').addClass('has-error');
            },

            success: function (label) {
                label.closest('.form-group').removeClass('has-error');
                label.remove();
            },

            errorPlacement: function (error, element) {
                element.parent('div').append(error);
            },

            submitHandler: function (form) {
                form.submit();
            }
        });

        $('.form-horizontal input').keypress(function (e) {
            if (e.which == 13) {
                if ($('.form-horizontal').validate().form()) {
                    $('.form-horizontal').submit();
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