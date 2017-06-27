$(function () {
    makeRequest = function (action, method, data) {
        var alertContainer = $("#alertContainer"); // TODO: Deleted on markup

        var call = $.ajax({
            url: action,
            type: method,
            data: data,
            dataType: "json",
            success: function (r) {
                alertContainer.empty();
                errors = r.errors;

                if (errors.length != 0) {
                    var errMarkup = `<div class="alert alert-danger alert-dismissible" role="alert">
                        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        <i class="fa fa-exclamation-triangle"></i> `+ errors + `
                    </div>`;
                    alertContainer.html(errMarkup);
                }
            }, error: function (r) {
                console.log(r);
            }
        }).done(function (r) {
            if (r.status == 200) {
                if (r.data.redirect != null) {
                    location.href = r.data.redirect;
                }
            }
            console.log(r);
        });

        return call.then(function (r) {
            return r;
        });
    }

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
                    reader.readAsDataURL(photo.files[0]);
                }
            } else {
                preview.attr("src", defaultPhoto);
            }
        });
    }

    duringSubmitDo = function(instance) {
        var submitButton = $(instance).find(":submit");
        submitButton.prop("disabled", true);
        submitButton.html(`<i class="fa fa-spinner fa-pulse fa-spin"></i> ` + submitButton.data("loading-text"));
    }

    makeFormRequest = function (instance, data) {
        var formAlert = $(instance).find(".form-alert");
        var submitButton = $(instance).find(":submit");
        var oldText = submitButton.text();
        var contentType = null;
        var processData = true;
        var enctype = $(instance).prop("enctype");
        submitButton.prop("disabled", true);
        submitButton.html(`<i class="fa fa-spinner fa-pulse fa-spin"></i> ` + submitButton.data("loading-text"));
        var multipart = "multipart/form-data";

        if (enctype == multipart) {
            contentType = false;
            processData = false;
        } else {
            contentType = enctype;
        }

        var call = $.ajax({
            url: $(instance).attr("action"),
            type: $(instance).attr("method"),
            data: data,
            dataType: "json",
            contentType: contentType,
            processData: processData,
            success: function (r) {
                formAlert.empty();

                if (r.error != null) {
                    var errMarkup = `<div class="alert alert-danger alert-dismissible" role="alert">
                        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        <i class="fa fa-exclamation-triangle"></i> `+ r.error + `
                    </div>`;
                    formAlert.html(errMarkup);
                }
            }, error: function (r) {
                console.log(r);
            }
        }).done(function (r) {
            if (r.message != null) {
                var msgMarkup = `<div class="alert alert-success alert-dismissible" role="alert">
                    <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                    <i class="fa fa-check"></i> `+ r.message + `
                </div>`;
                formAlert.html(msgMarkup);
            }

            if (r.redirect != null) {
                location.href = r.redirect;
            }

            if (enctype == multipart) {
                $(instance).find(":file").each(function () {
                    $(this).val("");
                });
            }
            submitButton.prop("disabled", false);
            submitButton.html(oldText);
            console.log(r);
        });

        return call.then(function (r) {
            return r;
        });
    }
});
