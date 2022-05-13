package main

import "fmt"

func main(){
	fmt.Println("Hello Worlds")
}

/* How do we run the code in our project? 
	1) $ go build   - compiles a bunch of go source code files
	2) $ go run     - compiles and executes one or two files
	3) $ go fmt     - formats all the code in each file in the current directory
	4) $ go install - compiles and installs a package
	5) $ go get     - Downloads the raw source code of someone else's package
	6) $ go test    - runs any test associated with the current project
*/


/* What does 'package main' mean? 

	Package == Project == Workspace

	[!] When we see the word package and go we can thing of a package as being like a project or a workspace.
	[!] A package is a collection of common source code files.
	[!] A package can have many related files inside of it each file ending with a file extension of go


	! REQUIREMENT !

		The only requirement for every file inside of a package is at the very first line of each file must
		declare the package that it belongs to.

	
	Two Types of Packages: Executable & Reusable

		[!] When we specify the words package main, it means that we make an executable type package

		Executable -  	Generates a file that we can run

			package main -	Defines a package that can be compiled and then *executed*. Must havea func called 'main'
		
		
		Reusable   -  	Code used as 'helpers'. Good place to put reusable logic. we can think of that as being like code dependencies
						or libraries.

			package calculator -	Defines a package that can be used as a dependency (helper code)

			package uploader   -	Defines a package that can be used as a dependency (helper code)
		
*/


/* What does 'import "fmt" mean? 
	 
	[!] By default main package has no access to other libraries or other packages, we have to specifically use the import
	 statement to form a link from our package to the other ones

	 [!] fmt package is a part of the standart library of Go.
	 
	 https://golang.org/pkg

*/