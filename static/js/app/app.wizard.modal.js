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
        var formalEduHighestGradeCompletedIdVal = $("#formalEduHighestGradeCompletedId").select2("val");
        var formalEduCourseDegreeIdVal = $("#formalEduCourseDegreeId").select2("val");
        var formalEduSchoolUnivIdVal = $("#formalEduSchoolUnivId").select2("val") != null ? $("#formalEduSchoolUnivId").select2("val") : "";
        var formalEduSchoolUnivOtherVal = $("#formalEduSchoolUnivOther").val().toUpperCase();
        var formalEduSchoolUnivText = formalEduSchoolUnivOtherVal;
        var formalEduYearGradVal = $("#formalEduYearGrad").val();
        var formalEduLastAttendedVal = $("#formalEduLastAttended").val();

        if (!$("#formalEduSchoolUnivNotListed").prop("checked")) {
            formalEduSchoolUnivText = $("#formalEduSchoolUnivId").select2("data")[0].text;
        }

        switch ($(this).attr("data-action")) {
            case "add":
                var formalEduIndex = $("#formalEduTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + formalEduIndex + `">
                    <td class="formal-edu-checkbox">
                        <input type="checkbox" class="checkbox" id="formalEduCheckbox_` + formalEduIndex + `">
                    </td>
                    <td class="formal-edu-highest-grade-completed-id">
                        <span>` + $("#formalEduHighestGradeCompletedId").select2("data")[0].text + `</span>
                        <input type="hidden" name="formalEduHighestGradeCompletedId" value="` + formalEduHighestGradeCompletedIdVal + `">
                    </td>
                    <td class="formal-edu-course-degree-id">
                        <span>` + $("#formalEduCourseDegreeId").select2("data")[0].text + `</span>
                        <input type="hidden" name="formalEduCourseDegreeId" value="` + formalEduCourseDegreeIdVal + `">
                    </td>
                    <td class="formal-edu-school-univ-id">
                        <span>` + formalEduSchoolUnivText + `</span>
                        <input type="hidden" name="formalEduSchoolUnivId" value="` + formalEduSchoolUnivIdVal + `">
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
                tr.find(".formal-edu-highest-grade-completed-id").find("span").text($("#formalEduHighestGradeCompletedId").select2("data")[0].text);
                tr.find(".formal-edu-highest-grade-completed-id").find('input[name="formalEduHighestGradeCompletedId"]').val(formalEduHighestGradeCompletedIdVal);
                tr.find(".formal-edu-course-degree-id").find("span").text($("#formalEduCourseDegreeId").select2("data")[0].text);
                tr.find(".formal-edu-course-degree-id").find('input[name="formalEduCourseDegreeId"]').val(formalEduCourseDegreeIdVal);
                tr.find(".formal-edu-school-univ-id").find("span").text(formalEduSchoolUnivText);
                tr.find(".formal-edu-school-univ-id").find('input[name="formalEduSchoolUnivId"]').val(formalEduSchoolUnivIdVal);
                tr.find(".formal-edu-school-univ-id").find('input[name="formalEduSchoolUnivOther"]').val(formalEduSchoolUnivOtherVal);
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
            var formalEduHighestGradeCompletedIdVal = tr.find(".formal-edu-highest-grade-completed-id").find('input[name="formalEduHighestGradeCompletedId"]').val();
            var formalEduCourseDegreeIdVal = tr.find(".formal-edu-course-degree-id").find('input[name="formalEduCourseDegreeId"]').val();
            var formalEduSchoolUnivIdVal = tr.find(".formal-edu-school-univ-id").find('input[name="formalEduSchoolUnivId"]').val();
            var formalEduYearGradVal = tr.find(".formal-edu-year-grad").find('input[name="formalEduYearGrad"]').val();
            var formalEduLastAttendedVal = tr.find(".formal-edu-last-attended").find('input[name="formalEduLastAttended"]').val();

            $("#formalEduHighestGradeCompletedId").val(formalEduHighestGradeCompletedIdVal).trigger("change");
            $("#formalEduCourseDegreeId").val(formalEduCourseDegreeIdVal).trigger("change");

            if (formalEduSchoolUnivIdVal == "") {
                var formalEduSchoolUnivOtherVal = tr.find(".formal-edu-school-univ-id").find('input[name="formalEduSchoolUnivOther"]').val();
                $("#formalEduSchoolUnivNotListed").prop("checked", true).trigger("change");
                $("#formalEduSchoolUnivOther").val(formalEduSchoolUnivOtherVal).trigger("change");
            } else {
                $("#formalEduSchoolUnivId").val(formalEduSchoolUnivIdVal).trigger("change");
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
        $("#formalEduSchoolUnivId").removeAttr("data-parsley-required");
        $("#formalEduSchoolUnivId").val(null).trigger("change");
        $("#formalEduSchoolUnivId").attr("data-parsley-required", true);

        if ($("#formalEduSchoolUnivId").prop("disabled")) {
            $("#formalEduSchoolUnivId").prop("disabled", false);
            $("#formalEduSchoolUnivOther").prop("disabled", true);
        }
        $("#formalEduHighestGradeCompletedId").removeAttr("data-parsley-required");
        $("#formalEduHighestGradeCompletedId").val(null).trigger("change");
        $("#formalEduHighestGradeCompletedId").attr("data-parsley-required", true);
        $("#formalEduCourseDegreeId").removeAttr("data-parsley-required");
        $("#formalEduCourseDegreeId").val(null).trigger("change");
        $("#formalEduCourseDegreeId").attr("data-parsley-required", true);
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
            $("#formalEduSchoolUnivId").removeAttr("data-parsley-required");
            $("#formalEduSchoolUnivId").val(null).trigger("change");
            $("#formalEduSchoolUnivId").prop("disabled", true);
            $("#formalEduSchoolUnivOther").attr("data-parsley-required", true);
            $("#formalEduSchoolUnivOther").prop("disabled", false);
            $("#formalEduSchoolUnivOther").focus();
        } else {
            $("#formalEduSchoolUnivOther").removeAttr("data-parsley-required");
            $("#formalEduSchoolUnivOther").val("");
            $("#formalEduSchoolUnivOther").parsley().reset();
            $("#formalEduSchoolUnivOther").prop("disabled", true);
            $("#formalEduSchoolUnivId").attr("data-parsley-required", true);
            $("#formalEduSchoolUnivId").prop("disabled", false);
            $("#formalEduSchoolUnivId").focus();
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
        var proLicenseTitleIdVal = $("#proLicenseTitleId").select2("val");
        var proLicenseExpiryDateVal = $("#proLicenseExpiryDate").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var proLicenseIndex = $("#proLicenseTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + proLicenseIndex + `">
                    <td class="pro-license-checkbox">
                        <input type="checkbox" class="checkbox" id="proLicenseCheckbox_` + proLicenseIndex + `">
                    </td>
                    <td class="pro-license-title-id">
                        <span>` + $("#proLicenseTitleId").select2("data")[0].text + `</span>
                        <input type="hidden" name="proLicenseTitleId" value="` + proLicenseTitleIdVal + `">
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
                tr.find(".pro-license-title-id").find("span").text($("#proLicenseTitleId").select2("data")[0].text);
                tr.find(".pro-license-title-id").find('input[name="proLicenseTitleId"]').val(proLicenseTitleIdVal);
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
            var proLicenseTitleIdVal = tr.find(".pro-license-title-id").find('input[name="proLicenseTitleId"]').val();
            var proLicenseExpiryDateVal = tr.find(".pro-license-expiry-date").find('input[name="proLicenseExpiryDate"]').val();

            $("#proLicenseTitleId").val(proLicenseTitleIdVal).trigger("change");
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
        $("#proLicenseTitleId").removeAttr("data-parsley-required");
        $("#proLicenseTitleId").val(null).trigger("change");
        $("#proLicenseTitleId").attr("data-parsley-required", true);
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
        var eligTitleIdVal = $("#eligTitleId").select2("val");
        var eligYearTakenVal = $("#eligYearTaken").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var eligIndex = $("#eligTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + eligIndex + `">
                    <td class="elig-checkbox">
                        <input type="checkbox" class="checkbox" id="eligCheckbox_` + eligIndex + `">
                    </td>
                    <td class="elig-title-id">
                        <span>` + $("#eligTitleId").select2("data")[0].text + `</span>
                        <input type="hidden" name="eligTitleId" value="` + eligTitleIdVal + `">
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
                tr.find(".elig-title-id").find("span").text($("#eligTitleId").select2("data")[0].text);
                tr.find(".elig-title-id").find('input[name="eligTitleId"]').val(eligTitleIdVal);
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
            var eligTitleIdVal = tr.find(".elig-title-id").find('input[name="eligTitleId"]').val();
            var eligYearTakenVal = tr.find(".elig-year-taken").find('input[name="eligYearTaken"]').val().toUpperCase();

            $("#eligTitleId").val(eligTitleIdVal).trigger("change");
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
        $("#eligTitleId").removeAttr("data-parsley-required");
        $("#eligTitleId").val(null).trigger("change");
        $("#eligTitleId").attr("data-parsley-required", true);
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
    var selectAllCert = $("#selectAllCert");
    var delCertBtn = $("#delCertBtn");

    $("#certDateIssued").datetimepicker({
        format: "YYYY-MM"
    });

    $("#certDateIssued").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#certForm").parsley();
    $("#certForm").on("submit", function (e) {
        e.preventDefault();
        var certTitleIdVal = $("#certTitleId").select2("val");
        var certRatingVal = $("#certRating").val().toUpperCase();
        var certIssuedByVal = $("#certIssuedBy").val().toUpperCase();
        var certDateIssuedVal = $("#certDateIssued").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var certIndex = $("#certTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + certIndex + `">
                    <td class="certCheckbox_">
                        <input type="checkbox" class="checkbox" id="certCheckbox_` + certIndex + `">
                    </td>
                    <td class="cert-title-id">
                        <span>` + $("#certTitleId").select2("data")[0].text + `</span>
                        <input type="hidden" name="certTitleId" value="` + certTitleIdVal + `">
                    </td>
                    <td class="cert-rating">
                        <span>` + certRatingVal + `</span>
                        <input type="hidden" name="certRating" value="` + certRatingVal + `">
                    </td>
                    <td class="cert-issued-by">
                        <span>` + certIssuedByVal + `</span>
                        <input type="hidden" name="certIssuedBy" value="` + certIssuedByVal + `">
                    </td>
                    <td class="cert-date-issued">
                        <span>` + certDateIssuedVal + `</span>
                        <input type="hidden" name="certDateIssued" value="` + certDateIssuedVal + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="cert-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#certTable tbody").append(row);
                break;
            case "edit":
                var tr = $("#certTable tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".cert-title-id").find("span").text($("#certTitleId").select2("data")[0].text);
                tr.find(".cert-title-id").find('input[name="certTitleId"]').val(certTitleIdVal);
                tr.find(".cert-rating").find("span").text(certRatingVal);
                tr.find(".cert-rating").find('input[name="certRating"]').val(certRatingVal);
                tr.find(".cert-issued-by").find("span").text(certIssuedByVal);
                tr.find(".cert-issued-by").find('input[name="certIssuedBy"]').val(certIssuedByVal);
                tr.find(".cert-date-issued").find("span").text(certDateIssuedVal);
                tr.find(".cert-date-issued").find('input[name="certDateIssued"]').val(certDateIssuedVal);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".cert-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                selectAllCert.prop("checked", false);
            }

            if ($(".cert-checkbox input:checked").length == $(".cert-checkbox input").length) {
                selectAllCert.prop("checked", true);
            }

            if ($(".cert-checkbox input:checked").length == 0) {
                delCertBtn.prop("disabled", true);
            } else {
                delCertBtn.prop("disabled", false);
            }
        });

        $(".cert-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var certTitleIdVal = tr.find(".cert-title-id").find('input[name="certTitleId"]').val();
            var certRatingVal = tr.find(".cert-rating").find('input[name="certRating"]').val().toUpperCase();
            var certIssuedByVal = tr.find(".cert-issued-by").find('input[name="certIssuedBy"]').val().toUpperCase();
            var certDateIssuedVal = tr.find(".cert-date-issued").find('input[name="certDateIssued"]').val().toUpperCase();

            $("#certTitleId").val(certTitleIdVal).trigger("change");
            $("#certRating").val(certRatingVal).trigger("change");
            $("#certIssuedBy").val(certIssuedByVal).trigger("change");
            $("#certDateIssued").val(certDateIssuedVal).trigger("change");
            $("#certForm").attr("data-edit-index", tr.data("index"));
            $("#certForm").attr("data-action", "edit");
            $("#certModal").modal("show");
        });
        selectAllCert.prop("checked", false);
        $("#certModal").modal("hide");
    });

    selectAllCert.on("change", function () {
        if ($(".cert-checkbox input").length > 0) {
            $(".cert-checkbox input").prop("checked", $(this).prop("checked"));
            delCertBtn.prop("disabled", !$(this).prop("checked"));
        }
    });

    delCertBtn.on("click", function () {
        $(".cert-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                selectAllCert.prop("checked", false);
            }
        });

        if ($("#certTable tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#certModal").on("hidden.bs.modal", function () {
        $("#certForm").removeAttr("data-action");
        $("#certTitleId").removeAttr("data-parsley-required");
        $("#certTitleId").val(null).trigger("change");
        $("#certTitleId").attr("data-parsley-required", true);
        $("#certRating").val("");
        $("#certRating").parsley().reset();
        $("#certIssuedBy").val("");
        $("#certIssuedBy").parsley().reset();
        $("#certDateIssued").val("");
        $("#certDateIssued").parsley().reset();
    });

    $("#addCertBtn").on("click", function () {
        $("#certForm").attr("data-action", "add");
        $("#certModal").modal("show");
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

    $("#workExpIsRelatedToFormalEdu").on("change", function () {
        $(this).val($(this).prop("checked"));
    });

    $("#workExpForm").parsley();
    $("#workExpForm").on("submit", function (e) {
        e.preventDefault();
        var workExpNameOfCompanyFirmVal = $("#workExpNameOfCompanyFirm").val().toUpperCase();
        var workExpAddressVal = $("#workExpAddress").val().toUpperCase();
        var workExpPositionHeldIdVal = $("#workExpPositionHeldId").select2("val");
        var workExpFromVal = $("#workExpFrom").val().toUpperCase();
        var workExpToVal = $("#workExpTo").val().toUpperCase();
        var workExpIsRelatedToFormalEduVal = $("#workExpIsRelatedToFormalEdu").val();
        var workExpIsRelatedToFormalEduText = $("#workExpIsRelatedToFormalEdu").val() == "true" ? "Yes" : "No";

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
                    <td class="work-exp-position-held-id">
                        <span>` + $("#workExpPositionHeldId").select2("data")[0].text + `</span>
                        <input type="hidden" name="workExpPositionHeldId" value="` + workExpPositionHeldIdVal + `">
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
                        <span>` + workExpIsRelatedToFormalEduText + `</span>
                        <input type="hidden" name="workExpIsRelatedToFormalEdu" value="` + workExpIsRelatedToFormalEduVal + `">
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
                tr.find(".work-exp-position-held-id").find("span").text($("#workExpPositionHeldId").select2("data")[0].text);
                tr.find(".work-exp-position-held-id").find('input[name="workExpPositionHeldId"]').val(workExpPositionHeldIdVal);
                tr.find(".work-exp-from").find("span").text(workExpFromVal);
                tr.find(".work-exp-from").find('input[name="workExpFrom"]').val(workExpFromVal);
                tr.find(".work-exp-to").find("span").text(workExpToVal);
                tr.find(".work-exp-to").find('input[name="workExpTo"]').val(workExpToVal);
                tr.find(".work-exp-related-to-formal-edu").find("span").text(workExpIsRelatedToFormalEduVal);
                tr.find(".work-exp-related-to-formal-edu").find('input[name="workExpIsRelatedToFormalEdu"]').val(workExpIsRelatedToFormalEduVal);
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
            var workExpPositionHeldIdVal = tr.find(".work-exp-position-held-id").find('input[name="workExpPositionHeldId"]').val();
            var workExpFromVal = tr.find(".work-exp-from").find('input[name="workExpFrom"]').val().toUpperCase();
            var workExpToVal = tr.find(".work-exp-to").find('input[name="workExpTo"]').val().toUpperCase();
            var workExpIsRelatedToFormalEduVal = tr.find(".work-exp-related-to-formal-edu").find('input[name="workExpIsRelatedToFormalEdu"]').val();

            $("#workExpNameOfCompanyFirm").val(workExpNameOfCompanyFirmVal).trigger("change");
            $("#workExpAddress").val(workExpAddressVal).trigger("change");
            $("#workExpPositionHeldId").val(workExpPositionHeldIdVal).trigger("change");
            $("#workExpFrom").val(workExpFromVal).trigger("change");
            $("#workExpTo").val(workExpToVal).trigger("change");
            $("#workExpIsRelatedToFormalEdu").val(workExpIsRelatedToFormalEduVal).trigger("change");
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
        $("#workExpPositionHeldId").removeAttr("data-parsley-required");
        $("#workExpPositionHeldId").val(null).trigger("change");
        $("#workExpPositionHeldId").attr("data-parsley-required", true);
        $("#workExpFrom").val("");
        $("#workExpFrom").parsley().reset();
        $("#workExpTo").val("");
        $("#workExpTo").parsley().reset();
        $("#workExpIsRelatedToFormalEdu").val(false).trigger("change");
        $("#workExpIsRelatedToFormalEdu").prop("checked", false);
    });

    $("#addWorkExpBtn").on("click", function () {
        $("#workExpForm").attr("data-action", "add");
        $("#workExpModal").modal("show");
    });
});