package internal

type Snapshot struct {
	Name    string
	Created string
	Size    int
}

type Snapshots []Snapshot

var dummyData = []Snapshot{
	{
		Name:    "snapshot1",
		Created: "2021-01-01T12:34:56",
		Size:    10,
	},
	{
		Name:    "snapshot2",
		Created: "2021-02-01T12:34:56",
		Size:    20,
	},
	{
		Name:    "snapshot3",
		Created: "2021-03-01T12:34:56",
		Size:    30,
	},
	{
		Name:    "snapshot4",
		Created: "2021-04-01T12:34:56",
		Size:    40,
	},
}

func (s Snapshots) TotalSize() int {
	total := 0
	for _, d := range s {
		total += d.Size
	}
	return total
}

func (s Snapshots) HasSnapshotWithName(name string) bool {
	for _, d := range s {
		if d.Name == name {
			return true
		}
	}
	return false
}

func GetSnapshots() Snapshots {
	return dummyData
}
