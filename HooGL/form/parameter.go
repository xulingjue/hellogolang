package form

func GetInt(req *http.Request,name string,def int) int{
	value := req.FormValue("name")
	result, err := strconv.Atoi(value)
	if err != nil{
		return def
	}
	return result;
}

func GetString(req *http.Request,name string,def string)  string{
	value := req.FormValue("name")
	if value == ""{
		return def
	}
	return value
}

func GetInt64(req *http.Request,name string,def int64)  int64{
	value := req.FormValue(name)
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil{
		return def
	}
	return result;
}