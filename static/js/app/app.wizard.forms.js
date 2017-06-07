$(function () {
    $("#create_registrant_form").parsley();
    $("#create_registrant_form").on("submit", function (e) {
        duringSubmitDo(this);
    });
});