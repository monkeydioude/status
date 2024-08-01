{{ define "index.tmpl" }}
<html class="bg-slate-900 text-cyan-100">
<head>
	<link href="/status/resources/main.css" rel="stylesheet" />
	<script src="/status/resources/main.js"></script>
</head>
	<body>
		<div class="main m-auto w-4/5 flex flex-col">
			<div class="m-auto bg-slate-800 w-full rounded flex h-7 items-center">
				<div class="w-1/4"></div>
				<div class="w-1/2 text-center">Status</div>
				<div class="w-1/4 flex justify-end pr-1">
					<label data-auto-reload-container class="inline-flex items-center cursor-pointer my-1 hidden">
						<span class="text-sm font-medium text-gray-900 dark:text-gray-300 mr-3">Auto-Reload</span>
						<input type="checkbox" data-auto-reload value="" class="sr-only peer">
						<div
						class="relative w-9 h-5 bg-gray-900 peer-focus:outline-none rounded-full peer dark:bg-gray-900 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all dark:border-gray-600 peer-checked:bg-cyan-800">
						</div>
					</label>
				</div>
			</div>
			{{ range .}}
			<div class="flex m-auto w-full mt-1" data-service-name="{{ .Name }}">
				<h2 class="w-1/4 bg-cyan-900 rounded-l px-2 text-cyan-200 flex items-center">{{ .Name }}</h2>
				<div class="w-3/4 flex flex-col bg-cyan-800 rounded-r">
					<div class="flex w-100">
						<div data-status class="flex grow flex-col">
							<div class="px-2 flex grow">Healthcheck</div>
							<div class="px-2">{{ if ne .Message "" }}<pre class="formatted-pre p-2 border-2 border-slate-800 mb-2 rounded">{{ .Message }}</pre> {{ end }}</div>
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