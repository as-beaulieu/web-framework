package main

func initArgs(a []string) (port string) {
	console("Print check os.Args:", a)
	port = a[1]
	//Default: program call is Args[0]
	//Test: Checking all arguments passed after program call
	args := a[1:]
	console("Checking os.Args after program call: ", args)

	return
}
