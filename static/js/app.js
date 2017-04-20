$(function () {
    ajaxCall = function (action, method, data) {
        $.ajax({
            url: action,
            type: method,
            data: data,
            dataType: "json",
            success: function (r) {
                if (r.data.redirect != null) {
                    window.location.href = r.data.redirect
                }
                console.log(r)
            }, error: function (r) {
                console.log(r)
            }
        });
    }

    validateImage = function (photo, preview, default_photo) {
        photo.change(function () {
            if (this.files && this.files[0]) {
                var reader = new FileReader();

                reader.onload = function (e) {
                    removeErrorMarkup(photo);
                    preview.attr("src", e.target.result);

                    preview.on("error", function () {
                        preview.attr("src", default_photo);
                        addErrorMarkup(photo, "The selected file is not a valid image");
                    });
                }
                reader.readAsDataURL(this.files[0]);
            }
        });
    }

    removeErrorMarkup = function (field) {
        var fc = field.parent();

        if (fc.hasClass("has-danger")) {
            fc.removeClass("has-danger");
        }

        if (field.hasClass("form-control-danger")) {
            field.removeClass("form-control-danger");
        }

        if (fc.find("div.form-control-feedback").get().length == 1) {
            fc.find("div.form-control-feedback").remove();
        }
    }

    addErrorMarkup = function (field, message) {
        var fc = field.parent();

        if (!fc.hasClass("has-danger")) {
            fc.addClass("has-danger");
        }

        if (!field.hasClass("form-control-danger")) {
            field.addClass("form-control-danger");
        }

        if (fc.find("div.form-control-feedback").get().length == 0) {
            fc.append(`<div class="form-control-feedback">` + message + `</div>`);
        }
    }

    validateForm = function (action, method, fields, data) {
        var alert = $("#alert");

        $.ajax({
            url: action,
            type: method,
            dataType: "json",
            data: data,
            success: function (r) {
                alert.empty();

                $.each(fields, function (k, v) {
                    var field = $("#" + v);
                    removeErrorMarkup(field);
                });
                errors = r.errors;

                if ($.inArray("photo", fields) == 0) {
                    var photo = $("#photo")[0];
                    var preview = $("#preview");

                    if (photo.files && photo.files[0]) {
                        var reader = new FileReader();

                        reader.onload = function (e) {
                            preview.attr("src", e.target.result);

                            preview.on("error", function () {
                                preview.attr("src", "/img/user/default.png"); // TODO: src must be dynamic
                                errors["photo"] = "The selected file is not a valid image";
                            });
                        }
                        reader.readAsDataURL(photo.files[0]);
                    }
                } // TODO: Must have separate function

                try {
                    if (Object.keys(errors).length != 0) {
                        $.each(errors, function (k, v) {
                            var field = $("#" + k);
                            addErrorMarkup(field, v);
                        });
                    } else {
                        if (r.data.message != null) {
                            var message_markup = `
                            <div class="alert alert-success alert-dismissible fade show" role="alert">
                                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <i class="fa fa-check"></i> `+ r.data.message + `
                            </div>`;
                            alert.html(message_markup);
                        }

                        if (r.data.redirect != null) {
                            window.location.href = r.data.redirect
                        }
                    }
                } catch (e) {
                    var err_markup = `
                    <div class="alert alert-danger alert-dismissible fade show" role="alert">
                        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        <i class="fa fa-exclamation-triangle"></i> `+ errors + `
                    </div>`;
                    alert.html(err_markup);
                }
                console.log(r);
            }, error: function (r) {
                console.log(r);
            }
        });
    }
});
