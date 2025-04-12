package openf1go

type Arg struct {
	Key   string
	Value string
}

func ArgBuilder(args []Arg) string {
	var url string

	if len(args) == 0 {
		return url
	}

	if len(args) > 0 {
		url += "?"
	}

	for i, arg := range args {
		if i != 0 {
			url += "&"
		}

		url += arg.Key + "=" + arg.Value
	}

	return url
}
