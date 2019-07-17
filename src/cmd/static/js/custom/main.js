/**
 * Created by nickly on 15/8/17.
 */
$(document).ready(function(){

    inputClick();
    classHidden();

    $(".web_classLeftBlockList_li").mouseenter(function(){

        $(".web_classLeftBlockList_li").removeClass("hover");
        $(this).addClass("hover");
        $(".web_classLeftBlock_hiddenBlock").fadeIn(100);
    })

    $(".web_classLeftBlockList_li").click(function(){
        $(".web_classLeftBlockList_li").removeClass("hover");
        $(this).addClass("hover");
        $(".web_classLeftBlock_hiddenBlock").fadeIn(100);
    })
    $(".web_classLeftBlock").bind("mouseleave",function(){
        $(".web_classLeftBlockList_li").removeClass("hover");
        $(".web_classLeftBlock_hiddenBlock").fadeOut(100);
    });
    $(".web_classLeftBlock_hiddenBlock").bind("click",function(){
        $(".web_classLeftBlockList_li").removeClass("hover");
        $(this).fadeOut(100);
    });



})

function classHidden(){

    if($(window).outerWidth() <= 640){
        $(".web_classLeftBlock_hiddenBlock").fadeIn(100);
        $(".web_classLeftBlock").unbind("mouseleave");
        $(".web_classLeftBlock_hiddenBlock").unbind("click");
        $(".web_classLeftBlockTitle").bind("click",function(){
            $(".web_classLeftBlock_hiddenBlock").fadeIn(100);
        })
    }else{
        $(".web_classLeftBlock").bind("mouseleave",function(){
            $(".web_classLeftBlockList_li").removeClass("hover");
            $(".web_classLeftBlock_hiddenBlock").fadeOut(100);
        });
        $(".web_classLeftBlock_hiddenBlock").bind("click",function(){
            $(".web_classLeftBlockList_li").removeClass("hover");
            $(this).fadeOut(100);
        });
        $(".web_classLeftBlockTitle").unbind("click");
    }
    $(window).resize(function(){
        if($(window).outerWidth() <= 640){
            $(".web_classLeftBlock_hiddenBlock").fadeIn(100);
            $(".web_classLeftBlock").unbind("mouseleave");
            $(".web_classLeftBlock_hiddenBlock").unbind("click");
            $(".web_classLeftBlockTitle").bind("click",function(){
                $(".web_classLeftBlock_hiddenBlock").fadeIn(100);
            })
        }else if($(window).outerWidth() > 640){
            $(".web_classLeftBlock_hiddenBlock").fadeOut(10);
            $(".web_classLeftBlock").bind("mouseleave",function(){
                $(".web_classLeftBlockList_li").removeClass("hover");
                $(".web_classLeftBlock_hiddenBlock").fadeOut(100);
            });
            $(".web_classLeftBlock_hiddenBlock").bind("click",function(){
                $(".web_classLeftBlockList_li").removeClass("hover");
                $(".web_classLeftBlock_hiddenBlock").fadeOut(100);
            });
            $(".web_classLeftBlockTitle").unbind("click");
        }
    })

}

function inputClick(obj){

    var temp;

    $(".startInput").focus(function(e){

        temp=e;

        if($(this).val()==$(this).attr('content')){
            $(this).val("");
        }
        $(".web_searchSelect").fadeIn(300);

        $(this).focusout(function(){
            if($(this).val()==""){
                $(this).val($(this).attr("content"));
            }
        })
    })


    $(".endInput").focus(function(e){

        temp=e;

        if($(this).val()==$(this).attr('content')){
            $(this).val("");
        }
        $(".web_searchSelect").fadeIn(300);

        $(this).focusout(function(){
            if($(this).val()==""){
                $(this).val($(this).attr("content"));
            }
        })


    })

    $(".web_searchSelect a").click(function(){

        if(temp.target.className=='endInput'){
            //console.log("33"+temp.target.className)
            $(".endInput").val($(this).html());

        }else if(temp.target.className=='startInput'){
            //console.log("44"+temp.target.className)
            $(".startInput").val($(this).html());
        }
        $(".web_searchSelect").fadeOut(300);
    })

    $(".web_searchBar").mouseleave(function(){
        $(".web_searchSelect").fadeOut(300);
    })

}