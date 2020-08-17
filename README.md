# Telegraf execd rand input example

This is a random number generator input plugin for Telegraf, meant to be compiled separately and used externally with telegraf's execd input plugin. It's an example of using the plugin shim for Go programs.

# Install Instructions

Download the repo somewhere

    $ git clone git@github.com:ssoroka/rand.git

build the "rand" binary

    $ go build -o rand cmd/main.go
    
 (if you're using windows, you'll want to give it an .exe extension)
 
    go build -o rand.exe cmd/main.go

You should be able to call this from telegraf now using execd:

```
[[inputs.execd]]
  command = ["/path/to/rand_binary"]
  signal = "none"
  
# sample output: write metrics to stdout
[[outputs.file]]
  files = ["stdout"]
```
