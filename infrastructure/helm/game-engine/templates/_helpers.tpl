{{/*
Expand the name of the chart.
*/}}
{{- define "game-engine.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
*/}}
{{- define "game-engine.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "game-engine.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "game-engine.labels" -}}
helm.sh/chart: {{ include "game-engine.chart" . }}
{{ include "game-engine.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: casino-gaming-platform
{{- end }}

{{/*
Selector labels
*/}}
{{- define "game-engine.selectorLabels" -}}
app.kubernetes.io/name: {{ include "game-engine.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Service labels
*/}}
{{- define "game-engine.serviceLabels" -}}
helm.sh/chart: {{ include "game-engine.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: casino-gaming-platform
app: {{ .serviceName }}
version: v1
component: service
{{- end }}

{{/*
Service selector labels
*/}}
{{- define "game-engine.serviceSelectorLabels" -}}
app: {{ .serviceName }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "game-engine.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "game-engine.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Common annotations
*/}}
{{- define "game-engine.podAnnotations" -}}
prometheus.io/scrape: "true"
prometheus.io/port: "8080"
prometheus.io/path: "/metrics"
checksum/config: {{ include (print $.Template.BasePath "/configmap.yaml") . | sha256sum }}
{{- end }}

{{/*
Image name
*/}}
{{- define "game-engine.image" -}}
{{- printf "%s/%s:%s" .Values.global.image.registry .serviceName .Values.global.image.tag }}
{{- end }}
