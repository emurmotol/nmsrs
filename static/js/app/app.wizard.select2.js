$(function () {
    $("#preferred_occupation").select2({
        placeholder: "Select preferred occupation",
        ajax: {
            url: "/api/positions",
            delay: 250,
            dataType: "json",
            data: function (params) {
                return {
                    q: params.term
                };
            },
            processResults: function (data) {
                return {
                    results: $.map(data.data.positions, function (lng) {
                        return {
                            id: lng.id,
                            text: lng.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#language_spoken").select2({
        placeholder: "Select language",
        ajax: {
            url: "/api/languages",
            delay: 250,
            dataType: "json",
            data: function (params) {
                return {
                    q: params.term
                };
            },
            processResults: function (data) {
                return {
                    results: $.map(data.data.languages, function (lng) {
                        return {
                            id: lng.id,
                            text: lng.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#religion").select2({
        placeholder: "Select a religion",
        ajax: {
            url: "/api/religions",
            delay: 250,
            dataType: "json",
            data: function (params) {
                return {
                    q: params.term
                };
            },
            processResults: function (data) {
                return {
                    results: $.map(data.data.religions, function (relig) {
                        return {
                            id: relig.id,
                            text: relig.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#province").select2({
        placeholder: "Select a province",
        ajax: {
            url: "/api/provinces",
            delay: 250,
            dataType: "json",
            data: function (params) {
                return {
                    q: params.term
                };
            },
            processResults: function (data) {
                return {
                    results: $.map(data.data.provinces, function (prov) {
                        return {
                            id: prov.code,
                            text: prov.desc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#city_municipality").select2({
        placeholder: "Select a city/municipality",
        ajax: {
            url: "/api/city-municipalities/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (data) {
                return {
                    results: $.map(data.data.city_municipalities, function (city_mun) {
                        return {
                            id: city_mun.code,
                            text: city_mun.desc + ", " + city_mun.province[0].desc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#barangay").select2({
        placeholder: "Select a barangay"
    });

    $("#city_municipality").on("change", function () {
        var selected = $(this).find(":selected");
        var city_municipality_code = $(this).val();
        $("#barangay").val("");

        $("#barangay").select2({
            placeholder: "Select a barangay",
            ajax: {
                url: "/api/city-municipalities/" + city_municipality_code + "/barangays",
                delay: 250,
                dataType: "json",
                data: function (params) {
                    return {
                        q: params.term
                    };
                },
                processResults: function (data) {
                    return {
                        results: $.map(data.data.barangays, function (brgy) {
                            return {
                                id: brgy.code,
                                text: brgy.desc
                            };
                        })
                    };
                },
                cache: true
            }
        });
        $("#barangay").prop("disabled", false);
    });

    $("#preferred_local_location").select2({
        placeholder: "Select preferred local location",
        ajax: {
            url: "/api/city-municipalities/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (data) {
                return {
                    results: $.map(data.data.city_municipalities, function (city_mun) {
                        return {
                            id: city_mun.code,
                            text: city_mun.desc + ", " + city_mun.province[0].desc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#preferred_overseas_location").select2({
        placeholder: "Select preferred overseas location",
        ajax: {
            url: "/api/countries?except=PHILIPPINES",
            delay: 250,
            dataType: "json",
            data: function (params) {
                return {
                    q: params.term
                };
            },
            processResults: function (data) {
                return {
                    results: $.map(data.data.countries, function (coun) {
                        return {
                            id: coun.id,
                            text: coun.name
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#unemployed_country").select2({
        placeholder: "Select a country"
    });

    $("#unemployed_current_status").select2({
        placeholder: "Select current status"
    });

    $("#unemployed_current_status").on("change", function () {
        if ($(this).val() == "5") {
            $("#unemployed_country").val("");
            $("#unemployed_country").text("");

            $("#unemployed_country").select2({
                placeholder: "Select a country",
                ajax: {
                    url: "/api/countries",
                    delay: 250,
                    dataType: "json",
                    data: function (params) {
                        return {
                            q: params.term
                        };
                    },
                    processResults: function (data) {
                        return {
                            results: $.map(data.data.countries, function (coun) {
                                return {
                                    id: coun.id,
                                    text: coun.name
                                };
                            })
                        };
                    },
                    cache: true
                }
            });
            $("#unemployed_country").prop("disabled", false);
            $("#unemployed_country").focus();
        } else {
            $("#unemployed_country").prop("disabled", true);
        }
    });

    $("#employment_status_radios").find("input[type=radio]").on("change", function () {
        $("#unemployed_current_status").val("");
        $("#unemployed_current_status").text("");
        $("#unemployed_country").val("");
        $("#unemployed_country").text("");

        var checked = $("input[name=employment_status]:checked");

        if ($(checked).val() == "3") {
            $("#unemployed_current_status").select2({
                placeholder: "Select current status",
                minimumResultsForSearch: Infinity,
                ajax: {
                    url: "/api/unemployed-statuses",
                    delay: 250,
                    dataType: "json",
                    processResults: function (data) {
                        return {
                            results: $.map(data.data.unemployed_statuses, function (un_emp_stat) {
                                return {
                                    id: un_emp_stat.id,
                                    text: un_emp_stat.name
                                };
                            })
                        };
                    },
                    cache: true
                }
            });
            $("#unemployed_current_status").prop("disabled", false);
            $("#unemployed_current_status").focus();
        } else {
            $("#unemployed_current_status").prop("disabled", true);
            $("#unemployed_country").prop("disabled", true);
        }
    });
});