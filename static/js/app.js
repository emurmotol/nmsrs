$(function () {
    quickRequest = function (action, method, data) {
        $.ajax({
            url: action,
            type: method,
            dataType: "json",
            data: data,
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

    $.fn.serializeObject = function () {
        var arr = this.serializeArray();
        var obj = {};

        for (var i = 0; i < arr.length; i++) {
            obj[arr[i].name] = arr[i].value;
        }
        return obj;
    }

    setCheckboxBoolValue = function (checkbox) {
        checkbox.on("change", function () {
            $(this).val($(this).is(":checked"));
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

    checkFileRequest = function (url, method, data, id) {
        var alert = $("#alert");

        $.ajax({
            url: url,
            type: method,
            dataType: "json",
            contentType: false,
            data: data,
            processData: false,
            success: function (r) {
                alert.empty();
                removeErrorMarkup($(id));
                errors = r.errors;

                try {
                    if (Object.keys(errors).length != 0) {
                        $.each(errors, function (k, v) {
                            var field = $("#" + k);
                            addErrorMarkup(field, v);
                        });
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

    makeRequest = function (action, method, fields, data, isMultipart) {
        var alert = $("#alert");
        var content_type = false

        if (!isMultipart) {
            content_type = "application/x-www-form-urlencoded; charset=UTF-8";
        }

        $.ajax({
            url: action,
            type: method,
            dataType: "json",
            contentType: content_type,
            data: data,
            processData: false,
            success: function (r) {
                alert.empty();

                $.each(fields, function (k, v) {
                    var field = $("#" + v);
                    removeErrorMarkup(field);
                });
                errors = r.errors;

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
