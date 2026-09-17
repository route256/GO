
Чтобы работать с переременными окружения, создайте у себя файлик `.env` со следующим содержимым:

```text
APP_HOST=localhost
APP_PORT=9999

ENVIRONMENT=DEV

CRYPTO_RULE=123456

LOG_LEVEL=DEBUG

REDIS_HOST=localhost
REDIS_PORT=6379
```

_(Да, его можно было закоммитить, но коммитить такие файлы негоже, поэтому оставим так)_


Запуск проекта:
```bash
go run main.go
```

Отправление запроса на наш сервер:
```bash
curl "http://localhost:9999/mine?data=strtomine"
```