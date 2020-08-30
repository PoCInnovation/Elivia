### e

## Description

E is a vocal assistant designed to be respectful of user privacy over theirs data.
The goal of this project is to be seen incorporated in /e/ os, a product from the [e foundation](https://e.foundation/), do to so it is supervised by [Gaël Duval](https://fr.wikipedia.org/wiki/Ga%C3%ABl_Duval).

E Vocal Assistant is composed of 3 parts, a front end built in Kotlin using gradle, a back-end generously inspired from [Olivia](https://github.com/olivia-ai/olivia)'s design, packaging an AI, and a CLI tool made for testing purpose.
Each part have it's own `README.md` where all the information relative to the said part will be detailed and explained more in depth.

 ## Installation

```shell
git clone https://github.com/PoCFrance/e
```



// todo luca -- build with android studio

## Quick Start

As a developer or for testing, you can simply start the back end in local and request on it.
-- for dev purpose it's the basic option since IP isn't configurable yet
first start the back-end

```shell
cd back
go run ./
```

Then lunch android studio in the `front` folder, and build the app
You can start the application once build is completed.

// todo luca -- update ip address (conf / raw code update)

## Maintainers

* [Theo Ardouin](https://github.com/Qwexta)
* [Luca Georges Francois](https://github.com/PixelFault-tech)

## Source

As said in the description, e is inspired from [Olivia](https://github.com/olivia-ai/olivia), yet it also uses a part of its code.
The project was built upon Olivia's IA and a fork from it's Back-end. Currently, the back has been totally updated to meet the requirement we had over flexibility and mutability, but we still use the [IA]() for the moment.

## License

This project is under MIT licence
