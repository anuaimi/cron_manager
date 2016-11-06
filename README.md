


# build

go build -o cron_manager main.go api.go

docker build -t cron_manager
docker run -it --rm --name cron_manager cron_manager


