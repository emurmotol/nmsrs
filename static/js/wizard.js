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

    $(".wizard ul li").on("click", function () {
        if ($(this).hasClass("disabled")) {
            return false;
        }
    });

    $("#preferred_occupation").select2({
        placeholder: "Select preferred occupation",
        ajax: {
            url: "/positions",
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

    $("#province").select2({
        placeholder: "Select a province",
        ajax: {
            url: "/provinces",
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
            url: "/city-municipalities/provinces",
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
    });

    $("#preferred_local_location").select2({
        placeholder: "Select preferred local location",
        ajax: {
            url: "/city-municipalities/provinces",
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
            $("#unemployed_current_status").prop("disabled", false);
            $("#unemployed_current_status").focus();
        } else {
            $("#unemployed_current_status").prop("disabled", true);
            $("#unemployed_country").prop("disabled", true);
        }
    });

    $("#civil_status_5").on("click", function () {
        $("#civil_status").prop("disabled", false);
        $("#civil_status").focus();
    });

    $("#civil_status_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=civil_status]:checked");

        if ($(checked).val() == "5") {
            $("input[name=civil_status_5]").prop("disabled", false);
            $("input[name=civil_status_5]").focus();
        } else {
            $("input[name=civil_status_5]").prop("disabled", true);
            $("input[name=civil_status_5]").val("");
        }
    });

    $("#disability_radios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=disability]:checked");

        if ($(checked).val() == "5") {
            $("input[name=disability_5]").prop("disabled", false);
            $("input[name=disability_5]").focus();
        } else {
            $("input[name=disability_5]").prop("disabled", true);
            $("input[name=disability_5]").val("");
        }
    });

    $("#disabled").on("change", function () {
        if ($(this).prop("checked")) {
            $("#disability_radios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", false);
            });
        } else {
            $("#disability_radios").find("input[type=radio]").each(function () {
                $("input[name=disability_5]").prop("disabled", true);
                $(this).prop("disabled", true);
                $(this).prop("checked", false);
            });
        }
    });

    $("#step_1_form").on("submit", function (e) {
        e.preventDefault();
        var submit_button = $(this).find(":submit");
        var old_text = submit_button.text();
        submit_button.prop("disabled", true);
        submit_button.html(`<i class="fa fa-spinner fa-pulse fa-spin"></i> Please wait...`);

        var validate_fields = [
            "family_name",
            "given_name",
            "middle_name",
            "birthdate",
            "password",
            "street_subdivision",
            "city_municipality",
            "province",
            "barangay",
            "place_of_birth",
            "religion",
            "sex",
            "age",
            "height",
            "weight",
            "landline",
            "mobile",
            "email",
        ];
        var civil_status = "";

        if ($("input[name=civil_status]:checked").val() == "5") {
            civil_status = $("input[name=civil_status_5]").val().toUpperCase();

            if ($.inArray("civil_status", validate_fields) == -1) {
                validate_fields.push("civil_status");
            }
        } else {
            civil_status = $("input[name=civil_status]:checked").data("name");

            if ($.inArray("civil_status", validate_fields) != -1) {
                validate_fields.pop("civil_status");
            }
        }

        if ($("input[name=civil_status]:checked").length == 0) {
            if ($.inArray("civil_status", validate_fields) == -1) {
                validate_fields.push("civil_status");
            }
        }

        console.log(validate_fields);

        var data = JSON.stringify({
            "personal_information": {
                "family_name": $("#family_name").val().toUpperCase(),
                "given_name": $("#given_name").val().toUpperCase(),
                "middle_name": $("#middle_name").val().toUpperCase(),
                "birthdate": $("#birthdate").val(),
                "password": $("#password").val().toUpperCase()
            },
            "basic_information": {
                "street_subdivision": $("#street_subdivision").val().toUpperCase(),
                "city_municipality": $("#city_municipality").text(),
                "province": $("#province").text(),
                "barangay": $("#barangay").text(),
                "place_of_birth": $("#place_of_birth").val().toUpperCase(),
                "religion": $("#religion").text(),
                "civil_status": {
                    "id": $("input[name=civil_status]:checked").val(),
                    "name": civil_status
                },
                "sex": $("input[name=sex]:checked").data("name"),
                "age": parseInt($("#age").val()),
                "height": parseFloat($("#height").val()),
                "weight": parseFloat($("#weight").val()),
                "landline": $("#landline").val(),
                "mobile": $("#mobile").val(),
                "email": $("#email").val(),
            }
        });

        $.ajax({
            url: $(this).attr("action"),
            type: $(this).attr("method"),
            data: data,
            dataType: "json",
            success: function (r) {
                $("#alert_container").empty();

                $.each(validate_fields, function (k, v) {
                    removeFormErrorMarkup(k);
                });
                errors = r.errors;

                try {
                    if (Object.keys(errors).length != 0) {
                        $.each(errors, function (k, v) {
                            addFormErrorMarkup(k, v);
                        });
                    }
                } catch (e) {
                    addAlertErrorMarkup(errors);
                }
            }, error: function (r) {
                console.log(r);
            }
        }).done(function (r) {
            if (r.status == 200) {
                if (r.data.message != null) {
                    var msg_markup = `<div class="alert alert-success alert-dismissible" role="alert">
                        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        <i class="fa fa-check"></i> `+ r.data.message + `
                    </div>`;
                    $("#alert_container").html(msg_markup);
                }

                if (r.data.proceed != null && r.data.proceed) {
                    nextStep();
                }
            }
            submit_button.prop("disabled", false);
            submit_button.html(old_text);
            console.log(r);
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