
#RooMate App


"RooMate a Smart Solution for Hotel Booking Room"

This Application created by :
@xrmzy
@maiing11
@imkhoirularifin
@Dimazz



## Installation

Install roomate with git clone

```bash
git clone https://link-to-project
```

## Environment variables

<details>
    <summary>Variables Defined in the project </summary>

    | Key                       | Value                    | Desc                                        |
    | ------------------------- | ------------------------ | ------------------------------------------- |
    | `API_PORT`                | `5000`                   | Port at which app runs                      |
    | `DB_USER`                 | `username`               | Database Username                           |
    | `DB_PASS`                 | `password`               | Database Password                           |
    | `DB_HOST`                 | `0.0.0.0`                | Database Host                               |
    | `DB_PORT`                 | `3000`                   | Database Port                               |
    | `DB_NAME`                 | `test`                   | Database Name                               |
    | `ACCESS_TOKEN_KEY`        | `secret`                 | Authentication                              |

</details>

### Run Locally

Clone the project

```bash
git clone https://link-to-project
```

Go to the project directory

```bash
cd roomate
```

Install dependencies

```bash
go mod download
go get -u
```

Start the server

```bash
go run .
```


## Documentation

Roomate Postman Collection

This Postman collection contains various endpoints to interact with the Roomate application. We also including the postman collection for this project so you can testing it on postman if you want.
Below are the instructions on how to use each request, not at all but at least we have some example for testing the API's.




```bash
CREATE ROLE

Endpoint: POST {{URL_PATH}}/roles

{
    "roleName": "test"
}

User
Endpoint: POST {{URL_PATH}}/users

{
    "name": "John",
    "email": "john@gmail.com",
    "password": "167916",
    "roleId": "1"
}

Customer
Endpoint: POST {{URL_PATH}}/customers

{
    "name": "Arifin Customer",
    "email": "arifin@gmail.com",
    "address": "Semarang",
    "phoneNumber": "911"
}


Room
Endpoint: POST {{URL_PATH}}/rooms
{
    "roomNumber": "02",
    "roomType": "President",
    "capacity": 3,
    "facility": "Laundry, Sauna, Bathub",
    "price": 900000
}

Create Service
Endpoint: POST {{URL_PATH}}/services
Body: json

{
    "name": "Dolby Atmos Sound System",
    "price": 200000
}


## Create Booking
Endpoint: POST {{URL_PATH}}/bookings
{
    "checkIn": "2023-01-10",
    "checkOut": "2023-01-11",
    "userId": "58d716cd-96e2-4150-b5ae-d0df7cfe979e",
    "customerId": "4c2b3a7a-e4d8-4b50-88e1-8ad4773e7379",
    "bookingDetails": [
        {
            "roomId": "R1277",
            "services": [
                {
                    "serviceId": "S88209"
                }
            ]
        }
    ]
}


## READ
Get All Users

Endpoint: GET {{URL_PATH}}/users
Get All Roles

Endpoint: GET {{URL_PATH}}/roles
Get All Customers

Endpoint: GET {{URL_PATH}}/customers

Get All Rooms
Endpoint: GET {{URL_PATH}}/rooms
Get All Services
Endpoint: GET {{URL_PATH}}/services

Get Role
Endpoint: GET {{URL_PATH}}/roles/2

Get User
Endpoint: GET {{URL_PATH}}/users

Get Customer

Endpoint: GET {{URL_PATH}}/customers/customer_id

Get Room
Endpoint: GET {{URL_PATH}}/rooms/room_id

Get Service
Endpoint: GET {{URL_PATH}}/services/service_id

```
## Deployment

To deploy this project to docker 


```bash
  docker compose up
```


## Implemented Features
- Routing (gin web framework)
- Environment Files
- Logging (logrus)
- Middlewares
- Database Setup (posgresql)
- Models Setup
- Repositories
- Implementing Basic CRUD Operation
- Authentication (JWT)
- Google sheet APi (downloading booking reports)
- Dockerize Application
- Unit testing

## Authors

- [@imkhoirularifin](https://github.com/imkhoirularifin)
- [@xrmzy ](https://github.com/xrmzy)
- [@maiing11 ](https://github.com/maiing11)
- [@Dimazz](https://github.com/Dimazzs)

## Todos
- [] File Upload Middleware examples. https://github.com/example

## Contributing

Please open issues if you want the template to add some features that is not in todos. 🙇‍♂️

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.

