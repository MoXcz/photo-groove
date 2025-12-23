web:
	elm make src/PhotoGallery.elm --output ./static/photos.js
	elm make src/PhotoFolders.elm --output ./static/folders.js
	elm make src/Main.elm --output ./static/app.js

server:
	go build -o ./tmp/srv && ./tmp/srv
