var resourceValidator = function() {
    var handleSubmit = function() {
        $('.form-horizontal').validate({
            errorElement : 'span',
            errorClass : 'help-block',
            focusInvalid : false,
            rules : {
                groupName : {
                    required : true,
                    minlength : 2,
                    maxlength : 50
                },
                groupDesc : {
                    maxlength : 200
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
                //submit changed resource status to server
                var resourceTree = $.fn.zTree.getZTreeObj('resourceTree');
                var nodes = resourceTree.getChangeCheckedNodes();
                var resStr = '';
                for (var i = 0; i < nodes.length; i++) {
                    resStr += nodes[i].type + "," + nodes[i].id + "," + nodes[i].checked;
                    if (i != nodes.length - 1) {
                        resStr += ";"
                    }
                }
                var hiddenField = document.createElement("input");
                hiddenField.setAttribute("type", "hidden");
                hiddenField.setAttribute("name", "resStr");
                hiddenField.setAttribute("value", resStr);
                form.appendChild(hiddenField);

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

var resourceTree = function() {
    var resourceTreeInit = function() {
        <!-- ztree settings -->
        var setting = {
            async: {
                autoParam: ["id=pid"],
                contentType: "application/x-www-form-urlencoded",
                dataFilter: transform_data,
                dataType: "text",
                enable: true,
                otherParam: {"resGrpId": $("#id").val()},
                type: "post",
                url: "/admin/permission/resource/tree"
            },
            check: {
                enable: true
            },
            data: {
                simpleData: {
                    enable: true,
                    idKey: "id",
                    pIdKey: "pid",
                    rootPId: null
                }
            },
            view: {
                showLine: false
            }
        };
        function transform_data(treeId, parentNode, data) {
            if (!data) return null;
            var menuLen = data.menus.length
            var btnLen = data.buttons.length;
            var treeData = [];
            for (var i = 0; i < menuLen; i++) {
                var node = {};
                node.id = data.menus[i].id;
                node.name = data.menus[i].menuName;
                node.pid = data.menus[i].parentId;
                node.isParent = !data.menus[i].leaf;
                node.checked = data.menuStatus[i];
                node.type = 'MENU';
                treeData.push(node);
            }
            for (var i = 0; i < btnLen; i++) {
                var node = {};
                node.id = data.buttons[i].id;
                node.name = data.buttons[i].buttonName;
                node.pid = data.buttons[i].menuId;
                node.isParent = false;
                node.checked = data.buttonStatus[i];
                node.type = 'BUTTON';
                treeData.push(node);
            }
            return treeData;
        }
        $.fn.zTree.init($("#resourceTree"), setting);
    }
    return {
        init: function() {
            resourceTreeInit();
        }
    }
}();