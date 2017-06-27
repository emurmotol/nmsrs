$(function () {
    var selectAllFormalEdu = $("#selectAllFormalEdu");
    var delFormalEduBtn = $("#delFormalEduBtn");

    $("#formalEduYearGrad").datetimepicker({
        viewMode: "years",
        format: "YYYY"
    });

    $("#formalEduYearGrad").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#formalEduLastAttended").datetimepicker({
        format: "YYYY-MM"
    });

    $("#formalEduLastAttended").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#formalEduForm").parsley();
    $("#formalEduForm").on("submit", function (e) {
        e.preventDefault();
        var formalEduHighestGradeCompletedVal = $("#formalEduHighestGradeCompleted").select2("val");
        var formalEduCourseDegreeVal = $("#formalEduCourseDegree").select2("val");
        var formalEduSchoolUnivVal = $("#formalEduSchoolUniv").select2("val") != null ? $("#formalEduSchoolUniv").select2("val") : "";
        var formalEduSchoolUnivOtherVal = $("#formalEduSchoolUnivOther").val().toUpperCase();
        var formalEduSchoolUnivText = formalEduSchoolUnivOtherVal;
        var formalEduYearGradVal = $("#formalEduYearGrad").val();
        var formalEduLastAttendedVal = $("#formalEduLastAttended").val();

        if (!$("#formalEduSchoolUnivNotListed").prop("checked")) {
            formalEduSchoolUnivText = $("#formalEduSchoolUniv").select2("data")[0].text;
        }

        switch ($(this).attr("data-action")) {
            case "add":
                var formalEduIndex = $("#formalEduTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + formalEduIndex + `">
                    <td class="formal-edu-checkbox">
                        <input type="checkbox" class="checkbox" id="formalEduCheckbox_` + formalEduIndex + `">
                    </td>
                    <td class="formal-edu-highest-grade-completed">
                        <span>` + $("#formalEduHighestGradeCompleted").select2("data")[0].text + `</span>
                        <input type="hidden" name="formalEduHighestGradeCompleted" value="` + formalEduHighestGradeCompletedVal + `">
                    </td>
                    <td class="formal-edu-course-degree">
                        <span>` + $("#formalEduCourseDegree").select2("data")[0].text + `</span>
                        <input type="hidden" name="formalEduCourseDegree" value="` + formalEduCourseDegreeVal + `">
                    </td>
                    <td class="formal-edu-school-univ">
                        <span>` + formalEduSchoolUnivText + `</span>
                        <input type="hidden" name="formalEduSchoolUniv" value="` + formalEduSchoolUnivVal + `">
                        <input type="hidden" name="formalEduSchoolUnivOther" value="` + formalEduSchoolUnivOtherVal + `">
                    </td>
                    <td class="formal-edu-year-grad">
                        <span>` + formalEduYearGradVal + `</span>
                        <input type="hidden" name="formalEduYearGrad" value="` + formalEduYearGradVal + `">
                    </td>
                    <td class="formal-edu-last-attended">
                        <span>` + formalEduLastAttendedVal + `</span>
                        <input type="hidden" name="formalEduLastAttended" value="` + formalEduLastAttendedVal + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="formal-edu-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#formalEduTable tbody").append(row);
                break;
            case "edit":
                var tr = $("#formalEduTable tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".formal-edu-highest-grade-completed").find("span").text($("#formalEduHighestGradeCompleted").select2("data")[0].text);
                tr.find(".formal-edu-highest-grade-completed").find('input[name="formalEduHighestGradeCompleted"]').val(formalEduHighestGradeCompletedVal);
                tr.find(".formal-edu-course-degree").find("span").text($("#formalEduCourseDegree").select2("data")[0].text);
                tr.find(".formal-edu-course-degree").find('input[name="formalEduCourseDegree"]').val(formalEduCourseDegreeVal);
                tr.find(".formal-edu-school-univ").find("span").text(formalEduSchoolUnivText);
                tr.find(".formal-edu-school-univ").find('input[name="formalEduSchoolUniv"]').val(formalEduSchoolUnivVal);
                tr.find(".formal-edu-school-univ").find('input[name="formalEduSchoolUnivOther"]').val(formalEduSchoolUnivOtherVal);
                tr.find(".formal-edu-year-grad").find("span").text(formalEduYearGradVal);
                tr.find(".formal-edu-year-grad").find('input[name="formalEduYearGrad"]').val(formalEduYearGradVal);
                tr.find(".formal-edu-last-attended").find("span").text(formalEduLastAttendedVal);
                tr.find(".formal-edu-last-attended").find('input[name="formalEduLastAttended"]').val(formalEduLastAttendedVal);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".formal-edu-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                selectAllFormalEdu.prop("checked", false);
            }

            if ($(".formal-edu-checkbox input:checked").length == $(".formal-edu-checkbox input").length) {
                selectAllFormalEdu.prop("checked", true);
            }

            if ($(".formal-edu-checkbox input:checked").length == 0) {
                delFormalEduBtn.prop("disabled", true);
            } else {
                delFormalEduBtn.prop("disabled", false);
            }
        });

        $(".formal-edu-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var formalEduHighestGradeCompletedVal = tr.find(".formal-edu-highest-grade-completed").find('input[name="formalEduHighestGradeCompleted"]').val();
            var formalEduCourseDegreeVal = tr.find(".formal-edu-course-degree").find('input[name="formalEduCourseDegree"]').val();
            var formalEduSchoolUnivVal = tr.find(".formal-edu-school-univ").find('input[name="formalEduSchoolUniv"]').val();
            var formalEduYearGradVal = tr.find(".formal-edu-year-grad").find('input[name="formalEduYearGrad"]').val();
            var formalEduLastAttendedVal = tr.find(".formal-edu-last-attended").find('input[name="formalEduLastAttended"]').val();

            $("#formalEduHighestGradeCompleted").val(formalEduHighestGradeCompletedVal).trigger("change");
            $("#formalEduCourseDegree").val(formalEduCourseDegreeVal).trigger("change");

            if (formalEduSchoolUnivVal == "") {
                var formalEduSchoolUnivOtherVal = tr.find(".formal-edu-school-univ").find('input[name="formalEduSchoolUnivOther"]').val();
                $("#formalEduSchoolUnivNotListed").prop("checked", true).trigger("change");
                $("#formalEduSchoolUnivOther").val(formalEduSchoolUnivOtherVal).trigger("change");
            } else {
                $("#formalEduSchoolUniv").val(formalEduSchoolUnivVal).trigger("change");
            }
            $("#formalEduYearGrad").val(formalEduYearGradVal).trigger("change");
            $("#formalEduLastAttended").val(formalEduLastAttendedVal).trigger("change");
            $("#formalEduForm").attr("data-edit-index", tr.data("index"));
            $("#formalEduForm").attr("data-action", "edit");
            $("#formalEduModal").modal("show");
        });
        selectAllFormalEdu.prop("checked", false);
        $("#formalEduModal").modal("hide");
    });

    selectAllFormalEdu.on("change", function () {
        if ($(".formal-edu-checkbox input").length > 0) {
            $(".formal-edu-checkbox input").prop("checked", $(this).prop("checked"));
            delFormalEduBtn.prop("disabled", !$(this).prop("checked"));
        }
    });

    delFormalEduBtn.on("click", function () {
        $(".formal-edu-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                selectAllFormalEdu.prop("checked", false);
            }
        });

        if ($("#formalEduTable tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#formalEduModal").on("hidden.bs.modal", function () {
        $("#formalEduForm").removeAttr("data-action");
        $("#formalEduSchoolUniv").removeAttr("data-parsley-required");
        $("#formalEduSchoolUniv").val(null).trigger("change");
        $("#formalEduSchoolUniv").attr("data-parsley-required", true);

        if ($("#formalEduSchoolUniv").prop("disabled")) {
            $("#formalEduSchoolUniv").prop("disabled", false);
            $("#formalEduSchoolUnivOther").prop("disabled", true);
        }
        $("#formalEduHighestGradeCompleted").removeAttr("data-parsley-required");
        $("#formalEduHighestGradeCompleted").val(null).trigger("change");
        $("#formalEduHighestGradeCompleted").attr("data-parsley-required", true);
        $("#formalEduCourseDegree").removeAttr("data-parsley-required");
        $("#formalEduCourseDegree").val(null).trigger("change");
        $("#formalEduCourseDegree").attr("data-parsley-required", true);
        $("#formalEduSchoolUnivNotListed").prop("checked", false).trigger("change");
        $("#formalEduYearGrad").val("");
        $("#formalEduYearGrad").parsley().reset();
        $("#formalEduLastAttended").val("");
        $("#formalEduLastAttended").parsley().reset();
    });

    $("#addFormalEduBtn").on("click", function () {
        $("#formalEduForm").attr("data-action", "add");
        $("#formalEduModal").modal("show");
    });

    $("#formalEduSchoolUnivNotListed").on("change", function () {
        $(this).val($(this).prop("checked"));

        if ($(this).prop("checked")) {
            $("#formalEduSchoolUniv").removeAttr("data-parsley-required");
            $("#formalEduSchoolUniv").val(null).trigger("change");
            $("#formalEduSchoolUniv").prop("disabled", true);
            $("#formalEduSchoolUnivOther").attr("data-parsley-required", true);
            $("#formalEduSchoolUnivOther").prop("disabled", false);
            $("#formalEduSchoolUnivOther").focus();
        } else {
            $("#formalEduSchoolUnivOther").removeAttr("data-parsley-required");
            $("#formalEduSchoolUnivOther").val("");
            $("#formalEduSchoolUnivOther").parsley().reset();
            $("#formalEduSchoolUnivOther").prop("disabled", true);
            $("#formalEduSchoolUniv").attr("data-parsley-required", true);
            $("#formalEduSchoolUniv").prop("disabled", false);
            $("#formalEduSchoolUniv").focus();
        }
    });
    var selectAllProLicense = $("#selectAllProLicense");
    var delProLicenseBtn = $("#delProLicenseBtn");

    $("#proLicenseExpiryDate").datetimepicker({
        format: "YYYY-MM"
    });

    $("#proLicenseExpiryDate").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#proLicenseForm").parsley();
    $("#proLicenseForm").on("submit", function (e) {
        e.preventDefault();
        var proLicenseTitleVal = $("#proLicenseTitle").select2("val");
        var proLicenseExpiryDateVal = $("#proLicenseExpiryDate").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var proLicenseIndex = $("#proLicenseTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + proLicenseIndex + `">
                    <td class="pro-license-checkbox">
                        <input type="checkbox" class="checkbox" id="proLicenseCheckbox_` + proLicenseIndex + `">
                    </td>
                    <td class="pro-license-title">
                        <span>` + $("#proLicenseTitle").select2("data")[0].text + `</span>
                        <input type="hidden" name="proLicenseTitle" value="` + proLicenseTitleVal + `">
                    </td>
                    <td class="pro-license-expiry-date">
                        <span>` + proLicenseExpiryDateVal + `</span>
                        <input type="hidden" name="proLicenseExpiryDate" value="` + proLicenseExpiryDateVal + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="pro-license-expiry-date-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#proLicenseTable tbody").append(row);
                break;
            case "edit":
                var tr = $("#proLicenseTable tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".pro-license-title").find("span").text($("#proLicenseTitle").select2("data")[0].text);
                tr.find(".pro-license-title").find('input[name="proLicenseTitle"]').val(proLicenseTitleVal);
                tr.find(".pro-license-expiry-date").find("span").text(proLicenseExpiryDateVal);
                tr.find(".pro-license-expiry-date").find('input[name="proLicenseExpiryDate"]').val(proLicenseExpiryDateVal);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".pro-license-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                selectAllProLicense.prop("checked", false);
            }

            if ($(".pro-license-checkbox input:checked").length == $(".pro-license-checkbox input").length) {
                selectAllProLicense.prop("checked", true);
            }

            if ($(".pro-license-checkbox input:checked").length == 0) {
                delProLicenseBtn.prop("disabled", true);
            } else {
                delProLicenseBtn.prop("disabled", false);
            }
        });

        $(".pro-license-expiry-date-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var proLicenseTitleVal = tr.find(".pro-license-title").find('input[name="proLicenseTitle"]').val();
            var proLicenseExpiryDateVal = tr.find(".pro-license-expiry-date").find('input[name="proLicenseExpiryDate"]').val();

            $("#proLicenseTitle").val(proLicenseTitleVal).trigger("change");
            $("#proLicenseExpiryDate").val(proLicenseExpiryDateVal).trigger("change");
            $("#proLicenseForm").attr("data-edit-index", tr.data("index"));
            $("#proLicenseForm").attr("data-action", "edit");
            $("#proLicenseModal").modal("show");
        });
        selectAllProLicense.prop("checked", false);
        $("#proLicenseModal").modal("hide");
    });

    selectAllProLicense.on("change", function () {
        if ($(".pro-license-checkbox input").length > 0) {
            $(".pro-license-checkbox input").prop("checked", $(this).prop("checked"));
            delProLicenseBtn.prop("disabled", !$(this).prop("checked"));
        }
    });

    delProLicenseBtn.on("click", function () {
        $(".pro-license-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                selectAllProLicense.prop("checked", false);
            }
        });

        if ($("#proLicenseTable tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#proLicenseModal").on("hidden.bs.modal", function () {
        $("#proLicenseForm").removeAttr("data-action");
        $("#proLicenseTitle").removeAttr("data-parsley-required");
        $("#proLicenseTitle").val(null).trigger("change");
        $("#proLicenseTitle").attr("data-parsley-required", true);
        $("#proLicenseExpiryDate").val("");
        $("#proLicenseExpiryDate").parsley().reset();
    });

    $("#addProLicenseBtn").on("click", function () {
        $("#proLicenseForm").attr("data-action", "add");
        $("#proLicenseModal").modal("show");
    });
    var selectAllElig = $("#selectAllElig");
    var deleteEligBtn = $("#deleteEligBtn");

    $("#eligYearTaken").datetimepicker({
        viewMode: "years",
        format: "YYYY-MM"
    });

    $("#eligYearTaken").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#eligForm").parsley();
    $("#eligForm").on("submit", function (e) {
        e.preventDefault();
        var eligTitleVal = $("#eligTitle").select2("val");
        var eligYearTakenVal = $("#eligYearTaken").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var eligIndex = $("#eligTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + eligIndex + `">
                    <td class="elig-checkbox">
                        <input type="checkbox" class="checkbox" id="eligCheckbox_` + eligIndex + `">
                    </td>
                    <td class="elig-title">
                        <span>` + $("#eligTitle").select2("data")[0].text + `</span>
                        <input type="hidden" name="eligTitle" value="` + eligTitleVal + `">
                    </td>
                    <td class="elig-year-taken">
                        <span>` + eligYearTakenVal + `</span>
                        <input type="hidden" name="eligYearTaken" value="` + eligYearTakenVal + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="elig-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#eligTable tbody").append(row);
                break;
            case "edit":
                var tr = $("#eligTable tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".elig-title").find("span").text($("#eligTitle").select2("data")[0].text);
                tr.find(".elig-title").find('input[name="eligTitle"]').val(eligTitleVal);
                tr.find(".elig-year-taken").find("span").text(eligYearTakenVal);
                tr.find(".elig-year-taken").find('input[name="eligYearTaken"]').val(eligYearTakenVal);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".elig-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                selectAllElig.prop("checked", false);
            }

            if ($(".elig-checkbox input:checked").length == $(".elig-checkbox input").length) {
                selectAllElig.prop("checked", true);
            }

            if ($(".elig-checkbox input:checked").length == 0) {
                deleteEligBtn.prop("disabled", true);
            } else {
                deleteEligBtn.prop("disabled", false);
            }
        });

        $(".elig-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var eligTitleVal = tr.find(".elig-title").find('input[name="eligTitle"]').val();
            var eligYearTakenVal = tr.find(".elig-year-taken").find('input[name="eligYearTaken"]').val().toUpperCase();

            $("#eligTitle").val(eligTitleVal).trigger("change");
            $("#eligYearTaken").val(eligYearTakenVal).trigger("change");
            $("#eligForm").attr("data-edit-index", tr.data("index"));
            $("#eligForm").attr("data-action", "edit");
            $("#eligModal").modal("show");
        });
        selectAllElig.prop("checked", false);
        $("#eligModal").modal("hide");
    });

    selectAllElig.on("change", function () {
        if ($(".elig-checkbox input").length > 0) {
            $(".elig-checkbox input").prop("checked", $(this).prop("checked"));
            deleteEligBtn.prop("disabled", !$(this).prop("checked"));
        }
    });

    deleteEligBtn.on("click", function () {
        $(".elig-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                selectAllElig.prop("checked", false);
            }
        });

        if ($("#eligTable tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#eligModal").on("hidden.bs.modal", function () {
        $("#eligForm").removeAttr("data-action");
        $("#eligTitle").removeAttr("data-parsley-required");
        $("#eligTitle").val(null).trigger("change");
        $("#eligTitle").attr("data-parsley-required", true);
        $("#eligYearTaken").val("");
        $("#eligYearTaken").parsley().reset();
    });

    $("#addEligBtn").on("click", function () {
        $("#eligForm").attr("data-action", "add");
        $("#eligModal").modal("show");
    });
    var selectAllTraining = $("#selectAllTraining");
    var delTrainingBtn = $("#delTrainingBtn");

    $("#trainingForm").parsley();
    $("#trainingForm").on("submit", function (e) {
        e.preventDefault();
        var trainingNameOfTrainingVal = $("#trainingNameOfTraining").val().toUpperCase();
        var trainingSkillsAcquiredVal = $("#trainingSkillsAcquired").val().toUpperCase();
        var trainingPeriodOfTrainingExpVal = $("#trainingPeriodOfTrainingExp").val().toUpperCase();
        var trainingCertReceivedVal = $("#trainingCertReceived").val().toUpperCase();
        var trainingIssuingSchoolAgencyVal = $("#trainingIssuingSchoolAgency").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var trainingIndex = $("#trainingTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + trainingIndex + `">
                    <td class="training-checkbox">
                        <input type="checkbox" class="checkbox" id="trainingCheckbox_` + trainingIndex + `">
                    </td>
                    <td class="training-name-of-training">
                        <span>` + trainingNameOfTrainingVal + `</span>
                        <input type="hidden" name="trainingNameOfTraining" value="` + trainingNameOfTrainingVal + `">
                    </td>
                    <td class="training-skills-acquired">
                        <span>` + trainingSkillsAcquiredVal + `</span>
                        <input type="hidden" name="trainingSkillsAcquired" value="` + trainingSkillsAcquiredVal + `">
                    </td>
                    <td class="training-period-of-training-exp">
                        <span>` + trainingPeriodOfTrainingExpVal + `</span>
                        <input type="hidden" name="trainingPeriodOfTrainingExp" value="` + trainingPeriodOfTrainingExpVal + `">
                    </td>
                    <td class="training-cert-received">
                        <span>` + trainingCertReceivedVal + `</span>
                        <input type="hidden" name="trainingCertReceived" value="` + trainingCertReceivedVal + `">
                    </td>
                    <td class="training-issuing-school-agency">
                        <span>` + trainingIssuingSchoolAgencyVal + `</span>
                        <input type="hidden" name="trainingIssuingSchoolAgency" value="` + trainingIssuingSchoolAgencyVal + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="training-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#trainingTable tbody").append(row);
                break;
            case "edit":
                var tr = $("#trainingTable tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".training-name-of-training").find("span").text(trainingNameOfTrainingVal);
                tr.find(".training-name-of-training").find('input[name="trainingNameOfTraining"]').val(trainingNameOfTrainingVal);
                tr.find(".training-skills-acquired").find("span").text(trainingSkillsAcquiredVal);
                tr.find(".training-skills-acquired").find('input[name="trainingSkillsAcquired"]').val(trainingSkillsAcquiredVal);
                tr.find(".training-period-of-training-exp").find("span").text(trainingPeriodOfTrainingExpVal);
                tr.find(".training-period-of-training-exp").find('input[name="trainingPeriodOfTrainingExp"]').val(trainingPeriodOfTrainingExpVal);
                tr.find(".training-cert-received").find("span").text(trainingCertReceivedVal);
                tr.find(".training-cert-received").find('input[name="trainingCertReceived"]').val(trainingCertReceivedVal);
                tr.find(".training-issuing-school-agency").find("span").text(trainingIssuingSchoolAgencyVal);
                tr.find(".training-issuing-school-agency").find('input[name="trainingIssuingSchoolAgency"]').val(trainingIssuingSchoolAgencyVal);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".training-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                selectAllTraining.prop("checked", false);
            }

            if ($(".training-checkbox input:checked").length == $(".training-checkbox input").length) {
                selectAllTraining.prop("checked", true);
            }

            if ($(".training-checkbox input:checked").length == 0) {
                delTrainingBtn.prop("disabled", true);
            } else {
                delTrainingBtn.prop("disabled", false);
            }
        });

        $(".training-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var trainingNameOfTrainingVal = tr.find(".training-name-of-training").find('input[name="trainingNameOfTraining"]').val();
            var trainingSkillsAcquiredVal = tr.find(".training-skills-acquired").find('input[name="trainingSkillsAcquired"]').val();
            var trainingPeriodOfTrainingExpVal = tr.find(".training-period-of-training-exp").find('input[name="trainingPeriodOfTrainingExp"]').val();
            var trainingCertReceivedVal = tr.find(".training-cert-received").find('input[name="trainingCertReceived"]').val();
            var trainingIssuingSchoolAgencyVal = tr.find(".training-issuing-school-agency").find('input[name="trainingIssuingSchoolAgency"]').val();

            $("#trainingNameOfTraining").val(trainingNameOfTrainingVal).trigger("change");
            $("#trainingSkillsAcquired").val(trainingSkillsAcquiredVal).trigger("change");
            $("#trainingPeriodOfTrainingExp").val(trainingPeriodOfTrainingExpVal).trigger("change");
            $("#trainingCertReceived").val(trainingCertReceivedVal).trigger("change");
            $("#trainingIssuingSchoolAgency").val(trainingIssuingSchoolAgencyVal).trigger("change");
            $("#trainingForm").attr("data-edit-index", tr.data("index"));
            $("#trainingForm").attr("data-action", "edit");
            $("#trainingModal").modal("show");
        });
        selectAllTraining.prop("checked", false);
        $("#trainingModal").modal("hide");
    });

    selectAllTraining.on("change", function () {
        if ($(".training-checkbox input").length > 0) {
            $(".training-checkbox input").prop("checked", $(this).prop("checked"));
            delTrainingBtn.prop("disabled", !$(this).prop("checked"));
        }
    });

    delTrainingBtn.on("click", function () {
        $(".training-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                selectAllTraining.prop("checked", false);
            }
        });

        if ($("#trainingTable tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#trainingModal").on("hidden.bs.modal", function () {
        $("#trainingForm").removeAttr("data-action");
        $("#trainingNameOfTraining").val("");
        $("#trainingNameOfTraining").parsley().reset();
        $("#trainingSkillsAcquired").val("");
        $("#trainingSkillsAcquired").parsley().reset();
        $("#trainingPeriodOfTrainingExp").val("");
        $("#trainingPeriodOfTrainingExp").parsley().reset();
        $("#trainingCertReceived").val("");
        $("#trainingCertReceived").parsley().reset();
        $("#trainingIssuingSchoolAgency").val("");
        $("#trainingIssuingSchoolAgency").parsley().reset();
    });

    $("#addTrainingBtn").on("click", function () {
        $("#trainingForm").attr("data-action", "add");
        $("#trainingModal").modal("show");
    });
    var selectAllCertOfCompetence = $("#selectAllCertOfCompetence");
    var delCertOfCompetenceBtn = $("#delCertOfCompetenceBtn");

    $("#certOfCompetenceDateIssued").datetimepicker({
        format: "YYYY-MM"
    });

    $("#certOfCompetenceDateIssued").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#certOfCompetenceForm").parsley();
    $("#certOfCompetenceForm").on("submit", function (e) {
        e.preventDefault();
        var certOfCompetenceTitleVal = $("#certOfCompetenceTitle").select2("val");
        var certOfCompetenceRatingVal = $("#certOfCompetenceRating").val().toUpperCase();
        var certOfCompetenceIssuedByVal = $("#certOfCompetenceIssuedBy").val().toUpperCase();
        var certOfCompetenceDateIssuedVal = $("#certOfCompetenceDateIssued").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var certOfCompetenceIndex = $("#certOfCompetenceTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + certOfCompetenceIndex + `">
                    <td class="certOfCompetenceCheckbox_">
                        <input type="checkbox" class="checkbox" id="certOfCompetenceCheckbox_` + certOfCompetenceIndex + `">
                    </td>
                    <td class="cert-of-competence-title">
                        <span>` + $("#certOfCompetenceTitle").select2("data")[0].text + `</span>
                        <input type="hidden" name="certOfCompetenceTitle" value="` + certOfCompetenceTitleVal + `">
                    </td>
                    <td class="cert-of-competence-rating">
                        <span>` + certOfCompetenceRatingVal + `</span>
                        <input type="hidden" name="certOfCompetenceRating" value="` + certOfCompetenceRatingVal + `">
                    </td>
                    <td class="cert-of-competence-issued-by">
                        <span>` + certOfCompetenceIssuedByVal + `</span>
                        <input type="hidden" name="certOfCompetenceIssuedBy" value="` + certOfCompetenceIssuedByVal + `">
                    </td>
                    <td class="cert-of-competence-date-issued">
                        <span>` + certOfCompetenceDateIssuedVal + `</span>
                        <input type="hidden" name="certOfCompetenceDateIssued" value="` + certOfCompetenceDateIssuedVal + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="cert-of-competence-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#certOfCompetenceTable tbody").append(row);
                break;
            case "edit":
                var tr = $("#certOfCompetenceTable tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".cert-of-competence-title").find("span").text($("#certOfCompetenceTitle").select2("data")[0].text);
                tr.find(".cert-of-competence-title").find('input[name="certOfCompetenceTitle"]').val(certOfCompetenceTitleVal);
                tr.find(".cert-of-competence-rating").find("span").text(certOfCompetenceRatingVal);
                tr.find(".cert-of-competence-rating").find('input[name="certOfCompetenceRating"]').val(certOfCompetenceRatingVal);
                tr.find(".cert-of-competence-issued-by").find("span").text(certOfCompetenceIssuedByVal);
                tr.find(".cert-of-competence-issued-by").find('input[name="certOfCompetenceIssuedBy"]').val(certOfCompetenceIssuedByVal);
                tr.find(".cert-of-competence-date-issued").find("span").text(certOfCompetenceDateIssuedVal);
                tr.find(".cert-of-competence-date-issued").find('input[name="certOfCompetenceDateIssued"]').val(certOfCompetenceDateIssuedVal);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".cert-of-competence-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                selectAllCertOfCompetence.prop("checked", false);
            }

            if ($(".cert-of-competence-checkbox input:checked").length == $(".cert-of-competence-checkbox input").length) {
                selectAllCertOfCompetence.prop("checked", true);
            }

            if ($(".cert-of-competence-checkbox input:checked").length == 0) {
                delCertOfCompetenceBtn.prop("disabled", true);
            } else {
                delCertOfCompetenceBtn.prop("disabled", false);
            }
        });

        $(".cert-of-competence-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var certOfCompetenceTitleVal = tr.find(".cert-of-competence-title").find('input[name="certOfCompetenceTitle"]').val();
            var certOfCompetenceRatingVal = tr.find(".cert-of-competence-rating").find('input[name="certOfCompetenceRating"]').val().toUpperCase();
            var certOfCompetenceIssuedByVal = tr.find(".cert-of-competence-issued-by").find('input[name="certOfCompetenceIssuedBy"]').val().toUpperCase();
            var certOfCompetenceDateIssuedVal = tr.find(".cert-of-competence-date-issued").find('input[name="certOfCompetenceDateIssued"]').val().toUpperCase();

            $("#certOfCompetenceTitle").val(certOfCompetenceTitleVal).trigger("change");
            $("#certOfCompetenceRating").val(certOfCompetenceRatingVal).trigger("change");
            $("#certOfCompetenceIssuedBy").val(certOfCompetenceIssuedByVal).trigger("change");
            $("#certOfCompetenceDateIssued").val(certOfCompetenceDateIssuedVal).trigger("change");
            $("#certOfCompetenceForm").attr("data-edit-index", tr.data("index"));
            $("#certOfCompetenceForm").attr("data-action", "edit");
            $("#certOfCompetenceModal").modal("show");
        });
        selectAllCertOfCompetence.prop("checked", false);
        $("#certOfCompetenceModal").modal("hide");
    });

    selectAllCertOfCompetence.on("change", function () {
        if ($(".cert-of-competence-checkbox input").length > 0) {
            $(".cert-of-competence-checkbox input").prop("checked", $(this).prop("checked"));
            delCertOfCompetenceBtn.prop("disabled", !$(this).prop("checked"));
        }
    });

    delCertOfCompetenceBtn.on("click", function () {
        $(".cert-of-competence-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                selectAllCertOfCompetence.prop("checked", false);
            }
        });

        if ($("#certOfCompetenceTable tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#certOfCompetenceModal").on("hidden.bs.modal", function () {
        $("#certOfCompetenceForm").removeAttr("data-action");
        $("#certOfCompetenceTitle").removeAttr("data-parsley-required");
        $("#certOfCompetenceTitle").val(null).trigger("change");
        $("#certOfCompetenceTitle").attr("data-parsley-required", true);
        $("#certOfCompetenceRating").val("");
        $("#certOfCompetenceRating").parsley().reset();
        $("#certOfCompetenceIssuedBy").val("");
        $("#certOfCompetenceIssuedBy").parsley().reset();
        $("#certOfCompetenceDateIssued").val("");
        $("#certOfCompetenceDateIssued").parsley().reset();
    });

    $("#addCertOfCompetenceBtn").on("click", function () {
        $("#certOfCompetenceForm").attr("data-action", "add");
        $("#certOfCompetenceModal").modal("show");
    });
    var selectAllWorkExp = $("#selectAllWorkExp");
    var delWorkExpBtn = $("#delWorkExpBtn");

    $("#workExpFrom").datetimepicker({
        format: "YYYY-MM"
    });

    $("#workExpFrom").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#workExpTo").datetimepicker({
        format: "YYYY-MM"
    });

    $("#workExpTo").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#workExpRelatedToFormalEdu").on("change", function () {
        $(this).val($(this).prop("checked"));
    });

    $("#workExpForm").parsley();
    $("#workExpForm").on("submit", function (e) {
        e.preventDefault();
        var workExpNameOfCompanyFirmVal = $("#workExpNameOfCompanyFirm").val().toUpperCase();
        var workExpAddressVal = $("#workExpAddress").val().toUpperCase();
        var workExpPositionHeldVal = $("#workExpPositionHeld").select2("val");
        var workExpFromVal = $("#workExpFrom").val().toUpperCase();
        var workExpToVal = $("#workExpTo").val().toUpperCase();
        var workExpRelatedToFormalEduVal = $("#workExpRelatedToFormalEdu").val();

        switch ($(this).attr("data-action")) {
            case "add":
                var workExpIndex = $("#workExpTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + workExpIndex + `">
                    <td class="work-exp-checkbox">
                        <input type="checkbox" class="checkbox" id="workExpCheckbox_` + workExpIndex + `">
                    </td>
                    <td class="work-exp-name-of-company-firm">
                        <span>` + workExpNameOfCompanyFirmVal + `</span>
                        <input type="hidden" name="workExpNameOfCompanyFirm" value="` + workExpNameOfCompanyFirmVal + `">
                    </td>
                    <td class="work-exp-address">
                        <span>` + workExpAddressVal + `</span>
                        <input type="hidden" name="workExpAddress" value="` + workExpAddressVal + `">
                    </td>
                    <td class="work-exp-position-held">
                        <span>` + $("#workExpPositionHeld").select2("data")[0].text + `</span>
                        <input type="hidden" name="workExpPositionHeld" value="` + workExpPositionHeldVal + `">
                    </td>
                    <td class="work-exp-from">
                        <span>` + workExpFromVal + `</span>
                        <input type="hidden" name="workExpFrom" value="` + workExpFromVal + `">
                    </td>
                    <td class="work-exp-to">
                        <span>` + workExpToVal + `</span>
                        <input type="hidden" name="workExpTo" value="` + workExpToVal + `">
                    </td>
                    <td class="work-exp-related-to-formal-edu">
                        <span>` + workExpRelatedToFormalEduVal + `</span>
                        <input type="hidden" name="workExpRelatedToFormalEdu" value="` + workExpRelatedToFormalEduVal + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="work-exp-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#workExpTable tbody").append(row);
                break;
            case "edit":
                var tr = $("#workExpTable tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".work-exp-name-of-company-firm").find("span").text(workExpNameOfCompanyFirmVal);
                tr.find(".work-exp-name-of-company-firm").find('input[name="workExpNameOfCompanyFirm"]').val(workExpNameOfCompanyFirmVal);
                tr.find(".work-exp-address").find("span").text(workExpAddressVal);
                tr.find(".work-exp-address").find('input[name="workExpAddress"]').val(workExpAddressVal);
                tr.find(".work-exp-position-held").find("span").text($("#workExpPositionHeld").select2("data")[0].text);
                tr.find(".work-exp-position-held").find('input[name="workExpPositionHeld"]').val(workExpPositionHeldVal);
                tr.find(".work-exp-from").find("span").text(workExpFromVal);
                tr.find(".work-exp-from").find('input[name="workExpFrom"]').val(workExpFromVal);
                tr.find(".work-exp-to").find("span").text(workExpToVal);
                tr.find(".work-exp-to").find('input[name="workExpTo"]').val(workExpToVal);
                tr.find(".work-exp-related-to-formal-edu").find("span").text(workExpRelatedToFormalEduVal);
                tr.find(".work-exp-related-to-formal-edu").find('input[name="workExpRelatedToFormalEdu"]').val(workExpRelatedToFormalEduVal);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".work-exp-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                selectAllWorkExp.prop("checked", false);
            }

            if ($(".work-exp-checkbox input:checked").length == $(".work-exp-checkbox input").length) {
                selectAllWorkExp.prop("checked", true);
            }

            if ($(".work-exp-checkbox input:checked").length == 0) {
                delWorkExpBtn.prop("disabled", true);
            } else {
                delWorkExpBtn.prop("disabled", false);
            }
        });

        $(".work-exp-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var workExpNameOfCompanyFirmVal = tr.find(".work-exp-name-of-company-firm").find('input[name="workExpNameOfCompanyFirm"]').val().toUpperCase();
            var workExpAddressVal = tr.find(".work-exp-address").find('input[name="workExpAddress"]').val().toUpperCase();
            var workExpPositionHeldVal = tr.find(".work-exp-position-held").find('input[name="workExpPositionHeld"]').val();
            var workExpFromVal = tr.find(".work-exp-from").find('input[name="workExpFrom"]').val().toUpperCase();
            var workExpToVal = tr.find(".work-exp-to").find('input[name="workExpTo"]').val().toUpperCase();
            var workExpRelatedToFormalEduVal = tr.find(".work-exp-related-to-formal-edu").find('input[name="workExpRelatedToFormalEdu"]').val();

            $("#workExpNameOfCompanyFirm").val(workExpNameOfCompanyFirmVal).trigger("change");
            $("#workExpAddress").val(workExpAddressVal).trigger("change");
            $("#workExpPositionHeld").val(workExpPositionHeldVal).trigger("change");
            $("#workExpFrom").val(workExpFromVal).trigger("change");
            $("#workExpTo").val(workExpToVal).trigger("change");
            $("#workExpRelatedToFormalEdu").val(workExpRelatedToFormalEduVal).trigger("change");
            $("#workExpForm").attr("data-edit-index", tr.data("index"));
            $("#workExpForm").attr("data-action", "edit");
            $("#workExpModal").modal("show");
        });
        selectAllWorkExp.prop("checked", false);
        $("#workExpModal").modal("hide");
    });

    selectAllWorkExp.on("change", function () {
        if ($(".work-exp-checkbox input").length > 0) {
            $(".work-exp-checkbox input").prop("checked", $(this).prop("checked"));
            delWorkExpBtn.prop("disabled", !$(this).prop("checked"));
        }
    });

    delWorkExpBtn.on("click", function () {
        $(".work-exp-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                selectAllWorkExp.prop("checked", false);
            }
        });

        if ($("#workExpTable tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#workExpModal").on("hidden.bs.modal", function () {
        $("#workExpForm").removeAttr("data-action");
        $("#workExpNameOfCompanyFirm").val("");
        $("#workExpNameOfCompanyFirm").parsley().reset();
        $("#workExpAddress").val("");
        $("#workExpAddress").parsley().reset();
        $("#workExpPositionHeld").removeAttr("data-parsley-required");
        $("#workExpPositionHeld").val(null).trigger("change");
        $("#workExpPositionHeld").attr("data-parsley-required", true);
        $("#workExpFrom").val("");
        $("#workExpFrom").parsley().reset();
        $("#workExpTo").val("");
        $("#workExpTo").parsley().reset();
        $("#workExpRelatedToFormalEdu").val(false).trigger("change");
        $("#workExpRelatedToFormalEdu").prop("checked", false);
    });

    $("#addWorkExpBtn").on("click", function () {
        $("#workExpForm").attr("data-action", "add");
        $("#workExpModal").modal("show");
    });
});