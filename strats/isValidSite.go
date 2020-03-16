package strats

func isValidSite(data Site) bool {
	if data == Attacker {
		return true
	}
	if data == Defender {
		return true
	}

	return false
}
