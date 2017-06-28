$(function () {
    var tabPane = $(".tab-pane");

    function navigateTo(index) {
        $(".form-nav .prev-tab").toggle(index > 0);
        var atTheEnd = index >= tabPane.length - 1;
        $(".form-nav .next-tab").toggle(!atTheEnd);
        $(".form-nav [type=submit]").toggle(atTheEnd);
    }

    function activeIndex() {
        return tabPane.index(tabPane.filter(".active"));
    }

    $(".form-nav .prev-tab").on("click", function () {
        navigateTo(activeIndex() - 1);
        $(".wizard .nav-pills li.active").prev().find('a[data-toggle="tab"]').tab("show");
    });

    $(".form-nav .next-tab").on("click", function () {
        if ($("#createRegistrantForm").parsley("").validate("block-" + activeIndex())) {
            navigateTo(activeIndex() + 1);
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