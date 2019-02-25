<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	{{if .IsDisplay}}
    	<em>{{.Content}}</em>
    {{else}}
    	<em>{{.Content2}}</em>
    {{end}}


    {{range .Users}}
    	{{.Username}} {{$.len}}<br>
    {{end}}
</body>
</html>