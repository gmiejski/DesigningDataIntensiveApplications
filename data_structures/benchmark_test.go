package data_structures

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const hashIndexBenchmarks = "/tmp/ddia/benchmarks/hashindex/"
const hashFileBenchmarks = "/tmp/ddia/benchmarks/filedb/"

var person = Person{
	Name:    "Greg",
	Surname: "Misiaosbf",
	Kids:    []string{"Alex", "Natasha"},
}

// TODO make benchmarks comparable between itself
// TODO prepare benchmarks for more stable evaluation

func BenchmarkSavingFileDB(b *testing.B) {
	b.StopTimer()
	index, err := newFileDB(hashFileBenchmarks)
	fmt.Print("dsaasdasd")
	//defer cleanup(hashFileBenchmarks)
	require.NoError(b, err)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = index.Save(1, person)
	}
}

func BenchmarkSavingHashIndex(b *testing.B) {
	b.StopTimer()
	index, err := newFileDB(hashFileBenchmarks)
	fmt.Print("dsaasdasd")
	//defer cleanup(hashFileBenchmarks)
	require.NoError(b, err)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = index.Save(1, person)
	}
}

func BenchmarkHashIndex(b *testing.B) {
	b.StopTimer()
	index, err := newHashIndex(hashIndexBenchmarks)
	defer cleanup(hashIndexBenchmarks)
	require.NoError(b, err)

	dbSize := 10000000
	for i := 0; i < dbSize; i++ {
		err := index.Save(i, person)
		if err != nil {
			fmt.Sprintln(err.Error())
		}
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		index.Find(14)
	}
}

func BenchmarkFileDB(b *testing.B) {
	b.StopTimer()
	index, err := newFileDB(hashFileBenchmarks)
	defer cleanup(hashFileBenchmarks)
	require.NoError(b, err)

	dbSize := 10000
	for i := 0; i < dbSize; i++ {
		err := index.Save(i, person)
		if err != nil {
			fmt.Sprintln(err.Error())
		}
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		index.Find(15)
	}
}

func TestTimeOfRead(t *testing.T) {
	// given
	d := "/tmp/ddia/benchmarks/timed"
	index, err := newHashIndex(d)
	defer cleanup(d)
	require.NoError(t, err)

	dbSize := 100000
	for i := 0; i < dbSize; i++ {
		err := index.Save(i, person)
		if err != nil {
			fmt.Sprintln(err.Error())
		}
	}
	// when
	start := time.Now()
	index.Find(15)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	// then
}
