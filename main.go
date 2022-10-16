package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Item struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	InfoHash string `json:"info_hash"`
	Leechers string `json:"leechers"`
	Seeders  string `json:"seeders"`
	NumFiles string `json:"num_files"`
	Size     string `json:"size"`
	Username string `json:"username"`
	Added    string `json:"added"`
	Status   string `json:"status"`
	Category string `json:"category"`
	Imdb     string `json:"imdb"`
}

func main() {
	fmt.Println("Welcome to TPB searcher")
	args := os.Args[1:]
	searchQuery := strings.Join(args, " ")
	fmt.Println("Searching for", searchQuery)
	resp, err := http.Get("https://apibay.org/q.php?q=" + url.QueryEscape(searchQuery))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var items []Item

	if err := json.Unmarshal([]byte(body), &items); err != nil {
		panic(err)
	}

	for i, item := range items {
		fmt.Println(i, item.Name)
	}

	fmt.Print("Please select torrent: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	index, err := strconv.Atoi(input.Text())

	if err != nil {
		panic(err)
	}

	chosenItem := items[index]
	fmt.Println("You chose:", chosenItem.Name)
	url := fmt.Sprintf("magnet:?xt=urn:btih:%s&dn=%s", chosenItem.InfoHash, url.QueryEscape(strings.TrimSpace(chosenItem.Name)))
	command := "open " + fmt.Sprintf("'%s'", url)
	err = exec.Command("bash", "-c", command).Run()

	if err != nil {
		fmt.Println(err)
	}
}
