window.ParsleyConfig = {
    successClass: "has-success",
    errorClass: "has-error",
    classHandler: function () {
        return this.$element.closest(".form-group");
    },
    errorsWrapper: `<span class="help-block"></span>`,
    errorTemplate: "<span></span>"
};