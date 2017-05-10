$(function () {
    $.fn.Parsley.Defaults = {
        excluded: "input[type=button], input[type=submit], input[type=reset], input[type=hidden], :disabled",
        successClass: "has-success",
        errorClass: "has-error",
        classHandler: function (elem) {
            return elem.$element.closest(".form-group");
        },
        errorsContainer: function (elem) {},
        errorsWrapper: `<span class="help-block"></span>`,
        errorTemplate: '<span></span>'
    };

    var alert = $("#alert_container");

    setCheckboxBoolValue = function (checkbox) {
        checkbox.on("change", function () {
            $(this).val($(this).is(":checked"));
        });
    }

    removeFormErrorMarkup = function (k) {
        var elem = null;

        if ($("#" + k).length) {
            elem = $("#" + k);
        } else {
            elem = $(`input[name=` + k + `]`);
        }
        var fp = null;

        if (elem.prop("type") == "radio") {
            fp = elem.parent().parent().parent();
        } else {
            fp = elem.parent();
        }

        if (fp.hasClass("has-error")) {
            fp.removeClass("has-error");
        }

        if (fp.find("span.help-block").get().length == 1) {
            fp.find("span.help-block").remove();
        }
    }

    addFormErrorMarkup = function (k, message) {
        var elem = null;

        if ($("#" + k).length) {
            elem = $("#" + k);
        } else {
            elem = $(`input[name=` + k + `]`);
        }
        var fp = null;

        if (elem.prop("type") == "radio") {
            fp = elem.parent().parent().parent();
        } else {
            fp = elem.parent();
        }

        if (!fp.hasClass("has-error")) {
            fp.addClass("has-error");
        }

        if (fp.find("span.help-block").get().length == 0) {
            fp.append(`<span class="help-block">` + message + `</span >`);
        }
    }

    addAlertErrorMarkup = function (error) {
        var err_markup = `<div class="alert alert-danger alert-dismissible" role="alert">
            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
            <i class="fa fa-exclamation-triangle"></i> `+ error + `
        </div>`;
        alert.html(err_markup);
    }

    makeRequest = function (action, method, data) {
        var call = $.ajax({
            url: action,
            type: method,
            data: data,
            dataType: "json",
            success: function (r) {
                alert.empty();
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

    validateImage = function (elem, max_upload_size) {
        elem.on("change", function () {
            checkFileRequest("/check/file/image/" + $(this)[0].id, "POST", this)
            var preview = $(this).parent().find("#preview");
            var old_photo = preview.attr("src");

            if (this.files[0].size > max_upload_size) {
                preview.attr("src", old_photo);
                return
            } // TODO: First check file size and on server side

            if (this.files && this.files[0]) {
                var reader = new FileReader();

                reader.onload = function (e) {
                    preview.attr("src", e.target.result);

                    preview.on("error", function () {
                        preview.attr("src", old_photo);
                    });
                }
                reader.readAsDataURL(photo.files[0]);
            } else {
                preview.attr("src", old_photo);
            }
        });
    }

    checkFileRequest = function (url, method, input) {
        var data = new FormData();
        data.append(input.id, input.files[0]);

        $.ajax({
            url: url,
            type: method,
            data: data,
            dataType: "json",
            contentType: false,
            processData: false,
            success: function (r) {
                alert.empty();
                removeFormErrorMarkup(input.id);
                errors = r.errors;

                try {
                    if (Object.keys(errors).length != 0) {
                        $.each(errors, function (k, v) {
                            addFormErrorMarkup(k, v);
                        });
                    }
                } catch (e) {
                    addAlertErrorMarkup(errors);
                }
            }, error: function (r) {
                console.log(r);
            }
        }).done(function (r) {
            console.log(r);
        });
    }

    makeFormRequest = function (form, validate_fields) {
        var submit_button = $(form).find(":submit");
        var old_text = submit_button.text();
        var content_type = null;
        var data = null;
        var enctype = $(form).prop("enctype");
        submit_button.prop("disabled", true);
        submit_button.html(`<i class="fa fa-spinner fa-pulse fa-spin"></i> Please wait...`);

        if (enctype == "multipart/form-data") {
            content_type = false;
            data = new FormData(form);
        } else {
            content_type = enctype;
            data = $(form).serialize();
        }

        var call = $.ajax({
            url: $(form).attr("action"),
            type: $(form).attr("method"),
            data: data,
            dataType: "json",
            contentType: content_type,
            processData: false,
            success: function (r) {
                alert.empty();

                $.each(validate_fields, function (k, v) {
                    removeFormErrorMarkup(k);
                });
                errors = r.errors;

                try {
                    if (Object.keys(errors).length != 0) {
                        $.each(errors, function (k, v) {
                            addFormErrorMarkup(k, v);
                        });
                    }
                } catch (e) {
                    addAlertErrorMarkup(errors);
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
                    alert.html(msg_markup);
                }

                if (r.data.redirect != null) {
                    window.location.href = r.data.redirect;
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
