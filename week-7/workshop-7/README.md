# Workshop 7

Запускаем сервис core:
```bash
cd core
go run core.go
```

Компилируем и запускаем бинарь майнера:
```bash
cd miner
go build -o miner
./miner -data=hello -rule=999
```

Поднимаем наши контейнеры:
(после того, как будет составлен `docker-compose.yml`)
```bash
docker-compose up --build
```
В `docker-compose.yml` должно быть лишь 3 сервиса (это если с графаной), и все образы из коробки.

