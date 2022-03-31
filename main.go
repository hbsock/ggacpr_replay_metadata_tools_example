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

func readReplayAndOutputAsCsv(filepath string, w *csv.Writer) {

	file, err := os.Open(filepath)
	check(err)
	defer file.Close()

	data, err := metadata.GetReplayMetaData(file)
	check(err)



	err = w.Write(metadata.GetReplayMetadataHeaders())
	check(err)

	err = w.Write(data.ToStringSlice())
	check(err)

}

func main() {

	w := csv.NewWriter(os.Stdout)

	readReplayAndOutputAsCsv("/home/hanbinsock/programman/test_replay_20220127_1829_Klantsmurfen_RO_vs_Nibnab_JA.ggr",
		w,
	)

	w.Flush()
	check(w.Error())
}
