web:
	elm make src/PhotoGroove.elm --output ./static/app.js

server:
	go build -o ./tmp/srv && ./tmp/srv
