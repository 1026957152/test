/**
 * Created by zohu on 3/21/2015.
 */
var coalProductValidator = function() {

    var handleSubmit = function() {
        $("#coalProductForm").validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                coalTypeLists:{
                    required:true
                },
                initialAmount:{
                    required:true,
                    number:true,
                    remote:{// verify the new coalInitialAmount is correct or not
                        type: "POST",
                        url: "/usercenter/others/coal/product/verifyCoalAmount",
                        async: true,
                        data: {
                            id: function () {
                                return $('#id').val();
                            },
                            initialAmount: function() {
                                return $('#initialAmount').val();
                            }
                        },
                        dataType: "json"
                    }
                },
                inspectionAgent:{
                    required:true
                },
                inspectionStandard:{
                    required:true
                },
                note:{
                    required:true,
                    maxlength:200
                },
                companyUser:{
                    required:true
                }
            },
            messages : {
                initialAmount:{
                    required: "商品库存数量必须输入",
                    number: "商品库存数量必须输入数字",
                    remote: "您修改的库存数量不符合要求，请重新输入"
                },
                inspectionAgent:{
                    required: "必须填写煤炭监察机构"
                },
                inspectionStandard:{
                    required: "必须填写煤炭检验标准"
                },
                note:{
                    maxlength:"备注必须控制在200字以内"
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
        init : function() {
            handleSubmit();
        }
    };

}();