/*
While if or if-else statements are very useful, sometimes they are a little uncomfortable.
switch statement is very useful if one needs to perform more checks for a value or if there are multiple cases for one check.
In this task you need to make a function that takes a string, that can contain a month name.
If a month name is given, the season has to be returned. Otherwise return "invalid input: " with the input appended to it.
Months:
winter: jan, feb, dec
spring: mar, apr, may
summer: jun, jul, aug
autumn: sep, oct, nov
*/

package sprint

func Season(month string) string {
	switch month {
	case "jan", "feb", "dec":
		return "winter"
	case "mar", "apr", "may":
		return "spring"
	case "jun", "jul", "aug":
		return "summer"
	case "sep", "oct", "nov":
		return "autumn"
	default:
		return "invalid input: " + month
	}
}
