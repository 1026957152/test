/**
 * Created by zohu on 4/4/2015.
 */
var contractTmplValidator = function() {
    var paramSettings = {
        txtMinLen:0,
        txtMaxLen:200
    };

    var handleSubmit = function(settings) {
        $.extend(paramSettings, settings);
        $("#contractTemplateForm").validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                templateName:{
                    required:true
                },
                variety:{
                    maxlength:paramSettings.txtMaxLen,
                    minlength:paramSettings.txtMinLen
                },
                quality:{
                    maxlength:paramSettings.txtMaxLen,
                    minlength:paramSettings.txtMinLen
                },
                deliveryWay:{
                    maxlength:paramSettings.txtMaxLen,
                    minlength:paramSettings.txtMinLen
                },
                price:{
                    maxlength:paramSettings.txtMaxLen,
                    minlength:paramSettings.txtMinLen
                },
                paymentMean:{
                    maxlength:paramSettings.txtMaxLen,
                    minlength:paramSettings.txtMinLen
                },
                checkMean:{
                    maxlength:paramSettings.txtMaxLen,
                    minlength:paramSettings.txtMinLen
                },
                tradeDispute:{
                    maxlength:paramSettings.txtMaxLen,
                    minlength:paramSettings.txtMinLen
                },
                otherClause:{
                    maxlength:paramSettings.txtMaxLen,
                    minlength:paramSettings.txtMinLen
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
        init : function(settings) {
            handleSubmit(settings);
        }
    };

}();