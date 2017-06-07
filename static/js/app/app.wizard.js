$(function () {
    setCheckboxBoolValue($("#alfw"));

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

    $("#civil_stat_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=civil_stat_id]:checked");

        if ($(checked).val() == "5") {
            $("#civil_stat_other").prop("disabled", false);
            $("#civil_stat_other").prop("placeholder", "Simplify");
            $("#civil_stat_other").focus();
        } else {
            $("#civil_stat_other").val("");
            $("#civil_stat_other").prop("disabled", true);
            $("#civil_stat_other").removeAttr("placeholder");
        }
    });

    $("#emp_stat_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=emp_stat_id]:checked");

        if ($(checked).val() == "3") {
            $("#un_emp_stat_id").val("");
            loadUnEmpStat();
            $("#un_emp_stat_id").prop("disabled", false);
            $("#un_emp_stat_id").focus();
        } else {
            $("#un_emp_stat_id").val("");
            $("#un_emp_stat_id").select2();
            $("#un_emp_stat_id").prop("disabled", true);
            $("#toc_id").val("");
            $("#toc_id").select2();
            $("#toc_id").prop("disabled", true);
        }
    });

    $("#disability_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=disability_id]:checked");

        if ($(checked).val() == "5") {
            $("#disability_other").prop("disabled", false);
            $("#disability_other").prop("placeholder", "Simplify");
            $("#disability_other").focus();
        } else {
            $("#disability_other").val("");
            $("#disability_other").prop("disabled", true);
            $("#disability_other").removeAttr("placeholder");
        }
    });

    $("#disabled").on("change", function () {
        if ($(this).prop("checked")) {
            $("#disability_radios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", false);
            });
        } else {
            $("#disability_radios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", true);
                $(this).prop("checked", false);
            });
            $("#disability_other").val("");
            $("#disability_other").prop("disabled", true);
            $("#disability_other").removeAttr("placeholder");
        }
    });
});