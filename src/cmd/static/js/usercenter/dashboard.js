/**
 * Created by silence yuan on 2015/9/1.
 */

var template = $("#messages").html();

renderShorpCartTable();

function renderShorpCartTable(){
    $("#cart").html();
    //alert($("#cart").html());
    var req = $.ajax({
        url: '/usercenter/coaldeal/topList',
        type: 'post'
    });
    req.done(function (data) {
      //  alert(JSON.stringify(data));
        var html = new Array();
        if (data) {
          //  $("#recentMessageCount").html(data.recentCount);
            data.content.forEach(function(row){
                html.push(formatByObject(template, row));

            });
        }


        $("#cart").html(html.join(""));

    });
};


function renderTruckTable(){
    $("#messages").html();
    var url = $("#messageUri").val();
   // alert(url);
    var req = $.ajax({
        url:  '/usercenter/message/top-message-List',
        type: 'post'
    });
    req.done(function (data) {
        var html = new Array();
        if (data) {
            $("#recentMessageCount").html(data.recentCount);
            data.items.forEach(function(row){
                html.push(formatByObject(template, row));

            });
        }


        $("#messages").html(html.join(""));

    });
};



var formatByObject = function (s, o) {
    for (x in o) {
        var v = o[x];
        if (v == null) {
            v = 'fghnjmkl';
        }
        s = s.replace(new RegExp('\\{' + x + '\\}', 'g'), v);
    }
    return s;
};

renderTruckTable();

var isMessageLoad = false;
$("#messageDropdown").on("click", function() {
    if(!isMessageLoad) {
        isMessageLoad = true;
    }

})