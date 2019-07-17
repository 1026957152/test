/**
 * Created by zohu on 4/6/2015.
 */
/**
 * Created by zohu on 4/4/2015.
 */
var deliveryAddrValidator = function() {

    var handleSubmit = function() {
        $("#deliveryAddrForm").validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                name:{
                    required:true,
                    maxlength:30
                },
                address:{
                    required:true,
                    maxlength:30
                }
            },
            messages : {
                required:"该字段必须输入",
                maxlength:"该字段最多输入不超过30个字"
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