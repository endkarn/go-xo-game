run:
	cd game/ && go run main.go

atdd:
	newman run atdd/api/xogame.json -e atdd/environment/postman_environment.json