$(function () {
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

    $("#brgy_id").select2();
    $("#brgy_id").on("change", function () {
        $("#place_of_birth").focus();
    });

    $("#city_mun_id").on("change", function () {
        var city_mun_id = $(this).val();
        var data = $(this).select2("data")[0];

        $("#prov_id").prop("data-id", data.prov_id);
        $("#prov_id").val(data.prov_desc);

        $("#brgy_id").val("");
        $("#brgy_id").select2({
            placeholder: "SELECT BARANGAY",
            ajax: {
                url: "/api/citymuns/" + city_mun_id + "/barangays",
                delay: 250,
                dataType: "json",
                processResults: function (r) {
                    console.log(r);
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
        }).prop("disabled", false);
        $("#brgy_id").focus();
    });

    $("#un_emp_stat_id").select2();
    loadUnEmpStat = function () {
        $("#un_emp_stat_id").select2({
            placeholder: "SELECT UNEMPLOYED STATUS",
            ajax: {
                url: "/api/unempstats",
                delay: 250,
                dataType: "json",
                processResults: function (r) {
                    console.log(r);
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
            },
            minimumResultsForSearch: Infinity
        });
    }

    $("#toc_id").select2();
    $("#un_emp_stat_id").on("change", function () {
        if ($(this).val() == "5") {
            $("#toc_id").val("");
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
            }).prop("disabled", false);
            $("#toc_id").focus();
        } else {
            $("#toc_id").val("");
            $("#toc_id").select2();
            $("#toc_id").prop("disabled", true);
        }
    });
});