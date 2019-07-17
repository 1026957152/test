/**
 * Created by zohu on 3/21/2015.
 */
var coalTypeValidator = function() {
    var paramSettings = {
        coalParamsMin:0,
        coalParamsMax:10000
    };

    var handleSubmit = function(settings) {
        $.extend(paramSettings, settings);
        $("#coalTypeForm").validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                productName:{
                    required:true,
                    maxlength:25
                },
                productArea:{
                    required:true,
                    maxlength:25
                },
                qnetarMin:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                category:{
                    required:true
                },
                granularity:{
                    required:true
                },

                prefecture:{
                    required:true
                },
                county:{
                    required:true
                },
                province:{
                    required:true
                },

                qnetarMax:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                aarMin:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                aarMax:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                starMin:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                starMax:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                modMin:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                modMax:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                varMin:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                varMax:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                crcMin:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                crcMin:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                crcMax:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                gMin:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                gMax:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                mtMin:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                mtMax:{
                    number:true,
                    min:paramSettings.coalParamsMin,
                    max:paramSettings.coalParamsMax
                },
                note:{
                    required: false,
                    maxlength:200
                },
                companyUser:{
                    required:true
                }
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