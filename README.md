# Delivery

## How to run

```
make up
```

## Task

Task is to develop a complete back-end system for delivery application.

## Stack

- Go.
- Postgres.
- Kafka.
- Redis.
- ELK (in the future).


## Sub-domains

- Accounts(contains auth, managing account info).
- Menu(contains serving the current menus of restaurants to the user).
- Ordering(manages the state of orders, keeps history of orders).
- Restaurant(manages the info about restaurant, its menu and orders for this restaurant).
- Courier(manages the info about couriers, their work-hours, current status, and orders that they have accepted to deliver).


## Decomposition

```
AccountsService(manages account info)
    --> AccountDB
    --> MessageQueue(Kafka)

AuthService(manages authorization and authentication)
    --> SessionStorage (Redis)
    --> MessageQueue(Kafka)

MenuService(manages menus)
    --> MenuDB
    --> MessageQueue(Kafka)

OrderingService(manages orders life-cycle)
    --> OrderDB
    --> MessageQueue(Kafka)

RestaurantService(manages restaurant)
    --> RestaurantDB
    --> MessageQueue(Kafka)

CourierService
    --> CourierDB
    --> MessageQueue(Kafka)
```

## Services

### AccountsService

AccountService contains several important features, like CRUD to account profile,
validating password and etc.

### AuthService

AuthService contains authorization and authentication to the system.
AuthService does not have access to AccountDB and it uses API of AccountsService.
The service keeps current active session in redis, just as uuid-session_token.

### MenuService

MenuService is responsible for storing the current menus of restaurants, and
provide CRUD operation for menus and dishes.

### OrderingService

OrderingService is responsible for managing the life-cycle of orders, its creation and etc. .
In addition, it keeps the state of dishes, as they were in the moment of order creation.

### RestaurantService

This service is managing the restaurant info, including the menu, using MenuService.
It provides interface for CRUD on restaurant info, accepting/rejecting/stopping the order.


### CourierService

CourierService is responsible for managing info about courier.
It provides API for accepting delivery of order, completing it, or failed it.
