# Go (Golang) Developer Basic 2024-09

_Сафаров Алексей Расимович_  
email: <safarov-ar@yandex.ru>  
г.Тверь

# Конвертер валют

### Описание:

В этом проекте студентам предлагается написать программу командной строки, которая конвертирует сумму из одной валюты в другую. Для получения курсов валют программа может использовать стороннее API, например, API открытых данных Центрального банка или сторонних сервисов.

### Требования:

- Программа должна принимать ввод суммы и валюты, из которой нужно конвертировать, а также валюты, в которую нужно сконвертировать.
- Программа должна использовать сторонний API для получения курсов валют.
- Реализована обработка ошибок, связанных с соединением с API или некорректными входными данными.
- Результаты конвертации должны быть точными и учитывать валютные курсы.

### Развертывание

Развертывание сервиса должно осуществляться с использованием docker compose в директории с проектом.

### Тестирование

Написаны юнит-тесты на core логику приложения. Плюсом будут тесты на транспортном уровне и на уровне хранения.

### Критерии оценивания

Максимум - 15 баллов (при условии выполнения обязательных требований):

- Реализован алгоритм - 2 балла.
- Реализовано разделение на слои (транспортный, хранения и т.д.) - 2 балла.
- Реализовано API сервиса - 2 балла.
- Реализован консольный пользовательский интерфейс - 2 балла.
- Написаны юнит-тесты - 1 балл.
- Написаны интеграционные тесты - 2 балла.
- Тесты адекватны и полностью покрывают функциональность - 1 балл.
- Понятность и чистота кода - до 3 баллов.

Зачёт от 10 баллов

docker compose build
docker compose run app
