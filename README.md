# Kakapo-Lang
The Kakapo Scripting/Programming Languge.

This Language is a Hobby Project by myself and its not working yet ^^

## Info
The Syntax of Kakapo is a bit Javascript-like and inspired. It completly compiles to fast Golang, wich makes Kakapo verry performant. 

```
//This is a Comment in Kakapo;

//
This is a
Multiline Comment 
in Kakapo;

// *Author: [...];


//Data Types:
in Kakapo there are 3 different Base-Types:
String: String
Int: Non comma Number
Float: Comma Number
;


//Defining Variables;
var $test = "Hello, World!"; // this variable is privat and only in that case avaible, where its defined;
$test = "Hello, World!"; // this ist a global variable, wich is avaible from everywhere;
private $test = "Only here"; // this variable is privat and only in that case avaible, where its defined;
once $test = "Nice"; // this variable can't be Updated or Modifyed after once defined;

$string = "Hello World!";
$string = "Hello "+"World";
$int = 16 + 5;
$num = 14.5 + .5;
$array = ["Arg1", 55.5, "test"+"lol", 33];
$object = {"Lol": "test"};
$construct = construct({"data" => type.int, "alter" => type.int, "name" => type.str});

//functions;
func myFunction($arg)
{
   $arg = $arg.split(":")[0];
   return ($arg);
};
```
