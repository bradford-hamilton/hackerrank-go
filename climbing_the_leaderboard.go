package main

// https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem
// scores: an array of integers that represent leaderboard scores
// alice: an array of integers that represent Alice's scores
// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	places := []scoreAndPlace{{score: scores[0], place: 1}}

	for i := 1; i < len(scores); i++ {
		if scores[i] < scores[i-1] {
			places = append(places, scoreAndPlace{score: scores[i], place: places[i-1].place + 1})
		} else {
			places = append(places, scoreAndPlace{score: scores[i], place: places[i-1].place})
		}
	}

	// scores = [100, 100, 50, 40, 40, 20, 10]
	// places = [{100 1} {100 1} {50 2} {40 3} {40 3} {20 4} {10 5}]
	// alices = [5 25 50 120]
	alicesPlaces := make([]int32, len(alice))

	for i := 0; i < len(alice); i++ {
		for j := 0; j < len(places); j++ {
			if j == len(places)-1 {
				if alice[i] < places[j].score {
					alicesPlaces[i] = places[j].place + 1
					break
				} else if alice[i] >= places[j].score {
					alicesPlaces[i] = places[j].place
					break
				}
			}
			if alice[i] >= places[j].score {
				alicesPlaces[i] = places[j].place
				break
			} else {
				continue
			}
		}
	}

	return alicesPlaces
}

type scoreAndPlace struct {
	score int32
	place int32
}
