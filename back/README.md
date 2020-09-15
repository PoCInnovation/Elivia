# Elivia - Backend

## Description

As explained in the main README, Elia is a vocal assistant, and this is her back end.
In this repository, you will find the IA itself, the packages added to the IA, and the websocket server linking it with the front end

## Quick start

all the project is made in go, so nothing more complicated than running it, go will download all needed dependencies

```shell
go run ./
```

## Docker

You can start the back end directly into a docker container with the following commands

```shell
docker build -t go-elivia:latest .
docker run -e "PORT=8080" -p 8080:8080 go-elivia
```

## Features

### IA

As said right before, the main part about Elia, is the IA, it was extracted from [Olivia's](https://github.com/olivia-ai/the-math-behind-a-neural-network). For all documentation about the way it works, we strongly recommend checking their repository.

### Packages

One thing we added upon it, is the **package** system, Elia is able to load, compile, and train package at run time allowing you to add, remove or edit response at run time.
To say it briefly, its the bread and butter of Elia, you can format response and trigger, call your own function, and safely extract data from the sentences that triggered your package.

**Packages** seems complicated but they aren't that much of deal, really. Plus there is a full [documentation](https://github.com/PoCFrance/e/blob/master/back/PACKAGES.md) on how to write them, which features are implemented and the best practices around them. 

### Websocket

Elia uses Websocket to communicate with the front-end, the only route used is the serve one, allowing you to handle a sentence through the IA.

## What's next ?

Elia is a PoC project, so the focus may change a bitt depending of the team taking over the project in the next months. but the main features left to implement are

* Enhance the IA - The current one works well, but with a **Text Data Vectorization** algorithm extracting meaning could have better result thus, leading to less misunderstanding error
* Make it a run time - I lied when I said it was a run time interpreter of packages. It does, indeed, compile and run every package at run time, but it doesn't reload itself yet. This is an easy update that would allow a permanently running soft without any restart needed.
* Package data - We would like to implement a tool making user's data manipulation easy inside of package, thus, you would be able to remember previous action from the user and act upon them.
* Packages, packages and again, packages - We want to add enough of them so user would be able to choose which one they want and load only the one needed when making request with a user thus having a "store" of packages from which user would build their own little IA.

As you can see, we have plenty of ideas and this is only a small grasp of what we plan to do.