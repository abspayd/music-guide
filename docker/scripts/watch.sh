#!/bin/sh
npx tailwindcss -i ./static/css/index-src.css -o ./static/css/index.css --watch=always </dev/null &
templ generate --watch &
air
