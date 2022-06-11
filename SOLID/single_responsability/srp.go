package single_responsability

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// the responsability of the Journal struct is to manage entries
// NOT to deal with persistence
type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(txt string) int {

	entry := fmt.Sprintf("%d: %s", j.GetEntriesCount()+1, txt)
	j.entries = append(j.entries, entry)
	return len(j.entries)
}
func (j *Journal) RemoveEntry(id int) {
	j.entries = append(j.entries[:id], j.entries[id+1:]...)
}

func (j *Journal) GetEntriesCount() int {
	return len(j.entries)
}

// implements Stringer
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// brakes separation of concerns
// i.e. God object
// implements persistence
// 		func (j *Journal) Save(filename string) {
// 			_ = ioutil.WriteFile(filename, []byte(j.String()), os.ModeAppend)
// 		}

// To not brake the SRP make Save() just a package function:
func Save(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}
