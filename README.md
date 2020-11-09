### Elivia


<p align="center">
  <img src="https://1.bp.blogspot.com/-U9xZXxJL0jM/XVflXA-y7rI/AAAAAAAAmfE/Sl3U6tDPetg2TwBgce39GqxI_n7d0bRBwCLcBGAs/s1600/serveimage.png" width=30%"/>
  <img src="../.github/poclogo.jpg" width=30%"/>
</p>

Elivia is an open-source project resulting from a partnership between the [PoC R&D center](https://github.com/PoCInnovation) and the [/e/ Foundation](https://e.foundation/),
along with the supervision of [Gaël Duval](https://fr.wikipedia.org/wiki/Ga%C3%ABl_Duval).

The goal of this project is to create a personal assistant that respects the privacy of its users, that can run on the /e/ operating system.

Elivia is composed of two parts :

- A frontend mobile application. It is currently written in Kotlin and can be built with Android Studio.

- A backend application. It is currently written in Go and is built on top of the [Olivia](https://github.com/olivia-ai/olivia) project. It contains an AI and a CLI tool made for testing purposes.

You can learn more about each one of these two parts in the corresponding folder located at thr root of this repository.

## Installation

#### Requirements

- **Frontend**

    You need [Android Studio](https://developer.android.com/studio) to build the frontend application. When you'll open the project, it will automatically sync all required dependencies.
    As long as you can build the application targetting the Android 4.0 to 9.0 platform, it will work.

    Learn more about the frontend [here](https://github.com/PoCInnovation/Elivia/blob/master/front/README.md).

- **Backend**

    You need to have go or Docker installed to run the Olivia based AI & server.
    Learn more about the backend [here](https://github.com/PoCInnovation/Elivia/blob/master/back/README.md).

 ## Installation

First you have to clone the repository.

```shell
git clone git@github.com:PoCInnovation/Elivia.git
```

## Quick Start


- **Backend**

You can run the backend on your local environment and make request to the server. Please note that the target IP is not configurable yet.

```shell
cd back
go run ./
```

- **Frontend**

In the first place, you need to open the `MainActivity` file located in `app/java/com.poc.elivia/` and set your server local IP address at line 26.

To build the frontend application, open the front/Elivia folder in Android Studio.
In the `Build` tab, select `Build Bundle(s) / APK(s)` and then `Build APK(s)`.

Once the build has finished, you can copy the apk to your smartphone and install it by opening it in your file explorer.

## Authors

* [Theo Ardouin](https://github.com/Qwexta)
* [Luca Georges Francois](https://github.com/PtitLuca)

## Source

As mentionned in the project description, Elivia backend is built on top of the [Olivia](https://github.com/olivia-ai/olivia) project and more precisely the backend.
Because it is still in development, we may update the current version of the AI.

## License

This project is under MIT licence
