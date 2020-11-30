package convpack

func FToM(f Feets) Meters { return Meters(f / 3.2808) }
func MToF(m Meters) Feets { return Feets(m * 3.2808) }

func PToK(p Pounds) Kilograms { return Kilograms(p / 2.2046) }
func KToP(k Kilograms) Pounds { return Pounds(k * 2.2046) }