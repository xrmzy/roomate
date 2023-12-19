
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


## Configuration

To run this project, you will need to add the following environment variables to your .env file

`API_PORT=1111`

`DB_HOST=localhost`

`DB_PORT=5432`

`DB_NAME=db_roomate`

`DB_USER=root`

`DB_PASSWORD=root`

`DB_DRIVER=postgres`


`ACCESS_TOKEN_PRIVATE_KEY = GET THIS WHEN YOU CREATE USER`

`ACCESS_TOKEN_EXPIRED_IN=15m`

`ACCESS_TOKEN_MAXAGE=15`

`REFRESH_TOKEN_PRIVATE_KEY= GET THIS WHEN YOU RUN OUT OF TIME`

`REFRESH_TOKEN_EXPIRED_IN=60m`

`REFRESH_TOKEN_MAXAGE=60`

## Run Locally

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


## Authors

- [@imkhoirularifin](https://github.com/imkhoirularifin)
- [@xrmzy ](https://github.com/xrmzy)
- [@maiing11 ](https://github.com/maiing11)
- [@Dimazz](https://github.com/Dimazzs)



## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.

