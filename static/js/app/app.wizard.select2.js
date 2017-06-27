$(function () {
    $("#workExpPositionHeld").select2({
        placeholder: "SELECT POSITION",
        ajax: {
            url: "/api/positions",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#otherSkills").select2({
        placeholder: "SELECT SKILL(S)",
        ajax: {
            url: "/api/otherskills",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#certOfCompetenceTitle").select2({
        placeholder: "SELECT CERTIFICATE TITLE",
        ajax: {
            url: "/api/certificates",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#eligTitle").select2({
        placeholder: "SELECT ELIGIBILITY TITLE",
        ajax: {
            url: "/api/eligibilities",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#proLicenseTitle").select2({
        placeholder: "SELECT LICENSE TITLE",
        ajax: {
            url: "/api/licenses",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#formalEduSchoolUniv").select2({
        placeholder: "SELECT SCHOOL/UNIVERSITY",
        ajax: {
            url: "/api/schools",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#formalEduCourseDegree").select2({
        placeholder: "SELECT COURSE/DEGREE",
        ajax: {
            url: "/api/courses",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#formalEduHighestGradeCompleted").select2({
        placeholder: "SELECT HIGHEST GRADE COMPLETED",
        ajax: {
            url: "/api/edulevels",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empPrefOccs").select2({
        placeholder: "SELECT PREFERRED OCCUPATION(S)",
        ajax: {
            url: "/api/positions",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empPrefLocalLoc").select2({
        placeholder: "SELECT PREFERRED LOCAL LOCATION",
        ajax: {
            url: "/api/citymuns/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.cityMunId,
                            text: data.cityMunDesc + ", " + data.provDesc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empPrefOverseasLoc").select2({
        placeholder: "SELECT PREFERRED OVERSEAS LOCATION",
        ajax: {
            url: "/api/countries",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#langs").select2({
        placeholder: "SELECT LANGUAGE(S)",
        ajax: {
            url: "/api/languages",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#basicInfoReligion").select2({
        placeholder: "SELECT RELIGION",
        ajax: {
            url: "/api/religions",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#basicInfoBarangay").select2({ placeholder: "SELECT BARANGAY" });
    $("#basicInfoBarangay").on("change", function () {
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
                            id: data.cityMunId,
                            text: data.cityMunDesc + ", " + data.provDesc,
                            provId: data.provId,
                            provDesc: data.provDesc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#basicInfoCityMunId").on("change", function () {
        var data = $(this).select2("data")[0];
        $("input[name=basicInfoCityMun]").val(data.text);
        $("#basicInfoProvince").val(data.provDesc);

        $("#basicInfoBarangay").removeAttr("data-parsley-required");
        $("#basicInfoBarangay").val(null).trigger("change");
        $("#basicInfoBarangay").attr("data-parsley-required", true);
        $("#basicInfoBarangay").select2({
            placeholder: "SELECT BARANGAY",
            ajax: {
                url: "/api/citymuns/" + $("#basicInfoCityMunId").select2("val") + "/barangays",
                delay: 250,
                dataType: "json",
                processResults: function (r) {
                    return {
                        results: $.map(r, function (data) {
                            return {
                                id: data.desc,
                                text: data.desc
                            };
                        })
                    };
                },
                cache: true
            }
        });
        $("#basicInfoBarangay").prop("disabled", false);
        $("#basicInfoBarangay").focus();
    });

    $("#empTeminatedOverseasCountry").select2({
        placeholder: "SELECT COUNTRY GOT TERMINATED",
        ajax: {
            url: "/api/countries",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empUnEmpStat").select2({
        placeholder: "SELECT UNEMPLOYED STATUS",
        ajax: {
            url: "/api/unempstats",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.id,
                            text: data.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#empUnEmpStat").on("change", function () {
        if ($(this).val() == "TERMINATED/LAID OFF, OVERSEAS") {
            $("#empTeminatedOverseasCountry").attr("data-parsley-required", true);
            $("#empTeminatedOverseasCountry").prop("disabled", false);
            $("#empTeminatedOverseasCountry").focus();
        } else {
            $("#empTeminatedOverseasCountry").removeAttr("data-parsley-required");
            $("#empTeminatedOverseasCountry").val(null).trigger("change");
            $("#empTeminatedOverseasCountry").prop("disabled", true);
        }
    });

    $("select").on("change", function () {
        var instance = $(this).parsley();
        if (instance.isValid()) {
            instance.reset();
        }
    });
});