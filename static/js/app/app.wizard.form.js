$(function () {


    $("#empIsActivelyLookingForWork").on("change", function () {
        $(this).val($(this).prop("checked"));
    });

    $("#iAccept").on("change", function () {
        $(this).val($(this).prop("checked"));
    });

    $("#personalInfoBirthdate").datetimepicker({
        viewMode: "years",
        format: "YYYY-MM-DD"
    });

    $("#personalInfoBirthdate").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#empPassportNumberExpiryDate").datetimepicker({
        viewMode: "years",
        format: "YYYY-MM"
    });

    $("#empPassportNumberExpiryDate").on("dp.change", function () {
        $(this).parsley().validate();
    });

    $("#basicInfoCivilStatIdRadios").find("input[type=radio]").on("change", function () {
        if ($("input[name=basicInfoCivilStatId]:checked").val() == "594cb5fd472e11263c3291aa") {
            $("#basicInfoCivilStatOther").attr("data-parsley-required", true);
            $("#basicInfoCivilStatOther").prop("disabled", false);
            $("#basicInfoCivilStatOther").focus();
        } else {
            $("#basicInfoCivilStatOther").removeAttr("data-parsley-required");
            $("#basicInfoCivilStatOther").val("");
            $("#basicInfoCivilStatOther").parsley().reset();
            $("#basicInfoCivilStatOther").prop("disabled", true);
        }
    });

    $("#empStatIdRadios").find("input[type=radio]").on("change", function () {
        if ($("input[name=empStatId]:checked").val() == "594cb674472e11263c32992f") {
            $("#empUnEmpStatId").val(null).trigger("change");
            $("#empUnEmpStatId").attr("data-parsley-required", true);
            $("#empUnEmpStatId").prop("disabled", false);
            $("#empUnEmpStatId").focus();
        } else {
            $("#empUnEmpStatId").removeAttr("data-parsley-required");
            $("#empUnEmpStatId").val(null).trigger("change");
            $("#empUnEmpStatId").prop("disabled", true);
            $("#empTeminatedOverseasCountryId").removeAttr("data-parsley-required");
            $("#empTeminatedOverseasCountryId").val(null).trigger("change");
            $("#empTeminatedOverseasCountryId").prop("disabled", true);
        }
    });

    $("#disabIdRadios").find("input[type=radio]").on("change", function () {
        if ($("input[name=disabId]:checked").val() == "594cb622472e11263c329906") {
            $("#disabOther").attr("data-parsley-required", true);
            $("#disabOther").prop("disabled", false);
            $("#disabOther").focus();
        } else {
            $("#disabOther").removeAttr("data-parsley-required");
            $("#disabOther").val("");
            $("#disabOther").parsley().reset();
            $("#disabOther").prop("disabled", true);
        }
    });

    $("#disabIsDisabled").on("change", function () {
        $(this).val($(this).prop("checked"));
        $("#disabId_0").parsley().reset();

        if ($(this).prop("checked")) {
            $("#disabId_0").attr("data-parsley-required", true);
            $("#disabIdRadios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", false);
            });
        } else {
            $("#disabId_0").removeAttr("data-parsley-required");
            $("#disabIdRadios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", true);
                $(this).prop("checked", false);
            });
            $("#disabOther").removeAttr("data-parsley-required");
            $("#disabOther").val(null).trigger("change");
            $("#disabOther").prop("disabled", true);
        }
    });

    $("#otherSkillNotListed").on("change", function () {
        $(this).val($(this).prop("checked"));

        if ($(this).prop("checked")) {
            $("#otherSkillIds").val(null).trigger("change");
            $("#otherSkillIds").prop("disabled", true);
            $("#otherSkillOther").attr("data-parsley-required", true);
            $("#otherSkillOther").prop("disabled", false);
            $("#otherSkillOther").focus();
        } else {
            $("#otherSkillOther").removeAttr("data-parsley-required");
            $("#otherSkillOther").val("");
            $("#otherSkillOther").parsley().reset();
            $("#otherSkillOther").prop("disabled", true);
            $("#otherSkillIds").prop("disabled", false);
            $("#otherSkillIds").focus();
        }
    });

    $("#registeredAt").datetimepicker({
        defaultDate: new Date(),
        format: "YYYY-MM-DD"
    });

    $("#registeredAt").on("dp.change", function () {
        $(this).parsley().validate();
    });

    previewImage("#personalInfoPhoto");

    $("#personalInfoGivenName").on("keyup", function () {
        $("#personalInfoPassword").val($(this).val().replace(/\s/g, "-") + "-" + $("#personalInfoFamilyName").val().replace(/\s/g, "-"));
    });

    $("#personalInfoFamilyName").on("keyup", function () {
        $("#personalInfoPassword").val($("#personalInfoGivenName").val().replace(/\s/g, "-") + "-" + $(this).val().replace(/\s/g, "-"));
    });

    $("#createRegistrantForm").parsley();
    $("#createRegistrantForm").on("submit", function () {
        $("#personalInfoFamilyName").val($("#personalInfoFamilyName").val().toUpperCase());
        $("#personalInfoGivenName").val($("#personalInfoGivenName").val().toUpperCase());
        $("#personalInfoMiddleName").val($("#personalInfoMiddleName").val().toUpperCase());
        $("#personalInfoPassword").val($("#personalInfoPassword").val().toUpperCase());
        $("#basicInfoStSub").val($("#basicInfoStSub").val().toUpperCase());
        $("#basicInfoProvinceId").val($("#basicInfoProvinceId").data("id"));
        $("#basicInfoPlaceOfBirth").val($("#basicInfoPlaceOfBirth").val().toUpperCase());
        $("#basicInfoCivilStatOther").val($("#basicInfoCivilStatOther").val().toUpperCase());
        $("#basicInfoEmail").val($("#basicInfoEmail").val().toLowerCase());
        $("#disabOther").val($("#disabOther").val().toUpperCase());
        var formalEduArr = [];

        $("#formalEduTable tbody tr").each(function () {
            formalEduArr.push({
                "highestGradeCompletedId": $(this).find(".formal-edu-highest-grade-completed-id").data("highest-grade-completed-id"),
                "courseDegreeId": $(this).find(".formal-edu-course-degree-id").data("course-degree-id"),
                "schoolUnivId": $(this).find(".formal-edu-school-univ-id").data("school-univ-id"),
                "schoolUnivOther": $(this).find(".formal-edu-school-univ-id").data("school-univ-other"),
                "yearGrad": $(this).find(".formal-edu-year-grad").data("year-grad"),
                "lastAttended": $(this).find(".formal-edu-last-attended").data("last-attended")
            });
        });
        $("#formalEduJson").val(JSON.stringify(formalEduArr));
        var proLicenseArr = [];

        $("#proLicenseTable tbody tr").each(function () {
            proLicenseArr.push({
                "titleId": $(this).find(".pro-license-title-id").data("title-id"),
                "expiryDate": $(this).find(".pro-license-expiry-date").data("expiry-date")
            });
        });
        $("#proLicenseJson").val(JSON.stringify(proLicenseArr));
        var eligArr = [];

        $("#eligTable tbody tr").each(function () {
            eligArr.push({
                "titleId": $(this).find(".elig-title-id").data("title-id"),
                "yearTaken": $(this).find(".elig-year-taken").data("year-taken")
            });
        });
        $("#eligJson").val(JSON.stringify(eligArr));
        var trainingArr = [];

        $("#trainingTable tbody tr").each(function () {
            trainingArr.push({
                "nameOfTraining": $(this).find(".training-name-of-training").data("name-of-training"),
                "skillsAcquired": $(this).find(".training-skills-acquired").data("skills-acquired"),
                "periodOfTrainingExp": $(this).find(".training-period-of-training-exp").data("period-of-training-exp"),
                "certReceived": $(this).find(".training-cert-received").data("cert-received"),
                "issuingSchoolAgency": $(this).find(".training-issuing-school-agency").data("issuing-school-agency")
            });
        });
        $("#trainingJson").val(JSON.stringify(trainingArr));
        var certArr = [];

        $("#certTable tbody tr").each(function () {
            certArr.push({
                "titleId": $(this).find(".cert-title-id").data("title-id"),
                "rating": $(this).find(".cert-rating").data("rating"),
                "issuedBy": $(this).find(".cert-issued-by").data("issued-by"),
                "dateIssued": $(this).find(".cert-date-issued").data("date-issued")
            });
        });
        $("#certJson").val(JSON.stringify(certArr));
        var workExpArr = [];

        $("#workExpTable tbody tr").each(function () {
            workExpArr.push({
                "nameOfCompanyFirm": $(this).find(".work-exp-name-of-company-firm").data("name-of-company-firm"),
                "address": $(this).find(".work-exp-address").data("address"),
                "positionHeldId": $(this).find(".work-exp-position-held-id").data("position-held-id"),
                "from": $(this).find(".work-exp-from").data("from"),
                "to": $(this).find(".work-exp-to").data("to"),
                "isRelatedToFormalEdu": $(this).find(".work-exp-is-related-to-formal-edu").data("is-related-to-formal-edu")
            });
        });
        $("#workExpJson").val(JSON.stringify(workExpArr));
        duringSubmitDo(this);
    });
});