$(function () {
    $("#workExpPositionHeldId").select2({
        placeholder: "SELECT POSITION",
        ajax: {
            url: "/api/positions",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#otherSkillIds").select2({
        placeholder: "SELECT SKILL(S)",
        ajax: {
            url: "/api/otherskills",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#certTitleId").select2({
        placeholder: "SELECT CERTIFICATE TITLE",
        ajax: {
            url: "/api/certificates",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#eligTitleId").select2({
        placeholder: "SELECT ELIGIBILITY TITLE",
        ajax: {
            url: "/api/eligibilities",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#proLicenseTitleId").select2({
        placeholder: "SELECT LICENSE TITLE",
        ajax: {
            url: "/api/licenses",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#formalEduSchoolUnivId").select2({
        placeholder: "SELECT SCHOOL/UNIVERSITY",
        ajax: {
            url: "/api/schools",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#formalEduCourseDegreeId").select2({
        placeholder: "SELECT COURSE/DEGREE",
        ajax: {
            url: "/api/courses",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#formalEduHighestGradeCompletedId").select2({
        placeholder: "SELECT HIGHEST GRADE COMPLETED",
        ajax: {
            url: "/api/edulevels",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empPrefOccIds").select2({
        placeholder: "SELECT PREFERRED OCCUPATION(S)",
        ajax: {
            url: "/api/positions",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empPrefLocalLocId").select2({
        placeholder: "SELECT PREFERRED LOCAL LOCATION",
        ajax: {
            url: "/api/citymuns/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.desc + ", " + data.province[0].desc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empPrefOverseasLocId").select2({
        placeholder: "SELECT PREFERRED OVERSEAS LOCATION",
        ajax: {
            url: "/api/countries",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#langIds").select2({
        placeholder: "SELECT LANGUAGE(S)",
        ajax: {
            url: "/api/languages",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#basicInfoReligionId").select2({
        placeholder: "SELECT RELIGION",
        ajax: {
            url: "/api/religions",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#basicInfoBarangayId").select2({ placeholder: "SELECT BARANGAY" });
    $("#basicInfoBarangayId").on("change", function () {
        $("#basicInfoPlaceOfBirth").focus();
    });

    $("#basicInfoCityMunId").select2({
        placeholder: "SELECT CITY/MUNICIPALITY",
        ajax: {
            url: "/api/citymuns/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.desc + ", " + data.province[0].desc,
                            provId: data.province[0]._id,
                            provDesc: data.province[0].desc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#basicInfoCityMunId").on("change", function () {
        var data = $(this).select2("data")[0];
        $("#BasicInfoProvinceId").attr("data-id", data.provId);
        $("#BasicInfoProvinceId").val(data.provDesc);

        $("#basicInfoBarangayId").removeAttr("data-parsley-required");
        $("#basicInfoBarangayId").val(null).trigger("change");
        $("#basicInfoBarangayId").attr("data-parsley-required", true);
        $("#basicInfoBarangayId").select2({
            placeholder: "SELECT BARANGAY",
            ajax: {
                url: "/api/citymuns/" + $("#basicInfoCityMunId").select2("val") + "/barangays",
                delay: 250,
                dataType: "json",
                processResults: function (r) {
                    return {
                        results: $.map(r, function (data) {
                            return {
                                id: data._id,
                                text: data.desc
                            };
                        })
                    };
                },
                cache: true
            }
        });
        $("#basicInfoBarangayId").prop("disabled", false);
        $("#basicInfoBarangayId").focus();
    });

    $("#empTeminatedOverseasCountryId").select2({
        placeholder: "SELECT COUNTRY GOT TERMINATED",
        ajax: {
            url: "/api/countries",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empUnEmpStatId").select2({
        placeholder: "SELECT UNEMPLOYED STATUS",
        ajax: {
            url: "/api/unempstats",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data._id,
                            text: data.value
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empUnEmpStatId").on("change", function () {
        if ($(this).select2("val") == "594cb702472e11263c32c495") {
            $("#empTeminatedOverseasCountryId").attr("data-parsley-required", true);
            $("#empTeminatedOverseasCountryId").prop("disabled", false);
            $("#empTeminatedOverseasCountryId").focus();
        } else {
            $("#empTeminatedOverseasCountryId").removeAttr("data-parsley-required");
            $("#empTeminatedOverseasCountryId").val(null).trigger("change");
            $("#empTeminatedOverseasCountryId").prop("disabled", true);
        }
    });

    $("select").on("change", function () {
        var instance = $(this).parsley();
        if (instance.isValid()) {
            instance.reset();
        }
    });
});