# Quest Client

React приложение с аккордеон-меню для отображения процедур розыска посылок.

## Установка зависимостей

```bash
npm install @radix-ui/react-accordion
```

## Запуск

```bash
npm run dev
```

Приложение будет доступно по адресу `http://localhost:5173`

## Настройка API

По умолчанию приложение подключается к `http://localhost:8000/api/v1`.

Для изменения URL API создайте файл `.env` в корне проекта `client`:

```
VITE_API_URL=http://localhost:8000/api/v1
```

## Структура

- `src/components/procedures-accordion.tsx` - основной компонент аккордеона
- `src/components/ui/accordion.tsx` - компонент Accordion из shadcn-ui
- `src/lib/api.ts` - API клиент для получения данных
- `src/types/procedure.ts` - TypeScript типы для процедур
