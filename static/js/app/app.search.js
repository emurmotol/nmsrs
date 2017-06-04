$(function () {
    $("#q").val($.url("?q"));

    var searchablePaths = [
        "/users",
        "/registrants"
    ];

    switch (searchablePaths.indexOf($.url("path"))) {
        case 0:
            $("#q").closest("form").prop("action", $.url("path"));
            $("#q").prop("placeholder", "Search users");
            break;
        case 1:
            $("#q").closest("form").prop("action", $.url("path"));
            $("#q").prop("placeholder", "Search registrants");
            break;
        default:
            $("#q").prop("placeholder", "Search");
    }

    $("#q").typeahead({
        source: function (query, proccess) {
            $.getJSON("/api/search", { q: query }, function (r) {
                console.log(r);
            }).fail(function (r) {
                console.log(r.responseText);
            }).done(function (r) {
                return proccess(r);
            });
        },
        afterSelect: function (obj) {
            var form = $("#q").closest("form");
            var path = $.url("path");

            switch (obj.type) {
                case "User":
                    if (path != "/users") {
                        form.prop("action", "/users");
                    }
                    break;
                case "Registrant":
                    if (path != "/registrants") {
                        form.prop("action", "/registrants");
                    }
                    break;
            }
            form.submit();
        },
        items: 6,
        autoSelect: false,
        fitToElement: true
    });
});
