# Homework 7



Добавить сервис по построению произвольных отчетов вашего бота.

Сервис должен быть реализован как отдельный микросервис, который запускается вместе с остальными в вашем компоуз файле. Пример компоуз файла для кафки, вы найдете по ссылке к материалам воркшопа.

Взаимодействие с сервисом построения отчетов осуществляется следующим образом:
1. пользователь выбирает в телеграм боте построить n отчет
2. сервис телеграм бота отправляет (продюсит) запрос на построение отчета в кафку
3. сервис построения отчетов потребляет (консюмит) сообщение из кафки, формирует нужный отчет
4. вызывает сервис телеграм бота по gRPC с результатами отчета
5. результаты возвращаются пользователю в телеграм

Задания на бонусы 💎
- добавить в интерцепторы gRPC метрики на вызываемые методы
- добавить метрики в продюсере и консюмере
- создать на стороне сервера [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) с возможностью вызова по gRPC и с помощью RESTful API
- добавить [валидацию](https://github.com/bufbuild/protoc-gen-validate) запросов путем добавления плагина в proto файле