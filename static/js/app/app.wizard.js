$(function () {
    var tabPane = $(".tab-pane");

    function navigateTo(index) {
        tabPane.removeClass("current").eq(index).addClass("current");
        $(".form-nav .prev-tab").toggle(index > 0);
        var at_the_end = index >= tabPane.length - 1;
        $(".form-nav .next-tab").toggle(!at_the_end);
        $(".form-nav [type=submit]").toggle(at_the_end);
    }

    function currentIndex() {
        return tabPane.index(tabPane.filter(".current"));
    }

    $(".form-nav .prev-tab").click(function () {
        navigateTo(currentIndex() - 1);
        $(".wizard .nav-pills li.active").prev().find('a[data-toggle="tab"]').tab("show");
    });

    $(".form-nav .next-tab").click(function () {
        if ($("#create_registrant_form").parsley().validate("block-" + currentIndex())) {
            navigateTo(currentIndex() + 1);

            var li = $(".wizard .nav-pills li.active");
            li.next().removeClass("disabled");
            li.next().find('a[data-toggle="tab"]').tab("show");
        }
    });

    tabPane.each(function (index, pane) {
        $(pane).find(":input").attr("data-parsley-group", "block-" + index);
    });
    navigateTo(0);

    $(".wizard ul li").on("click", function () {
        if ($(this).hasClass("disabled")) {
            return false;
        }
    });
});