#!/bin/sh
# npm run tailwind:build:watch &
npx tailwindcss -i ./views/static/css/input.css -o ./views/static/css/output.css --watch=always </dev/null &
templ generate --watch &
air
