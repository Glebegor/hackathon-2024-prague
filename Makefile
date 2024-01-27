rundb:
	sudo docker run --name=ton-event-db -e POSTGRES_PASSWORD="123321" -p 5555:5432 --rm -d postgres
