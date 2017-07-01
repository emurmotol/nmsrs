$(function () {
    $("#workExpPositionHeldHexId").select2({
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

    $("#otherSkillHexIds").select2({
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

    $("#certTitleHexId").select2({
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

    $("#eligTitleHexId").select2({
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

    $("#proLicenseTitleHexId").select2({
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

    $("#formalEduSchoolUnivHexId").select2({
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

    $("#formalEduCourseDegreeHexId").select2({
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

    $("#formalEduHighestGradeCompletedHexId").select2({
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

    $("#empPrefOccHexIds").select2({
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

    $("#empPrefLocalLocHexId").select2({
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

    $("#empPrefOverseasLocHexId").select2({
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

    $("#langHexIds").select2({
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

    $("#basicInfoReligionHexId").select2({
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

    $("#basicInfoBarangayHexId").select2({ placeholder: "SELECT BARANGAY" });
    $("#basicInfoBarangayHexId").on("change", function () {
        $("#basicInfoPlaceOfBirth").focus();
    });

    $("#basicInfoCityMunHexId").select2({
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

    $("#basicInfoCityMunHexId").on("change", function () {
        var data = $(this).select2("data")[0];
        $("#basicInfoProvinceHexId").attr("data-hex-id", data.provId);
        $("#basicInfoProvinceHexId").val(data.provDesc);

        $("#basicInfoBarangayHexId").removeAttr("data-parsley-required");
        $("#basicInfoBarangayHexId").val(null).trigger("change");
        $("#basicInfoBarangayHexId").attr("data-parsley-required", true);
        $("#basicInfoBarangayHexId").select2({
            placeholder: "SELECT BARANGAY",
            ajax: {
                url: "/api/citymuns/" + $("#basicInfoCityMunHexId").select2("val") + "/barangays",
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
        $("#basicInfoBarangayHexId").prop("disabled", false);
        $("#basicInfoBarangayHexId").focus();
    });

    $("#empTeminatedOverseasCountryHexId").select2({
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

    $("#empUnEmpStatHexId").select2({
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

    $("#empUnEmpStatHexId").on("change", function () {
        if ($(this).select2("val") == "594cb702472e11263c32c495") {
            $("#empTeminatedOverseasCountryHexId").attr("data-parsley-required", true);
            $("#empTeminatedOverseasCountryHexId").prop("disabled", false);
            $("#empTeminatedOverseasCountryHexId").focus();
        } else {
            $("#empTeminatedOverseasCountryHexId").removeAttr("data-parsley-required");
            $("#empTeminatedOverseasCountryHexId").val(null).trigger("change");
            $("#empTeminatedOverseasCountryHexId").prop("disabled", true);
        }
    });

    $("select").on("change", function () {
        var instance = $(this).parsley();
        if (instance.isValid()) {
            instance.reset();
        }
    });
});