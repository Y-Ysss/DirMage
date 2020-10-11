# DirMage
Directory Bookmark Manager for easy access to directories on the console.

## Get the binary
Clone the repository and build the binary file from downloaded repository.

``` bash
go build .
```

## Run executable
To run DirMage, place the configuration file and directories list file in the same directory as  the binary. These two files are created when you first start up.

``` bash
dirmage 
```

### Commandline arguments
There are three options currently being implemented.
- add
- edit
- remove

```
default (No argument)
    Switch to directory select mode
    
add -a
    Switch to directotry add mode
edit -e
    Switch to directory edit mode
remove -r
    Switch to directory remove mode
```

## Directories List
Directories list file is written in JSON format. 
Directory information should be grouped by operating systems.

directories.json
``` json
{
  "Linux": [],
  "Windows": [
    {
      "name": "C-Drive",
      "description": "",
      "path": "C:\\",
      "enabled": true
    },
    {
      "name": "UserDirectory",
      "description": "User Home",
      "path": "%USERPROFILE%",
      "enabled": false
    }
  ]
}
```

## Config Options
Config file is written in TOML format.
TOML's syntax primarily consists of `key = "value"` pairs, `[section names]`, and `# comments`.

- References : [TOML Official website](https://toml.io/)

config.toml
``` toml
[data]
dirs_file = "my-directories.json"

[selector]
page_size = 15
selector_text = "{$Name} <{$Path}>"
edit_selector_text = "{$Enabled}{$Name} <{$Path}>"
enabled_text = [" ", "*"]

[prompt]
prompt_text = "\n{32}[{$DirName}] {0}{36}{$WorkingDir} {37}{$Git}{37}\n$ "
```

### Replaceable text
|Scope|Replaceable|Means|
|:-:|:-|:-|
|Selector|`{$Name}`|Directory name |
||`{$Desc}`|Directory description|
||`{$Path}`|Directory path|
||`{$Enabled}`|Display directory info in select mode|
|Terminal|`{$DirName}`|Directory name of the selected item|
||`{$WorkingDir}`|The working directory of the terminal opened by the selected item|
||`{$Git}`|Display the Git information when if the working directory is local repository|
|Global|`{Number}`|Color escape sequences|

### Color escape sequences
`{31}` replace to `\e[31m` it means `Foreground color red`

- Reference : [ANSI Escape Code](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors)

Example
```
{31}Foreground Red
{42}Background Green
{93;42}Foreground Yellow and Background Green
```


## Attributions
### Library
- [AlecAivazis/Survey](https://github.com/AlecAivazis/survey) [MIT License]

## License
MIT License Copyright (c) 2020 Y-Ysss