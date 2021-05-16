# **Запуск**
sudo docker-compose up  
**документация swagger**  
http://127.0.0.1:8080/swagger/index.html  

### **Тестирование**
### добавление нового пользователя 
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/add-user \
  --header 'Content-Type: application/json' \
  --data '{
	"name":"Ted"
}'
```
### добавление лайка 
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/add-like \
  --header 'Content-Type: application/json' \
  --data '{
	"video_id":"122",
	"user_id":"1",
	"create":"2021-05-15 15:04:05"
}'
```
### добавление коммента
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/add-comment \
  --header 'Content-Type: application/json' \
  --data '{
	"video_id":"122",
	"text":"some text",
	"user_id":"1",
	"create":"2021-05-15 15:04:05"
}'
```

### просмотр активных пользователей за сегодня
  ```bash
curl --request GET \
  --url http://127.0.0.1:8080/api/v1/get-active-today
```