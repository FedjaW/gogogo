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
