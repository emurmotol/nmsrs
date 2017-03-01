{{define "auth"}}
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{block "title" .}}{{end}}</title>
    {{block "style" .}}{{end}}
</head>
<body>
{{template "content" .}}
{{block "script" .}}{{end}}
</body>
</html>
{{end}}