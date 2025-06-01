# Forum Project Modules

Этот репозиторий содержит модули, которые используются в проекте форума.

## Структура

```
modules/
├── auth/              # Модуль аутентификации
├── forum/            # Модуль форума
├── proto/            # Общие прото-файлы
├── common/           # Общие утилиты и интерфейсы
└── go.mod           # Общий файл зависимостей
```

## Использование

1. Клонировать репозиторий:
```bash
git clone https://github.com/yourusername/forum-modules.git
```

2. Добавить в go.mod основного проекта:
```go
require github.com/yourusername/forum-modules v0.1.0
```

3. Импортировать нужные пакеты:
```go
import (
    "github.com/yourusername/forum-modules/auth"
    "github.com/yourusername/forum-modules/forum"
    "github.com/yourusername/forum-modules/proto"
    "github.com/yourusername/forum-modules/common"
)
```

## Модули

### Auth Module
Модуль аутентификации и авторизации.

### Forum Module
Модуль форума с темами и комментариями.

### Proto Module
Общие прото-файлы для всех сервисов.

### Common Module
Общие утилиты, интерфейсы и константы. 