{{ define "base" }}
  {{ range .Styles }}
    <link rel="stylesheet" type="text/css" href="/static/css/forms/{{.}}" media="all" />
  {{ end }}
  {{ range .Scrypts }}
    <script type="text/javascript" src="/static/js/forms/{{.}}"></script>
  {{ end }}
  {{ template "form" .Content }}
{{ end }}