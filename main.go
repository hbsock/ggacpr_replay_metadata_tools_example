package main

import (
	"fmt"
	"os"
	"hbsock/ggacpr_replay_metadata_tools/pkg"
	"encoding/csv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hi world")

	file, err := os.Open("/home/hanbinsock/programman/test_replay_20220127_1829_Klantsmurfen_RO_vs_Nibnab_JA.ggr")
	check(err)
	defer file.Close()

	data, err := metadata.GetReplayMetaData(file)
	check(err)


	w := csv.NewWriter(os.Stdout)

	err = w.Write(metadata.GetReplayMetadataHeaders())
	check(err)

	err = w.Write(data.ToStringSlice())
	check(err)

	w.Flush()
	check(w.Error())
}
