* `GET /events`                    Get list of events
* `GET /events/<id>`               Get event by id
* `POST /events`                   Create event `auth_required`
* `PUT /events/<id>`               Update event `auth_required` `only_by_creator`
* `DELETE /events/<id>`            Delete event `auth_required` `only_by_creator`
* `POST /singup`                   Create user
* `POST /login`                    Authenticate user `auth_token`
* `POST /events/<id>/register`     Register user for event `auth_required`
* `DELETE /events/<id>/register`   Cancel registration for event `auth_required`


## create user
```
POST   http://localhost:8080/signup HTTP/1.1
content-type: application/json

{
    "email": "user_1@mail.com",
    "password": "password"
}
```

## log in as user
```
POST   http://localhost:8080/login HTTP/1.1
content-type: application/json

{
    "email": "user_1@mail.com",
    "password": "password"
}
```

## create event
```
POST   http://localhost:8080/events HTTP/1.1
content-type: application/json
Authorization: <token>

{
    "title": "Wacken Open Air",
    "description": "metal",
    "location": "Germany",
    "date_time": "2025-07-31T15:01:00.000Z"
}
```

## update event
```
PUT   http://localhost:8080/events/<id> HTTP/1.1
content-type: application/json
Authorization: <token>

{
    "title": "Viva Braslav",
    "description": "pop",
    "location": "Belarus",
    "date_time": "2025-07-31T15:01:00.000Z"
}
```

## get event
```
GET   http://localhost:8080/events/<id> HTTP/1.1
```

## get events
```
GET   http://localhost:8080/events HTTP/1.1
```

## delete event
```
DELETE   http://localhost:8080/events/<id> HTTP/1.1
Authorization: <token>
```

## create registration
```
POST   http://localhost:8080/events/<id>/registration HTTP/1.1
Authorization: <token>
```

## delete registration
```
DELETE   http://localhost:8080/events/<id>/registration HTTP/1.1
Authorization: <token>
```
