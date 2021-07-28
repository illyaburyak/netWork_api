.PHONY : pre ext db prod dev stop editor
pre:
	cp -n .env.example .env
	cp -n .env.example .env.docker
	sed -i 's/localhost/postgres/' .env.docker

ext:
	docker-compose exec -u postgres postgres psql -d test -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp";'

clear:
	docker-compose exec -u postgres postgres psql -d test -c 'DROP SCHEMA public CASCADE;CREATE SCHEMA public;'

db: pre
	docker-compose up -d postgres

prod: db ext
	docker-compose up -d network_api

editor:
	docker-compose up -d editor

dev: db ext
	go run .

reflex: db ext
	reflex --verbose -r "\.go" -s -- sh -c "go run ./"

stop:
	docker-compose down