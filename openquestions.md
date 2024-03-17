# Strings

Ein string kann nicht nil sein. Was sagt das über string aus?
Ist ein string ein referenztyp oder ein valuetyp?
Ist ein string immutable?

# Types

Was genau ist der Unterschied zwischen

`type a int32` und `type a = int32`

# Functions vs Methods

Go differentiates between functions and methods.

- [ ] be able to explain the difference

A (versuch, unsicher): Eine Methode ist quasi die extension methode auf einen datentyp???

func (x MyDataType) Foo() {} // Methode
func Foo() {} // Function

Aus Sicht des Aufrufers:

http.Post() // sieht aus wie eine Methode ist es aber nicht weil Post eine `func Post()` ist und keine func (http Http) Post() ist.

Hier ist http kein Datentyp sondern ein Package und die Function Post ist zugaenglich weil sie mit einem Grossbuchstaben beginnt.

**Const**

Weird or no? We have a const but we set some value dynamically... WTF???

How does it work under the hood?

```
const apiUrl = "https://cex.io/api/ticker/%s/EUR"
```
