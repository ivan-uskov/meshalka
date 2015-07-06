{{ define "base" }}
<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>{{ .Title }}</title>
    <link rel="stylesheet" type="text/css" href="/static/css/styles.css" media="all" />
    {{ range .Styles }}
      <link rel="stylesheet" type="text/css" href="/static/css/{{.}}" media="all" />
    {{ end }}
  </head>
  <body>
    <div class="main_container">
      <div class="header white_element">
        <a href="/" class="go_home_link" title="home">Make your cocktail</a>
        {{ template "header_elements" .Content }}
      </div>
      <div class="content">
        {{ template "content" .Content }}
      </div>
    </div>
    {{ range .Scrypts }}
      <script type="text/javascript" src="/static/js/{{.}}"></script>
    {{ end }}
  </body>
</html>
{{ end }}