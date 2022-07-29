package main

import (
	"fmt"
	"sort"

	sort57 "57/the-go-programming-language/ch7/sort/lib"
)

func main() {
	fmt.Println("original order")
	sort57.PrintTracks(sort57.Tracks)

	fmt.Println("\nsort by artist")
	sort.Sort(sort57.ByArtist(sort57.Tracks))
	sort57.PrintTracks(sort57.Tracks)

	fmt.Println("\nsort by artist in reverse order")
	sort.Sort(sort.Reverse(sort57.ByArtist(sort57.Tracks)))
	sort57.PrintTracks(sort57.Tracks)

	fmt.Println("\nsort by year")
	sort.Sort(sort57.ByYear(sort57.Tracks))
	sort57.PrintTracks(sort57.Tracks)

	fmt.Println("\nsort by title, then year, and lastly length")
	sort.Sort(sort57.CustomSort{
		T: sort57.Tracks,
		LessCustom: func(x, y *sort57.Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			} else if x.Year != y.Year {
				return x.Year < y.Year
			} else if x.Length != y.Length {
				return x.Length < y.Length
			}
			return false
		}})
	sort57.PrintTracks(sort57.Tracks)
}
