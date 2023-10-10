# ntpclient

Программа осуществляет вывод точного времени с использованием протокола ```ntp```.

# Установка:

```bash
git clone git@github.com:bersen66/ntpclient.git
cd ntpclient
go install
```

# Использование:

```bash
ntpclient time # Вывод текущего времени
ntpclient time --server=ntp0.NL.net # Используя свой сервер
```
