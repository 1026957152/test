var personValidator = function() {
    var handleSubmit = function() {
        $('#truckForm').validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                realName : {
                    required : true,
                    minlength : 2,
                    maxlength : 20
                },
                nickName : {
                    maxlength : 50
                },
                mobileNumber : {
                    rangelength : [11,11],
                    digits : true
                },
                department : {
                    maxlength : 50
                },
                position : {
                    maxlength : 50
                },
                note : {
                    maxlength : 200
                }
            },
            messages : {
                mobileNumber : "请输入11位手机号"
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