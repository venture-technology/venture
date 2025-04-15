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
		docker pull $(IMAGE_NAME):latest \
		&& (docker stop venture-api || true) \
		&& (docker rm venture-api || true) \
		&& docker run -d \
			--name venture-api \
			--restart always \
			-v ~/config:/app/config \
			-p 9999:9999 \
			$(IMAGE_NAME):latest \
	'

qa-deploy-docker:
	ssh -t root@$(HOST_REMOTE_SERVER_IP) '\
		docker pull $(IMAGE_NAME):latest \
		&& (docker stop venture-api-staging || true) \
		&& (docker rm venture-api-staging || true) \
		&& docker run -d \
			--name venture-api-staging \
			--restart always \
			-v ~/config:/app/config \
			-p 9998:9998 \
			$(IMAGE_NAME):latest \
	'

migrateup:
	go run cmd/db/migrate_up/main.go

migratedown:
	go run cmd/db/migrate_down/main.go

migrateforce:
	go run cmd/db/migrate_force/main.go

send-config-json:
	rsync config/development.json root@$(HOST_REMOTE_SERVER_IP):~/config
