$(function () {
    var select_all_formal_edu = $("#select_all_formal_edu");
    var delete_formal_edu_button = $("#delete_formal_edu_button");

    $("#formal_edu_yg").datetimepicker({
        viewMode: "years",
        format: "YYYY"
    });

    $("#formal_edu_yg").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#formal_edu_la").datetimepicker({
        format: "YYYY-MM"
    });

    $("#formal_edu_la").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#formal_edu_form").parsley();
    $("#formal_edu_form").on("submit", function (e) {
        e.preventDefault();
        var formal_edu_hgc_id_val = $("#formal_edu_hgc_id").select2("val");
        var formal_edu_cd_id_val = $("#formal_edu_cd_id").select2("val");
        var formal_edu_su_id_val = $("#formal_edu_su_id").select2("val") != null ? $("#formal_edu_su_id").select2("val") : "";
        var formal_edu_su_other_val = $("#formal_edu_su_other").val().toUpperCase();
        var formal_edu_su_text = formal_edu_su_other_val;
        var formal_edu_yg_val = $("#formal_edu_yg").val();
        var formal_edu_la_val = $("#formal_edu_la").val();

        if (!$("#formal_edu_sunl").prop("checked")) {
            formal_edu_su_text = $("#formal_edu_su_id").select2("data")[0].text;
        }

        switch ($(this).attr("data-action")) {
            case "add":
                var formal_edu_index = 1 + $("#formal_edu_table tbody tr").length++;
                var row = `
                <tr data-index="` + formal_edu_index + `">
                    <td class="formal-edu-checkbox">
                        <input type="checkbox" class="checkbox" id="formal_edu_checkbox_` + formal_edu_index + `">
                    </td>
                    <td class="formal-edu-hgc">
                        <span>` + $("#formal_edu_hgc_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="formal_edu_hgc_id" value="` + formal_edu_hgc_id_val + `">
                    </td>
                    <td class="formal-edu-cd">
                        <span>` + $("#formal_edu_cd_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="formal_edu_cd_id" value="` + formal_edu_cd_id_val + `">
                    </td>
                    <td class="formal-edu-su">
                        <span>` + formal_edu_su_text + `</span>
                        <input type="hidden" name="formal_edu_su_id" value="` + formal_edu_su_id_val + `">
                        <input type="hidden" name="formal_edu_su_other" value="` + formal_edu_su_other_val + `">
                    </td>
                    <td class="formal-edu-yg">
                        <span>` + formal_edu_yg_val + `</span>
                        <input type="hidden" name="formal_edu_yg" value="` + formal_edu_yg_val + `">
                    </td>
                    <td class="formal-edu-la">
                        <span>` + formal_edu_la_val + `</span>
                        <input type="hidden" name="formal_edu_la" value="` + formal_edu_la_val + `">
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
                tr.find(".formal-edu-hgc").find("span").text($("#formal_edu_hgc_id").select2("data")[0].text);
                tr.find(".formal-edu-hgc").find('input[name="formal_edu_hgc_id"]').val(formal_edu_hgc_id_val);
                tr.find(".formal-edu-cd").find("span").text($("#formal_edu_cd_id").select2("data")[0].text);
                tr.find(".formal-edu-cd").find('input[name="formal_edu_cd_id"]').val(formal_edu_cd_id_val);
                tr.find(".formal-edu-su").find("span").text(formal_edu_su_text);
                tr.find(".formal-edu-su").find('input[name="formal_edu_su_id"]').val(formal_edu_su_id_val);
                tr.find(".formal-edu-su").find('input[name="formal_edu_su_other"]').val(formal_edu_su_other_val);
                tr.find(".formal-edu-yg").find("span").text(formal_edu_yg_val);
                tr.find(".formal-edu-yg").find('input[name="formal_edu_yg"]').val(formal_edu_yg_val);
                tr.find(".formal-edu-la").find("span").text(formal_edu_la_val);
                tr.find(".formal-edu-la").find('input[name="formal_edu_la"]').val(formal_edu_la_val);
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
            var formal_edu_hgc_id_val = tr.find(".formal-edu-hgc").find('input[name="formal_edu_hgc_id"]').val();
            var formal_edu_cd_id_val = tr.find(".formal-edu-cd").find('input[name="formal_edu_cd_id"]').val();
            var formal_edu_su_id_val = tr.find(".formal-edu-su").find('input[name="formal_edu_su_id"]').val();
            var formal_edu_yg_val = tr.find(".formal-edu-yg").find('input[name="formal_edu_yg"]').val();
            var formal_edu_la_val = tr.find(".formal-edu-la").find('input[name="formal_edu_la"]').val();

            $("#formal_edu_hgc_id").val(parseInt(formal_edu_hgc_id_val)).trigger("change");
            $("#formal_edu_cd_id").val(parseInt(formal_edu_cd_id_val)).trigger("change");

            if (formal_edu_su_id_val == "") {
                var formal_edu_su_other_val = tr.find(".formal-edu-su").find('input[name="formal_edu_su_other"]').val();
                $("#formal_edu_sunl").prop("checked", true).trigger("change");
                $("#formal_edu_su_other").val(formal_edu_su_other_val).trigger("change");
            } else {
                $("#formal_edu_su_id").val(parseInt(formal_edu_su_id_val)).trigger("change");
            }
            $("#formal_edu_yg").val(formal_edu_yg_val).trigger("change");
            $("#formal_edu_la").val(formal_edu_la_val).trigger("change");
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
        $("#formal_edu_su_id").removeAttr("data-parsley-required");
        $("#formal_edu_su_id").val(null).trigger("change");
        $("#formal_edu_su_id").attr("data-parsley-required", true);

        if ($("#formal_edu_su_id").prop("disabled")) {
            $("#formal_edu_su_id").prop("disabled", false);
            $("#formal_edu_su_other").prop("disabled", true);
        }
        $("#formal_edu_hgc_id").removeAttr("data-parsley-required");
        $("#formal_edu_hgc_id").val(null).trigger("change");
        $("#formal_edu_hgc_id").attr("data-parsley-required", true);
        $("#formal_edu_cd_id").removeAttr("data-parsley-required");
        $("#formal_edu_cd_id").val(null).trigger("change");
        $("#formal_edu_cd_id").attr("data-parsley-required", true);
        $("#formal_edu_sunl").prop("checked", false).trigger("change");
        $("#formal_edu_yg").val("");
        $("#formal_edu_yg").parsley().reset();
        $("#formal_edu_la").val("");
        $("#formal_edu_la").parsley().reset();
    });

    $("#add_formal_edu_button").on("click", function () {
        $("#formal_edu_form").attr("data-action", "add");
        $("#formal_edu_modal").modal("show");
    });

    $("#formal_edu_sunl").on("change", function () {
        $(this).val($(this).prop("checked"));

        if ($(this).prop("checked")) {
            $("#formal_edu_su_id").removeAttr("data-parsley-required");
            $("#formal_edu_su_id").val(null).trigger("change");
            $("#formal_edu_su_id").prop("disabled", true);
            $("#formal_edu_su_other").attr("data-parsley-required", true);
            $("#formal_edu_su_other").prop("disabled", false);
            $("#formal_edu_su_other").focus();
        } else {
            $("#formal_edu_su_other").removeAttr("data-parsley-required");
            $("#formal_edu_su_other").val("");
            $("#formal_edu_su_other").parsley().reset();
            $("#formal_edu_su_other").prop("disabled", true);
            $("#formal_edu_su_id").attr("data-parsley-required", true);
            $("#formal_edu_su_id").prop("disabled", false);
            $("#formal_edu_su_id").focus();
        }
    });
    var select_all_pro_license = $("#select_all_pro_license");
    var delete_pro_license_button = $("#delete_pro_license_button");

    $("#pro_license_ed").datetimepicker({
        format: "YYYY-MM"
    });

    $("#pro_license_ed").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#pro_license_form").parsley();
    $("#pro_license_form").on("submit", function (e) {
        e.preventDefault();
        var pro_license_t_id_val = $("#pro_license_t_id").select2("val");
        var pro_license_ed_val = $("#pro_license_ed").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var pro_license_index = 1 + $("#pro_license_table tbody tr").length++;
                var row = `
                <tr data-index="` + pro_license_index + `">
                    <td class="pro-license-checkbox">
                        <input type="checkbox" class="checkbox" id="pro_license_checkbox_` + pro_license_index + `">
                    </td>
                    <td class="pro-license-t">
                        <span>` + $("#pro_license_t_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="pro_license_t_id" value="` + pro_license_t_id_val + `">
                    </td>
                    <td class="pro-license-ed">
                        <span>` + pro_license_ed_val + `</span>
                        <input type="hidden" name="pro_license_ed" value="` + pro_license_ed_val + `">
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
                tr.find(".pro-license-t").find("span").text($("#pro_license_t_id").select2("data")[0].text);
                tr.find(".pro-license-t").find('input[name="pro_license_t_id"]').val(pro_license_t_id_val);
                tr.find(".pro-license-ed").find("span").text(pro_license_ed_val);
                tr.find(".pro-license-ed").find('input[name="pro_license_ed"]').val(pro_license_ed_val);
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
            var pro_license_t_id_val = tr.find(".pro-license-t").find('input[name="pro_license_t_id"]').val();
            var pro_license_ed_val = tr.find(".pro-license-ed").find('input[name="pro_license_ed"]').val();

            $("#pro_license_t_id").val(parseInt(pro_license_t_id_val)).trigger("change");
            $("#pro_license_ed").val(pro_license_ed_val).trigger("change");
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
        $("#pro_license_t_id").removeAttr("data-parsley-required");
        $("#pro_license_t_id").val(null).trigger("change");
        $("#pro_license_t_id").attr("data-parsley-required", true);
        $("#pro_license_ed").val("");
        $("#pro_license_ed").parsley().reset();
    });

    $("#add_pro_license_button").on("click", function () {
        $("#pro_license_form").attr("data-action", "add");
        $("#pro_license_modal").modal("show");
    });
    var select_all_elig = $("#select_all_elig");
    var delete_elig_button = $("#delete_elig_button");

    $("#elig_yt").datetimepicker({
        viewMode: "years",
        format: "YYYY-MM"
    });

    $("#elig_yt").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#elig_form").parsley();
    $("#elig_form").on("submit", function (e) {
        e.preventDefault();
        var elig_t_id_val = $("#elig_t_id").select2("val");
        var elig_yt_val = $("#elig_yt").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var elig_index = 1 + $("#elig_table tbody tr").length++;
                var row = `
                <tr data-index="` + elig_index + `">
                    <td class="elig-checkbox">
                        <input type="checkbox" class="checkbox" id="elig_checkbox_` + elig_index + `">
                    </td>
                    <td class="elig-t">
                        <span>` + $("#elig_t_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="elig_t_id" value="` + elig_t_id_val + `">
                    </td>
                    <td class="elig-yt">
                        <span>` + elig_yt_val + `</span>
                        <input type="hidden" name="elig_yt" value="` + elig_yt_val + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="elig-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#elig_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#elig_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".elig-t").find("span").text($("#elig_t_id").select2("data")[0].text);
                tr.find(".elig-t").find('input[name="elig_t_id"]').val(elig_t_id_val);
                tr.find(".elig-yt").find("span").text(elig_yt_val);
                tr.find(".elig-yt").find('input[name="elig_yt"]').val(elig_yt_val);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".elig-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_elig.prop("checked", false);
            }

            if ($(".elig-checkbox input:checked").length == $(".elig-checkbox input").length) {
                select_all_elig.prop("checked", true);
            }

            if ($(".elig-checkbox input:checked").length == 0) {
                delete_elig_button.prop("disabled", true);
            } else {
                delete_elig_button.prop("disabled", false);
            }
        });

        $(".elig-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var elig_t_id_val = tr.find(".elig-t").find('input[name="elig_t_id"]').val();
            var elig_yt_val = tr.find(".elig-yt").find('input[name="elig_yt"]').val().toUpperCase();

            $("#elig_t_id").val(parseInt(elig_t_id_val)).trigger("change");
            $("#elig_yt").val(elig_yt_val).trigger("change");
            $("#elig_form").attr("data-edit-index", tr.data("index"));
            $("#elig_form").attr("data-action", "edit");
            $("#elig_modal").modal("show");
        });
        select_all_elig.prop("checked", false);
        $("#elig_modal").modal("hide");
    });

    select_all_elig.on("change", function () {
        if ($(".elig-checkbox input").length > 0) {
            $(".elig-checkbox input").prop("checked", $(this).prop("checked"));
            delete_elig_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_elig_button.on("click", function () {
        $(".elig-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_elig.prop("checked", false);
            }
        });

        if ($("#elig_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#elig_modal").on("hidden.bs.modal", function () {
        $("#elig_form").removeAttr("data-action");
        $("#elig_t_id").removeAttr("data-parsley-required");
        $("#elig_t_id").val(null).trigger("change");
        $("#elig_t_id").attr("data-parsley-required", true);
        $("#elig_yt").val("");
        $("#elig_yt").parsley().reset();
    });

    $("#add_elig_button").on("click", function () {
        $("#elig_form").attr("data-action", "add");
        $("#elig_modal").modal("show");
    });
    var select_all_vttare = $("#select_all_vttare");
    var delete_vttare_button = $("#delete_vttare_button");

    $("#vttare_form").parsley();
    $("#vttare_form").on("submit", function (e) {
        e.preventDefault();
        var vttare_not_val = $("#vttare_not").val().toUpperCase();
        var vttare_sa_val = $("#vttare_sa").val().toUpperCase();
        var vttare_pote_val = $("#vttare_pote").val().toUpperCase();
        var vttare_cr_val = $("#vttare_cr").val().toUpperCase();
        var vttare_isa_val = $("#vttare_isa").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var vttare_index = 1 + $("#vttare_table tbody tr").length++;
                var row = `
                <tr data-index="` + vttare_index + `">
                    <td class="vttare-checkbox">
                        <input type="checkbox" class="checkbox" id="vttare_checkbox_` + vttare_index + `">
                    </td>
                    <td class="vttare-not">
                        <span>` + vttare_not_val + `</span>
                        <input type="hidden" name="vttare_not" value="` + vttare_not_val + `">
                    </td>
                    <td class="vttare-sa">
                        <span>` + vttare_sa_val + `</span>
                        <input type="hidden" name="vttare_sa" value="` + vttare_sa_val + `">
                    </td>
                    <td class="vttare-pote">
                        <span>` + vttare_pote_val + `</span>
                        <input type="hidden" name="vttare_pote" value="` + vttare_pote_val + `">
                    </td>
                    <td class="vttare-cr">
                        <span>` + vttare_cr_val + `</span>
                        <input type="hidden" name="vttare_cr" value="` + vttare_cr_val + `">
                    </td>
                    <td class="vttare-isa">
                        <span>` + vttare_isa_val + `</span>
                        <input type="hidden" name="vttare_isa" value="` + vttare_isa_val + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="vttare-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#vttare_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#vttare_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".vttare-not").find("span").text(vttare_not_val);
                tr.find(".vttare-not").find('input[name="vttare_not"]').val(vttare_not_val);
                tr.find(".vttare-sa").find("span").text(vttare_sa_val);
                tr.find(".vttare-sa").find('input[name="vttare_sa"]').val(vttare_sa_val);
                tr.find(".vttare-pote").find("span").text(vttare_pote_val);
                tr.find(".vttare-pote").find('input[name="vttare_pote"]').val(vttare_pote_val);
                tr.find(".vttare-cr").find("span").text(vttare_cr_val);
                tr.find(".vttare-cr").find('input[name="vttare_cr"]').val(vttare_cr_val);
                tr.find(".vttare-isa").find("span").text(vttare_isa_val);
                tr.find(".vttare-isa").find('input[name="vttare_isa"]').val(vttare_isa_val);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".vttare-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_vttare.prop("checked", false);
            }

            if ($(".vttare-checkbox input:checked").length == $(".vttare-checkbox input").length) {
                select_all_vttare.prop("checked", true);
            }

            if ($(".vttare-checkbox input:checked").length == 0) {
                delete_vttare_button.prop("disabled", true);
            } else {
                delete_vttare_button.prop("disabled", false);
            }
        });

        $(".vttare-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var vttare_not_val = tr.find(".vttare-not").find('input[name="vttare_not"]').val();
            var vttare_sa_val = tr.find(".vttare-sa").find('input[name="vttare_sa"]').val();
            var vttare_pote_val = tr.find(".vttare-pote").find('input[name="vttare_pote"]').val();
            var vttare_cr_val = tr.find(".vttare-cr").find('input[name="vttare_cr"]').val();
            var vttare_isa_val = tr.find(".vttare-isa").find('input[name="vttare_isa"]').val();

            $("#vttare_not").val(vttare_not_val).trigger("change");
            $("#vttare_sa").val(vttare_sa_val).trigger("change");
            $("#vttare_pote").val(vttare_pote_val).trigger("change");
            $("#vttare_cr").val(vttare_cr_val).trigger("change");
            $("#vttare_isa").val(vttare_isa_val).trigger("change");
            $("#vttare_form").attr("data-edit-index", tr.data("index"));
            $("#vttare_form").attr("data-action", "edit");
            $("#vttare_modal").modal("show");
        });
        select_all_vttare.prop("checked", false);
        $("#vttare_modal").modal("hide");
    });

    select_all_vttare.on("change", function () {
        if ($(".vttare-checkbox input").length > 0) {
            $(".vttare-checkbox input").prop("checked", $(this).prop("checked"));
            delete_vttare_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_vttare_button.on("click", function () {
        $(".vttare-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_vttare.prop("checked", false);
            }
        });

        if ($("#vttare_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#vttare_modal").on("hidden.bs.modal", function () {
        $("#vttare_form").removeAttr("data-action");
        $("#vttare_not").val("");
        $("#vttare_not").parsley().reset();
        $("#vttare_sa").val("");
        $("#vttare_sa").parsley().reset();
        $("#vttare_pote").val("");
        $("#vttare_pote").parsley().reset();
        $("#vttare_cr").val("");
        $("#vttare_cr").parsley().reset();
        $("#vttare_isa").val("");
        $("#vttare_isa").parsley().reset();
    });

    $("#add_vttare_button").on("click", function () {
        $("#vttare_form").attr("data-action", "add");
        $("#vttare_modal").modal("show");
    });
    var select_all_coc = $("#select_all_coc");
    var delete_coc_button = $("#delete_coc_button");

    $("#coc_di").datetimepicker({
        format: "YYYY-MM"
    });

    $("#coc_di").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#coc_form").parsley();
    $("#coc_form").on("submit", function (e) {
        e.preventDefault();
        var coc_t_id_val = $("#coc_t_id").select2("val");
        var coc_r_val = $("#coc_r").val().toUpperCase();
        var coc_ib_val = $("#coc_ib").val().toUpperCase();
        var coc_di_val = $("#coc_di").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var coc_index = 1 + $("#coc_table tbody tr").length++;
                var row = `
                <tr data-index="` + coc_index + `">
                    <td class="coc-checkbox">
                        <input type="checkbox" class="checkbox" id="coc_checkbox_` + coc_index + `">
                    </td>
                    <td class="coc-t">
                        <span>` + $("#coc_t_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="coc_t_id" value="` + coc_t_id_val + `">
                    </td>
                    <td class="coc-r">
                        <span>` + coc_r_val + `</span>
                        <input type="hidden" name="coc_r" value="` + coc_r_val + `">
                    </td>
                    <td class="coc-ib">
                        <span>` + coc_ib_val + `</span>
                        <input type="hidden" name="coc_ib" value="` + coc_ib_val + `">
                    </td>
                    <td class="coc-di">
                        <span>` + coc_di_val + `</span>
                        <input type="hidden" name="coc_di" value="` + coc_di_val + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="coc-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#coc_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#coc_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".coc-t").find("span").text($("#coc_t_id").select2("data")[0].text);
                tr.find(".coc-t").find('input[name="coc_t_id"]').val(coc_t_id_val);
                tr.find(".coc-r").find("span").text(coc_r_val);
                tr.find(".coc-r").find('input[name="coc_r"]').val(coc_r_val);
                tr.find(".coc-ib").find("span").text(coc_ib_val);
                tr.find(".coc-ib").find('input[name="coc_ib"]').val(coc_ib_val);
                tr.find(".coc-di").find("span").text(coc_di_val);
                tr.find(".coc-di").find('input[name="coc_di"]').val(coc_di_val);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".coc-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_coc.prop("checked", false);
            }

            if ($(".coc-checkbox input:checked").length == $(".coc-checkbox input").length) {
                select_all_coc.prop("checked", true);
            }

            if ($(".coc-checkbox input:checked").length == 0) {
                delete_coc_button.prop("disabled", true);
            } else {
                delete_coc_button.prop("disabled", false);
            }
        });

        $(".coc-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var coc_t_id_val = tr.find(".coc-t").find('input[name="coc_t_id"]').val();
            var coc_r_val = tr.find(".coc-r").find('input[name="coc_r"]').val().toUpperCase();
            var coc_ib_val = tr.find(".coc-ib").find('input[name="coc_ib"]').val().toUpperCase();
            var coc_di_val = tr.find(".coc-di").find('input[name="coc_di"]').val().toUpperCase();

            $("#coc_t_id").val(parseInt(coc_t_id_val)).trigger("change");
            $("#coc_r").val(coc_r_val).trigger("change");
            $("#coc_ib").val(coc_ib_val).trigger("change");
            $("#coc_di").val(coc_di_val).trigger("change");
            $("#coc_form").attr("data-edit-index", tr.data("index"));
            $("#coc_form").attr("data-action", "edit");
            $("#coc_modal").modal("show");
        });
        select_all_coc.prop("checked", false);
        $("#coc_modal").modal("hide");
    });

    select_all_coc.on("change", function () {
        if ($(".coc-checkbox input").length > 0) {
            $(".coc-checkbox input").prop("checked", $(this).prop("checked"));
            delete_coc_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_coc_button.on("click", function () {
        $(".coc-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_coc.prop("checked", false);
            }
        });

        if ($("#coc_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#coc_modal").on("hidden.bs.modal", function () {
        $("#coc_form").removeAttr("data-action");
        $("#coc_t_id").removeAttr("data-parsley-required");
        $("#coc_t_id").val(null).trigger("change");
        $("#coc_t_id").attr("data-parsley-required", true);
        $("#coc_r").val("");
        $("#coc_r").parsley().reset();
        $("#coc_ib").val("");
        $("#coc_ib").parsley().reset();
        $("#coc_di").val("");
        $("#coc_di").parsley().reset();
    });

    $("#add_coc_button").on("click", function () {
        $("#coc_form").attr("data-action", "add");
        $("#coc_modal").modal("show");
    });
    var select_all_work_exp = $("#select_all_work_exp");
    var delete_work_exp_button = $("#delete_work_exp_button");

    $("#work_exp_f").datetimepicker({
        format: "YYYY-MM"
    });

    $("#work_exp_f").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#work_exp_t").datetimepicker({
        format: "YYYY-MM"
    });

    $("#work_exp_t").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#work_exp_rtfe").on("change", function () {
        $(this).val($(this).prop("checked"));
    });

    $("#work_exp_form").parsley();
    $("#work_exp_form").on("submit", function (e) {
        e.preventDefault();
        var work_exp_nocf_val = $("#work_exp_nocf").val().toUpperCase();
        var work_exp_a_val = $("#work_exp_a").val().toUpperCase();
        var work_exp_ph_id_val = $("#work_exp_ph_id").select2("val");
        var work_exp_f_val = $("#work_exp_f").val().toUpperCase();
        var work_exp_t_val = $("#work_exp_t").val().toUpperCase();
        var work_exp_rtfe_val = $("#work_exp_rtfe").val();

        switch ($(this).attr("data-action")) {
            case "add":
                var work_exp_index = 1 + $("#work_exp_table tbody tr").length++;
                var row = `
                <tr data-index="` + work_exp_index + `">
                    <td class="work-exp-checkbox">
                        <input type="checkbox" class="checkbox" id="work_exp_checkbox_` + work_exp_index + `">
                    </td>
                    <td class="work-exp-nocf">
                        <span>` + work_exp_nocf_val + `</span>
                        <input type="hidden" name="work_exp_nocf" value="` + work_exp_nocf_val + `">
                    </td>
                    <td class="work-exp-a">
                        <span>` + work_exp_a_val + `</span>
                        <input type="hidden" name="work_exp_a" value="` + work_exp_a_val + `">
                    </td>
                    <td class="work-exp-ph">
                        <span>` + $("#work_exp_ph_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="work_exp_ph_id" value="` + work_exp_ph_id_val + `">
                    </td>
                    <td class="work-exp-f">
                        <span>` + work_exp_f_val + `</span>
                        <input type="hidden" name="work_exp_f" value="` + work_exp_f_val + `">
                    </td>
                    <td class="work-exp-t">
                        <span>` + work_exp_t_val + `</span>
                        <input type="hidden" name="work_exp_t" value="` + work_exp_t_val + `">
                    </td>
                    <td class="work-exp-rtfe">
                        <span>` + work_exp_rtfe_val + `</span>
                        <input type="hidden" name="work_exp_rtfe" value="` + work_exp_rtfe_val + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="work-exp-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#work_exp_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#work_exp_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".work-exp-nocf").find("span").text(work_exp_nocf_val);
                tr.find(".work-exp-nocf").find('input[name="work_exp_nocf"]').val(work_exp_nocf_val);
                tr.find(".work-exp-a").find("span").text(work_exp_a_val);
                tr.find(".work-exp-a").find('input[name="work_exp_a"]').val(work_exp_a_val);
                tr.find(".work-exp-ph").find("span").text($("#work_exp_ph_id").select2("data")[0].text);
                tr.find(".work-exp-ph").find('input[name="work_exp_ph_id"]').val(work_exp_ph_id_val);
                tr.find(".work-exp-f").find("span").text(work_exp_f_val);
                tr.find(".work-exp-f").find('input[name="work_exp_f"]').val(work_exp_f_val);
                tr.find(".work-exp-t").find("span").text(work_exp_t_val);
                tr.find(".work-exp-t").find('input[name="work_exp_t"]').val(work_exp_t_val);
                tr.find(".work-exp-rtfe").find("span").text(work_exp_rtfe_val);
                tr.find(".work-exp-rtfe").find('input[name="work_exp_rtfe"]').val(work_exp_rtfe_val);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".work-exp-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_work_exp.prop("checked", false);
            }

            if ($(".work-exp-checkbox input:checked").length == $(".work-exp-checkbox input").length) {
                select_all_work_exp.prop("checked", true);
            }

            if ($(".work-exp-checkbox input:checked").length == 0) {
                delete_work_exp_button.prop("disabled", true);
            } else {
                delete_work_exp_button.prop("disabled", false);
            }
        });

        $(".work-exp-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var work_exp_nocf_val = tr.find(".work-exp-nocf").find('input[name="work_exp_nocf"]').val().toUpperCase();
            var work_exp_a_val = tr.find(".work-exp-a").find('input[name="work_exp_a"]').val().toUpperCase();
            var work_exp_ph_id_val = tr.find(".work-exp-ph").find('input[name="work_exp_ph_id"]').val();
            var work_exp_f_val = tr.find(".work-exp-f").find('input[name="work_exp_f"]').val().toUpperCase();
            var work_exp_t_val = tr.find(".work-exp-t").find('input[name="work_exp_t"]').val().toUpperCase();
            var work_exp_rtfe_val = tr.find(".work-exp-rtfe").find('input[name="work_exp_rtfe"]').val();

            $("#work_exp_nocf").val(work_exp_nocf_val).trigger("change");
            $("#work_exp_a").val(work_exp_a_val).trigger("change");
            $("#work_exp_ph_id").val(parseInt(work_exp_ph_id_val)).trigger("change");
            $("#work_exp_f").val(work_exp_f_val).trigger("change");
            $("#work_exp_t").val(work_exp_t_val).trigger("change");
            $("#work_exp_rtfe").val(work_exp_rtfe_val).trigger("change");
            $("#work_exp_form").attr("data-edit-index", tr.data("index"));
            $("#work_exp_form").attr("data-action", "edit");
            $("#work_exp_modal").modal("show");
        });
        select_all_work_exp.prop("checked", false);
        $("#work_exp_modal").modal("hide");
    });

    select_all_work_exp.on("change", function () {
        if ($(".work-exp-checkbox input").length > 0) {
            $(".work-exp-checkbox input").prop("checked", $(this).prop("checked"));
            delete_work_exp_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_work_exp_button.on("click", function () {
        $(".work-exp-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_work_exp.prop("checked", false);
            }
        });

        if ($("#work_exp_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#work_exp_modal").on("hidden.bs.modal", function () {
        $("#work_exp_form").removeAttr("data-action");
        $("#work_exp_nocf").val("");
        $("#work_exp_nocf").parsley().reset();
        $("#work_exp_a").val("");
        $("#work_exp_a").parsley().reset();
        $("#work_exp_ph_id").removeAttr("data-parsley-required");
        $("#work_exp_ph_id").val(null).trigger("change");
        $("#work_exp_ph_id").attr("data-parsley-required", true);
        $("#work_exp_f").val("");
        $("#work_exp_f").parsley().reset();
        $("#work_exp_t").val("");
        $("#work_exp_t").parsley().reset();
        $("#work_exp_rtfe").val(false).trigger("change");
        $("#work_exp_rtfe").prop("checked", false);
    });

    $("#add_work_exp_button").on("click", function () {
        $("#work_exp_form").attr("data-action", "add");
        $("#work_exp_modal").modal("show");
    });
});