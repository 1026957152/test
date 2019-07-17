$(function(){
    var preRegister = function (){
        var emailReg = /^([a-z0-9\+_\-]+)(\.[a-z0-9\+_\-]+)*@([a-z0-9\-]+\.)+[a-z]{2,6}$/i;

        $('#email').off('keyup').on('keyup',function (e) {
            var email = $('#email').val().trim();
            if (emailReg.test(email)) {
                $('#registerBtn').removeAttr('disabled');
            }
            else {
                $('#registerBtn').attr('disabled', 'disabled');
            }
        });
        $('#email').change(function (e) {
            var email = $('#email').val().trim();
            if (emailReg.test(email)) {
                $('#registerBtn').removeAttr('disabled');
            }
            else {
                $('#registerBtn').attr('disabled', 'disabled');
            }
        });


        $('#registerBtn').off('click').on('click', function () {
            var email = $('#email').val().trim();
            var req = $.ajax({
                type: 'post',
                url: '/homepage/user/sendPreRegisterMail',
                data: {email: email}
            });
            req.done(function (data) {
                if (data.status) {
                    $('#email_0').html(email);
                    $('#default_tip').hide();
                    $('#success_tip').show();
                    $('#email').off('keyup');

                    $('#email_0').off('click').on('click', function (e) {
                        e.preventDefault();
                        e.stopPropagation();
                        var url = getURL(email);
                        if (url != '') {
                            window.location.href = 'http://' + url;
                        }
                        else {
                            alert('对不起，未找到对应的邮箱登录地址，请自行登录邮箱查看');
                        }
                    });
                }
                else {
                    $('#default_tip').show();
                    $('#success_tip').hide();
                    alert(data.message);
                }
            });
        });
        var getURL = function ($mail) {
            $t = $mail.split('@')[1];
            $t = $t.toLowerCase();
            if ($t == '163.com') {
                return 'mail.163.com';
            } else if ($t == 'vip.163.com') {
                return 'vip.163.com';
            } else if ($t == '126.com') {
                return 'mail.126.com';
            } else if ($t == 'qq.com' || $t == 'vip.qq.com' || $t == 'foxmail.com') {
                return 'mail.qq.com';
            } else if ($t == 'outlook.com') {
                return 'outlook.com';
            } else if ($t == 'gmail.com') {
                return 'mail.google.com';
            } else if ($t == 'sohu.com') {
                return 'mail.sohu.com';
            } else if ($t == 'tom.com') {
                return 'mail.tom.com';
            } else if ($t == 'vip.sina.com') {
                return 'vip.sina.com';
            } else if ($t == 'sina.com.cn' || $t == 'sina.com') {
                return 'mail.sina.com.cn';
            } else if ($t == 'tom.com') {
                return 'mail.tom.com';
            } else if ($t == 'yahoo.com.cn' || $t == 'yahoo.cn') {
                return 'mail.cn.yahoo.com';
            } else if ($t == 'tom.com') {
                return 'mail.tom.com';
            } else if ($t == 'yeah.net') {
                return 'www.yeah.net';
            } else if ($t == '21cn.com') {
                return 'mail.21cn.com';
            } else if ($t == 'hotmail.com') {
                return 'www.hotmail.com';
            } else if ($t == 'sogou.com') {
                return 'mail.sogou.com';
            } else if ($t == '188.com') {
                return 'www.188.com';
            } else if ($t == '139.com') {
                return 'mail.10086.cn';
            } else if ($t == '189.cn') {
                return 'webmail15.189.cn/webmail';
            } else if ($t == 'wo.com.cn') {
                return 'mail.wo.com.cn/smsmail';
            } else if ($t == '139.com') {
                return 'mail.10086.cn';
            } else {
                return '';
            }
        };
    }();
});