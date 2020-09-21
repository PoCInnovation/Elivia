# Elivia - Packages

**Packages** are the **Elivia**'s addon they will make your life easier when contributing to thee application and are totally independent from the rest of **Elivia**
Each **Package** holds several **Modules** and each module have its own trigger and response.
In the future, modules from the same package will share the same data.

## Format

Package all have the same architecture tree
first you have to create a folder ***package_name*** inside the *package* folder
inside this, you will have

* Your go file and architecture that **must** build.
* A *res* folder

The res folder will hold a *locales* folder itself containing a folder for each of the languages (we are currently supporting English -tagged **en**- and French -tagged **fr**- )

each of these language folder will hold a **response.json** and a **triggers.json**

## Code Architecture

For the moment only Go is supported, we may add a support for all language compiling to ELF.
Note that you can update the [loader](https://github.com/PoCInnovation/Elivia/blob/master/back/plugins/package.go) function to change the way the program load functions from external libraries

Each Module will need an entry point, it been a function with this exact prototype

```go
func Module(sentence string, entries map[string]string) (string, map[string]interface{})
```

In the **sentence** variable you will find the request that triggered your module, the **entry** will be explained inside the **Triggers** part
The Return type has two things.
First : the **Tag**, it will choose a random response from the one defined in the `response.json`

second : the **Map** you will be able to feed all the information you want to transfer to you front end, there is only one type of data you have to feed, and we will explain it in the **Response** part

## Triggers

This is where the majority of the configuration is. it define when and how the Module will be called. as well as the module itself
it is recommended to check the [sms](https://github.com/PoCInnovation/Elivia/blob/master/back/package/sms/res/locales/en/triggers.json) package as it serve as example.
In this json, you will have an Array of Modules.
Each module is defined by an **Array** of **Module Type** thus, a module can be triggered by multiple ways, and have various catchphrase without having issues to define them all.

A **Module Type** has these fields :

* **entries** - **Entry Type Array**  - A possibly null field containing all the resources for extracting meaning from the sentence (i.e. text to send and contact inside the sms module) **entries** can be found inside the module's function when call, in the form of a map.
* **pattern** - **String Array** - Key word needed to match this module
* **olivia-feed** - **String Array** - Real life full sentence serving as example for the current IA, will be deprecated when the IA will be reworked

We will also define the **Entry Type** straight away:

* **name** - **String** - The name of the entry (the one you will have to query in the map)
* **parser** - **String** - There is 3 parsers existing, **after**, **before**, and **between ** depending on which type of data you want to retrieve.
* **resources** - Custom Object - depends on the parser type
  * The after and before will need the **key** - **String** - and the **x** - **Int** - field, the first being the the match from which you extract and the second been the number of occurrences you skip before counting it a match (note you can have negative number, thus `"x": -1` will lead to seek the last occurrences)
  * The between parser will need two - **String** -, **after** and **before** and two - **Int** -, **x** and **y**, and will get all the text matching between the two field
    It works as if you had a union between the after parser with the **after** and **x** value, and a before parser with the **before** and **y** value.

To summarize it all, there is a little example of a valid triggers.json file 

```json
[
    {
		"FunctionName/Module": [
			{
                "entries":[
                    {
                        "name":"",
                        "parser":"",
                        "ressources":{}
                    }
                ],
			   "pattern":[
                    "-"
                ],
                "olivia-feed":[
                    "- - - -",
                    "- - - - -"
                ]
			}
		]
	}
]
```

## Response

`Response.json` is pretty straight forward, you have a list of Tag - **String** -, associated with message - **String Array** -
When you return from your Go function, you have to specify a Tag, it will serve to pick one of the messages array, and will pick a answer at random

Messages aren't totally basic string, they can be formatted with data from your code.
To do so, you have to precise which part have to be filled with  this marking `%`  before and after  (i.e. : `hello, %variable%` ). 

remember in the code part when we said we would come back on the map, this is precisely the other use case we have of it
When returning with a tag implementing formatted variable, you will have to add a field in the map with the text to replace with.
for example, the  `hello, %variable%` will need you to return something like this

```go
return "success", map[string]interface{}{
		"variable": "Elivia",
	}
```

The Front End will then receive `hello, Elivia`.