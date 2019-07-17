/**
 * Created by zohu on 4/11/2015.
 */

var coalDemandValidator = function() {
    var handleSubmit = function() {
        $('.form-horizontal').validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                title : {
                    required : true,
                    maxlength : 100
                },
                demandQuantity :{
                    required : true,
                    number : true,
                    min :1
                },
                deliveryStartDate : {
                    required : true,
                    dateISO:false
                },
                deliveryEndDate : {
                    required : true,
                    dateISO:false
                },
                price : {
                    number : true,
                    min : 1
                },
                depositRate : {
                    number : true,
                    min : 1,
                    max : 100
                },
                note : {
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
                validateAll(e);
            }
        });
        $("#submitBtn").on("click", function(e) {
            validateAll(e);
        });

        var validateAll = function(e){
            if ($('.form-horizontal').validate().form()
                && validateCoalType()
                && validateDeliveryAddress()
                && validateContract()) {
                $('.form-horizontal').submit();
            } else {
                e.preventDefault();
                return false;
            }
        }
        var validateCoalType = function(){
            var b = $("#coalTypeList").html() != '';
            if (!b) {
                alert("请先新建煤种！");
            }
            return b;
        }
        var validateDeliveryAddress = function(){
            var b = $("#deliveryAddressList").html() != '';
            if (!b) {
                alert("请先添加收货地点！");
            }
            return b;
        }
        var validateContract = function(){
            var b = $("#contractDetailDiv").html() != '';
            if (!b) {
                alert("请先添加合同！");
            }
            return b;
        }

    }
    return {
        init : function() {
            handleSubmit();
        }
    };
}();