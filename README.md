# regex_go
OpenFAAS function to check Regex (implemented in Go)
  
**Example:**
```
I/P:
{
"data": "The moon is our natural satellite, i.e. it rotates around the Earth!",
"regex": "(\\b[^\\s]+\\b)"
}


O/P:
{
  "match": true,
  "matches": [
    "The",
    "moon",
    "is",
    "our",
    "natural",
    "satellite",
    "i.e",
    "it",
    "rotates",
    "around",
    "the",
    "Earth"
  ]
}

```
