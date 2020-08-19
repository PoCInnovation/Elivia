# Elivia Plugins

## format

![Screenshot from 2020-08-04 15-36-29](/home/qwexta/Pictures/Screenshot from 2020-08-04 15-36-29.png)

All the plugins are located in the "package" directory.
They all have the same architecture :

* `.go` file making the code of the module
* a resource folder `res` that holds the triggers and response of all languages

### triggers.json

holds a json array of handlers
those handlers are relative to a specific function of the plugins although multiple handlers with different pattern can be link to the same function
the object is defined as follow

```json
{
    "entries": [] // optional will be explain more in depth after
	"pattern": [] // the format of the sentence you want to match
	"olivia-feed": [] // temporary and mendary, helps the IA to recognize your sentence from various exemple
	"callback": "" // the function you want to be called then this pattern is uncountered
}
```



the "pattern", "feed", and "callback" may seems straight forward, yet, the entries needs a bit more explainations


### Entries

Entry helps to retreive element from the string you match, note that you should leave a blank in the pattern if you want to catch something there.

entries are composed of 3 elements

```json
{
    "name":""
	"parser":""
	"resources": {}
}
```

the name is string you will use to retreive that entry within your code.
the parser is the function that will be called to find out your entry, 3 of them are available nowadays :

* after : gets the rest of the string after the match
  it uses a specific ressourse to define the match requieres

  ```json
  "ressources":{
      "key":"" // the needle you want to find, will not be included in the selection
  	"x":0 	// the xth occurence of the key before considering it a match
  			// note that you can set it with negative value to get the x nth last 				// element
  }
  ```

* before : works the same as after but gets everything before, match not included

* between : gets everything between the 2 matchs
  it also define a resource format to match its need

  ```json
  "ressources":{
  	"after": "", // The key for the starting match
      "x": 0,		 // the x for the starting match

      "before": "",// the key for the starting match
      "y": 0		 // the x for the starting match
  }
  ```

  ​