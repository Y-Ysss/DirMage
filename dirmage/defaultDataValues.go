package dirmage

const defaultConfigTomlData string = `# DirMage Config
# https://github.com/Y-Ysss/DirMage

# TOML's syntax primarily consists of
#     key = "value" pairs, [section names], and # comments.
# TOML Official website https://toml.io/

[data]
dirs_file = "directories.json"

[selector]
page_size = 15
selector_text = "{$Name} <{$Path}>"
edit_selector_text = "{$Enabled}{$Name} <{$Path}>"
enabled_text = [" ", "*"]

[prompt]
prompt_text = "\n{32}[{$DirName}] {0}{36}{$WorkingDir} {37}{$Git}{37}\n$ "
`

const defaultDirectoriesJsonData string = `{
  "linux": [
    {
      "name": "Root",
      "description": "Root directory",
      "path": "/",
      "enabled": true
    },
    {
      "name": "UserDirectory",
      "description": "",
      "path": "$HOME",
      "enabled": false
    }
  ],
  "windows": [
    {
      "name": "C-Drive",
      "description": "",
      "path": "C:\\",
      "enabled": true
    },
    {
      "name": "UserDirectory",
      "description": "",
      "path": "%USERPROFILE%",
      "enabled": true
    }
  ]
}`
