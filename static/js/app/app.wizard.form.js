$(function () {
    $("#civil_stat_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=civil_stat_id]:checked");

        if ($(checked).val() == "5") {
            $("#civil_stat_other").prop("disabled", false);
            $("#civil_stat_other").prop("placeholder", "Simplify");
            $("#civil_stat_other").attr("data-parsley-required", "true");
            $("#civil_stat_other").focus();
        } else {
            $("#civil_stat_other").parsley().reset();
            $("#civil_stat_other").val("");
            $("#civil_stat_other").removeAttr("placeholder");
            $("#civil_stat_other").removeAttr("data-parsley-required");
            $("#civil_stat_other").prop("disabled", true);
        }
    });

    $("#emp_stat_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=emp_stat_id]:checked");

        if ($(checked).val() == "3") {
            $("#un_emp_stat_id").val(null).trigger("change");
            loadUnEmpStat();
            $("#un_emp_stat_id").prop("disabled", false);
            $("#un_emp_stat_id").attr("data-parsley-required", "true");
            $("#un_emp_stat_id").focus();
        } else {
            $("#un_emp_stat_id").parsley().reset();
            $("#un_emp_stat_id").select2();
            $("#un_emp_stat_id").removeAttr("data-parsley-required");
            $("#un_emp_stat_id").val(null).trigger("change");
            $("#un_emp_stat_id").prop("disabled", true);
            $("#toc_id").parsley().reset();
            $("#toc_id").select2();
            $("#toc_id").removeAttr("data-parsley-required");
            $("#toc_id").val(null).trigger("change");
            $("#toc_id").prop("disabled", true);
        }
    });

    $("#disability_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=disability_id]:checked");

        if ($(checked).val() == "5") {
            $("#disability_other").prop("disabled", false);
            $("#disability_other").prop("placeholder", "Simplify");
            $("#disability_other").attr("data-parsley-required", "true");
            $("#disability_other").focus();
        } else {
            $("#disability_other").val("");
            $("#disability_other").removeAttr("data-parsley-required");
            $("#disability_other").parsley().reset();
            $("#disability_other").prop("disabled", true);
            $("#disability_other").removeAttr("placeholder");
        }
    });

    $("#disabled").on("change", function () {
        $("#disability_1").parsley().reset();

        if ($(this).prop("checked")) {
            $("#disability_1").attr("data-parsley-required", "true");
            $("#disability_radios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", false);
            });
        } else {
            $("#disability_1").removeAttr("data-parsley-required");
            $("#disability_radios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", true);
                $(this).prop("checked", false);
            });
            $("#disability_other").parsley().reset();
            $("#disability_other").val("");
            $("#disability_other").removeAttr("data-parsley-required");
            $("#disability_other").prop("disabled", true);
            $("#disability_other").removeAttr("placeholder");
        }
    });

    $("#create_registrant_form").parsley();
    $("#create_registrant_form").on("submit", function () {
        duringSubmitDo(this);
    });
    previewImage($("#photo"));
});