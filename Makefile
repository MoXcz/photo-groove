web:
	elm make src/PhotoGroove.elm --output ./static/app.js
	elm make src/PhotoFolders.elm --output ./static/b.js

server:
	go build -o ./tmp/srv && ./tmp/srv
