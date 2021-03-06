package welcome

import "fmt"

const (
	version = "alpha-0.1"

	welcomeText = `
 ######   ##     ##    ###    ##    ##  ######    #######   ######  
##    ##  ###   ###   ## ##   ###   ## ##    ##  ##     ## ##    ## 
##        #### ####  ##   ##  ####  ## ##        ##     ## ##       
##   #### ## ### ## ##     ## ## ## ## ##   #### ##     ##  ######  
##    ##  ##     ## ######### ##  #### ##    ##  ##     ##       ## 
##    ##  ##     ## ##     ## ##   ### ##    ##  ##     ## ##    ## 
 ######   ##     ## ##     ## ##    ##  ######    #######   ######
`
)

func Welcome() {
	fmt.Print(welcomeText)
}

func Version() string {
	return version
}
