window.ParsleyConfig = {
    successClass: "has-success",
    errorClass: "has-error",
    classHandler: function (elem) {
        return elem.$element.closest(".form-group");
    },
    errorsContainer: function (elem) { },
    errorsWrapper: `<span class="help-block"></span>`,
    errorTemplate: '<span></span>'
};