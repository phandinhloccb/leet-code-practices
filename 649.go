package main

import "fmt"

func predictPartyVictory(senate string) string {
    n := len(senate)
    qR := []int{}
    qD := []int{}

    for key, value := range senate {
        if value == 'R' {
            qR = append(qR, key)
        } else {
            qD = append(qD, key)
        }
    }    

    for len(qR) > 0 && len(qD) > 0 {
        r := qR[0]
        d := qD[0]
        qR = qR[1:]
        qD = qD[1:]

        if r < d {
            qR = append(qR, r + n)
        } else {
            qD = append(qD, d + n)
        }
    }

    if len(qR) > 0 {
        return "Radiant"
    }
    return "Dire"
}

func main(){
	fmt.Println(predictPartyVictory("RD"))
}