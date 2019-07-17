/**
 * Created by zhongkw on 1/8/2015.
 */
(function () {
    var mods = [], version = parseFloat(seajs.version);
    define(["jquery"], function (require, exports, module) {
        var uri = module.uri || module.id,
            m = uri.split('?')[0].match(/^(.+\/)([^\/]*?)(?:\.js)?$/i),
            root = m && m[1],
            name = m && ('./' + m[2]),
            i = 0,
            len = mods.length,
            curr, args;

        //unpack
        for (; i < len; i++) {
            args = mods[i];
            if (typeof args[0] === 'string') {
                name === args[0] && ( curr = args[2] );
                args[0] = root + args[0].replace('./', '');
                (version > 1.0) && define.apply(this, args);
            }
        }
        mods = [];
        require.get = require;
        return typeof curr === 'function' ? curr.apply(this, arguments) : require;
    });
    define.pack = function () {
        mods.push(arguments);
        (version > 1.0) || define.apply(null, arguments);
    };
})();
define.pack("./ajax", ["jquery"], function (require, exports, module) {
    var $ = require('jquery');
    var mod = {};

    mod.sendEmail = function (opt) {
        var defer = $.Deferred();
        $.ajax({
            type: 'post',
            url: window.ROOT_DIR + '/user/reset_pwd',
            data: opt,
            dataType: 'json',
            success: function (data) {
                defer.resolve(data);
            },
            error: function (err) {
                defer.reject(JSON.parse(err.responseText));
            }
        });
        return defer;
    };

    //发送注册邮件
    mod.sendRegisterEmail = function (opt) {
        var defer = $.Deferred();
        $.ajax({
            type: 'post',
            url: window.ROOT_DIR + '/user/verify_email',
            data: opt,
            dataType: 'json',
            success: function (data) {
                defer.resolve(data);
            },
            error: function (err) {
                defer.reject(JSON.parse(err.responseText));
            }
        });
        return defer;
    };

    //发送密码重置邮件
    mod.resetPassword = function (opt) {
        var defer = $.Deferred();
        $.ajax({
            type: 'put',
            url: window.ROOT_DIR + '/user/reset_pwd',
            data: opt,
            dataType: 'json',
            success: function (data) {
                defer.resolve(data);
            },
            error: function (err) {
                defer.reject(JSON.parse(err.responseText));
            }
        });
        return defer;
    }

    mod.registerUser = function (opt) {
        var defer = $.Deferred();
        $.ajax({
            type: 'post',
            url: window.ROOT_DIR + '/auto/user/register_user',
            data: opt,
            dataType: 'json',
            success: function (data) {
                defer.resolve(data);
            },
            error: function (err) {
                defer.reject(JSON.parse(err.responseText));
            }
        });
        return defer;
    };

    module.exports = mod;
});
define.pack("./bind", ["jquery", "./ajax"], function (require, exports, module) {
    var $ = require('jquery'),
        ajax = require('./ajax');

    var mod = {};
    mod.djsFlag = 0;

    mod.emailReg = /^([a-z0-9\+_\-]+)(\.[a-z0-9\+_\-]+)*@([a-z0-9\-]+\.)+[a-z]{2,6}$/i;

    mod.userType = 1;

    mod.resetBindInit = function () {
        var _this = this;

        if (_this.emailReg.test($('#email').val().trim())) {
            $('#sendEmailBtn').removeAttr('disabled');
        }

        _this.sendEmailBind();
    };

    mod.doResetBindInit = function () {
        var _this = this;
        _this.resetPasswordBind();
    };

    mod.resetPasswordBind = function () {
        var pwd1, pwd2;

        $('#pwd_1').off('keyup').on('keyup', function (e) {
            pwd1 = $('#pwd_1').val().trim();
            if (pwd1.length < 6) {
                $('#resetBtn').html('密码过短');
                $('#resetBtn').attr('disabled', 'disabled');
                return;
            }
            else {
                $('#resetBtn').html('完成');
            }
            if (pwd1 == pwd2) {
                $('#resetBtn').html('完成');
                $('#resetBtn').removeAttr('disabled');
            }
        });

        $('#pwd_2').off('keyup').on('keyup', function (e) {
            pwd2 = $('#pwd_2').val().trim();
            if (pwd2.length < 6) {
                $('#resetBtn').html('密码过短');
                $('#resetBtn').attr('disabled', 'disabled');
                return;
            }
            if (pwd2 != pwd1) {
                $('#resetBtn').html('密码不一致');
                $('#resetBtn').attr('disabled', 'disabled');
            }
            else {
                $('#resetBtn').html('完成');
                $('#resetBtn').removeAttr('disabled');
            }
        });

        $('#resetBtn').off('click').on('click', function (e) {
            var opt = {
                email: window.EMAIL,
                time: window.TIME,
                pwd_token: window.TOKEN,
                new_password: pwd1,
                new_password_confirm: pwd2
            };
            $('#resetBtn').html('提交中...').attr('disabled', 'disabled');
            ajax.resetPassword(opt).done(function (data) {
                if (data.status) {
                    $('#resetBtn').removeAttr('disabled').html('修改成功');
                    window.location.href = ROOT_DIR;
                }
                else {
                    alert(data.data.message);
                    $('#resetBtn').html('完成').removeAttr('disabled');
                }
            }).fail(function (err) {
                alert(err.data.message);
                $('#resetBtn').html('完成').removeAttr('disabled');
            });

        });
    };

    mod.sendEmailBind = function () {
        var _this = this;

        setInterval(function () {
            var email = $('#email').val().trim();
            if (_this.emailReg.test(email) && _this.djsFlag == 0) {
                $('#sendEmailBtn').removeAttr('disabled');
            }
            else {
                $('#sendEmailBtn').attr('disabled', 'disabled');
            }
        }, 300);

        $('#email').off('keyup').on('keyup', function (e) {
            var email = $('#email').val().trim();
            if (_this.emailReg.test(email)) {
                $('#sendEmailBtn').removeAttr('disabled');
            }
            else {
                $('#sendEmailBtn').attr('disabled', 'disabled');
            }
        });

        $('#sendEmailBtn').off('click').on('click', function () {
            var email = $('#email').val().trim();

            var opt = {
                email: email
            };

            ajax.sendEmail(opt).done(function (data) {
                if (data.status) {
                    //成功
                    $('#email_0').html(email);
                    $('#default_tip').hide();
                    $('#success_tip').show();
                    $('#email').off('keyup');
                    $('#sendEmailBtn').html('60 秒').off('click').attr('disabled', 'disabled');
                    _this.setDJS('sendEmailBtn');

                    //点击email进入对应邮箱
                    $('#email_0').off('click').on('click', function (e) {
                        e.preventDefault();
                        e.stopPropagation();
                        var url = _this.getURL(email);
                        if (url != '') {
                            window.location.href = 'http://' + url;
                        }
                        else {
                            alert('对不起，未找到对应的邮箱登录地址，请自行登录邮箱查看');
                        }
                    });
                }
                else {
                    alert(data.message);
                }
            }).fail(function (err) {
                alert(err.message);
            });
        });
    };

    mod.setDJS = function (domid) {
        var _this = this;
        var time = 60;
        _this.djsFlag = 1;
        var t = setInterval(function () {
            time--;
            $('#' + domid).html(time + ' 秒');
            if (time == 0) {
                $('#' + domid).html('发送').removeAttr('disabled');
                clearInterval(t);
                _this.djsFlag = 0;
                //_this.sendEmailBind();
                if (domid == 'sendEmailBtn') {
                    _this.sendEmailBind();
                }
                else {
                    _this.registerBind();
                }
            }
        }, 1000);
    };


    mod.registerBind = function () {
        var _this = this;
        setInterval(function () {
            var email = $('#email').val().trim();
            if (_this.emailReg.test(email) && _this.djsFlag == 0) {
                $('#registerBtn').removeAttr('disabled');
            }
            else {
                $('#registerBtn').attr('disabled', 'disabled');
            }
        }, 300);

        $('#email').off('keyup').on('keyup', function (e) {
            var email = $('#email').val().trim();
            if (_this.emailReg.test(email)) {
                $('#registerBtn').removeAttr('disabled');
            }
            else {
                $('#registerBtn').attr('disabled', 'disabled');
            }
        });

        $('#registerBtn').off('click').on('click', function () {
            var email = $('#email').val().trim();
            var opt = {
                email: email
            };
            ajax.sendRegisterEmail(opt).done(function (data) {
                if (data.status) {
                    $('#email_0').html(email);
                    $('#default_tip').hide();
                    $('#success_tip').show();
                    $('#email').off('keyup');
                    $('#registerBtn').html('60 秒').off('click').attr('disabled', 'disabled');
                    _this.setDJS('registerBtn');

                    $('#email_0').off('click').on('click', function (e) {
                        e.preventDefault();
                        e.stopPropagation();
                        var url = _this.getURL(email);
                        if (url != '') {
                            window.location.href = 'http://' + url;
                        }
                        else {
                            alert('对不起，未找到对应的邮箱登录地址，请自行登录邮箱查看');
                        }
                    });
                }
                else {
                    alert(data.message);
                }
            }).fail(function (err) {
                alert(err.message);
            });
        });

    };

    mod.verifyBind = function () {
        var _this = this;
        var pwd1, pwd2;

        $('#pwd_1').off('keyup').on('keyup', function (e) {
            pwd1 = $('#pwd_1').val().trim();
            if (pwd1.length < 6) {
                $('#resetBtn').html('密码过短');
                $('#resetBtn').attr('disabled', 'disabled');
                return;
            }
            else {
                $('#resetBtn').html('完成');
            }
            if (pwd1 == pwd2) {
                $('#resetBtn').html('完成');
                $('#resetBtn').removeAttr('disabled');
            }
        });

        $('#pwd_2').off('keyup').on('keyup', function (e) {
            pwd2 = $('#pwd_2').val().trim();
            if (pwd2.length < 6) {
                $('#resetBtn').html('密码过短');
                $('#resetBtn').attr('disabled', 'disabled');
                return;
            }
            if (pwd2 != pwd1) {
                $('#resetBtn').html('密码不一致');
                $('#resetBtn').attr('disabled', 'disabled');
            }
            else {
                $('#resetBtn').html('完成');
                $('#resetBtn').removeAttr('disabled');
            }
        });

        $('#resetBtn').off('click').on('click', function (e) {
            var opt = {
                email: window.EMAIL,
                time: window.TIME,
                register_token: window.TOKEN,
                //type:mod.userType,
                new_password: pwd1,
                new_password_confirm: pwd2
            };
            $('#resetBtn').html('提交中...').attr('disabled', 'disabled');
            ajax.registerUser(opt).done(function (data) {
                if (data.status) {
                    $('#resetBtn').removeAttr('disabled').html('注册成功');
                    window.location.href = ROOT_DIR;
                }
                else {
                    alert(data.data.message);
                    $('#resetBtn').html('完成').removeAttr('disabled');
                }
            }).fail(function (err) {
                alert(err.data.message);
                $('#resetBtn').html('完成').removeAttr('disabled');
            });

        });
    };

    mod.getURL = function ($mail) {
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

    module.exports = mod;
});
define.pack("./init", ["jquery", "./tmpl", "./bind"], function (require, exports, module) {
    var $ = require('jquery'),
        tmpl = require('./tmpl'),
        bind = require('./bind');

    var mod = {};

    mod.init = function () {

        var pageUrl = {
            'reset': 'resetpwd',
            'doReset': 'pwdemail',
            'preRegister': 'preRegister',
            'verify': 'register_detail'//'verify/email'
        };
        var loc = window.location.href;
        var htmlStr = '';

        if (loc.indexOf(pageUrl.reset) > 0) {
            htmlStr = tmpl.resetPassword({useremail: USEREMAIL});
            $('#p-body').html(htmlStr);
            bind.resetBindInit();
        }
        if (loc.indexOf(pageUrl.doReset) > 0) {
            htmlStr = tmpl.doResetPassword({useremail: USEREMAIL});
            $('#p-body').html(htmlStr);
            bind.doResetBindInit();
        }
        if (loc.indexOf(pageUrl.preRegister) > 0) {
            htmlStr = tmpl.register({useremail: USEREMAIL});
            $('#p-body').html(htmlStr);
            bind.registerBind();
        }
        if (loc.indexOf(pageUrl.verify) > 0) {
            htmlStr = tmpl.verifyemail({useremail: USEREMAIL});
            $('#p-body').html(htmlStr);
            bind.verifyBind();

            $('#company').trigger('click');
        }
    };

    module.exports = mod;
});
//tmpl file list:
//resetpwd/src/main.tmpl.html
define.pack("./tmpl", [], function (require, exports, module) {
    var tmpl = {
        'resetPassword': function (data) {

            var __p = [], _p = function (s) {
                __p.push(s)
            };
            __p.push('    <div class="container">\r\n\
        <div class="text-center">\r\n\
            <h2>重置密码</h2>\r\n\
            <p id="default_tip" style="display: ;" class="info">CoalValue 将发送一封密码重置邮件到你的邮箱<br />点击按钮即可重置密码</p>\r\n\
            <p id="success_tip" style="display:none;" class="info">密码重置邮件已发送到 <a style="text-decoration: underline;" href="javascript:void 0;" id="email_0"></a>，请在24小时内激活完成注册。<br>\r\n\
                如果没有收到，请检查垃圾邮件<br>如果还是没有收到，请重新填写邮箱</p>\r\n\
            <div class="form-inline">\r\n\
                <div class="form-group">\r\n\
                    <label class="sr-only">输入你的邮箱</label>\r\n\
                    <input ');
            if (data.useremail != '') {
                __p.push('readonly');
            }
            __p.push(' id="email" value="');
            _p(data.useremail);
            __p.push('" class="form-control" placeholder="输入你的邮箱" />\r\n\
                </div>\r\n\
                <button disabled id="sendEmailBtn" type="submit" class="btn btnBlack">发送</button>\r\n\
            </div>\r\n\
        </div>\r\n\
    </div>');

            return __p.join("");
        },

        'doResetPassword': function (data) {

            var __p = [], _p = function (s) {
                __p.push(s)
            };
            __p.push('    <div class="retrieveOuter wrapper">\r\n\
    <div class="retrieveInner">\r\n\
        <h2>重置密码</h2>\r\n\
        <p class="info">创建密码并确认即可完成密码重置</p>\r\n\
        <div class="form-horizontal">\r\n\
            <div class="form-group">\r\n\
                <label class="sr-only">创建新密码</label>\r\n\
                <input type="password" id="pwd_1" class="form-control" placeholder="创建新密码" />\r\n\
            </div>\r\n\
            <div class="form-group">\r\n\
                <label class="sr-only">重新输入密码</label>\r\n\
                <input type="password" id="pwd_2" class="form-control" placeholder="重新输入密码" />\r\n\
            </div>\r\n\
            <div class="form-group">\r\n\
                <button style="width: 120px;" disabled id="resetBtn" type="submit" class="btn btnBlack">完成</button>\r\n\
            </div>\r\n\
        </div>\r\n\
    </div>\r\n\
    </div>');

            return __p.join("");
        },

        'verifyemail': function (data) {

            var __p = [], _p = function (s) {
                __p.push(s)
            };
            __p.push('    <div class="retrieveOuter wrapper">\r\n\
        <div class="retrieveInner">\r\n\
            <h2>设置密码</h2>\r\n\
            <p class="info">创建密码并确认即可完成注册</p>\r\n\
            <div class="form-horizontal">\r\n\
                <div class="form-group">\r\n\
                    <label class="sr-only">创建新密码</label>\r\n\
                    <input type="password" id="pwd_1" class="form-control" placeholder="创建新密码" />\r\n\
                </div>\r\n\
                <div class="form-group">\r\n\
                    <label class="sr-only">重新输入密码</label>\r\n\
                    <input type="password" id="pwd_2" class="form-control" placeholder="重新输入密码" />\r\n\
                </div>\r\n\
                <div class="form-group">\r\n\
                    <button style="width: 120px;" disabled id="resetBtn" type="submit" class="btn btnBlack">完成</button>\r\n\
                </div>\r\n\
            </div>\r\n\
        </div>\r\n\
    </div>');

            return __p.join("");
        },

        'register': function (data) {

            var __p = [], _p = function (s) {
                __p.push(s)
            };
            __p.push('    <div class="container">\r\n\
        <div class="text-center">\r\n\
            <h2>注 册</h2>\r\n\
            <p id="default_tip" style="display: ;" class="info">CoalValue 将发送一封验证邮件到你的邮箱，此邮箱将作为登录用户名<br />点击按钮，注册账号</p>\r\n\
            <p id="success_tip" style="display:none;" class="info">注册验证邮件已发送到 <a style="text-decoration: underline;" href="javascript:void 0;" id="email_0"></a><br>\r\n\
                如果没有收到，请检查垃圾邮件<br>如果还是没有收到，请重新填写邮箱</p>\r\n\
            <div class="form-inline">\r\n\
                <div class="form-group">\r\n\
                    <label class="sr-only">输入你的邮箱</label>\r\n\
                    <input id="email" value="');
            _p(data.useremail);
            __p.push('" class="form-control" placeholder="输入你的邮箱" />\r\n\
                </div>\r\n\
                <button disabled id="registerBtn" type="submit" class="btn btnBlack">发送</button>\r\n\
            </div>\r\n\
        </div>\r\n\
    </div>');

            return __p.join("");
        }
    };
    return tmpl;
});
