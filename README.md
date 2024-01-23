Сборка изображения
docker build -t spotify .

Запуск контейнера
docker run --env-file .env -p 8080:8080 -d --name spotifyapp spotify
