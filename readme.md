# Генерация конфигурации для авторизационного прокси

Чтобы сгенерировать `nginx.conf` выполните

    make > nginx.conf

Результат работы программы выводится в `STDOUT`, а весь вывод сообщений о компиляции перенаправлен в `STDERR`.
Таким образом команда выше создаст валидный `nginx.conf`.

По умолчанию для настройки сервисов используется файл `services.yml`, а для настройки роутинга `locations.yml`.
Переопределить эти значения можно с помощью переменных окружения `SERVICES_PATH` и `LOCATIONS_PATH`.

Если расширение конфигурационного файла отличается от `yml` (например `conf`), то он будет считаться частью конфигурации *nginx*
и вставлен в вывод без изменений.

    SERVICES_PATH=services.conf make > nginx.conf

Пример consul template:

```
upstream service-name {
  <<< range service "full-service-name" >>>
  server <<< .Address >>>:<<< .Port >>>;
  <<< end >>>
}
```

Шаблон `nginx.tmpl` написан с использованием [go template](https://golang.org/pkg/text/template/).
