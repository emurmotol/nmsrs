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
        viewMode: "years",
        format: "YYYY-MM"
    });

    $("#formalEduLastAttended").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#formalEduForm").parsley();
    $("#formalEduForm").on("submit", function (e) {
        e.preventDefault();
        var formalEduHighestGradeCompletedId = $("#formalEduHighestGradeCompletedId").select2("val");
        var formalEduCourseDegreeId = $("#formalEduCourseDegreeId").select2("val");
        var formalEduSchoolUnivId = $("#formalEduSchoolUnivId").select2("val") != null ? $("#formalEduSchoolUnivId").select2("val") : "";
        var formalEduSchoolUnivOther = $("#formalEduSchoolUnivOther").val().toUpperCase();
        var formalEduSchoolUnivText = formalEduSchoolUnivOther;
        var formalEduYearGrad = $("#formalEduYearGrad").val();
        var formalEduLastAttended = $("#formalEduLastAttended").val();

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
                    <td class="formal-edu-highest-grade-completed-id" data-highest-grade-completed-id="`+ formalEduHighestGradeCompletedId + `">
                        ` + $("#formalEduHighestGradeCompletedId").select2("data")[0].text + `
                    </td>
                    <td class="formal-edu-course-degree-id" data-course-degree-id="`+ formalEduCourseDegreeId + `">
                        ` + $("#formalEduCourseDegreeId").select2("data")[0].text + `
                    </td>
                    <td class="formal-edu-school-univ-id" data-school-univ-id="`+ formalEduSchoolUnivId + `" data-school-univ-other="` + formalEduSchoolUnivOther + `">
                        ` + formalEduSchoolUnivText + `
                    </td>
                    <td class="formal-edu-year-grad" data-year-grad="`+ formalEduYearGrad + `">
                        ` + formalEduYearGrad + `
                    </td>
                    <td class="formal-edu-last-attended" data-last-attended="`+ formalEduLastAttended + `">
                        ` + formalEduLastAttended + `
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
                tr.find(".formal-edu-highest-grade-completed-id").text($("#formalEduHighestGradeCompletedId").select2("data")[0].text);
                tr.find(".formal-edu-highest-grade-completed-id").data("highest-grade-completed-id", formalEduHighestGradeCompletedId);
                tr.find(".formal-edu-course-degree-id").text($("#formalEduCourseDegreeId").select2("data")[0].text);
                tr.find(".formal-edu-course-degree-id").data("course-degree-id", formalEduCourseDegreeId);
                tr.find(".formal-edu-school-univ-id").text(formalEduSchoolUnivText);
                tr.find(".formal-edu-school-univ-id").data("school-univ-id", formalEduSchoolUnivId);
                tr.find(".formal-edu-school-univ-id").data("school-univ-other", formalEduSchoolUnivOther);
                tr.find(".formal-edu-year-grad").text(formalEduYearGrad);
                tr.find(".formal-edu-year-grad").data("year-grad", formalEduYearGrad);
                tr.find(".formal-edu-last-attended").text(formalEduLastAttended);
                tr.find(".formal-edu-last-attended").data("last-attended", formalEduLastAttended);
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
            var formalEduHighestGradeCompletedId = tr.find(".formal-edu-highest-grade-completed-id").data("highest-grade-completed-id");
            var formalEduCourseDegreeId = tr.find(".formal-edu-course-degree-id").data("course-degree-id");
            var formalEduSchoolUnivId = tr.find(".formal-edu-school-univ-id").data("school-univ-id");
            var formalEduYearGrad = tr.find(".formal-edu-year-grad").data("year-grad");
            var formalEduLastAttended = tr.find(".formal-edu-last-attended").data("last-attended");
            var formalEduSchoolUnivNotListed = $("#formalEduSchoolUnivNotListed").val();

            $("#formalEduHighestGradeCompletedId").val(formalEduHighestGradeCompletedId).trigger("change");
            $("#formalEduCourseDegreeId").val(formalEduCourseDegreeId).trigger("change");

            if (formalEduSchoolUnivId == "") {
                var formalEduSchoolUnivOther = tr.find(".formal-edu-school-univ-id").data("school-univ-other");
                $("#formalEduSchoolUnivOther").val(formalEduSchoolUnivOther).trigger("change");
            }

            if (formalEduSchoolUnivNotListed == "true") {
                $("#formalEduSchoolUnivNotListed").prop("checked", true).trigger("change");
            } else {
                $("#formalEduSchoolUnivId").val(formalEduSchoolUnivId).trigger("change");
                $("#formalEduSchoolUnivNotListed").prop("checked", false).trigger("change");
            }
            $("#formalEduYearGrad").val(formalEduYearGrad).trigger("change");
            $("#formalEduLastAttended").val(formalEduLastAttended).trigger("change");
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
        viewMode: "years",
        format: "YYYY-MM"
    });

    $("#proLicenseExpiryDate").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#proLicenseForm").parsley();
    $("#proLicenseForm").on("submit", function (e) {
        e.preventDefault();
        var proLicenseTitleId = $("#proLicenseTitleId").select2("val");
        var proLicenseExpiryDate = $("#proLicenseExpiryDate").val();

        switch ($(this).attr("data-action")) {
            case "add":
                var proLicenseIndex = $("#proLicenseTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + proLicenseIndex + `">
                    <td class="pro-license-checkbox">
                        <input type="checkbox" class="checkbox" id="proLicenseCheckbox_` + proLicenseIndex + `">
                    </td>
                    <td class="pro-license-title-id" data-title-id="`+ proLicenseTitleId + `">
                        ` + $("#proLicenseTitleId").select2("data")[0].text + `
                    </td>
                    <td class="pro-license-expiry-date" data-expiry-date="`+ proLicenseExpiryDate + `">
                        ` + proLicenseExpiryDate + `
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
                tr.find(".pro-license-title-id").text($("#proLicenseTitleId").select2("data")[0].text);
                tr.find(".pro-license-title-id").data("title-id", proLicenseTitleId);
                tr.find(".pro-license-expiry-date").text(proLicenseExpiryDate);
                tr.find(".pro-license-expiry-date").data("expiry-date", proLicenseExpiryDate);
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
            var proLicenseTitleId = tr.find(".pro-license-title-id").data("title-id");
            var proLicenseExpiryDate = tr.find(".pro-license-expiry-date").data("expiry-date");

            $("#proLicenseTitleId").val(proLicenseTitleId).trigger("change");
            $("#proLicenseExpiryDate").val(proLicenseExpiryDate).trigger("change");
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
        var eligTitleId = $("#eligTitleId").select2("val");
        var eligYearTaken = $("#eligYearTaken").val();

        switch ($(this).attr("data-action")) {
            case "add":
                var eligIndex = $("#eligTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + eligIndex + `">
                    <td class="elig-checkbox">
                        <input type="checkbox" class="checkbox" id="eligCheckbox_` + eligIndex + `">
                    </td>
                    <td class="elig-title-id" data-title-id="`+ eligTitleId + `">
                        ` + $("#eligTitleId").select2("data")[0].text + `
                    </td>
                    <td class="elig-year-taken" data-year-taken="`+ eligYearTaken + `">
                        ` + eligYearTaken + `
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
                tr.find(".elig-title-id").text($("#eligTitleId").select2("data")[0].text);
                tr.find(".elig-title-id").data("title-id", eligTitleId);
                tr.find(".elig-year-taken").text(eligYearTaken);
                tr.find(".elig-year-taken").data("year-taken", eligYearTaken);
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
            var eligTitleId = tr.find(".elig-title-id").data("title-id");
            var eligYearTaken = tr.find(".elig-year-taken").data("year-taken");

            $("#eligTitleId").val(eligTitleId).trigger("change");
            $("#eligYearTaken").val(eligYearTaken).trigger("change");
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
        var trainingNameOfTraining = $("#trainingNameOfTraining").val().toUpperCase();
        var trainingSkillsAcquired = $("#trainingSkillsAcquired").val().toUpperCase();
        var trainingPeriodOfTrainingExp = $("#trainingPeriodOfTrainingExp").val().toUpperCase();
        var trainingCertReceived = $("#trainingCertReceived").val().toUpperCase();
        var trainingIssuingSchoolAgency = $("#trainingIssuingSchoolAgency").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var trainingIndex = $("#trainingTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + trainingIndex + `">
                    <td class="training-checkbox">
                        <input type="checkbox" class="checkbox" id="trainingCheckbox_` + trainingIndex + `">
                    </td>
                    <td class="training-name-of-training" data-name-of-training="`+ trainingNameOfTraining + `">
                        ` + trainingNameOfTraining + `
                    </td>
                    <td class="training-skills-acquired" data-skills-acquired="`+ trainingSkillsAcquired + `">
                        ` + trainingSkillsAcquired + `
                    </td>
                    <td class="training-period-of-training-exp" data-period-of-training-exp="`+ trainingPeriodOfTrainingExp + `">
                        ` + trainingPeriodOfTrainingExp + `
                    </td>
                    <td class="training-cert-received" data-cert-received="`+ trainingCertReceived + `">
                        ` + trainingCertReceived + `
                    </td>
                    <td class="training-issuing-school-agency" data-issuing-school-agency="`+ trainingIssuingSchoolAgency + `">
                        ` + trainingIssuingSchoolAgency + `
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
                tr.find(".training-name-of-training").text(trainingNameOfTraining);
                tr.find(".training-name-of-training").data("name-of-training", trainingNameOfTraining);
                tr.find(".training-skills-acquired").text(trainingSkillsAcquired);
                tr.find(".training-skills-acquired").data("skills-acquired", trainingSkillsAcquired);
                tr.find(".training-period-of-training-exp").text(trainingPeriodOfTrainingExp);
                tr.find(".training-period-of-training-exp").data("period-of-training-exp", trainingPeriodOfTrainingExp);
                tr.find(".training-cert-received").text(trainingCertReceived);
                tr.find(".training-cert-received").data("cert-received", trainingCertReceived);
                tr.find(".training-issuing-school-agency").text(trainingIssuingSchoolAgency);
                tr.find(".training-issuing-school-agency").data("issuing-school-agency", trainingIssuingSchoolAgency);
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
            var trainingNameOfTraining = tr.find(".training-name-of-training").data("name-of-training");
            var trainingSkillsAcquired = tr.find(".training-skills-acquired").data("skills-acquired");
            var trainingPeriodOfTrainingExp = tr.find(".training-period-of-training-exp").data("period-of-training-exp");
            var trainingCertReceived = tr.find(".training-cert-received").data("cert-received");
            var trainingIssuingSchoolAgency = tr.find(".training-issuing-school-agency").data("issuing-school-agency");

            $("#trainingNameOfTraining").val(trainingNameOfTraining).trigger("change");
            $("#trainingSkillsAcquired").val(trainingSkillsAcquired).trigger("change");
            $("#trainingPeriodOfTrainingExp").val(trainingPeriodOfTrainingExp).trigger("change");
            $("#trainingCertReceived").val(trainingCertReceived).trigger("change");
            $("#trainingIssuingSchoolAgency").val(trainingIssuingSchoolAgency).trigger("change");
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
        viewMode: "years",
        format: "YYYY-MM"
    });

    $("#certDateIssued").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#certForm").parsley();
    $("#certForm").on("submit", function (e) {
        e.preventDefault();
        var certTitleId = $("#certTitleId").select2("val");
        var certRating = $("#certRating").val().toUpperCase();
        var certIssuedBy = $("#certIssuedBy").val().toUpperCase();
        var certDateIssued = $("#certDateIssued").val();

        switch ($(this).attr("data-action")) {
            case "add":
                var certIndex = $("#certTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + certIndex + `">
                    <td class="certCheckbox_">
                        <input type="checkbox" class="checkbox" id="certCheckbox_` + certIndex + `">
                    </td>
                    <td class="cert-title-id" data-title-id="`+ certTitleId + `">
                        ` + $("#certTitleId").select2("data")[0].text + `
                    </td>
                    <td class="cert-rating" data-rating="`+ certRating + `">
                        ` + certRating + `
                    </td>
                    <td class="cert-issued-by" data-issued-by="`+ certIssuedBy + `">
                        ` + certIssuedBy + `
                    </td>
                    <td class="cert-date-issued" data-date-issued="`+ certDateIssued + `">
                        ` + certDateIssued + `
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
                tr.find(".cert-title-id").text($("#certTitleId").select2("data")[0].text);
                tr.find(".cert-title-id").data("title-id", certTitleId);
                tr.find(".cert-rating").text(certRating);
                tr.find(".cert-rating").data("rating", certRating);
                tr.find(".cert-issued-by").text(certIssuedBy);
                tr.find(".cert-issued-by").data("issued-by", certIssuedBy);
                tr.find(".cert-date-issued").text(certDateIssued);
                tr.find(".cert-date-issued").data("date-issued", certDateIssued);
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
            var certTitleId = tr.find(".cert-title-id").data("title-id");
            var certRating = tr.find(".cert-rating").data("rating");
            var certIssuedBy = tr.find(".cert-issued-by").data("issued-by");
            var certDateIssued = tr.find(".cert-date-issued").data("date-issued");

            $("#certTitleId").val(certTitleId).trigger("change");
            $("#certRating").val(certRating).trigger("change");
            $("#certIssuedBy").val(certIssuedBy).trigger("change");
            $("#certDateIssued").val(certDateIssued).trigger("change");
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
        viewMode: "years",
        format: "YYYY-MM"
    });

    $("#workExpFrom").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#workExpTo").datetimepicker({
        viewMode: "years",
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
        var workExpNameOfCompanyFirm = $("#workExpNameOfCompanyFirm").val().toUpperCase();
        var workExpAddress = $("#workExpAddress").val().toUpperCase();
        var workExpPositionHeldId = $("#workExpPositionHeldId").select2("val");
        var workExpFrom = $("#workExpFrom").val();
        var workExpTo = $("#workExpTo").val();
        var workExpIsRelatedToFormalEdu = $("#workExpIsRelatedToFormalEdu").val();
        var workExpIsRelatedToFormalEduText = workExpIsRelatedToFormalEdu == "true" ? "Yes" : "No";

        switch ($(this).attr("data-action")) {
            case "add":
                var workExpIndex = $("#workExpTable tbody tr").index() + 1;
                var row = `
                <tr data-index="` + workExpIndex + `">
                    <td class="work-exp-checkbox">
                        <input type="checkbox" class="checkbox" id="workExpCheckbox_` + workExpIndex + `">
                    </td>
                    <td class="work-exp-name-of-company-firm" data-name-of-company-firm="`+ workExpNameOfCompanyFirm + `">
                        ` + workExpNameOfCompanyFirm + `
                    </td>
                    <td class="work-exp-address" data-address="`+ workExpAddress + `">
                        ` + workExpAddress + `
                    </td>
                    <td class="work-exp-position-held-id" data-position-held-id="`+ workExpPositionHeldId + `">
                        ` + $("#workExpPositionHeldId").select2("data")[0].text + `
                    </td>
                    <td class="work-exp-from" data-from="`+ workExpFrom + `">
                        ` + workExpFrom + `
                    </td>
                    <td class="work-exp-to" data-to="`+ workExpTo + `">
                        ` + workExpTo + `
                    </td>
                    <td class="work-exp-related-to-formal-edu" data-related-to-formal-edu="`+ workExpIsRelatedToFormalEdu + `">
                        ` + workExpIsRelatedToFormalEduText + `
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
                tr.find(".work-exp-name-of-company-firm").text(workExpNameOfCompanyFirm);
                tr.find(".work-exp-name-of-company-firm").data("name-of-company-firm", workExpNameOfCompanyFirm);
                tr.find(".work-exp-address").text(workExpAddress);
                tr.find(".work-exp-address").data("address", workExpAddress);
                tr.find(".work-exp-position-held-id").text($("#workExpPositionHeldId").select2("data")[0].text);
                tr.find(".work-exp-position-held-id").data("position-held-id", workExpPositionHeldId);
                tr.find(".work-exp-from").text(workExpFrom);
                tr.find(".work-exp-from").data("from", workExpFrom);
                tr.find(".work-exp-to").text(workExpTo);
                tr.find(".work-exp-to").data("to", workExpTo);
                tr.find(".work-exp-related-to-formal-edu").text(workExpIsRelatedToFormalEduText);
                tr.find(".work-exp-related-to-formal-edu").data("related-to-formal-edu", workExpIsRelatedToFormalEdu);
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
            var workExpNameOfCompanyFirm = tr.find(".work-exp-name-of-company-firm").data("name-of-company-firm");
            var workExpAddress = tr.find(".work-exp-address").data("address");
            var workExpPositionHeldId = tr.find(".work-exp-position-held-id").data("position-held-id");
            var workExpFrom = tr.find(".work-exp-from").data("from");
            var workExpTo = tr.find(".work-exp-to").data("to");
            var workExpIsRelatedToFormalEdu = tr.find(".work-exp-related-to-formal-edu").data("related-to-formal-edu");

            $("#workExpNameOfCompanyFirm").val(workExpNameOfCompanyFirm).trigger("change");
            $("#workExpAddress").val(workExpAddress).trigger("change");
            $("#workExpPositionHeldId").val(workExpPositionHeldId).trigger("change");
            $("#workExpFrom").val(workExpFrom).trigger("change");
            $("#workExpTo").val(workExpTo).trigger("change");

            if (workExpIsRelatedToFormalEdu == "true") {
                $("#workExpIsRelatedToFormalEdu").prop("checked", true).trigger("change");
            } else {
                $("#workExpIsRelatedToFormalEdu").prop("checked", false).trigger("change");
            }
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