# Clean Arch Golang Best Practices

* credit-executor - основное приложение
* credit-adm-executor - приложение управления и контроля
* credit-shared-module - модуль для хранения общей логики или объектов
* credit-library - библиотека для унификации работы с системными утилитами или общих технических методов (что бы не копировать каждый раз), в реальном мире это отдельный репозиторий

## Справка по архитектуре
https://github.com/alfssobsd/notes/blob/main/golang/arch/golang_arch_description.md


## Основные правила при написание кода
1. Любой UseCase,Entrypoint, Dataprovider, Util должно быть легко замокать через конструктор для тестирования
2. Использование singleton нужно избегать
3. Если объект делят несколько gorutines (каждый запрос в echo рождает новую gorutine), то нужно, что бы объект был в монопольном доступе или же использовать мьютексы.
   1. Исключения - объекты которые не изменяются после старта приложения, стоит их передавать по значению (пример appconfig)
4. в Dataprovider слое может быть Provider, Gateway, Repository, ServiceAdapter и тд, главное, что бы точно было понятно для чего используется

## Основные примеры кода на которые следует обратить внимание
### DI
```
credit-executor/http_app_server.go
```
### Конвертирование DTO из одного слоя в другой
```
credit-adm-executor/entrypoints/http_controllers/loan_customer_adm_http_controller_dto.go
```

### Вынесение кода репозитория в общий модуль
```
credit-shared-module/dataproviders/main_db_provider
```

### Фоновые задачи это так же входные точки
```
credit-executor/entrypoints/background/heavyprocessor_watcher_bg_task.go
```

### Конфиг у каждого приложения свой хот и похож
```
credit-adm-executor/utils/appconfig/config.go
credit-executor/utils/appconfig/config.go
```

### Использование мютексов и пулов ресурсов
Так же стоит обратить внимание как сделан фасад
```
credit-shared-module/utils/heavyprocessor
```