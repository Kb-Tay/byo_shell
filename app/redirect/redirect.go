package redirect


// returns if redirect and the redirect path
func IsRedirect(args []string) (bool, string, int) {	
	if len(args) < 2 {
		return false, "", 0
	}

	for i, arg := range(args) {
		if (arg == ">" || arg == "1>") && i < len(args) - 1 {
			return true, args[i + 1], i	
		}
	}
	
	return false, "", 0
}
