# Structure Me

Structure me is a GoLang command line tool to help bootstrap your usual folder structure. It can be used for:
  - bootstrapping your usual folder structure for your programming workflow.
  - create a new folder with the proper structure to keep your files organized.


### TODO:
  - Implement an API to store and retrieve custom templates.
  - Organize current folder into the JSON template proper structure.

### Installation
To install, you have to have GoLang install and then run the following command:

```sh
$ cd structureMe
$ go build
```

### How To Use

First create a custom JSON file. `structure-template.json` is the default name but you can choose any name you want.
Then, to bootstrap your folder structure type:

```sh
$ .\structureMe.exe init -template=your-template-name.json
```

If you are not using a Windows machine, you can type `structureMe` directly provided you are in the same CWD.
This will create the folders for you, and you can start your work!
