package main

func main() {

}

// Exit with status code 1 in case of failure
func must(err error) {
	if err != nil {
		panic(err)
	}
}
