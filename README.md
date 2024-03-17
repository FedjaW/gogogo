# gogogo

gooooooooooooooo

**Init a go module**

`go mod init <name>`

**Run (it will search for the entrypoint main func in all files in the module)**

`go run .`

**Structure of go**
Workspaces -> Modules -> Packaged
Workspace are optional

**Difference between constants and immutable variables**
immutable variable is a variable so it has space in memory, memory that needs to be allocated. Immutable by a mechanism that is checking that no one is changing that variable.
In js a const is a immutable variable.

A constant will tell the compiler to copy the given initial value and paste it. (very simplified)

**read the doc**
`go doc <package> <method>`

**BTW** .... `this` is not a keyword in go. I am in heaven, beautiful, I LOVE IT!

**FormatOnSave in VSC**

"[go]": {
"editor.insertSpaces": true,
"editor.formatOnSave": true,
"editor.defaultFormatter": "golang.go"
},

**Goroutines and Channels**

To avoid deadlocks you have to close the channels before ending the program with `close(chan)`

**FOR THE MAC USERS**

USE LULU! It is a firewall that will protect you.
https://objective-see.org/products/lulu.html

**Biray**

Go is a binary, running a go app `go run .` will build a new app everytime in a temp directory.
From the os pov it is always a new app.
So if this app will make a call to the network lulu will pop up and ask for permission. Even if you allow it, next time it will aks again because it is a new app.... hahaha oh boy, isnt't dev sexy
