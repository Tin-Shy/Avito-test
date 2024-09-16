Инструкция к использованию(все команды должны вводиться в папке src):

    Для запуска:
    
1. Собираем Dockerfile "sudo docker build -t tender_service ."
2. Запускаем Dockerfile с портами 8080 "sudo docker run -p 8080:8080 tender_service"
3. Проверяем "http://localhost:8080/ping" или "curl -i http://localhost:8080/ping"    

Ожидаемый вывод:

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Mon, 16 Sep 2024 09:08:03 GMT    
Content-Length: 2

