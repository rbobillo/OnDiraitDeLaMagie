package internal

import "fmt"

var (
	/*ErrWizardsNotFounds is use when trying to find a Wizard in database
	 with is id but the id doesn't exist or is not found */
	ErrWizardsNotFounds    = fmt.Errorf("wizard not found")
)
