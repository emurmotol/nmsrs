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

    $("#personalInfoBirthdate").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#empPassportNumberExpiryDate").datetimepicker({
        format: "YYYY-MM"
    });

    $("#empPassportNumberExpiryDate").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#basicInfoCivilStatIdRadios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=basicInfoCivilStatId]:checked");

        if (checked[0].id == "basicInfoCivilStatId_4") {
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
        var checked = $("input[name=empStatId]:checked");

        if (checked[0].id == "empStatId_2") {
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
        var checked = $("input[name=disabId]:checked");

        if (checked[0].id == "disabId_4") {
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

    $("#registeredAt").on("dp.change", function() {
        $(this).parsley().validate();
    });

    previewImage("#personalInfoPhoto");

    $("#createRegistrantForm").parsley();
    $("#createRegistrantForm").on("submit", function () {
        duringSubmitDo(this);
    });
});