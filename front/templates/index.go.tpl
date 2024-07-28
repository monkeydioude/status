{{ define "index.tmpl" }}
<html class="bg-slate-900 text-cyan-100">
<head>
	<script src="https://cdn.tailwindcss.com"></script>
	<link href="/status/resources/main.css" rel="stylesheet" />
	<script src="/status/resources/main.js"></script>
</head>
	<body>
		<div class="main m-auto w-4/5 flex flex-col">
			<div class="m-auto bg-slate-800 w-full text-center rounded">Status</div>
			{{ range .}}
			<div class="flex m-auto w-full mt-1" data-service-name="{{ .Name }}">
				<h2 class="w-1/3 bg-cyan-900 rounded-l px-2 text-cyan-200 flex items-center">{{ .Name }}</h2>
				<div class="w-full flex flex-col bg-cyan-800 rounded-r">
					<div class="flex w-full">
						<div data-status class="flex flex-col w-full">
							<div class="px-2 flex grow">Healthcheck</div>
							<div class="px-2">{{ .Message }}</div>
							{{ if ne .Status "" }}
							<div class="px-2">
								<div class="marker">Status</div>
								<div class="text"><pre class="formatted-pre w-full p-2 border-2 border-slate-800 mb-2 rounded">{{ .Status }}</pre></div>
							</div>
							{{ end }}
						</div>
						<div class="content-center justify-center px-2 {{ if eq .Health "OK" }}text-green-500{{ else }}text-red-500{{ end }} " data-healthcheck-status>{{ .Health }}</div>
					</div>
				</div>
			</div>
			{{ end }}
		</div>
	</body>
</html>
{{ end }}