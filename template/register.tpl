<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Login</title>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/app.css" rel="stylesheet">
</head>
<body>
<div class="container">
    <div class="row align-items-center">
        <div class="col-lg-3 mx-lg-auto">
            <a href="/template/home.tpl">
                <img src="/static/img/logo/dole-logo.png" height="192" class="mx-auto d-block logo" alt="Dole Logo">
            </a>
            <form>
                <div class="form-group">
                    <input type="text" class="form-control" name="name" placeholder="Name">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" name="username" placeholder="Username">
                </div>
                <div class="form-group">
                    <input type="password" class="form-control" name="password" placeholder="Password">
                </div>
                <div class="form-group">
                    <input type="password" class="form-control" name="confirm_password" placeholder="Confirm Password">
                </div>
                <button type="submit" class="btn btn-outline-primary btn-block">Register</button>
            </form>
        </div>
    </div>
</div>
<script src="/static/js/jquery-3.1.1.slim.min.js"></script>
<script src="/static/js/tether.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
<script src="/static/js/app.js"></script>
</body>
</html>