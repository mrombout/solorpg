# Gimme

Gimme gives you random things.

## Usage

```
$ gimme quest
Rescue the princess from a highly secured castle.
```

```
$ gimme npc
Giant turtle
```

```
$ gimme place
Glittering carnival
```

```
$ gimme enemy
Zombie
```

```
$ gimme thing
Rusty sword
```

```
$ gimme skill
Heraldry
```

## Custom tables

Gimme makes use of tables and dice rolls in order to generate random things.

Apart from the tables that `gimme` comes with, you can also create your own by creating a generator file.
A generator file is a text file in a special format that `gimme` understands.
It contains one or more tables that `gimme` will run in sequence in order to generate a random thing.

### Table definition

A table is a finite list of things that can be generated.
It is simply a numbered list that is rolled on with the given dice.

A table is defined as follows:

```
table: (d6) "The NPC feels..." [feels]
1. "Angry"
2. "Sad"
3. "Bored"
4. "Happy"
5. "Livid"
6. "Sexy" 
```

The table definition above defines a table with 6 options, rolled on with a `d6`.
The result of what is being rolled will be stored in the variable `feels`, which we can use later to print the output.

### Template definition

The result of multiple tables to not always create a coherent sentence or thing.
Using an output template you can tweak the output to be a little more interesting and well written.
All table definitions store their result in a variable, those variables can then be used inside a template definition.
There are also a number of functions available that you can use to further tweak the result.

```

```

#### Functions

##### .An (word) / .A (word)

The `.An`/`.A` functions determine which indefinite article to use for the given word.
Both functions are the same, they are copies just to make the template read a bit better.

```
{{.An "plane"}} -> "a plane"
{{.A "plane"}} -> "a plane"
{{.An "orange"}} -> "an orange"
{{.A "orange"}} -> "an orange"
```

##### .Capitalize (word)

The `.Capitalize` function capitalizes the first letter of the given word.

### Generation definition

A generator is a collection of one or more table definitions and finally an output transformation.
The `gimme` tool will simple read each table, roll on it and remember the result to construct a random thing later.

```
d3 The NPC is a... [gender]
1. male
2. female
3. unknown
d4 The NPC's occupation is... [occupation]
1. knight
2. peasant
3. blacksmith
>>>
{{.Capitalize .An gender}} {{occupation}}
```

The generator above defines two tables to determine the NPC's `gender` and `occupation` using a `d3` for the first, and a `d4` for the second.
Finally, after the `>>>` and output template is defined to turn the information that was being rolled into one thing.

Say the first table rolls a `2` (Female) and the second tables rolls a `1` (Knight), the result of the template would then be:

```
A female knight
```
