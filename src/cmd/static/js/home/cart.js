/**
 * Created by silence yuan on 2015/9/1.
 */

renderShorpCartTable();

function renderShorpCartTable(){

    var template = $("#cart").html();
    var req = $.ajax({
        url: '/homepage/shopCart',
        type: 'post'
    });
    req.done(function (data) {
        var html = new Array();
       // alert(template);
        //alert(JSON.stringify(data));
        if (data) {
           // ;
            $("#cartItemCount").html(data.itemCount);
            //  $("#recentMessageCount").html(data.recentCount);
            data.cartItems.forEach(function(row){
                html.push(formatByObject(template, row));

            });
        }


        $("#cart").html(html.join(""));

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


