# streamer

## Для запуска nats-streaming и postgres
```
docker-compose up -d
```

## Для запуска консюмера из корня:
```
go run ./cmd/consumer
```

## Для запуска паблишера из корня:
```
go run ./cmd/publisher
```

## Фронт доступен по адресу: 
```
127.0.0.1:8081
```