package main

// func handlerFunc(f func(http.ResponseWriter, *http.Request)) (string, error) {
// 	s := httptest.NewServer(http.HandlerFunc(f))
// 	defer s.Close()
// 	res, err := http.Get(s.URL)

// 	if err != nil {
// 		return "", fmt.Errorf("err %s", err)
// 	}

// 	rbody, err := ioutil.ReadAll(res.Body)
// 	defer res.Body.Close()

// 	if err != nil {
// 		return "", fmt.Errorf("err %s", err)
// 	}
// 	if res.StatusCode != 200 {
// 		return "", fmt.Errorf("a response code is not 200")
// 	}

// 	return string(rbody), nil
// }

// func TestRoot(t *testing.T) {
// 	root, err := handlerFunc(RootHandler)

// 	if err != nil {
// 		t.Errorf("err: %v", err)
// 	}

// 	if string(root) != "Hello world! Welcome Go App!" {
// 		t.Errorf("a response is not Hello world! Welcome Go App!: %v", root)
// 	}
// }

// func TestHealthCheck(t *testing.T)  {
// 	healthCheck, err := handlerFunc(HealthCheck)

// 	if err != nil {
// 		t.Errorf("err: %v", err)
// 	}

// 	if string(healthCheck) != "ok" {
// 		t.Errorf("a response is ok: %v", string(healthCheck))
// 	}
// }