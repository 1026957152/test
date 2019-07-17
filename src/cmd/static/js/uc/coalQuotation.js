/**
 * Created by zohu on 4/11/2015.
 */

var coalQuotationValidator = function() {
    var handleSubmit = function() {
        $('.form-horizontal').validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                unitPrice :{
                    required : true,
                    number : true,
                    digits:true,
                    min :1
                },
                supplyNum :{
                    required : true,
                    number : true,
                    digits:true,
                    min :1
                },
                transExpense :{
                    required : true,
                    number : true,
                    digits:true,
                    min :1
                },
                quotationNote :{
                    maxlength : 500
                }
            },
            messages : {

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