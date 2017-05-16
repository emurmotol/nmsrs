$(function () {
    var alert_container = $("#alert_container");

    setCheckboxBoolValue = function (checkbox) {
        checkbox.on("change", function () {
            $(this).val($(this).is(":checked"));
        });
    }

    addAlertErrorMarkup = function (error) {
        var err_markup = `<div class="alert alert-danger alert-dismissible" role="alert">
            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
            <i class="fa fa-exclamation-triangle"></i> `+ error + `
        </div>`;
        alert_container.html(err_markup);
    }

    makeRequest = function (action, method, data) {
        var call = $.ajax({
            url: action,
            type: method,
            data: data,
            dataType: "json",
            success: function (r) {
                alert_container.empty();
                errors = r.errors;

                if (errors.length != 0) {
                    addAlertErrorMarkup(errors);
                }
            }, error: function (r) {
                console.log(r);
            }
        }).done(function (r) {
            if (r.status == 200) {
                if (r.data.redirect != null) {
                    window.location.href = r.data.redirect;
                }
            }
            console.log(r)
        });

        return call.then(function (r) {
            return r;
        });
    }

    previewImage = function (elem) {
        elem.on("change", function () {
            var preview = $(this).parent().find("#preview");
            var default_photo = preview.data("default-photo");
            var maxMB = parseInt($(this).attr("data-parsley-file-max-megabytes")) * 1000000;

            if (this.files[0].size > maxMB) {
                preview.attr("src", default_photo);
                return
            } // TODO: 2nd check file size

            if (this.files && this.files[0]) {
                var reader = new FileReader();

                reader.onload = function (e) {
                    preview.attr("src", e.target.result);

                    preview.on("error", function () {
                        preview.attr("src", default_photo);
                    });
                }
                reader.readAsDataURL(photo.files[0]);
            } else {
                preview.attr("src", default_photo);
            }
        });
    }

    makeFormRequest = function (form, data) {
        var submit_button = $(form).find(":submit");
        var old_text = submit_button.text();
        var content_type = null;
        var enctype = $(form).prop("enctype");
        submit_button.prop("disabled", true);
        submit_button.html(`<i class="fa fa-spinner fa-pulse fa-spin"></i> ` + submit_button.data("loading-text"));

        if (enctype == "multipart/form-data") {
            content_type = false;
        } else {
            content_type = enctype;
        }

        var call = $.ajax({
            url: $(form).attr("action"),
            type: $(form).attr("method"),
            data: data,
            dataType: "json",
            contentType: content_type,
            processData: false,
            success: function (r) {
                alert_container.empty();

                if (r.data.error != null) {
                    addAlertErrorMarkup(r.data.error);
                }
            }, error: function (r) {
                console.log(r);
            }
        }).done(function (r) {
            if (r.status == 200) {
                if (r.data.message != null) {
                    var msg_markup = `<div class="alert alert-success alert-dismissible" role="alert">
                        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        <i class="fa fa-check"></i> `+ r.data.message + `
                    </div>`;
                    alert_container.html(msg_markup);
                }

                if (r.data.redirect != null) {
                    location.href = r.data.redirect;
                }

                $(form).find(":file").each(function () {
                    $(this).val("");
                });
            }
            submit_button.prop("disabled", false);
            submit_button.html(old_text);
            console.log(r);
        });

        return call.then(function (r) {
            return r;
        });
    }
});
