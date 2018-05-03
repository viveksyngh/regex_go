import re
import json

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    raw = json.load(req)
    pattern = re.compile(raw["regex"])
    data = raw["data"]
    result = {}
    result["match"] = False
    result["matches"] = re.findall(patter, string)
    if len(result["matches"]) > 0 {
        result["match"] = True
    }
    return json.dumps(result)
