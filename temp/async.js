function UserAuthorityCheck() {
    var returnData = $.ajax({
        async: true,
        type: "POST",
        url: "http://localhost:80/myApp/MyWebService.asmx" + "/" + "MyUserInRole",
        contentType: "application/json; charset=utf-8",
        dataType: "json"
    });
    return returnData.then(function (response) {
        return response.d === 'success';
    }, function () {
        return false;
    });
}

$(document).ready(function () {
    var checkPromise = UserAuthorityCheck();
    checkPromise.then(isUserAuthority, function () { isUserAuthority(false); });

    function isUserAuthority(response) {
        $('#btnsearch').prop('disabled', !response);
    }
});