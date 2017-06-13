$(function () {
    var select_all_formal_edu = $("#select_all_formal_edu");
    var delete_formal_edu_button = $("#delete_formal_edu_button");

    $("#formal_edu_form").parsley();
    $("#formal_edu_form").on("submit", function (e) {
        e.preventDefault();
        var school_univ_text = $("#school_univ_other").val().toUpperCase();
        var school_univ_id_val = $("#school_univ_id").select2("val") != null ? $("#school_univ_id").select2("val") : "";

        if (!$("#sunl").prop("checked")) {
            school_univ_text = $("#school_univ_id").select2("data")[0].text;
        }

        switch ($(this).attr("data-action")) {
            case "add":
                var formal_edu_index = 1 + $("#formal_edu_table tbody tr").length++;
                var row = `
                <tr data-index="` + formal_edu_index + `">
                    <td class="formal-edu-checkbox">
                        <input type="checkbox" class="checkbox" id="formal_edu_checkbox_` + formal_edu_index + `">
                    </td>
                    <td class="high-grade-comp">
                        <span>` + $("#high_grade_comp_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="high_grade_comp_id[]" value="` + $("#high_grade_comp_id").select2("val") + `">
                    </td>
                    <td class="course-degree">
                        <span>` + $("#course_degree_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="course_degree_id[]" value="` + $("#course_degree_id").select2("val") + `">
                    </td>
                    <td class="school-univ">
                        <span>` + school_univ_text + `</span>
                        <input type="hidden" name="school_univ_id[]" value="` + school_univ_id_val + `">
                        <input type="hidden" name="school_univ_other[]" value="` + $("#school_univ_other").val().toUpperCase() + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="formal-edu-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#formal_edu_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#formal_edu_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".high-grade-comp").find("span").text($("#high_grade_comp_id").select2("data")[0].text);
                tr.find(".high-grade-comp").find('input[name="high_grade_comp_id[]"]').val($("#high_grade_comp_id").select2("val"));
                tr.find(".course-degree").find("span").text($("#course_degree_id").select2("data")[0].text);
                tr.find(".course-degree").find('input[name="course_degree_id[]"]').val($("#course_degree_id").select2("val"));
                tr.find(".school-univ").find("span").text(school_univ_text);
                tr.find(".school-univ").find('input[name="school_univ_id[]"]').val(school_univ_id_val);
                tr.find(".school-univ").find('input[name="school_univ_other[]"]').val($("#school_univ_other").val());
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".formal-edu-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_formal_edu.prop("checked", false);
            }

            if ($(".formal-edu-checkbox input:checked").length == $(".formal-edu-checkbox input").length) {
                select_all_formal_edu.prop("checked", true);
            }

            if ($(".formal-edu-checkbox input:checked").length == 0) {
                delete_formal_edu_button.prop("disabled", true);
            } else {
                delete_formal_edu_button.prop("disabled", false);
            }
        });

        $(".formal-edu-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var high_grade_comp_id_val = tr.find(".high-grade-comp").find('input[name="high_grade_comp_id[]"]').val();
            var course_degree_id_val = tr.find(".course-degree").find('input[name="course_degree_id[]"]').val();
            school_univ_id_val = tr.find(".school-univ").find('input[name="school_univ_id[]"]').val();

            $("#high_grade_comp_id").val(parseInt(high_grade_comp_id_val)).trigger("change");
            $("#course_degree_id").val(parseInt(course_degree_id_val)).trigger("change");

            if (school_univ_id_val == "") {
                var school_univ_other_val = tr.find(".school-univ").find('input[name="school_univ_other[]"]').val();
                $("#sunl").prop("checked", true).trigger("change");
                $("#school_univ_other").val(school_univ_other_val).trigger("change");
            } else {
                $("#school_univ_id").val(parseInt(school_univ_id_val)).trigger("change");
            }
            $("#formal_edu_form").attr("data-edit-index", tr.data("index"));
            $("#formal_edu_form").attr("data-action", "edit");
            $("#formal_edu_modal").modal("show");
        });
        select_all_formal_edu.prop("checked", false);
        $("#formal_edu_modal").modal("hide");
    });

    select_all_formal_edu.on("change", function () {
        if ($(".formal-edu-checkbox input").length > 0) {
            $(".formal-edu-checkbox input").prop("checked", $(this).prop("checked"));
            delete_formal_edu_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_formal_edu_button.on("click", function () {
        $(".formal-edu-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_formal_edu.prop("checked", false);
            }
        });

        if ($("#formal_edu_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#formal_edu_modal").on("hidden.bs.modal", function () {
        $("#formal_edu_form").removeAttr("data-action");

        if ($("#school_univ_id").prop("disabled")) {
            $("#school_univ_id").prop("disabled", false);
            $("#school_univ_other").prop("disabled", true);
        }
        $("#high_grade_comp_id").removeAttr("data-parsley-required");
        $("#high_grade_comp_id").val(null).trigger("change");
        $("#high_grade_comp_id").attr("data-parsley-required", true);
        $("#course_degree_id").removeAttr("data-parsley-required");
        $("#course_degree_id").val(null).trigger("change");
        $("#course_degree_id").attr("data-parsley-required", true);
        $("#sunl").prop("checked", false).trigger("change");
    });

    $("#add_formal_edu_button").on("click", function () {
        $("#formal_edu_form").attr("data-action", "add");
        $("#formal_edu_modal").modal("show");
    });

    $("#sunl").on("change", function () {
        $("#school_univ_id").removeAttr("data-parsley-required");
        $("#school_univ_id").val(null).trigger("change");

        if ($(this).prop("checked")) {
            $("#school_univ_id").prop("disabled", true);
            $("#school_univ_other").attr("data-parsley-required", true);
            $("#school_univ_other").prop("disabled", false);
            $("#school_univ_other").focus();
        } else {
            $("#school_univ_other").removeAttr("data-parsley-required");
            $("#school_univ_other").val(null).trigger("change");
            $("#school_univ_other").prop("disabled", true);
            $("#school_univ_id").attr("data-parsley-required", true);
            $("#school_univ_id").prop("disabled", false);
            $("#school_univ_id").focus();
        }
    });
    var select_all_pro_license = $("#select_all_pro_license");
    var delete_pro_license_button = $("#delete_pro_license_button");

    $("#pro_license_form").parsley();
    $("#pro_license_form").on("submit", function (e) {
        e.preventDefault();
        var plt_id_val = $("#plt_id").select2("val");

        switch ($(this).attr("data-action")) {
            case "add":
                var pro_license_index = 1 + $("#pro_license_table tbody tr").length++;
                var row = `
                <tr data-index="` + pro_license_index + `">
                    <td class="pro-license-checkbox">
                        <input type="checkbox" class="checkbox" id="pro_license_checkbox_` + pro_license_index + `">
                    </td>
                    <td class="plt">
                        <span>` + $("#plt_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="plt_id[]" value="` + $("#plt_id").select2("val") + `">
                    </td>
                    <td class="pled">
                        <span>` + $("#pled").val().toUpperCase() + `</span>
                        <input type="hidden" name="pled[]" value="` + $("#pled").val().toUpperCase() + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="pro-license-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#pro_license_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#pro_license_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".plt").find("span").text($("#plt_id").select2("data")[0].text);
                tr.find(".plt").find('input[name="plt_id[]"]').val($("#plt_id").select2("val"));
                tr.find(".pled").find("span").text($("#pled").val());
                tr.find(".pled").find('input[name="pled[]"]').val($("#pled").val());
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".pro-license-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_pro_license.prop("checked", false);
            }

            if ($(".pro-license-checkbox input:checked").length == $(".pro-license-checkbox input").length) {
                select_all_pro_license.prop("checked", true);
            }

            if ($(".pro-license-checkbox input:checked").length == 0) {
                delete_pro_license_button.prop("disabled", true);
            } else {
                delete_pro_license_button.prop("disabled", false);
            }
        });

        $(".pro-license-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var plt_id_val = tr.find(".plt").find('input[name="plt_id[]"]').val();
            var pled_val = tr.find(".pled").find('input[name="pled[]"]').val();

            $("#plt_id").val(parseInt(plt_id_val)).trigger("change");
            $("#pled").val(pled_val).trigger("change");
            $("#pro_license_form").attr("data-edit-index", tr.data("index"));
            $("#pro_license_form").attr("data-action", "edit");
            $("#pro_license_modal").modal("show");
        });
        select_all_pro_license.prop("checked", false);
        $("#pro_license_modal").modal("hide");
    });

    select_all_pro_license.on("change", function () {
        if ($(".pro-license-checkbox input").length > 0) {
            $(".pro-license-checkbox input").prop("checked", $(this).prop("checked"));
            delete_pro_license_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_pro_license_button.on("click", function () {
        $(".pro-license-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_pro_license.prop("checked", false);
            }
        });

        if ($("#pro_license_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#pro_license_modal").on("hidden.bs.modal", function () {
        $("#pro_license_form").removeAttr("data-action");
        $("#plt_id").removeAttr("data-parsley-required");
        $("#plt_id").val(null).trigger("change");
        $("#plt_id").attr("data-parsley-required", true);
        $("#pled").val("");
        $("#pled").parsley().reset();
    });

    $("#add_pro_license_button").on("click", function () {
        $("#pro_license_form").attr("data-action", "add");
        $("#pro_license_modal").modal("show");
    });
});