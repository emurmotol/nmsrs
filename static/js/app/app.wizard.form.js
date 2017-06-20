$(function () {
    $("#alfw").on("change", function () {
        $(this).val($(this).prop("checked"));
    });

    $("#i_accept").on("change", function () {
        $(this).val($(this).prop("checked"));
    });

    $("#birthdate").datetimepicker({
        viewMode: "years",
        format: "YYYY-MM-DD"
    });

    $("#birthdate").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#pned").datetimepicker({
        format: "YYYY-MM"
    });

    $("#pned").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#civil_stat_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=civil_stat_id]:checked");

        if ($(checked).val() == "5") {
            $("#civil_stat_other").attr("data-parsley-required", true);
            $("#civil_stat_other").prop("disabled", false);
            $("#civil_stat_other").focus();
        } else {
            $("#civil_stat_other").removeAttr("data-parsley-required");
            $("#civil_stat_other").val("");
            $("#civil_stat_other").parsley().reset();
            $("#civil_stat_other").prop("disabled", true);
        }
    });

    $("#emp_stat_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=emp_stat_id]:checked");

        if ($(checked).val() == "3") {
            $("#un_emp_stat_id").val(null).trigger("change");
            $("#un_emp_stat_id").attr("data-parsley-required", true);
            $("#un_emp_stat_id").prop("disabled", false);
            $("#un_emp_stat_id").focus();
        } else {
            $("#un_emp_stat_id").removeAttr("data-parsley-required");
            $("#un_emp_stat_id").val(null).trigger("change");
            $("#un_emp_stat_id").prop("disabled", true);
            $("#toc_id").removeAttr("data-parsley-required");
            $("#toc_id").val(null).trigger("change");
            $("#toc_id").prop("disabled", true);
        }
    });

    $("#disability_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=disability_id]:checked");

        if ($(checked).val() == "5") {
            $("#disability_other").attr("data-parsley-required", true);
            $("#disability_other").prop("disabled", false);
            $("#disability_other").focus();
        } else {
            $("#disability_other").removeAttr("data-parsley-required");
            $("#disability_other").val("");
            $("#disability_other").parsley().reset();
            $("#disability_other").prop("disabled", true);
        }
    });

    $("#disabled").on("change", function () {
        $(this).val($(this).prop("checked"));
        $("#disability_1").parsley().reset();

        if ($(this).prop("checked")) {
            $("#disability_1").attr("data-parsley-required", true);
            $("#disability_radios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", false);
            });
        } else {
            $("#disability_1").removeAttr("data-parsley-required");
            $("#disability_radios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", true);
                $(this).prop("checked", false);
            });
            $("#disability_other").removeAttr("data-parsley-required");
            $("#disability_other").val(null).trigger("change");
            $("#disability_other").prop("disabled", true);
        }
    });

    $("#skill_nl").on("change", function () {
        $(this).val($(this).prop("checked"));

        if ($(this).prop("checked")) {
            $("#skill_ids").val(null).trigger("change");
            $("#skill_ids").prop("disabled", true);
            $("#skill_other").attr("data-parsley-required", true);
            $("#skill_other").prop("disabled", false);
            $("#skill_other").focus();
        } else {
            $("#skill_other").removeAttr("data-parsley-required");
            $("#skill_other").val("");
            $("#skill_other").parsley().reset();
            $("#skill_other").prop("disabled", true);
            $("#skill_ids").prop("disabled", false);
            $("#skill_ids").focus();
        }
    });

    $("#registered_at").datetimepicker({
        defaultDate: new Date(),
        format: "YYYY-MM-DD"
    });

    $("#registered_at").on("dp.change", function() {
        $(this).parsley().validate();
    });

    previewImage($("#photo"));

    $("#create_registrant_form").parsley();
    $("#create_registrant_form").on("submit", function () {
        $("#prov_id").val($("#prov_id").data("id"));
        duringSubmitDo(this);
    });
});