package lists

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
// Wenn pos außerhalb des Bereichs der Liste liegt, wird die
// ursprüngliche Liste zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func RemoveElement(list []int, pos int) []int {
	// TODO
	ziel := []int{}
	if Empty(list) {
		return ziel
	}

	if list[0] == pos {

		return RemoveElement(ziel, pos-1)
	}

	return RemoveElement(append(ziel, list[0]), pos-1)
}
