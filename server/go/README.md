# Installation

Install golang
Set GOPATH:
    cd to directory of this file
    export GOPATH=$(pwd)

go get github.com/gorilla/mux
go get golang.org/x/net/context
go get golang.org/x/oauth2
go get golang.org/x/oauth2/facebook
go get github.com/asaskevich/govalidator
go get github.com/lib/pq

# Run
### Make sure to run in path with view folder
cd go
go build main && ./main

# Sample API Commands:
{"action": "login", "token_type": "fb", "token": "EAAEfJLEyKMkBAA2YTLjO0Cxh5c7DEmQyIaXgYjzdideZAlRaRAkNP6mV4kumn8z8djYXDXd3IvOk5nZAzwZByzemZBCkqXRmM7pmLIOoXyG3m5mPxIQO3pZBp9qTflncXQnhzU4dT3ZCK7wS956wPOIHcs437C7ubKOmXZAQuKp3bZBTIzwDF4t8"}
{"action": "create_card", "card_data": "lalal", "token_type": "fb", "token": "EAAEfJLEyKMkBAFevPSArRGSaxCabmyzLjx200OdMuQXsZCmZCMbh3pqnA15ZA03yNrCyAGHhUzZAPOEzqa8iJv1rZCrk79RwbU46Al4FFxYlbbrLUiX6WV8VZA0RbET8xzzaJmL7RBPKpnEE8N0xKO3Ngzp9aAIvwS1emPwOnjZAFd76x6zugjC"}
{"action": "create_deck", "deck_name": "lalal", "token_type": "fb", "token": "EAAEfJLEyKMkBAFevPSArRGSaxCabmyzLjx200OdMuQXsZCmZCMbh3pqnA15ZA03yNrCyAGHhUzZAPOEzqa8iJv1rZCrk79RwbU46Al4FFxYlbbrLUiX6WV8VZA0RbET8xzzaJmL7RBPKpnEE8N0xKO3Ngzp9aAIvwS1emPwOnjZAFd76x6zugjC"}
{"action": "create_collection", "collection_name": "collec", "token_type": "fb", "token": "EAAEfJLEyKMkBAFevPSArRGSaxCabmyzLjx200OdMuQXsZCmZCMbh3pqnA15ZA03yNrCyAGHhUzZAPOEzqa8iJv1rZCrk79RwbU46Al4FFxYlbbrLUiX6WV8VZA0RbET8xzzaJmL7RBPKpnEE8N0xKO3Ngzp9aAIvwS1emPwOnjZAFd76x6zugjC"}
{"action": "list_cards", "token_type": "fb", "token": "EAAEfJLEyKMkBAAIxzrFlvmv4O8gBxLCfMHnux6bNvGmdR2GZCb0LVMGvVLjS3TQJlXRZBrq0UFlR6e7p2jbIW7Ywk1A70ZC8LEZAto3OgRIthQoYxOsF2Bta4DhGT8DfmTPtPLz6idO5iy8M2KrZAwDYVnNhgIeAkXcMCZC14s8dxtDQ5ig9VD"}
{"action": "list_decks", "token_type": "fb", "token": "EAAEfJLEyKMkBAAIxzrFlvmv4O8gBxLCfMHnux6bNvGmdR2GZCb0LVMGvVLjS3TQJlXRZBrq0UFlR6e7p2jbIW7Ywk1A70ZC8LEZAto3OgRIthQoYxOsF2Bta4DhGT8DfmTPtPLz6idO5iy8M2KrZAwDYVnNhgIeAkXcMCZC14s8dxtDQ5ig9VD"}
{"action": "list_collections", "token_type": "fb", "token": "EAAEfJLEyKMkBAAIxzrFlvmv4O8gBxLCfMHnux6bNvGmdR2GZCb0LVMGvVLjS3TQJlXRZBrq0UFlR6e7p2jbIW7Ywk1A70ZC8LEZAto3OgRIthQoYxOsF2Bta4DhGT8DfmTPtPLz6idO5iy8M2KrZAwDYVnNhgIeAkXcMCZC14s8dxtDQ5ig9VD"}
{"action": "replace_card", "card_id": 2, "card_data": "wowow", "token_type": "fb", "token": "EAAEfJLEyKMkBAFQZCu4542ev5CzBCEZCX7MCcTs0XiSI2HlwjuJzvIi9PbFZAV7Ne0lKBi8ZAiKpiAvacVWq6s5If7Y13JpK2ZCr8ldtzRYoBZAej0fpLKPbZBDMBvEgr5kxTeyfVQv30stCjEizxrS7c0pPcylMJxrZCDI0Srv7gqVfRsSjqnCo"}
{"action": "rename_deck", "deck_id": 2, "deck_name": "My awesome dek", "token_type": "fb", "token": "EAAEfJLEyKMkBAFQZCu4542ev5CzBCEZCX7MCcTs0XiSI2HlwjuJzvIi9PbFZAV7Ne0lKBi8ZAiKpiAvacVWq6s5If7Y13JpK2ZCr8ldtzRYoBZAej0fpLKPbZBDMBvEgr5kxTeyfVQv30stCjEizxrS7c0pPcylMJxrZCDI0Srv7gqVfRsSjqnCo"}
{"action": "rename_collection", "collection_id": 2, "collection_name": "My awesome dek", "token_type": "fb", "token": "EAAEfJLEyKMkBAFQZCu4542ev5CzBCEZCX7MCcTs0XiSI2HlwjuJzvIi9PbFZAV7Ne0lKBi8ZAiKpiAvacVWq6s5If7Y13JpK2ZCr8ldtzRYoBZAej0fpLKPbZBDMBvEgr5kxTeyfVQv30stCjEizxrS7c0pPcylMJxrZCDI0Srv7gqVfRsSjqnCo"}
{"action": "add_card_to_deck", "card_id": 2, "deck_id": 1, "token_type": "fb", "token": "EAAEfJLEyKMkBAFQZCu4542ev5CzBCEZCX7MCcTs0XiSI2HlwjuJzvIi9PbFZAV7Ne0lKBi8ZAiKpiAvacVWq6s5If7Y13JpK2ZCr8ldtzRYoBZAej0fpLKPbZBDMBvEgr5kxTeyfVQv30stCjEizxrS7c0pPcylMJxrZCDI0Srv7gqVfRsSjqnCo"}
{"action": "add_deck_to_collection", "collection_id": 1, "deck_id": 1, "token_type": "fb", "token": "EAAEfJLEyKMkBAJwXeXZCoqhcmD7qGsq59Y9IsvPep4DZCRMN18zz4j7ZCZBvHYxX5HmYBZAO7DzU0aq5fF4YqeVwMuZAEjU6GBzJJoZCXF6XvIZARtofk4Hgiu8UXfInFPvtZB8h3ZA5QnNLgwReu6ZBJpHRkGDBZBSzRTZC28vtyM7dAcltWYgIakN9G"}