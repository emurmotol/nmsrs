$(function () {
    // $("#preferred_occupation").select2({
    //     placeholder: "Select preferred occupation",
    //     ajax: {
    //         url: "/api/positions",
    //         delay: 250,
    //         dataType: "json",
    //         data: function (params) {
    //             return {
    //                 q: params.term
    //             };
    //         },
    //         processResults: function (data) {
    //             return {
    //                 results: $.map(data.data.positions, function (lng) {
    //                     return {
    //                         id: lng.id,
    //                         text: lng.name
    //                     };
    //                 })
    //             };
    //         },
    //         cache: true
    //     }
    // });

    // $("#language_spoken").select2({
    //     placeholder: "Select language",
    //     ajax: {
    //         url: "/api/languages",
    //         delay: 250,
    //         dataType: "json",
    //         data: function (params) {
    //             return {
    //                 q: params.term
    //             };
    //         },
    //         processResults: function (data) {
    //             return {
    //                 results: $.map(data.data.languages, function (lng) {
    //                     return {
    //                         id: lng.id,
    //                         text: lng.name
    //                     };
    //                 })
    //             };
    //         },
    //         cache: true
    //     }
    // });

    $("#religion").select2({
        placeholder: "Select religion",
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

    // $("#province").select2({
    //     placeholder: "Select province",
    //     ajax: {
    //         url: "/api/provinces",
    //         delay: 250,
    //         dataType: "json",
    //         data: function (params) {
    //             return {
    //                 q: params.term
    //             };
    //         },
    //         processResults: function (data) {
    //             return {
    //                 results: $.map(data.data.provinces, function (prov) {
    //                     return {
    //                         id: prov.code,
    //                         text: prov.desc
    //                     };
    //                 })
    //             };
    //         },
    //         cache: true
    //     }
    // });

    $("#city_mun").select2({
        placeholder: "Select city/municipality",
        ajax: {
            url: "/api/citymuns/provinces",
            delay: 250,
            dataType: "json",
            processResults: function (r) {
                return {
                    results: $.map(r, function (data) {
                        return {
                            id: data.city_mun_code,
                            text: data.city_mun_desc + ", " + data.prov_desc,
                            prov_desc: data.prov_desc
                        };
                    })
                };
            },
            cache: true
        }
    });

    $("#city_mun").on("change", function () {
        var city_mun_code = $(this).val();
        var prov_desc = $(this).select2("data")[0].prov_desc;
        $("#province").val(prov_desc);

        $("#barangay").select2({
            placeholder: "Select barangay",
            ajax: {
                url: "/api/citymuns/" + city_mun_code + "/barangays",
                delay: 250,
                dataType: "json",
                processResults: function (r) {
                    return {
                        results: $.map(r, function (data) {
                            return {
                                id: data.code,
                                text: data.desc
                            };
                        })
                    };
                },
                cache: true
            }
        }).prop("disabled", false);
    });

    // $("#barangay").select2({
    //     placeholder: "Select barangay"
    // });

    // $("#preferred_local_location").select2({
    //     placeholder: "Select preferred local location",
    //     ajax: {
    //         url: "/api/city-municipalities/provinces",
    //         delay: 250,
    //         dataType: "json",
    //         processResults: function (data) {
    //             return {
    //                 results: $.map(data.data.city_municipalities, function (city_mun) {
    //                     return {
    //                         id: city_mun.code,
    //                         text: city_mun.desc + ", " + city_mun.province[0].desc
    //                     };
    //                 })
    //             };
    //         },
    //         cache: true
    //     }
    // });

    // $("#preferred_overseas_location").select2({
    //     placeholder: "Select preferred overseas location",
    //     ajax: {
    //         url: "/api/countries?except=PHILIPPINES",
    //         delay: 250,
    //         dataType: "json",
    //         data: function (params) {
    //             return {
    //                 q: params.term
    //             };
    //         },
    //         processResults: function (data) {
    //             return {
    //                 results: $.map(data.data.countries, function (coun) {
    //                     return {
    //                         id: coun.id,
    //                         text: coun.name
    //                     };
    //                 })
    //             };
    //         },
    //         cache: true
    //     }
    // });

    // $("#unemployed_country").select2({
    //     placeholder: "Select country"
    // });

    // $("#unemployed_current_status").select2({
    //     placeholder: "Select current status"
    // });

    // $("#unemployed_current_status").on("change", function () {
    //     if ($(this).val() == "5") {
    //         $("#unemployed_country").val("");
    //         $("#unemployed_country").text("");

    //         $("#unemployed_country").select2({
    //             placeholder: "Select country",
    //             ajax: {
    //                 url: "/api/countries",
    //                 delay: 250,
    //                 dataType: "json",
    //                 data: function (params) {
    //                     return {
    //                         q: params.term
    //                     };
    //                 },
    //                 processResults: function (data) {
    //                     return {
    //                         results: $.map(data.data.countries, function (coun) {
    //                             return {
    //                                 id: coun.id,
    //                                 text: coun.name
    //                             };
    //                         })
    //                     };
    //                 },
    //                 cache: true
    //             }
    //         });
    //         $("#unemployed_country").prop("disabled", false);
    //         $("#unemployed_country").focus();
    //     } else {
    //         $("#unemployed_country").prop("disabled", true);
    //     }
    // });

    // $("#employment_status_radios").find("input[type=radio]").on("change", function () {
    //     $("#unemployed_current_status").val("");
    //     $("#unemployed_current_status").text("");
    //     $("#unemployed_country").val("");
    //     $("#unemployed_country").text("");

    //     var checked = $("input[name=employment_status]:checked");

    //     if ($(checked).val() == "3") {
    //         $("#unemployed_current_status").select2({
    //             placeholder: "Select current status",
    //             minimumResultsForSearch: Infinity,
    //             ajax: {
    //                 url: "/api/unemployed-statuses",
    //                 delay: 250,
    //                 dataType: "json",
    //                 processResults: function (data) {
    //                     return {
    //                         results: $.map(data.data.unemployed_statuses, function (un_emp_stat) {
    //                             return {
    //                                 id: un_emp_stat.id,
    //                                 text: un_emp_stat.name
    //                             };
    //                         })
    //                     };
    //                 },
    //                 cache: true
    //             }
    //         });
    //         $("#unemployed_current_status").prop("disabled", false);
    //         $("#unemployed_current_status").focus();
    //     } else {
    //         $("#unemployed_current_status").prop("disabled", true);
    //         $("#unemployed_country").prop("disabled", true);
    //     }
    // });
});