# gogogo

gooooooooooooooo

**Init a go module**

`go mod init <name>`

**Run (it will search for the entrypoint main func in all files in the module)**

`go run .`

**Build**

Will build for your operating system that you are on

`go build .`

Will build for other operating system

`env GOOS=target-OS`\
`env GOARCH=target-architecture go build. `

**Install**

will install your go app to be used from everywhere

`go install .`

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

# Dear future me - Weather

Today's date is `March 17th, 2024`. You find yourself at home on a serene Sunday, engrossed in programming with `Go`. Following an engaging course spanning approximately 8 hours. Despite the weather forecast predicting otherwise, the sun graces the day with its warm presence.

According to wetter.com:

Wetter\
Sonntag, 14:00 Uhr\
Stark bewölkt


**Candy**

If you're transferring `less than 1500 bytes` over the network (http), it's all the same, right? That's called the `TCP packet`. That's the `MTU`,  the `maximum transmission unit`. So, whether you're transferring `one byte or 1499, it's pretty much all the same`.
