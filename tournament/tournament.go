package tournament

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type T struct {
	played int
	won int
	drawn int
	lost int
	points int
}

func Tally(in io.Reader, out io.Writer) error {
	all, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	scores := map[string]*T{}
	for _, line := range strings.Split(string(all), "\n") {
		if line == "" || line[0] == '#' {
			continue
		}
		items := strings.Split(line, ";")
		fmt.Println(line)
		if len(items) != 3 {
			return fmt.Errorf("len")
		}
		a := items[0]
		b := items[1]
		if _, ok := scores[a]; !ok {
			scores[a] = &T{}
		}
		if _, ok := scores[b]; !ok {
			scores[b] = &T{}
		}
		result := items[2]
		if result == "draw" {
			scores[a].drawn++
			scores[b].drawn++
			scores[a].points++
			scores[b].points++
		} else if result == "win" || result == "loss" {
			var w, l string
			if result == "win" {
				w = a
				l = b
			} else {
				w = b
				l = a
			}
			scores[w].won++
			scores[l].lost++
			scores[w].points += 3
		} else {
			return fmt.Errorf("bad")
		}
		scores[a].played++
		scores[b].played++
	}
	teams := make([]string, 0)
	for t := range scores {
		teams = append(teams, t)
	}
	sort.Slice(teams, func(i, j int) bool {
		a := teams[i]
		b := teams[j]
		if scores[a].points == scores[b].points {
			return a < b
		} else {
			return scores[a].points > scores[b].points
		}
	})
	writer := bufio.NewWriter(out)
	_, err = writer.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}
	for _, t := range teams {
		s := scores[t]
		_, err = writer.WriteString(fmt.Sprintf("%s%s| %2d | %2d | %2d | %2d | %2d\n", t, strings.Repeat(" ", 31 - len(t)), s.played, s.won, s.drawn, s.lost, s.points))
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}
