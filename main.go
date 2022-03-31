package main

import (
	"os"
	"hbsock/ggacpr_replay_metadata_tools/pkg"
	"encoding/csv"
	"io/ioutil"
	"path/filepath"
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




	err = w.Write(data.ToStringSlice())
	check(err)

}

func main() {

	w := csv.NewWriter(os.Stdout)

	err := w.Write(metadata.GetReplayMetadataHeaders())
	check(err)

	replays_dir := "/home/hanbinsock/programman/ggacr_replays"
	files, err := ioutil.ReadDir(replays_dir)
	check(err)

	for _, file := range files {
		if !file.IsDir() {
			replay_file_path := filepath.Join( replays_dir, file.Name() )
			readReplayAndOutputAsCsv(
				replay_file_path,
				w,
			)
		}
	}

	/*
	readReplayAndOutputAsCsv("/home/hanbinsock/programman/test_replay_20220127_1829_Klantsmurfen_RO_vs_Nibnab_JA.ggr",
		w,
	)
	*/

	w.Flush()
	check(w.Error())
}
