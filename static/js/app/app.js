$(function () {
    previewImage = function (elem) {
        $(elem).on("change", function () {
            var preview = $(this).parent().find("#preview");
            var defaultPhoto = preview.data("default-photo");
            var maxMB = parseInt($(this).attr("data-parsley-maxmegabytes")) * 1000000;

            if (this.files && this.files[0]) {
                if (this.files[0].size > maxMB) {
                    preview.attr("src", defaultPhoto);
                } else {
                    var reader = new FileReader();

                    reader.onload = function (e) {
                        preview.attr("src", e.target.result);

                        preview.on("error", function () {
                            preview.attr("src", defaultPhoto);
                        });
                    }
                    reader.readAsDataURL(this.files[0]);
                }
            } else {
                preview.attr("src", defaultPhoto);
            }
        });
    }

    duringSubmitDo = function(instance) {
        var submitBtn = $(instance).find(":submit");
        submitBtn.prop("disabled", true);
        submitBtn.html(`<i class="fa fa-spinner fa-pulse fa-spin"></i> ` + submitBtn.data("loading-text"));
    }
});
