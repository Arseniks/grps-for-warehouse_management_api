# JSON-RPC Warehouse management API

## Описание проекта 
    JSON-RPC сервер для работы с API управления товарами на складе

## Клонирование проекта
```bash
git clone https://github.com/Arseniks/jsonrpc_warehouse_management_api
```
```bash
cd jsonrpc_warehouse_management_api
```

## Документация

    Документация по аргументации использованных зависимостей лежит в директории docs

## Запуск проекта

### Скачать дополнительные зависимости при отсутствии
1. docker
2. docker-compose
3. postgres
4. golang-migrate

### Создать и заполнить файл .env
```bash
cp .env.template .env
```

### Выполнить команду для запуска проекта и готовой инфраструктуры в Docker
```bash
make up
```

### Выполнить команду для установки миграций БД
```bash
make up_migrations
```

## Коллекция Postman для тестирования работы API
[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/10955370-c2092ad9-10b8-49f5-a6e1-92d457fd81b0?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D10955370-c2092ad9-10b8-49f5-a6e1-92d457fd81b0%26entityType%3Dcollection%26workspaceId%3D6cf96601-bbe6-4949-a625-a9b929779df5)