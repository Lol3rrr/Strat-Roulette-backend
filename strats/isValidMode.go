package strats

func isValidMode(data GameMode) bool {
	if data == Bomb {
		return true
	}
	if data == Hostage {
		return true
	}
	if data == SecureArea {
		return true
	}

	return false
}
