package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var coins int
	var rounds int

	app := cli.NewApp()

	app.Name = "coinflip"
	app.Usage = "Run a simulation of a number of coins flipping"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "coins",
			Value:       20,
			Usage:       "Number of coins to flip",
			Destination: &coins,
		},

		cli.IntFlag{
			Name:        "rounds",
			Value:       1000,
			Usage:       "Number of rounds of flipping",
			Destination: &rounds,
		},
	}

	// Each coin may land on head or tail
	// bins[0] is incremented for no heads
	// bins[1] is incremented for 1 heads
	// bins[coins] is incremented for no tails
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Coins: %d\n", coins)
		fmt.Printf("Rounds: %d\n", rounds)

		// Set the bins up...
		bins := make([]int, coins+1)

		// Do all flips...
		for round := 0; round < rounds; round++ {
			heads := getFlipResult(coins, r)
			bins[heads]++
		}

		// Distribute output...
		nLines := 20
		lines := make([]int, nLines)
		for bin := 0; bin < len(bins); bin++ {
			lineForBin := int((float32(bin) / float32(len(bins))) * float32(len(lines)))
			lines[lineForBin] += bins[bin]
		}

		// Scale output...
		maxInLine := 0
		for line := 0; line < len(lines); line++ {
			if lines[line] > maxInLine {
				maxInLine = lines[line]
			}
		}

		// Display output...
		for line := 0; line < len(lines); line++ {
			strBar := strings.Repeat("*", (60*lines[line])/maxInLine)
			fmt.Printf("%4d: %s\n", line, strBar)
		}

		return nil
	}

	app.Run(os.Args)
}

func getFlipResult(coins int, r *rand.Rand) int {
	heads := 0
	for i := 0; i < coins; i++ {
		heads += r.Intn(2)
	}

	return heads
}
