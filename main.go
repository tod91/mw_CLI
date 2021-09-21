package main

//
//import "flag"
//
//var evalInput = flag.String("eval", "", "evaluate math expression")
//	var validateInput = flag.Bool("val", false, "evaluate validate expression")
//	var errorsInput = flag.Bool("errors", false, "evaluate error expression")
//	flag.Parse()
//
//	//resp, err := http.Post("http://localhost:8080/evaluate")
//	//if err != nil {
//	//	log.Fatalln(err)
//	//}
//
//	//fmt.Println(*evalInput, *validateInput, *errorsInput)
//	////We Read the response body on the line below.
//	////body, err := ioutil.ReadAll(resp.Body)
//	//if err != nil {
//	//	log.Fatalln(err)
//	//}
//	//var posts []*Post
//	//err = json.Unmarshal(body, &posts)
//	//if err != nil {
//	//	panic(err.Error())
//	//}
//	//for _, p := range posts {
//	//	fmt.Printf("post %d has content '%s'\n", p.Id, p.Title)
//	//}
//
//}
//
//func isFlagPassed(name string) bool {
//	found := false
//	flag.Visit(func(f *flag.Flag) {
//		if f.Name == name {
//			found = true
//		}
//	})
//	return found
//}
