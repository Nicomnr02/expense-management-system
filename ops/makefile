soda:
	go install github.com/gobuffalo/pop/v6/soda@latest
sodacfg:
	soda g config	

db: 
	soda create -e "$(DB)"
dropdb:
	soda drop -e "$(DB)"

migration: #FN = FileName
	soda generate sql $(FN) -p ./migrations
migrate-up:
	soda migrate up
migrate-down:
	soda migrate down -e "$(ENV)"



