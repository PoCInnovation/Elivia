### Elivia

## Description

Elivia is a vocal assistant designed to be respectful of user privacy over theirs data.
The goal of this project is to be seen incorporated in /e/ os, a product from the [e foundation](https://e.foundation/), do to so it is supervised by [Gaël Duval](https://fr.wikipedia.org/wiki/Ga%C3%ABl_Duval).

E Vocal Assistant is composed of 3 parts, a front end built in Kotlin using gradle, a back-end generously inspired from [Olivia](https://github.com/olivia-ai/olivia)'s design, packaging an AI, and a CLI tool made for testing purpose.
Each part have it's own `README.md` where all the information relative to the said part will be detailed and explained more in depth.

 ## Installation

```shell
git clone git@github.com:PoCFrance/Elivia.git
```

You need [Android Studio](https://developer.android.com/studio) to build the frontend application. When you'll open the project, it will automatically sync all dependencies. 

## Quick Start

As a developer or for testing, you can simply start the back end in local and request on it.
-- for dev purpose it's the basic option since IP isn't configurable yet
first start the back-end

- **Backend**
```shell
cd back
go run ./
```
More information [here](https://github.com/PoCFrance/e/blob/master/back/README.md)

- **Frontend**

In the first place, you need to open the `MainActivity` file located in `app/java/com.poc.elivia/` and set your local IP address at line 26.

To build the frontend application, open the front/Elivia folder in Android Studio.
In the `Build` tab, select `Build Bundle(s) / APK(s)` and then `Build APK(s)`.

Once the build has finished, you can copy the apk to your smartphone and install it by opening it in your file explorer.

More information [here](https://github.com/PoCFrance/e/blob/master/front/README.md)

## Maintainers

* [Theo Ardouin](https://github.com/Qwexta)
* [Luca Georges Francois](https://github.com/PixelFault-tech)

## Source

As said in the description, e is inspired from [Olivia](https://github.com/olivia-ai/olivia), yet it also uses a part of its code.
The project was built upon Olivia's IA and a fork from it's Back-end. Currently, the back has been totally updated to meet the requirement we had over flexibility and mutability, but we still use the [IA]() for the moment.

## License

This project is under MIT licence
