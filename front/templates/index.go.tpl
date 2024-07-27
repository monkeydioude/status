{{ define "index.tmpl" }}
<html class="bg-slate-900 text-cyan-100">
<head>
	<script src="https://cdn.tailwindcss.com"></script>
	<link href="/resources/main.css" rel="stylesheet" />
</head>
	<body>
		<div class="main m-auto w-4/5 flex flex-col">
			<div class="m-auto bg-slate-800 w-full text-center rounded">Status</div>
			{{ range .}}
			<div class="flex m-auto w-1/3 mt-1" data-service-name="{{ .Name }}">
				<h2 class="w-1/3 bg-cyan-900 rounded-l px-2 text-cyan-200">{{ .Name }}</h2>
				<div class="flex flex-col bg-cyan-800 rounded-r">
					<div class="flex flex-col">
						<div class="flex">
							<div class="px-2">Healthcheck</div>
							<div class="px-2 {{ if eq .Health "OK" }}text-green-500{{ else }}text-red-500{{ end }} " data-healthcheck-status>{{ .Health }}</div>
						</div>
						<div class="px-2" data-healthcheck-message>{{ .Message }}</div>
					</div>
				</div>
			</div>
			{{ end }}
		</div>
	</body>
</html>
{{ end }}