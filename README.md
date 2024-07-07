# Web приложение для отображения электронной карты местности и дополнительных информационных слоёв

# Функциональные требования

- Электронная карта
- Должен быть предусмотрен интерфейс для управления слоями отображения на карте и информацией
- Сохранение в базу данных всей информации и реализация CRUD операций

3 слоя работы:
- 1 слой - возможность поставить и показать знаки на карте с возможностью установки координат мышью и с помощью ручного ввода (до 10000 знаков)
- 2 слой - возможность нанести полигоны мышкой и редактировать каждую точку вручную (до 20 точек на полигон)
- 3 слой - возможность нанести направленный граф маршрутов (до 30 рёбер на маршрут и до 1000 маршрутов)

# Не функциональные требования

Карта Яндекс или аналог по открытой лицензии
- Все используемые технологии должны поддерживать лицензию Apache 2.0 или MIT
- Технология отображения 2d
- Разместить на localhost
- Работа со слоями должна быть реализована через API типа REST

# Миграция базы данных

Поднятие
```sh
./migrate -path ./schema/ -database 'postgres://postgres:{POSTGRES_PASSWORD}@{POSTGRES_ADDRESS}/postgres?sslmode=disable' up
```

Попускание
```sh
./migrate -path ./schema/ -database 'postgres://postgres:{POSTGRES_PASSWORD}@{POSTGRES_ADDRESS}/postgres?sslmode=disable' down
```
