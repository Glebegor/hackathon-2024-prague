rundb:
	sudo docker run --name=tonsub-db -e POSTGRES_PASSWORD="123321" -p 5436:5432 --rm -d postgres
