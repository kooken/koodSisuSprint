/*
Write a Go function that takes an integer amount and a slice of integers denominations.
The function should use the values in denominations to pay exactly the amount.
The denominations can be used any number of times, but higher denominations should be preferred.
The function should return the resulting denominations as an array of integers in descending order.
If the payout cannot be made, return an empty array.
Expected function signature
func Payout(amount int, denominations []int) (payout []int) {

}
Usage
amount := 128
denominations := []int{1, 2, 5, 10, 20, 50, 100, 200}
Payout(amount, denominations)
>> [100 20 5 2 1]
*/

package sprint

func Payout(amount int, denominations []int) (payout []int) {
	for i := 0; i < len(denominations)-1; i++ {
		for j := i + 1; j < len(denominations); j++ {
			if denominations[i] < denominations[j] {
				denominations[i], denominations[j] = denominations[j], denominations[i] // swapping values
			}
		}
	}

	// highest denominations first
	for _, denom := range denominations {
		for amount >= denom {
			payout = append(payout, denom)
			amount -= denom
		}
	}

	// if amount could'nt be divided into denoms
	if amount != 0 {
		return []int{} // returning empty slice
	}
	return payout
}
