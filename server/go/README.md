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