window.ParsleyConfig = {
    inputs: "input, textarea, select, input[type=hidden], :hidden",
    excluded: "input[type=button], input[type=submit], input[type=reset]",
    // successClass: "has-success",
    successClass: "",
    errorClass: "has-error",
    classHandler: function (elem) {
        return elem.$element.closest(".form-group");
    },
    errorsWrapper: `<span class="help-block"></span>`,
    errorTemplate: "<span></span>"
};