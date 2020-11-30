package convpack

import "fmt"

type Feets float64
type Meters float64
type Pounds float64
type Kilograms float64

func (f Feets) String() string { return fmt.Sprintf("%gft", f) }
func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (p Pounds) String() string { return fmt.Sprintf("%gp", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%gkg", k) }