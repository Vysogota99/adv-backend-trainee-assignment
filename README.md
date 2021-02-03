# adv-backend-trainee-assignment

<h2>Инструкция по запуску</h2>
<h3>Структура проекта</h3>
<a href="https://github.com/golang-standards/project-layout">https://github.com/golang-standards/project-layout</a>
<br>
<ul>
    <li>build - содержит все необходимое для запуска и работы сервер: .env, Dockerfile для golang, миграции и файлы базы данных. После сборки сервера здесь появится скомпилированный файл для запуска;</li>
    <li>cmd - содержит main package;</li>
    <li>deployments - содержит файл для docker-compose.yml. Контейнеры запускаются здесь;</li>
    <li>internal - содержит пакеты для работы сервера.
        <ul>
            <li>models - модели сущностей;</li>
            <li>server - конфиг, сервер, роутер и обработчики;</li>
            <li>store - включает в себя две реализации интерфейса для работы с б.д: mock для тестирования http и postgres для работы сервера;</li>
        </ul>
    </li>
</ul>
<h3>Запуск контейнеров</h3>
<p>В директории ./deployments прописать:
    <br>
    <code>docker-compose up -d</code>
    <br>
    После этого соберутся и запустятся контейнеры для работы с go и postgres
</p>
<p>
 Посмотреть их названия можно командой:
    <br>
    <code>
        docker ps
    </code>
</p>

<h3>Создание пользователя для работы с базой данных </h3>
<ul>
    <li>Войти в контейнере c postgres в учетную запись postgres (пароль - "qwerty"):
    <br>
    <code>docker exec -it deployments_store_1 /bin/bash</code>
    <br>
        <code>psql -U postgres -p 5432 -h store</code>
    <br>
    </li>
    <li>Создать нового пользователя:
    <br>
        <code>CREATE ROLE user1 WITH PASSWORD 'password' LOGIN CREATEDB;</code>
    <br>
    </li>
    <li>Создать базы данных:
    <br>
        <code>CREATE DATABASE user1;
        <br>
        CREATE DATABASE app
        </code>
    <br>
    </li>
</ul>

<h3>Запуск миграций </h3>
<ul>
    <li>
        Зайти в контейнер с golang:
        <br>
        <code>docker exec -it deployments_backend_1 /bin/bash</code>
    </li>
    <li>В директории ./build запустить миграции командой:
    <br>
    <code>migrate -database ${POSTGRESQL_URL} -path ./migrations up</code>
</ul>


<h3>Запуск сервера</h3>
<ul>
    <li>
        В директории ./build необходимо создать .env файл и заполнить его по примеру .env_example(скопировать все из .env_example в .env)
        <br>
        <code>
            cp .env_example .env
        </code>
    </li>
    <li>
        В директории ./build необходимо прописать команду для сборки сервера:
        <br>
            <code>go build ../cmd/app/main.go</code>
        <br>
    </li>
    <li>
        Теперь его можно запустить командой:
        <br>
            <code>./main</code>
        <br>
    </li>
</ul>
<h3>API</h3>
<ul>
    <li>
        <h4>Метод сохранения статистики</h4>
        <p>
        POST /stat
        </p>
        <pre>
        пример тела запроса в формате JSON:
        {
            "date": "2010-03-09",
            "views": 30,
            "clicks": 30,
            "cost": 0.01
        }
        пример ответа:
        {
            "error": "",
            "result": "готово"
        }
        </pre>
    </li>
    <li>
        <h4>Метод показа статистики</h4>
        <p>
        GET /stat
        </p>
        <pre>
        пример запроса:
        http://127.0.0.1:3000/stat?from=2010-01-12&to=2011-02-12&orderBy=cpc

        from - дата начала периода (включительно)
        to - дата окончания периода (включительно)
        orderBy - поле для сортировки
        пример ответа:
        {
            "error": "",
            "result": [
                {
                    "date": "2010-03-09T00:00:00Z",
                    "views": 30,
                    "clicks": 30,
                    "cost": 0.01,
                    "cpc": 0.0003333333333333333,
                    "cpm": 0.3333333333333333
                },
                {
                    "date": "2010-03-07T00:00:00Z",
                    "views": 85,
                    "clicks": 13,
                    "cost": 0.01,
                    "cpc": 0.0007692307692307692,
                    "cpm": 0.11764705882352941
                },
                {
                    "date": "2010-03-05T00:00:00Z",
                    "views": 90,
                    "clicks": 10,
                    "cost": 0.01,
                    "cpc": 0.001,
                    "cpm": 0.1111111111111111
                },
                {
                    "date": "2010-03-06T00:00:00Z",
                    "views": 80,
                    "clicks": 10,
                    "cost": 0.01,
                    "cpc": 0.001,
                    "cpm": 0.125
                }
            ]
        }
</pre>
</li>
    <li>
        <h4>Метод удаления всей статистики</h4>
        <p>
        DELETE /stat
        <pre>
         пример ответа:
         {
            "error": "",
            "result": {
                "rows deleted": 5
            }
         }
        </pre>
    </li>
</ul>