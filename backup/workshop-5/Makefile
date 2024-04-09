
dev:
	go run workshop5/cmd/fibonacci -devel

prod:
	mkdir -p logs/data
	go run workshop5/cmd/fibonacci 2>&1 | tee logs/data/log.txt

.PHONY: logs
logs:
	mkdir -p logs/data
	touch logs/data/log.txt
	touch logs/data/offsets.yaml
	sudo chmod -R 777 logs/data
	cd logs && sudo docker compose up

.PHONY: tracing
tracing:
	cd tracing && sudo docker compose up

.PHONY: metrics
metrics:
	mkdir -p metrics/data
	sudo chmod -R 777 metrics/data
	cd metrics && sudo docker compose up

pull:
	sudo docker pull prom/prometheus
	sudo docker pull grafana/grafana-oss
	sudo docker pull ozonru/file.d:latest-linux-amd64
	sudo docker pull elasticsearch:7.17.6
	sudo docker pull graylog/graylog:4.3
	sudo docker pull jaegertracing/all-in-one:1.18
