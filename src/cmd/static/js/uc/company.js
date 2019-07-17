var companyValidator = function() {
    var handleSubmit = function() {
        $('#companyForm').validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                companyName : {
                    required : true,
                    maxlength : 100
                },
                companyType : {
                    maxlength : 50
                },
                phoneNumber : {
                    maxlength : 20
                },
                faxNumber : {
                    maxlength : 20
                },
                websiteUrl : {
                    maxlength : 100
                },
                country : {
                    maxlength : 20
                },
                city : {
                    maxlength : 20
                },
                location : {
                    maxlength : 200
                },
                companyDesc : {
                    maxlength : 2048
                },
                mainBusiness : {
                    maxlength : 1024
                },
                orgCodeCertification : {
                    maxlength : 100
                },
                orgCode : {
                    maxlength : 100
                },
                businessLicense : {
                    maxlength : 100
                },
                legalRepresentative : {
                    maxlength : 45
                },
                taxAccount : {
                    maxlength : 45
                },
                coalLicense : {
                    maxlength : 100
                },
                bankName : {
                    maxlength : 45
                },
                bankBranch : {
                    maxlength : 100
                },
                bankAccount : {
                    maxlength : 45
                },
                accountName : {
                    maxlength : 45
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