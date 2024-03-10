# Welcome to my order-service with nats-streaming!

## How to use

First of all you need to PostgreSql and nats-straming using Docker, then you need to run nats-streaming publisher to 
create stream. And finally you can run the main service which can accept requests to receive an order by his order_uid. 

#### To run PostgreSql and nats-streaming:
    docker compose up

#### To run nats-streaming publisher to publish ONE order:
    go run .\nats-streaming-publish\publisher.go

#### To run main service:
    go run .\order-service\cmd\main.go

Thank you for checking out my project! Feel free to reach out if you have any questions or feedback.


   



