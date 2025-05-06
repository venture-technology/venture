build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/venture-api cmd/api/server/main.go

send:
	rsync dist/venture-api root@$(HOST_REMOTE_SERVER_IP):~

send-service:
	rsync venture-api.service root@$(HOST_REMOTE_SERVER_IP):~

deploy: build send send-service
	ssh -t root@$(HOST_REMOTE_SERVER_IP) '\
		sudo mv ~/venture-api.service /etc/systemd/system \
		&& sudo systemctl enable venture-api \
		&& sudo systemctl restart venture-api \
	'

send-nginx-config:
	rsync nginx.conf root@$(HOST_REMOTE_SERVER_IP):~/config-venture

deploy-http-server:
	ssh -t root@$(HOST_REMOTE_SERVER_IP) '\
		sudo mv /etc/nginx/sites-available/site.conf /etc/nginx/sites-available/site.conf.bak \
		&& sudo mv ~/config-venture/nginx.conf /etc/nginx/sites-available/site.conf \
	'

deploy-nginx: send-nginx-config deploy-http-server
	ssh -t root@$(HOST_REMOTE_SERVER_IP) '\
		service restart nginx \
	'

prod-deploy-docker:
	ssh -t root@$(HOST_REMOTE_SERVER_IP) '\
		cd ~/infrastructure \
		&& docker compose pull venture-server \
		&& docker compose up --build -d \
	'

staging-deploy-docker:
	ssh -t root@$(HOST_REMOTE_SERVER_IP) '\
		cd ~/infrastructure \
		&& docker compose pull venture-server-staging \
		&& docker compose up --build -d \
	'

migrateup:
	go run cmd/db/migrate_up/main.go

migratedown:
	go run cmd/db/migrate_down/main.go

migrateforce:
	go run cmd/db/migrate_force/main.go

send-config-json:
	rsync config/development.json root@$(HOST_REMOTE_SERVER_IP):~/infrastructure/apis/venture

api:
	go run cmd/api/server/main.go
