window.Parsley.on("field:success", function() {
    var span = this.$element.closest(".help-block")

    if (span.length) {
        span.remove();
    }
});

// window.Parsley.on("field:error", function() {
//     var formGroup = this.$element.closest(".form-group")

//     if (this.$element.prop("id") == "pref_occ_ids") {
//         formGroup.addClass("has-error")
//     }
// });