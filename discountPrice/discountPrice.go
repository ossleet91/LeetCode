package main

// Final prices with a special discount in a shop
//
// Given the array prices where prices[i] is the price of the ith item
// in a shop. There is a special discount for items in the shop, if you
// buy the ith item, then you will receive a discount equivalent to
// prices[j] where j is the minimum index such that j > i and prices[j]
// <= prices[i], otherwise, you will not receive any discount at all.
//
// Return an array where the ith element is the final price you will pay
// for the ith item of the shop considering the special discount.
//
// Example:

//	Input: prices = [8,4,6,2,3]
//	Output: [4,2,4,2,3]
//	Explanation:
//		- For item 0 with price[0]=8 you will receive a discount
//		  equivalent to prices[1]=4, therefore, the final price
//		  you will pay is 8 - 4 = 4.
//		- For item 1 with price[1]=4 you will receive a discount
//		  equivalent to prices[3]=2, therefore, the final price
//		  you will pay is 4 - 2 = 2.
//		- For item 2 with price[2]=6 you will receive a discount
//		  equivalent to prices[3]=2, therefore, the final price
//		  you will pay is 6 - 2 = 4.
//		- For items 3 and 4 you will not receive any discount at
//		  all.
//
// Constraints:
//
//	- 1 <= prices.length <= 500
//	1 <= prices[i] <= 10^3
//
// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/description/

import "fmt"

func discountedPrice(price int, discounts []int) int {
	if len(discounts) == 0 {
		return price
	}

	for _, d := range discounts {
		if price >= d {
			return price - d
		}
	}

	return price
}

func discountPriceSlow(prices []int) []int {
	for i, p := range prices {
		prices[i] = discountedPrice(p, prices[i+1:])
	}

	return prices
}

func discountPrice(prices []int) []int {
	var stack []int

	for i, p := range prices {
		for len(stack) != 0 && prices[stack[len(stack)-1]] >= p {
			prices[stack[len(stack)-1]] -= prices[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return prices
}

func main() {
	fmt.Println("vim-go")
}
