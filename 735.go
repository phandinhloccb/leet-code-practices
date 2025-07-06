package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0, len(asteroids))

    for _, a := range asteroids {
        destroyed := false 

        for len(stack) > 0 && a < 0 && stack[len(stack) - 1] > 0 {
            top := stack[len(stack) - 1]

            if top < -a {
                stack = stack[:len(stack)-1]
                
            } else if top == -a {
                stack = stack[:len(stack) - 1]
                destroyed = true
                break

            } else {
                destroyed = true
                break
            }

        }

        if !destroyed {
            stack = append(stack, a)
        }
    }

	return stack
}

func main() {
	fmt.Println(asteroidCollision([]int{5,10,-15}))
}


/*
We are given an array asteroids of integers representing asteroids in a row. The indices of the asteriod in the array represent their relative position in space.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

 

Example 1:

Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
Example 2:

Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.
Example 3:

Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

*/