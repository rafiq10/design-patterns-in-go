package singleton

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var once sync.Once
var instance *singletonDB

type singletonDB struct {
	capitals map[string]int
}

func (db *singletonDB) GetPopulation(name string) int {
	return db.capitals[name]
}

// making it thread-safe - we don't want to instantiate to threads to instantiate the singleton
// we can use sync.Once or init() on the package level
// we need although laziness which is only guaranteed with sync.Once - not with init()
func GetSingletonDB() *singletonDB {
	once.Do(func() {
		caps, err := readData("singleton/capitals.txt")
		db := singletonDB{caps}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {
	pwd, _ := os.Getwd()
	absPath := filepath.Join(pwd, path)
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}
