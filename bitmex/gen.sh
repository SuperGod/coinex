swagger generate model -f ./swagger.json
swagger generate client -f ./swagger.json  -A APIClient --default-scheme=https
