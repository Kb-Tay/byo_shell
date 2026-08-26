package redirect

func IsRedirect(args []string) bool {	
	if len(args) == 3 && args[1] == ">" {
		return true	
	}
	
	return false
}
