# DND Generator v0.1

## Stuff generator
Генерирует вещи на основе количества и выбранных категорий подходит для торговцев и лута.

## Установка
### Требования
- PHP > 7
- PostgreSQL или MySQL (настраивается в конфиге)

### Процедура установки
- Развернуть базу
- Создать env.php файл из env.php.example
- Запустить public/index.php

### Запуск с помощью Docker
- `docker-compose up -d`
- `docker-compose down` чтобы погасить контейнеры

### Запуск бандлера
- `npm run build` чтобы просто собрать ассеты в режиме разработки
- `npm run prod` чтобы собрать в продакшн-режиме
- `npm run dev` бандлинг в режиме hot reload