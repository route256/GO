# Homework 4



**Предыстория**

На воркшопе мы разбирали на примере проектирования системы уведомлений. На практике потренировались в архитектуре маштабируемых и отказоустойчивых систем. Теперь нам нужно представить себя архитекторами которые должны взять приложение по учету расходов, представить как будет выглядеть его архитектура нашего телеграм бота при его маштабировании.

**Задание**

Нарисовать три схемы на (1000, 100 000, 1 000 000 пользователей). Сделать это можно в любом подходящем для этого редакторе (app.diagrams.net, miro.com, etc...).

При проектировании на схеме необходимо создать табличку где будем указывать следующие вводные, разберем на примере сервиса почтовых рассылок.

**Функциональные требования (зачем нужен сервис, какую проблему он решает):**
- отправляет почтовые рассылки
- позволяет тестировать а/б
- предоставляет статистику

**Не функциональные требования:**
- высокая скорость работы
- высокая отказоустойчивость

**Дополнительные требования:**
- возможность push нотификаций
- отслеживания доставки в реальном времени на grafana

**Нагрузка:**
- n RPS (средняя, максимальная)

**Оценка хранилища:**
- за год мы отправим n сообщений
- в год мы ожидаем прирост на n петабайт для хранения данных
- через n времени необходимо будет реплицировать базу, перевести запросы на чтение с асинхронной реплики для аналитики
- бекапы (x3 к размеру данных)

**Оценка размера оперативной памяти:**
- Базе данных требуется n ГБ на инстанс
- Кеш хранит 20% запросов за 24 часа, в среднем сообщение занимает n byte, необходимо n гигабайт памяти.

Вам необходимо сделать оценку по данному шаблону, с вашими цифрами и формулировками по сервису учета расходов. Табличку можно сделать общую на три схемы.

Ссылочка на графики / расчеты которые получились по итогу воркшопа https://miro.com/app/board/uXjVPN53G_A=/