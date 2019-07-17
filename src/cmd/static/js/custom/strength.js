/**
 * Created by zhongkw on 1/16/2015.
 */
;(function ( $, window, document, undefined ) {
    var pluginName = "strength",
        defaults = {
            strengthClass: 'strength',
            strengthMeterClass: 'strength_meter'
        };

    function Plugin( element, options ) {
        this.element = element;
        this.$elem = $(this.element);
        this.options = $.extend( {}, defaults, options );
        this.init();
    }

    Plugin.prototype = {
        init: function() {
            var characters = 0;
            var capitalLetters = 0;
            var loweLetters = 0;
            var number = 0;
            var special = 0;

            var upperCase= new RegExp('[A-Z]');
            var lowerCase= new RegExp('[a-z]');
            var numbers = new RegExp('[0-9]');
            var specialChars = new RegExp('([!,%,&,@,#,$,^,*,?,_,~])');

            function check_strength(thisval,thisid){
                if (thisval.length > 8) { characters = 1; } else { characters = 0; };
                if (thisval.match(upperCase)) { capitalLetters = 1} else { capitalLetters = 0; };
                if (thisval.match(lowerCase)) { loweLetters = 1}  else { loweLetters = 0; };
                if (thisval.match(numbers)) { number = 1}  else { number = 0; };
                if (thisval.match(specialChars)) { special = 1}  else { special = 0; };

                var total = characters + capitalLetters + loweLetters + number + special;

                get_total(total,thisid);
            }

            function get_total(total,thisid){
                var thismeter = $('div[data-meter="'+thisid+'"]');
                if(total == 0){
                    thismeter.removeClass();
                    $("#registerBtn").removeAttr('disabled');
                }else if (total <= 1) {
                    thismeter.removeClass();
                    thismeter.addClass('veryweak');
                    $("#registerBtn").attr('disabled', 'disabled');
                } else if (total == 2){
                    thismeter.removeClass();
                    thismeter.addClass('weak');
                    $("#registerBtn").attr('disabled', 'disabled');
                } else if(total == 3){
                    thismeter.removeClass();
                    thismeter.addClass('medium');
                    $("#registerBtn").removeAttr('disabled');
                } else {
                    thismeter.removeClass();
                    thismeter.addClass('strong');
                    $("#registerBtn").removeAttr('disabled');
                }
            }

            thisid = this.$elem.attr('id');

            this.$elem.addClass(this.options.strengthClass).attr('data-password',thisid).after('<div class="'+this.options.strengthMeterClass+'"><div data-meter="'+thisid+'"></div></div>');

            this.$elem.bind('keyup keydown', function(event) {
                thisval = $('#'+thisid).val();
                check_strength(thisval,thisid);
            });
        }
    };
    $.fn[pluginName] = function ( options ) {
        return this.each(function () {
            if (!$.data(this, "plugin_" + pluginName)) {
                $.data(this, "plugin_" + pluginName, new Plugin( this, options ));
            }
        });
    };
})( jQuery, window, document );