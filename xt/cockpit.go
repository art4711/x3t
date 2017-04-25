package xt

import (
	"encoding/csv"
	"log"
	"os"
)

func GetCockpits(n string, text Text) []Cockpit {
	f, err := os.Open(n)
	if err != nil {
		log.Fatal(err)
	}

	// It's not really a csv file, but this works, so why not.
	r := csv.NewReader(f)
	r.Comment = '/'
	r.Comma = ';'

	rec, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	inf := struct {
		Ver  int
		Nrec int
	}{}
	t := tParser{rec: rec, t: text}
	t.parseAll(&inf)

	cockpits := make([]Cockpit, inf.Nrec)
	for i := 0; i < inf.Nrec; i++ {
		r.FieldsPerRecord = 0
		rec, err := r.Read()
		if err != nil {
			log.Fatal(err)
		}
		t := tParser{rec: rec, t: text}
		t.parseAll(&cockpits[i])
	}
	return cockpits
}

type Cockpit struct {
	BodyFile               string
	PictureID              string
	RotX                   float64
	RotY                   float64
	RotZ                   float64
	GalaxySubtype          string
	Description            string
	SceneFile              string
	LaserMask              int
	Volume                 string
	ProductionRelValNPC    string
	PriceModifier1         string
	PriceModifier2         string
	WareClass              string
	ProductionRelValPlayer string
	MinNotoriety           string
	VideoID                string
	Skin                   string
	ObjectID               string
}

/*
0. Body file	Not used
1. Picture ID	Not used
2. Rotation X
3. Rotation Y
4. Rotation Z
5. Galaxy subtype	Not used
6. Description	Not used
7. Scene file	Scene containing the graphics of the object
8. Laser mask	Bit mask defining which lasers can be equipped
9. Volume
10. Production RelVal (NPC)	Price for NPCs (it's not really a price)
11. Price modifier PRI	Primary Price Modifier
12. Price modifier SEC	Secondary Price Modifier
13. Ware class	Class (size) of the object - affects which ships can carry it
14. Production RelVal (player)	Price for the player (it's not really a price)
15. Min. Notoriety	Minimal notoriety the player must have to be able to use this object
16. Video ID	Stream ID from Videos.txt containing the animation displayed in the Info screen
17. Skin	Index to Skins.txt
18. Object ID	Identifier of the object
*/