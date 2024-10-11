build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/venture-api cmd/api/main.go

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

