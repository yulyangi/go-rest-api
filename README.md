# Simple GO application without UI

## Endpoints
* `GET /events`                    Get list of events
* `GET /events/<id>`               Get event by id
* `POST /events`                   Create event `auth_required`
* `PUT /events/<id>`               Update event `auth_required` `only_by_creator`
* `DELETE /events/<id>`            Delete event `auth_required` `only_by_creator`
* `POST /singup`                   Create user
* `POST /login`                    Authenticate user `auth_token`
* `POST /events/<id>/register`     Register user for event `auth_required`
* `DELETE /events/<id>/register`   Cancel registration for event `auth_required`

## Build
```
go get ./...
go build .
```

## Run
`./go-res-api`

## Test
### You can use CURl or REST Client Extension for VSCode
#### create a user
```
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"email":"johndow@mail.com","password":"mypassword"}' \
     http://localhost:8080/signup
```

```
POST   http://localhost:8080/signup HTTP/1.1
content-type: application/json

{
    "email": "johndow@@mail.com",
    "password": "mypassword"
}
```

#### log in as user (token will be returned)
```
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"email":"johndow@mail.com","password":"mypassword"}' \
     http://localhost:8080/login
```

```
POST   http://localhost:8080/login HTTP/1.1
content-type: application/json

{
    "email": "johndow@@mail.com",
    "password":"mypassword"
}
```

#### create event
```
curl --header "Content-Type: application/json" \
     --header "Authorization: <token>" \
     --request POST \
     --data '{"title": "Wacken Open Air", "description": "heavy metal music festival", "location": "Wacken, Germany", "date_time": "2025-07-30T00:00:00.000Z"}' \
     http://localhost:8080/events
```

```
POST   http://localhost:8080/events HTTP/1.1
content-type: application/json
Authorization: <token>

 {
   "title": "Wacken Open Air",
   "description": "heavy metal music festival",
   "location": "Wacken, Germany",
   "date_time": "2025-07-30T00:00:00.000Z"
 }
```

#### update event
```
curl --header "Content-Type: application/json" \
     --header "Authorization: <token>" \
     --request PUT \
     --data '{"title": "Bloodstock Open Air", "description": "heavy metal music festival", "location": "Walton-on-Trent, Derbyshire, England", "date_time": "2025-08-07T00:00:00.000Z"}' \
     http://localhost:8080/events/<id>
```

```
PUT   http://localhost:8080/events/<id> HTTP/1.1
content-type: application/json
Authorization: <token>

{
    "title": "Bloodstock Open Air",
    "description": "heavy metal music festival",
    "location": "Walton-on-Trent, Derbyshire, England",
    "date_time": "2025-08-07T00:00:00.000Z"
}
```

#### get event
```
curl http://localhost:8080/events/<id>
```

```
GET   http://localhost:8080/events/<id> HTTP/1.1
```

#### get events
```
curl http://localhost:8080/events
```

```
GET   http://localhost:8080/events HTTP/1.1
```

#### delete event
```
curl --header "Authorization: Bearer <token>" \
     --request DELETE \
     http://localhost:8080/events/<id>
```

```
DELETE   http://localhost:8080/events/<id> HTTP/1.1
Authorization: <token>
```

#### create registration
```
curl --header "Authorization: Bearer <token>" \
     --request POST \
     http://localhost:8080/events/<id>/registration
```

```
POST   http://localhost:8080/events/<id>/registration HTTP/1.1
Authorization: <token>
```

#### delete registration
```
curl --header "Authorization: Bearer <token>" \
     --request DELETE \
     http://localhost:8080/events/<id>/registration
```

```
DELETE   http://localhost:8080/events/<id>/registration HTTP/1.1
Authorization: <token>
```
