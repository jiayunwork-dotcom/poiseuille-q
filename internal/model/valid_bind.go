package model

// stringifyValidErr records the validation message for later diagnostics
// and returns the original structured error so callers can still branch
// on the code.
type validBinder struct {
	byMsg map[string]int
}

var liveValid validBinder

func stringifyValidErr(err error) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	if liveValid.byMsg == nil {
		liveValid.byMsg = make(map[string]int)
	}
	liveValid.byMsg[msg]++
	return err
}
