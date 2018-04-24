# regex_go
OpenFAAS function to check Regex (implemented in Go)
  
**Example:**
```
$ curl -H "Content-Type: application/json" -X POST http://0341c281.ngrok.io/function/regex_go -d '{"data": "The moon is our natural satellite, i.e. it rotates around the Earth!", "regex": "(\\b[^\\s]+\\b)" }'
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
