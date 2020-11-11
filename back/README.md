# Elivia - Backend

## Description

As explained in the main README, Elivia is a vocal assistant and this is her back end.
In this repository, you will find the AI itself, the packages added to the AI and the websocket server linking it with the front end.

## Quick start

All the project has been wrote using go, so nothing more complicated than running it because go will download all required dependencies.

```shell
go run ./
```

## Docker

We also provided a Dockerfile if you want to run the backend in an isolated environment.
You can build the docker image and run the docker container using the following commands : 

```shell
docker build -t go-elivia:latest .
docker run -e "PORT=8080" -p 8080:8080 go-elivia
```

## Features

### AI

As said right before, the main part about Elivia is the IA and has been extracted from [Olivia's](https://github.com/olivia-ai/the-math-behind-a-neural-network). For further documentation about the way it works we strongly recommend checking the Olivia repository.

### Packages

One thing we added upon it is the **package** system. Elivia is able to load, compile and train package at runtime allowing you to add, remove or edit response at runtime.
Its the main operation of Elivia. You can format response and trigger, call your own function and safely extract data from the sentences that your package has triggered .

**Packages** seems complicated but they aren't really that much of deal. And there is a full [documentation](https://github.com/PoCInnovation/Elivia/blob/master/back/PACKAGES.md) on how to write them, which features are implemented and the best practices around them. 

### Websocket

Elivia uses Websocket to communicate with the frontend. The only route used is the serve one, allowing you to handle a sentence through the AI.

## What's next ?

Elivia is a PoC project, so the focus may change a bit depending of the team taking over the project in the next months, but the main features left to implement are :

* Enhance the AI. The current one works well, but with a **Text Data Vectorization** algorithm extracting meaning could have better result thus, leading to less misunderstanding error

* Make it a runtime. It is not really a runtime interpreter of packages. It does, indeed, compile and run every package at runtime, but it doesn't reload itself yet. This is an easy update that would allow a permanently running soft without any restart needed.

* Package data. We would like to implement a tool making user's data manipulation easy inside of package, thus, you would be able to remember previous action from the user and act upon them.

* Packages, packages and again, packages. We want to add enough of them so user would be able to choose which one they want and load only the one needed when making request with a user thus having a "store" of packages from which user would build their own little IA.

As you can see, we have plenty of ideas and this is only a small grasp of what we plan to do.