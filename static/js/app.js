$(function () {
    validateForm = function(url, f, d) {
        var err = $("#error");

        $.post(url, d, function (r) {
            err.empty();
            console.log(r)

            $.each(f, function (i, v) {
                var field = $("#" + v);
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
            });

            try {
                e = r.errors;

                if (e.length != 0) {
                    $.each(e, function (i, v) {
                        var field = $("#" + i);
                        var fc = field.parent();

                        if (!fc.hasClass("has-danger")) {
                            fc.addClass("has-danger");
                        }

                        if (!field.hasClass("form-control-danger")) {
                            field.addClass("form-control-danger");
                        }

                        if (fc.find("div.form-control-feedback").get().length == 0) {
                            fc.append(`<div class="form-control-feedback">` + v + `</div>`);
                        }
                    });
                } else {
                    // TODO: Success logic
                    var markup = `
                        <div class="alert alert-success alert-dismissible fade show" role="alert">
                            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                                <span aria-hidden="true">&times;</span>
                            </button>
                            `+r.data.message+`
                        </div>`;
                    err.html(markup);
                    // TODO: Process jwt token
                    console.log(r.data.token);
                }
            } catch (e) {
                var markup = `
                    <div class="alert alert-danger alert-dismissible fade show" role="alert">
                        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        `+ r.errors + `
                    </div>`;
                err.html(markup);
            }
        }, "json");
    }
});
