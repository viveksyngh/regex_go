## regex_go - OpenFAAS function to check Regex
[![OpenFaaS](https://img.shields.io/badge/openfaas-serverless-blue.svg)](https://www.openfaas.com)
[![Status](http://7d86c827.ngrok.io/function/of_badge_gen?user=s8sg&repo=regex_go&branch=master)](http://0341c281.ngrok.io)

This is build as a dummy app to check openfaas-cloud and Status integration with openfaas deployed at: 
`0341c281.ngrok.io` and can be reached at `http://0341c281.ngrok.io/function/s8sg-regex_go`
   
  
**Example:**
```
$ curl -H "Content-Type: application/json" -X POST http://0341c281.ngrok.io/function/s8sg-regex_go -d '{"data": "The moon is our natural satellite, i.e. it rotates around the Earth!", "regex": "(\\b[^\\s]+\\b)" }'
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

$ curl -H "Content-Type: application/json" -X POST http://0341c281.ngrok.io/function/s8sg-regex_py -d '{"data": "The moon is our natural satellite, i.e. it rotates around the Earth!", "regex": "(\\b[^\\s]+\\b)" }'
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
