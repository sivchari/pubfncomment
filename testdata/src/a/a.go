package a

func ok() {}

// OK is a function
func OK() {}

func NG() {} // want "public function \\(NG\\) should have comment"
