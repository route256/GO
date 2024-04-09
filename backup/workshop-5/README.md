## Logs

`make logs`

Graylog: http://127.0.0.1:7555/ (admin/admin)

System->Inputs, добавляем инпут типа GELF tcp, все значения по-умолчанию

## Metrics

`make metrics`

Prometheus: http://127.0.0.1:9090/

Grafana: http://127.0.0.1:3000/ (admin/admin)

При первом логине в Графану она попросит установить новый пароль, ставим.

Заходим в шестеренку слева, выбираем Data sources, добавляем Prometheus, адрес `http://prometheus:9090`

Нагрузку даем через hey: `hey -c 5 -z 10m "http://127.0.0.1:8080/fibonacci?n=5"`

## Tracing

`make tracing`

Jaeger: http://127.0.0.1:16686/
