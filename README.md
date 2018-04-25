# regex_go
OpenFAAS function to check Regex (implemented in Go)

This is build as a dummy app to check openfaas-cloud and Status integration with openfaas deployed at: 
`0341c281.ngrok.io` and can be reached at `http://0341c281.ngrok.io/function/regex_go`
  
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
