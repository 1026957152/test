var coalSupplyValidator = function() {
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

                moq : {
                    required : true,
                    number : true,
                    min : 1,
                    lessOrEq : '#initialQuantity'
                },
                price : {
                    required : true,
                    number : true,
                    min : 1
                },
                intentionPrice : {
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
                initialQuantity : "数量必须大于1且小于等于库存量"
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

$.validator.addMethod("lessOrEq",function(value,element,params){
    var target = $( params );
    return parseFloat(value) <= parseFloat(target.val());
},"起订量必须大于1且小于等于数量");