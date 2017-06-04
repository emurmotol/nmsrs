$(function () {
    $("#step_1_form").on("submit", function (e) {
        e.preventDefault();
        var civil_status = "";

        if ($("input[name=civil_status]:checked").val() == "5") {
            civil_status = $("input[name=civil_status_5]").val().toUpperCase();
        } else {
            civil_status = $("input[name=civil_status]:checked").data("name");
        }

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
        var promise = makeFormRequest(this, data);

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