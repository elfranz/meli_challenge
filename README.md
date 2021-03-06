# meli_challenge

Welcome! This is a challenge for Mercadolibre.

## Installation

* Clone this repo
```
git clone https://github.com/elfranz/meli_challenge
```
* Install [docker-compose](https://docs.docker.com/compose/install/).

## Usage

* Read the [Postman docs](https://documenter.getpostman.com/view/11207309/TVCiT6m9) for information on how to invoke the different endpoints.

* Change to the created directory
```
cd meli_challenge
```

* Build and start containers with
```docker-compose up --build```

You are now ready to use the app.

## Unit Testing

* Access the api container with
```
docker-compose exec app bash
```
* Inside the container do
```
cd src/api/app/items
go test
```

## Integration Testing

* Download and install [Postman](https://www.postman.com/)

* Go to the [docs](https://documenter.getpostman.com/view/11207309/TVCiT6m9).

* On the top right corner, click "Run in Postman", then select the desktop option (web won't work since this isn't deployed on a public server) and open the collection in your local Postman client.

* Make sure the containers are up, you can do ```docker-compose ps -a``` on your terminal to list the active processes, you should see the three containers up (db, nginx and app).

* Run the whole regression by running the collection in your postman client. To do this click on the arrow on the imported collection, then RUN and then RUN again.

## Additional

You can read about the development of this app on https://docs.google.com/document/d/1qNNp9GIijiL8KNwgjoOHmjyt6VaHal300Mc7DpgM1f0/edit?usp=sharing.
This was a really fun and elightening experience, it took me out of Ruby's comfort zone and challenged me to expand my knowledge about infrastructure and programming. I learned a lot about Go and Nginx in very little time so in general it was very positive despite the anxiety and stress the unit testing gave me :grin:.
