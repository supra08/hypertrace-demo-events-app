# Hypertrace Demo Events App
This app is designed to demonstrate the basic functionality of Hypertrace. It has a very simple architecture with mainly 2 services, the frontend and the backend. Both the services are equipped with opentracing+jaegar clients to send traces of all the requests to the collector and henceforth, for Hypertrace to detect them.
The app is designed to be highly configurable, so that it can be used in and to demostrate a variety of contexts. By the default, the app is configured to show events and comments and the users can add more events/comments. 
The app can be configured by -
* Editing the `events.json` to edit the default events and the comments.
* Editing the `services/frontend/web_assets/src/appConfig.json` to make the frontend display data in a way you like.

If there are changes in the frontend, build the project CLI tool first and then run `./<cli-binary> gen-frontend` to regenrate the frontend assets.

## Working with the demo app
* Start the app with `docker-compose up`.
* Start Hypertrace Docker Compose version by running `docker-compose up` from the hypertrace project.
* The app would be running at `localhost:3000`. Play around the app a bit to generate traces.
* Visit the Hypertrace UI at `localhost:2020` to check the traces.

## Hacking the demo app
* After you have done configuring the `events.json`, `appConfig.json` and regenerating the frontend, run `docker build -t supradeux/events-app .` and run the app.

The app docker image is available at `supradeux/events-app:latest`
