package apiserver

type ApiData struct {
	BandImages      []string
	MembersImages   []string
	LocationsImages []string
	YoutubeLinks    map[string][]string
}

// Initialise my ApiData struct fields.
func (data *ApiData) Initialise() {

	data.BandImages = []string{
		"https://i.postimg.cc/QMq6MD91/soja.png",
		"https://i.postimg.cc/xTgDzjyt/queen.png",
		"https://i.postimg.cc/DfQNVd4z/pink-floyd.png",
		"https://i.postimg.cc/jjrNrN7s/scorpions.png",
		"https://i.postimg.cc/Px3hj4yP/xxxtentacion.png",
		"https://i.postimg.cc/YjKgqmsf/mac-miller.png",
		"https://i.postimg.cc/L6GR53vM/joyner-lucas.png",
		"https://i.postimg.cc/rwBNh2Lp/kendrick-lamar.png",
		"https://i.postimg.cc/vHMBtPjd/acdc.png",
		"https://i.postimg.cc/DwXMfgv8/pearl-jam.png",
		"https://i.postimg.cc/KYYQkDFP/katy-perry.png",
		"https://i.postimg.cc/MKxLZRcD/rihanna.png",
		"https://i.postimg.cc/YScnVWwY/genesis.png",
		"https://i.postimg.cc/658FmZP5/phil-collins.png",
		"https://i.postimg.cc/j5GkVgmx/led-zeppelin.png",
		"https://i.postimg.cc/V65R4xfG/the-jimi-hendrix-experience.png",
		"https://i.postimg.cc/XqHCrqSQ/bee-gees.png",
		"https://i.postimg.cc/ry18TJBm/deep-purple.png",
		"https://i.postimg.cc/d11qwky0/aerosmith.png",
		"https://i.postimg.cc/633YcwDS/dire-straits.png",
		"https://i.postimg.cc/j2f96tfT/mamonas-assassinas.png",
		"https://i.postimg.cc/8Ph9jkDB/thirty-seconds-to-mars.png",
		"https://i.postimg.cc/HLLhZQMB/imagine-dragons.png",
		"https://i.postimg.cc/7hrXk7Cv/juice-wrld.png",
		"https://i.postimg.cc/W4QwGrwD/logic.png",
		"https://i.postimg.cc/SN3RNrSz/alec-benjamin.png",
		"https://i.postimg.cc/xdxThy4C/bobby-mcferrins.png",
		"https://i.postimg.cc/mkcc86ZB/r3hab.png",
		"https://i.postimg.cc/PJZDdFgK/post-malone.png",
		"https://i.postimg.cc/90W2fDSS/travis-scott.png",
		"https://i.postimg.cc/q7TQC2J1/j-cole.png",
		"https://i.postimg.cc/kGLYY48G/nickelback.png",
		"https://i.postimg.cc/QCrgPTCF/mob-deep.png",
		"https://i.postimg.cc/BQjDgpMm/guns-n-roses.png",
		"https://i.postimg.cc/xdX1rpNZ/nwa.png",
		"https://i.postimg.cc/Mpb0JZNn/u2.png",
		"https://i.postimg.cc/3rdP8zK8/arctic-monkeys.png",
		"https://i.postimg.cc/02CGL2Lb/fall-out-boy.png",
		"https://i.postimg.cc/sD0xBkWT/gorillaz.png",
		"https://i.postimg.cc/bY4Npxf5/eagles.png",
		"https://i.postimg.cc/JhbCpp8Z/linkin-park.png",
		"https://i.postimg.cc/t4qhNQvW/red-hot-chili-peppers.png",
		"https://i.postimg.cc/Y0vXsSFZ/eminem.png",
		"https://i.postimg.cc/MKjP0vqj/green-day.png",
		"https://i.postimg.cc/9f1JX9ts/metallica.png",
		"https://i.postimg.cc/ncB8Fds5/coldplay.png",
		"https://i.postimg.cc/T36zprgP/maroon-5.png",
		"https://i.postimg.cc/v8xkhPvB/twenty-one-pilots.png",
		"https://i.postimg.cc/XvD1cFGd/the-rolling-stones.png",
		"https://i.postimg.cc/NjddBTxS/muse.png",
		"https://i.postimg.cc/pL3nfLVy/foo-fighters.png",
		"https://i.postimg.cc/TYwyBDpp/the-chainsmokers.png",
	}

	data.MembersImages = []string{
		// Queen
		"https://i.postimg.cc/w1crjTtB/roger-meddows-taylor.png",
		"https://i.postimg.cc/vDcCtsKP/barry-mitchell.webp",
		"https://i.postimg.cc/9zKKxk5c/brian-may.webp",
		"https://i.postimg.cc/pmt7vnL7/freddie-mercury.webp",
		"https://i.postimg.cc/0KM3kYmg/john-daecon.webp",
		"https://i.postimg.cc/TLHFv7Kq/mike-grose.webp",
		"https://i.postimg.cc/w1crjTtB/roger-meddows-taylor.png",

		// Soja
		"https://i.postimg.cc/g20kgd2G/jacob-hemphill.webp",
		"https://i.postimg.cc/CKq02KTp/ryan-byrd-berty.jpg",
		"https://i.postimg.cc/pTHw1YJK/bob-jefferson.jpg",
		"https://i.postimg.cc/WpW2xfRP/hellman-escorcia.jpg",
		"https://i.postimg.cc/4djJMmMh/ken-brownell.jpg",
		"https://i.postimg.cc/3xjzQdyD/patrick-oshea.webp",
		"https://i.postimg.cc/Rq5pjB4R/rafael-rodriguez.jpg",

	}

	data.LocationsImages = []string{
		// Queen
		"https://i.postimg.cc/tgMZBt72/dunedin-new-zealand.jpg",
		"https://i.postimg.cc/K8QgMybX/north-carolina-usa.webp",
		"https://i.postimg.cc/sg3hxgt9/georgia-usa.webp",
		"https://i.postimg.cc/yYmR7Zr8/los-angeles-usa.jpg",
		"https://i.postimg.cc/v8zBhVrs/saitama-japan.webp",
		"https://i.postimg.cc/pdT5Q1Vm/nagoya-japan.jpg",
		"https://i.postimg.cc/J01hmrWv/penrose-new-zealand.jpg",
		"https://i.postimg.cc/x1qcc1bS/osaka-japan.jpg",
		"https://i.postimg.cc/4djJMmMh/ken-brownell.jpg",
	}

	data.YoutubeLinks = map[string][]string{
		"Queen": {
			"https://www.youtube.com/embed/-tJYN-eG1zk",
			"https://www.youtube.com/embed/fJ9rUzIMcZQ",
			"https://www.youtube.com/embed/HgzGwKwLmgM",			
		},
		"soja": {
			"https://www.youtube.com/embed/X572Mp_r46E",
			
		},
		"Katy Perry": {
			"https://www.youtube.com/embed/CevxZvSJLk8",
		},
	}
}
