function page() {
    var pageSettings = {
        page: 0,
        size: 10,
        placeHoler: 0,//0:{0}{1}style; else:{property} style
        rowTemplateDiv: '#rowTemplate',
        pageDataDiv: '#pageData',
        paginationDiv: '#pagination',
        params: function () {
        },
        rowData: function (row) {
            return row;
        },
        afterRenderTable: function() {

        }
    };
    var initPage = function (settings) {
        $.extend(pageSettings, settings);
        showPage(pageSettings.page);
    };
    var refreshPage = function() {
        showPage(pageSettings.page);
    }
    var showPage = function (i) {
        var goToPage = function (i) {
            pageSettings.page = i;
            pullData();
        };
        var pullData = function () {
            var pageData = {
                page: pageSettings.page,
                size: pageSettings.size
            };

            $.ajax({
                url: pageSettings.url,
                type: 'post',
                data: $.extend(pageData, pageSettings.params()),
                success: renderPage
            })
        };
        var renderPage = function (data) {
            renderTable(data);
            afterRenderTable();
            renderPagination(data);
            afterRenderPagination();
        };
        var afterRenderTable = function () {
            pageSettings.afterRenderTable();
        }
        var renderTable = function (data) {
            var content = data.content;
            var html = new Array();
            content.forEach(function (row) {
                var rowTemplate = $(pageSettings.rowTemplateDiv).html();
                var rowData = pageSettings.rowData(row);
                var rowHtml = "";
                if (pageSettings.placeHoler == 0) {
                    rowHtml = format(rowTemplate, rowData);
                } else {
                    rowHtml = formatByObject(rowTemplate, rowData);
                }
                html.push(rowHtml);
            });
            $(pageSettings.pageDataDiv).html(html.join(""));
        }
        var renderPagination = function (data) {
            if (data.totalPages <= 1) {
                $(pageSettings.paginationDiv).html('');
                return;
            }
            var html = new Array();
            html.push("<div class=\"pull-right pagination\">");

            var page_li = "<li class='{0}' value='{1}'><a href='javascript:void(0)'>{2}</a></li>";
            //render pre actions
            var prePage = data.number - 1;
            if (data.first) {
                html.push(format(page_li, ['page-first disabled', '0', '&lt;&lt;']));
                html.push(format(page_li, ['page-pre disabled', prePage, '&lt;']));
            } else {
                html.push(format(page_li, ['page-first', '0', '&lt;&lt;']));
                html.push(format(page_li, ['page-pre', prePage, '&lt;']));
            }

            var currentPage = data.number;
            var startPage = (currentPage - 2) < 0 ? 0 : (currentPage - 2);
            var endPage = (currentPage + 3) > data.totalPages ? data.totalPages : (currentPage + 3);
            if (endPage - startPage < 5) {
                if (startPage == 0) {
                    endPage = (startPage + 5) > data.totalPages ? data.totalPages : (startPage + 5);
                }
                if (endPage == data.totalPages) {
                    startPage = (endPage - 5) < 0 ? 0 : (endPage - 5);
                }
            }

            //render page numbers
            for (var i = startPage; i < endPage; i++) {
                if (i == data.number) {
                    html.push(format(page_li, ['page-number active disabled', i, i + 1]));
                } else {
                    html.push(format(page_li, ['page-number', i, i + 1]));
                }
            }

            //render next actions
            var nextPage = data.number + 1;
            if (data.last) {
                html.push(format(page_li, ['page-next disabled', nextPage, '&gt;']));
                html.push(format(page_li, ['page-last disabled', data.totalPages - 1, '&gt;&gt;']));
            } else {
                html.push(format(page_li, ['page-next', nextPage, '&gt;']));
                html.push(format(page_li, ['page-last', data.totalPages - 1, '&gt;&gt;']));
            }
            html.push("</div>");
            $(pageSettings.paginationDiv).html(html.join(""));
        }
        var afterRenderPagination = function () {
            $(pageSettings.paginationDiv).find("li").not('.disabled').each(function () {
                $(this).bind("click", function () {
                    goToPage(this.value);
                });
            });
        }
        var format = function (s, a) {
            for (var i = 0; i < a.length; i++) {
                var v = a[i];
                if (v == null) {
                    v = '&nbsp;';
                }
                s = s.replace(new RegExp('\\{' + i + '\\}', 'g'), v);
            }
            return s;
        }
        var formatByObject = function (s, o) {
            for (x in o) {
                var v = o[x];
                if (v == null) {
                    v = '&nbsp;';
                }
                s = s.replace(new RegExp('\\{' + x + '\\}', 'g'), v);
            }
            return s;
        }
        goToPage(i);
    }
    return {
        initPage: function (settings) {
            initPage(settings);
        },
        refreshPage: function () {
            refreshPage();
        },
        showPage: function (i) {
            showPage(i);
        }
    };
};