$(function () {
    $("#step_1_form").on("submit", function (e) {
        duringSubmitDo(this);
        nextStep();
    });

    $("#step_2_form").on("submit", function (e) {
        nextStep();
    });

    $("#step_3_form").on("submit", function (e) {
        nextStep();
    });

    $("#step_4_form").on("submit", function (e) {
        nextStep();
    });

    $("#step_5_form").on("submit", function (e) {
        nextStep();
    });

    $("#step_6_form").on("submit", function (e) {
        nextStep();
    });

    $("#step_7_form").on("submit", function (e) {
        nextStep();
    });

    $("#step_8_form").on("submit", function (e) {
        nextStep();
    });

    $("#step_9_form").on("submit", function (e) {
        console.log("FINISH");
    });

    $(".prev-step").on("click", function (e) {
        prevStep();
    });
});