package function

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Input struct {
	Data  string `json:"data"`
	Regex string `json:"regex"`
}

type Output struct {
	Match   bool     `json:"match"`
	Matches []string `json:"matches"`
}

const help = `i/p data format:
{
	"data": "my Test text with UPPERCASE",
	"regex": "[A-Z]"
}`

// Handle a serverless request
func Handle(req []byte) string {

	var ip Input

	err := json.Unmarshal(req, &ip)
	if err != nil {
		return fmt.Sprintf("Error: Failed to parse input : %v \n %s", err, help)
	}

	r, rerr := regexp.Compile(ip.Regex)
	if rerr != nil {
		return fmt.Sprintf("Error: Failed to compile regex %s : %v", ip.Regex, rerr)
	}

	var op Output
	op.Match = true
	op.Matches = r.FindAllString(ip.Data, -1)
	if len(op.Matches) == 0 {
		op.Match = false
	}

	mdata, merr := json.Marshal(op)
	if merr != nil {
		return fmt.Sprintf("Error: Failed to marshal output %v : %v", op, merr)
	}
	return string(mdata)
}
