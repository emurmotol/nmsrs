$(function () {
    $.fn.select2.defaults.set("theme", "bootstrap");
    var frag = $.url("#");
    // TODO: Navigate on page

    nextStep = function () {
        var li = $(".wizard .nav-pills li.active");
        li.next().removeClass("disabled");
        li.next().find('a[data-toggle="tab"]').tab("show");
    }

    prevStep = function () {
        $(".wizard .nav-pills li.active").prev().find('a[data-toggle="tab"]').tab("show");
    }

    $("#preferred_occupation").select2({
        placeholder: "Select preferred occupation",
        ajax: {
            url: "/languages", // TODO: Change to job positions
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

    $("#language_spoken").select2({
        placeholder: "Select language(s) spoken",
        ajax: {
            url: "/languages",
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
            url: "/religions",
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

    $("#city_municipality").select2({
        placeholder: "Select a city/municipality",
        ajax: {
            url: "/city-municipalities/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (data) {
                return {
                    results: $.map(data.data.city_municipalities_with_provinces, function (city_mun_with_prov) {
                        return {
                            id: city_mun_with_prov.code,
                            text: city_mun_with_prov.desc + ", " + city_mun_with_prov.province[0].desc
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
                url: "/city-municipalities/" + city_municipality_code + "/barangays",
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
        // $("#province").val(selected.data("province-desc"));
        // $("#province").attr("data-province-code", selected.data("province-code"));
    });

    $("#preferred_local_location").select2({
        placeholder: "Select preferred local location",
        ajax: {
            url: "/city-municipalities/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (data) {
                return {
                    results: $.map(data.data.city_municipalities_with_provinces, function (city_mun_with_prov) {
                        return {
                            id: city_mun_with_prov.code,
                            text: city_mun_with_prov.desc + ", " + city_mun_with_prov.province[0].desc
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
            url: "/countries?except=PHILIPPINES",
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

    $("#unemployed_current_status").select2({
        placeholder: "Select current status"
    });

    $("#unemployed_country").select2({
        placeholder: "Select a country"
    });

    $("#employment_status_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=employment_status]:checked");

        if (checked[0].id == "employment_status_unemployed") {
            $("#unemployed_current_status").select2({
                placeholder: "Select current status",
                minimumResultsForSearch: Infinity,
                ajax: {
                    url: "/unemployed-statuses",
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

            $("#unemployed_country").select2({
                placeholder: "Select a country",
                ajax: {
                    url: "/countries",
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
            $("#unemployed_current_status").prop("disabled", false);
            $("#unemployed_current_status").focus();
            $("#unemployed_country").prop("disabled", false);
        } else {
            $("#unemployed_current_status").prop("disabled", true);
            $("#unemployed_country").prop("disabled", true);
        }
    });

    $("#civil_status_other").on("click", function () {
        $("#civil_status").prop("disabled", false);
        $("#civil_status").focus();
    });

    $("#civil_status_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=civil_status]:checked");

        if (checked[0].id == "civil_status_other") {
            $("input[name=civil_status_other]").prop("disabled", false);
            $("input[name=civil_status_other]").focus();
        } else {
            $("input[name=civil_status_other]").prop("disabled", true);
        }
    });

    $("#disability_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=disability]:checked");

        if (checked[0].id == "disability_other") {
            $("input[name=disability_other]").prop("disabled", false);
            $("input[name=disability_other]").focus();
        } else {
            $("input[name=disability_other]").prop("disabled", true);
        }
    });

    $("#disabled").on("change", function () {
        if ($(this).prop("checked")) {
            $("#disability_radios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", false);
            });
        } else {
            $("#disability_radios").find("input[type=radio]").each(function () {
                $("input[name=disability_other]").prop("disabled", true);
                $(this).prop("disabled", true);
                $(this).prop("checked", false);
            });
        }
    });

    $("#step_1_form").on("submit", function (e) {
        e.preventDefault();
        var promise = makeFormRequest(this, []);

        promise.then(function (r) {
            if (r.data.proceed != null && r.data.proceed) {
                nextStep();
            }
        });
    });

    $("#step_2_form").on("submit", function (e) {
        e.preventDefault();
        nextStep();
    });

    $("#step_3_form").on("submit", function (e) {
        e.preventDefault();
        nextStep();
    });

    $("#step_4_form").on("submit", function (e) {
        e.preventDefault();
        nextStep();
    });

    $("#step_5_form").on("submit", function (e) {
        e.preventDefault();
        nextStep();
    });

    $("#step_6_form").on("submit", function (e) {
        e.preventDefault();
        nextStep();
    });

    $("#step_7_form").on("submit", function (e) {
        e.preventDefault();
        nextStep();
    });

    $("#step_8_form").on("submit", function (e) {
        e.preventDefault();
        nextStep();
    });

    $("#step_9_form").on("submit", function (e) {
        e.preventDefault();
    });

    $(".prev-step").on("click", function (e) {
        e.preventDefault();
        prevStep();
    });
});