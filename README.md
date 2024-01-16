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
make service-up
```

### Выполнить команду для установки миграций БД
```bash
make up_migrations
```
