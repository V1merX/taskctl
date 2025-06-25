# Task Manager HTTP API

## Описание
Сервис для управления долгоиграющими I/O-bound задачами с HTTP API. Каждая задача выполняется от 10 до 30 секунд (случайное время). Все данные хранятся в памяти.

## Требования
- Go 1.24.2+
- Для тестирования: curl/httpie/Postman

## Установка и запуск
1. Клонировать репозиторий:
git clone https://github.com/V1merX/taskctl.git
cd taskctl

2. Запустить сервер:
go run cmd/main.go

Сервер будет доступен по адресу: http://localhost:8080

## Конфигурация
Файл конфигурации config.json должен находиться в папке configs:
{
    "http-server": {
        "host": "localhost", 
        "port": 8080
    }
}

## API Endpoints

### Создание задачи
POST /tasks

Пример запроса:
curl -X POST http://localhost:8080/tasks

Пример ответа:
{
    "task_id": "task_abc123"
}

### Получение статуса задачи
GET /tasks/{task_id}

Пример запроса:
curl http://localhost:8080/tasks/task_abc123

Пример ответа:
{
    "data": {
        "id": "task_abc123",
        "status": "in_progress", 
        "result": null,
        "duration": 5.231,
        "created_at": "2025-06-25T15:04:05Z"
    }
}

Поля ответа:
- status:
  - pending - задача создана
  - in_progress - выполняется
  - completed - завершена успешно
- duration: время выполнения в секундах (обновляется в реальном времени)
- created_at: время создания задачи

### Удаление задачи
DELETE /tasks/{task_id}

Пример запроса: 
curl -X DELETE http://localhost:8080/tasks/task_abc123

Коды ответа:
- 200 OK - задача удалена
- 404 Not Found - задача не найдена

## Особенности работы
1. Время выполнения каждой задачи случайно и составляет от 10 до 30 секунд
2. Поле duration показывает сколько секунд прошло с момента создания задачи
3. После завершения в поле result появляется значение "created"
4. Все данные хранятся в памяти и пропадают при перезапуске сервера

## Пример рабочего цикла
1. Создаем задачу:
POST /tasks → {"task_id": "task_123"}

2. Проверяем статус (во время выполнения):
GET /tasks/task_123 → 
{
    "data": {
        "status": "in_progress",
        "duration": 7.891
    }
}

3. Проверяем статус (после завершения):
GET /tasks/task_123 → 
{
    "data": {
        "status": "completed",
        "result": "created", 
        "duration": 15.342
    }
}