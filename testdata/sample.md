# RunDoc Sample Documentation

Это пример исполняемой документации. Блоки кода с аннотацией `run` будут проверены.

## Пример 1: Bash с точным совпадением

```bash run
echo "Hello, RunDoc!"
# expect: Hello, RunDoc!
```

## Пример 2: Bash + арифметика
```bash run
echo 1+1
# expect: 1+1
```