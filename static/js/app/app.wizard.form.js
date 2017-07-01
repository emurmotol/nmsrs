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

    $("#basicInfoCivilStatHexIdRadios").find("input[type=radio]").on("change", function () {
        if ($("input[name=basicInfoCivilStatHexId]:checked").val() == "594cb5fd472e11263c3291aa") {
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

    $("#empStatHexIdRadios").find("input[type=radio]").on("change", function () {
        if ($("input[name=empStatHexId]:checked").val() == "594cb674472e11263c32992f") {
            $("#empUnEmpStatHexId").val(null).trigger("change");
            $("#empUnEmpStatHexId").attr("data-parsley-required", true);
            $("#empUnEmpStatHexId").prop("disabled", false);
            $("#empUnEmpStatHexId").focus();
        } else {
            $("#empUnEmpStatHexId").removeAttr("data-parsley-required");
            $("#empUnEmpStatHexId").val(null).trigger("change");
            $("#empUnEmpStatHexId").prop("disabled", true);
            $("#empTeminatedOverseasCountryHexId").removeAttr("data-parsley-required");
            $("#empTeminatedOverseasCountryHexId").val(null).trigger("change");
            $("#empTeminatedOverseasCountryHexId").prop("disabled", true);
        }
    });

    $("#disabHexIdRadios").find("input[type=radio]").on("change", function () {
        if ($("input[name=disabHexId]:checked").val() == "594cb622472e11263c329906") {
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
        $("#disabHexId_0").parsley().reset();

        if ($(this).prop("checked")) {
            $("#disabHexId_0").attr("data-parsley-required", true);
            $("#disabHexIdRadios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", false);
            });
        } else {
            $("#disabHexId_0").removeAttr("data-parsley-required");
            $("#disabHexIdRadios").find("input[type=radio]").each(function () {
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
            $("#otherSkillHexIds").val(null).trigger("change");
            $("#otherSkillHexIds").prop("disabled", true);
            $("#otherSkillOther").attr("data-parsley-required", true);
            $("#otherSkillOther").prop("disabled", false);
            $("#otherSkillOther").focus();
        } else {
            $("#otherSkillOther").removeAttr("data-parsley-required");
            $("#otherSkillOther").val("");
            $("#otherSkillOther").parsley().reset();
            $("#otherSkillOther").prop("disabled", true);
            $("#otherSkillHexIds").prop("disabled", false);
            $("#otherSkillHexIds").focus();
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
        $("#basicInfoProvinceHexId").val($("#basicInfoProvinceHexId").data("hex-id"));
        $("#basicInfoPlaceOfBirth").val($("#basicInfoPlaceOfBirth").val().toUpperCase());
        $("#basicInfoCivilStatOther").val($("#basicInfoCivilStatOther").val().toUpperCase());
        $("#basicInfoEmail").val($("#basicInfoEmail").val().toLowerCase());
        $("#disabOther").val($("#disabOther").val().toUpperCase());
        var formalEduArr = [];

        $("#formalEduTable tbody tr").each(function () {
            formalEduArr.push({
                "highestGradeCompletedHexId": $(this).find(".formal-edu-highest-grade-completed-hex-id").data("highest-grade-completed-hex-id"),
                "courseDegreeHexId": $(this).find(".formal-edu-course-degree-hex-id").data("course-degree-hex-id"),
                "schoolUnivHexId": $(this).find(".formal-edu-school-univ-hex-id").data("school-univ-hex-id"),
                "schoolUnivOther": $(this).find(".formal-edu-school-univ-hex-id").data("school-univ-other"),
                "yearGrad": $(this).find(".formal-edu-year-grad").data("year-grad"),
                "lastAttended": $(this).find(".formal-edu-last-attended").data("last-attended")
            });
        });
        $("#formalEduJson").val(JSON.stringify(formalEduArr));
        var proLicenseArr = [];

        $("#proLicenseTable tbody tr").each(function () {
            proLicenseArr.push({
                "titleHexId": $(this).find(".pro-license-title-hex-id").data("title-hex-id"),
                "expiryDate": $(this).find(".pro-license-expiry-date").data("expiry-date")
            });
        });
        $("#proLicenseJson").val(JSON.stringify(proLicenseArr));
        var eligArr = [];

        $("#eligTable tbody tr").each(function () {
            eligArr.push({
                "titleHexId": $(this).find(".elig-title-hex-id").data("title-hex-id"),
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
                "titleHexId": $(this).find(".cert-title-hex-id").data("title-hex-id"),
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
                "positionHeldHexId": $(this).find(".work-exp-position-held-hex-id").data("position-held-hex-id"),
                "from": $(this).find(".work-exp-from").data("from"),
                "to": $(this).find(".work-exp-to").data("to"),
                "isRelatedToFormalEdu": $(this).find(".work-exp-is-related-to-formal-edu").data("is-related-to-formal-edu")
            });
        });
        $("#workExpJson").val(JSON.stringify(workExpArr));
        duringSubmitDo(this);
    });
});