$(function () {
    nextStep = function () {
        var li = $(".wizard .nav-pills li.active");
        li.next().removeClass("disabled");
        li.next().find('a[data-toggle="tab"]').tab("show");
    }

    prevStep = function () {
        $(".wizard .nav-pills li.active").prev().find('a[data-toggle="tab"]').tab("show");
    }

    $(".wizard ul li").on("click", function () {
        if ($(this).hasClass("disabled")) {
            return false;
        }
    });

    $("#civil_status_5").on("click", function () {
        $("#civil_status").prop("disabled", false);
        $("#civil_status").focus();
    });

    $("#civil_status_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=civil_status]:checked");

        if ($(checked).val() == "5") {
            $("input[name=civil_status_5]").prop("disabled", false);
            $("input[name=civil_status_5]").focus();
        } else {
            $("input[name=civil_status_5]").prop("disabled", true);
            $("input[name=civil_status_5]").val("");
        }
    });

    $("#disability_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=disability]:checked");

        if ($(checked).val() == "5") {
            $("input[name=disability_5]").prop("disabled", false);
            $("input[name=disability_5]").focus();
        } else {
            $("input[name=disability_5]").prop("disabled", true);
            $("input[name=disability_5]").val("");
        }
    });

    $("#disabled").on("change", function () {
        if ($(this).prop("checked")) {
            $("#disability_radios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", false);
            });
        } else {
            $("#disability_radios").find("input[type=radio]").each(function () {
                $("input[name=disability_5]").prop("disabled", true);
                $(this).prop("disabled", true);
                $(this).prop("checked", false);
            });
        }
    });
});