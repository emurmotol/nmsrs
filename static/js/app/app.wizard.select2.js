$(function () {
    $("#school_univ_id").select2({
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

    $("#course_degree_id").select2({
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

    $("#high_grade_comp_id").select2({
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

    $("#pref_occ_ids").select2({
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

    $("#pref_local_loc_id").select2({
        placeholder: "SELECT PREFERRED LOCAL LOCATION",
        ajax: {
            url: "/api/citymuns/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.city_mun_id,
                            text: data.city_mun_desc + ", " + data.prov_desc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#pref_overseas_loc_id").select2({
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

    $("#language_ids").select2({
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

    $("#religion_id").select2({
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

    $("#brgy_id").select2({ placeholder: "SELECT BARANGAY" });
    $("#brgy_id").on("change", function () {
        $("#place_of_birth").focus();
    });

    $("#city_mun_id").select2({
        placeholder: "SELECT CITY/MUNICIPALITY",
        ajax: {
            url: "/api/citymuns/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.city_mun_id,
                            text: data.city_mun_desc + ", " + data.prov_desc,
                            prov_code: data.prov_id,
                            prov_desc: data.prov_desc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#city_mun_id").on("change", function () {
        var data = $(this).select2("data")[0];

        $("#prov_id").prop("data-id", data.prov_id);
        $("#prov_id").val(data.prov_desc);

        $("#brgy_id").removeAttr("data-parsley-required");
        $("#brgy_id").val(null).trigger("change");
        $("#brgy_id").attr("data-parsley-required", true);
        $("#brgy_id").select2({
            placeholder: "SELECT BARANGAY",
            ajax: {
                url: "/api/citymuns/" + $("#city_mun_id").select2("val") + "/barangays",
                delay: 250,
                dataType: "json",
                processResults: function (r) {
                    return {
                        results: $.map(r, function (data) {
                            return {
                                id: data.id,
                                text: data.desc
                            };
                        })
                    };
                },
                cache: true
            }
        });
        $("#brgy_id").prop("disabled", false);
        $("#brgy_id").focus();
    });

    $("#toc_id").select2({
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

    $("#un_emp_stat_id").select2({
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

    $("#un_emp_stat_id").on("change", function () {
        if ($(this).val() == "5") {
            $("#toc_id").attr("data-parsley-required", true);
            $("#toc_id").prop("disabled", false);
            $("#toc_id").focus();
        } else {
            $("#toc_id").removeAttr("data-parsley-required");
            $("#toc_id").val(null).trigger("change");
            $("#toc_id").prop("disabled", true);
        }
    });

    $("select").on("change", function() {
        var instance = $(this).parsley();
        if (instance.isValid()) {
            instance.reset();
        }
    }); // todo: temporary fix for select2 on change not working properly
});