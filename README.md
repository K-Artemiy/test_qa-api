# test_qa-api  
REST API сервис для вопросов и ответов 

##### Используемые технологии:
1. ___Go 1.25.1___
2. ___PostgreSQL___
3. ___net/http___ для реализации HTTP API сервера
4. ___GORM___ для взаимодействия с БД
5. ___Goose___ для реализации миграций 
6. ___Docker + docker-compose___
7. ___httptest + testify___ для тестов



### Реализованные модели:
1. ___Question___ – вопрос: 
    - id: int 
    - text: str (текст вопроса) 
    - created_at: datetime 
2. ___Answer___ – ответ на вопрос: 
    - id: int 
    - question_id: int (ссылка на Question) 
    - user_id: str (идентификатор пользователя, например uuid) 
    - text: str (текст ответа) 
    - created_at: datetime


### Методы API:
1. ___Вопросы (Questions)___: 
- <span style='color: green;'>GET /questions/ </span>— список всех вопросов 
- <span style='color: green;'>POST /questions/ </span>— создать новый вопрос 

    *Тело запроса:*
```json: {"text": "..."}```
- <span style='color: green;'>GET /questions/{id} </span>— получить вопрос и все ответы на него 
- <span style='color: green;'>DELETE /questions/{id} </span>— удалить вопрос (вместе с ответами) 
2. ___Ответы (Answers)___: 
- <span style='color: green;'>POST /questions/{id}/answers/ </span>— добавить ответ к вопросу 

    *Тело запроса:*
```json: {"text": "...", "user_id": "..."}```
- <span style='color: green;'>GET /answers/{id}</span> — получить конкретный ответ 
- <span style='color: green;'>DELETE /answers/{id}</span> — удалить ответ 
---
### Запуск проекта

#### 1. Установка

Склонируйте репозиторий: 
```git clone https://github.com/K-Artemiy/test_qa-api```

И перейдите в каталог проекта:
```cd test_qa-api```

#### 2. Копирование переменных окружения

Скопируйте переменные окружения в этот каталог:
```cp .env.example .env```

#### 3. Запуск сервиса

Запустите сервис (для этого дополнительно необходимо запустить Docker Desktop):

```docker-compose up --build```

После запуска:
- API будет доступно по адресу: http://localhost:8080
- PostgreSQL будет работать в контейнере db

Миграции применяются автоматически при старте контейнера приложения.
Тесты также запускаются автоматически после старта контейнера приложения.

#### Просмотр логов тестирования

После запуска сервера и тестов, логи тестирования будут доступны по следующей команде:
```docker-compose logs tests```
