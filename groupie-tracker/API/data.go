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

	data.YoutubeLinks = map[string][]string{
		"queen": {
			"https://www.youtube.com/embed/-tJYN-eG1zk",
			"https://www.youtube.com/embed/fJ9rUzIMcZQ",
			"https://www.youtube.com/embed/HgzGwKwLmgM",
		},
		"soja": {
			"https://www.youtube.com/embed/ejfd2_ayxl0",  // rest of my life
			"https://www.youtube.com/embed/b9pw9dzbtoa",  // true love
			"https://www.youtube.com/embed/mpaahfmny4w",  // i believe (featuring michael franti)
		},		
		"pink floyd": {
			"https://www.youtube.com/embed/DLOth-BuCNY",  // Comfortably Numb
			"https://www.youtube.com/embed/YR5ApYxkU-U",  // Wish You Were Here
			"https://www.youtube.com/embed/-0kcet4aPpQ",  // Another Brick in the Wall (Part 2)
		},		
		"scorpions": {
			"https://www.youtube.com/embed/n4RjJKxsamQ",  // wind of change
			"https://www.youtube.com/embed/lrEAlhGtaLM",  // still loving you
			"https://www.youtube.com/embed/rB8HudfbaTE",  // rock you like a hurricane
		},
		"xxxtentacion": {
			"https://www.youtube.com/embed/ipvKvtgqJyc",  // look at me!
			"https://www.youtube.com/embed/pgN-vvVVxMA",  // sad!
			"https://www.youtube.com/embed/fGqdipSPwZ0",  // moonlight
		},
		"mac miller": {
			"https://www.youtube.com/embed/SSOoeY5vI3w",  // self care
			"https://www.youtube.com/embed/RmTfg2haOZk",  // diablo
			"https://www.youtube.com/embed/7xzU9Qqdqww",  // good news
		},
		"joyner lucas": {
			"https://www.youtube.com/embed/h1JBmRtESQU",  // i'm not racist
			"https://www.youtube.com/embed/87bCN1y4JPw",  // adhd
			"https://www.youtube.com/embed/KUM3l4vSddY",  // isis ft. logic
		},
		"kendrick lamar": {
			"https://www.youtube.com/embed/8aShfolR6w8",  // humble
			"https://www.youtube.com/embed/tvTRZJ-4EyI",  // dna
			"https://www.youtube.com/embed/Yt8wfGJ6HTQ",  // alright
		},
		"acdc": {
			"https://www.youtube.com/embed/v2AC41dglnM",  // thunderstruck
			"https://www.youtube.com/embed/gEPmA3USJdI",  // highway to hell
			"https://www.youtube.com/embed/pAgnJDJN4VA",  // back in black
		},
		"pearl jam": {
			"https://www.youtube.com/embed/ms2DDtX88do",  // alive
			"https://www.youtube.com/embed/9SPjIZx4wS8",  // jeremy
			"https://www.youtube.com/embed/hs8y3kneqrs",  // even flow
		},
		
		"katy perry": {
			"https://www.youtube.com/embed/CevxZvSJLk8",  // roar
			"https://www.youtube.com/embed/0KSOMA3QBU0",  // dark horse
			"https://www.youtube.com/embed/QGJuMBdaqIw",  // firework
		},
		"rihanna": {
			"https://www.youtube.com/embed/J3UjJ4wKLkg",  // diamonds
			"https://www.youtube.com/embed/2XtIKp3jwM8",  // work
			"https://www.youtube.com/embed/uelHwf8o7_U",  // love the way you lie (ft. eminem)
		},
		"genesis": {
			"https://www.youtube.com/embed/sry5l4ad-p8",  // invisible touch
			"https://www.youtube.com/embed/7uw1yZlMY5E",  // i can't dance
			"https://www.youtube.com/embed/qHT3Av_jIuw",  // land of confusion
		},
		"phil collins": {
			"https://www.youtube.com/embed/YkADj0TPrJA",  // in the air tonight
			"https://www.youtube.com/embed/1n_YniLzjBM",  // another day in paradise
			"https://www.youtube.com/embed/SidxJz94Svs",  // against all odds
		},
		"led zeppelin": {
			"https://www.youtube.com/embed/qyivczZI5pw",  // stairway to heaven
			"https://www.youtube.com/embed/xbhCPt6PZIU",  // whole lotta love
			"https://www.youtube.com/embed/JzDGd5ItLus",  // immigrant song
		},
		"the jimi hendrix experience": {
			"https://www.youtube.com/embed/fVyVIsvQoaE",  // all along the watchtower
			"https://www.youtube.com/embed/qFfnlYbFEiE",  // purple haze
			"https://www.youtube.com/embed/sAQIuEMl4AE",  // hey joe
		},
		"bee gees": {
			"https://www.youtube.com/embed/I_izvAbhExY",  // stayin' alive
			"https://www.youtube.com/embed/A3b9gOtQoq4",  // how deep is your love
			"https://www.youtube.com/embed/ekxtsWyfXdg",  // night fever
		},
		"deep purple": {
			"https://www.youtube.com/embed/zUwEIt9ez7M",  // smoke on the water
			"https://www.youtube.com/embed/W0PaMOy48qM",  // highway star
			"https://www.youtube.com/embed/IkN6rc9RNeI",  // child in time
		},
		"aerosmith": {
			"https://www.youtube.com/embed/FM9htniBBaA",  // dream on
			"https://www.youtube.com/embed/JRl3p7sf9G8",  // i don't want to miss a thing
			"https://www.youtube.com/embed/4y_x9JdG45k",  // crazy
		},
		"dire straits": {
			"https://www.youtube.com/embed/wTP2RUD_cL0",  // sultans of swing
			"https://www.youtube.com/embed/lAD6Obi7Cag",  // money for nothing
			"https://www.youtube.com/embed/ami6noWCoZs",  // brothers in arms
		},
		"mamonas assassinas": {
			"https://www.youtube.com/embed/bmrMQPSwevA",  // pelados em santos
			"https://www.youtube.com/embed/ABJt01iCHaE",  // robocop gay
			"https://www.youtube.com/embed/bO9kFfwJlYk",  // sabão crá-crá
		},
		"thirty seconds to mars": {
			"https://www.youtube.com/embed/8yvGCAvOAfM",  // the kill
			"https://www.youtube.com/embed/LmBRNnR8YHY",  // closer to the edge
			"https://www.youtube.com/embed/hTMrlHHVx8A",  // this is war
		},
		"imagine dragons": {
			"https://www.youtube.com/embed/ktvTqknDobU",  // radioactive
			"https://www.youtube.com/embed/IXXxciRUMzE",  // demons
			"https://www.youtube.com/embed/ktvTqknDobU",  // believer
		},
		"juice wrld": {
			"https://www.youtube.com/embed/mzB1VGEGcSU",  // lucid dreams
			"https://www.youtube.com/embed/K_IYlHXU8RI",  // robbery
			"https://www.youtube.com/embed/F-k4V4EVTgo",  // all girls are the same
		},
		"logic": {
			"https://www.youtube.com/embed/Kb24RrHIbFk",  // 1-800-273-8255
			"https://www.youtube.com/embed/jzpy7R8Nj4A",  // homicide
			"https://www.youtube.com/embed/Jv3xQ6uIhzE",  // everyday
		},	
		"alec benjamin": {
			"https://www.youtube.com/embed/50VNCymT-Cs",  // let me down slowly
			"https://www.youtube.com/embed/t2RfriaxGsc",  // boy in the bubble
			"https://www.youtube.com/embed/MVGLhEB6BTs",  // outrunning karma
		},		
		"bobby mcferrin": {
			"https://www.youtube.com/embed/d-diB65scQU",  // don't worry, be happy
			"https://www.youtube.com/embed/MvguECtTgZ4",  // drive my car
			"https://www.youtube.com/embed/Mb7Cw7hN4DQ",  // thinking about your body
		},
		"r3hab": {
			"https://www.youtube.com/embed/J--MguAkl3o",  // this is how we party
			"https://www.youtube.com/embed/YvzgMBwsm1c",  // all around the world (la la la)
			"https://www.youtube.com/embed/YV5NzGt7Ka8",  // wrong move
		},
		"post malone": {
			"https://www.youtube.com/embed/ApXoWvfEYVU",  // rockstar
			"https://www.youtube.com/embed/UceaB4D0jpo",  // circles
			"https://www.youtube.com/embed/EyIvuigqDoA",  // sunflower
		},
		"travis scott": {
			"https://www.youtube.com/embed/tG-heK0kDX8",  // sicko mode
			"https://www.youtube.com/embed/QkRvcuZyi4o",  // goosebumps
			"https://www.youtube.com/embed/0hthFysVO7I",  // highest in the room
		},
		"j. cole": {
			"https://www.youtube.com/embed/eCGV26aj-mM",  // middle child
			"https://www.youtube.com/embed/ziOB9V_WI0c",  // no role modelz
			"https://www.youtube.com/embed/KiNVzZDUQW8",  // wet dreamz
		},
		"nickelback": {
			"https://www.youtube.com/embed/_nWoZbPq6-Y",  // how you remind me
			"https://www.youtube.com/embed/71Ll7Mycp2Y",  // photograph
			"https://www.youtube.com/embed/2m1fsHAI8Rk",  // rockstar
		},
		"mobb deep": {
			"https://www.youtube.com/embed/DoZ9ldMzX8o",  // shook ones, pt ii
			"https://www.youtube.com/embed/4qUOdCImJus",  // survival of the fittest
			"https://www.youtube.com/embed/qQhxXOOwg_o",  // hell on earth
		},
		"guns n' roses": {
			"https://www.youtube.com/embed/o1tj2zJ2Wvg",  // sweet child o' mine
			"https://www.youtube.com/embed/8SbUC-UaAxE",  // november rain
			"https://www.youtube.com/embed/pAgnJDJN4VA",  // paradise city
		},
		"nwa": {
			"https://www.youtube.com/embed/kzhqIf6G5JY",  // straight outta compton
			"https://www.youtube.com/embed/SjbPi00k_ME",  // fuck tha police
			"https://www.youtube.com/embed/AMv7CJuJ3R8",  // express yourself
		},
		"u2": {
			"https://www.youtube.com/embed/ftjEcrrf7r0",  // with or without you
			"https://www.youtube.com/embed/e3-5YC_oHjE",  // i still haven't found what i'm looking for
			"https://www.youtube.com/embed/ggFHKa3o4oI",  // beautiful day
		},
		"arctic monkeys": {
			"https://www.youtube.com/embed/bpOSxM0rNPM",  // do i wanna know?
			"https://www.youtube.com/embed/TuPw4iFgWyQ",  // why'd you only call me when you're high?
			"https://www.youtube.com/embed/gJgnRao1H8g",  // arabella
		},
		"fall out boy": {
			"https://www.youtube.com/embed/l9PxOanFjxQ",  // centuries
			"https://www.youtube.com/embed/LBr7kECsjcQ",  // sugar, we're goin down
			"https://www.youtube.com/embed/01VnOkMWlFk",  // thnks fr th mmrs
		},
		"gorillaz": {
			"https://www.youtube.com/embed/B0pkQgqbMGI",  // feel good inc.
			"https://www.youtube.com/embed/rB5Nbp_gmgQ",  // clint eastwood
			"https://www.youtube.com/embed/HyHNuVaZJ-k",  // on melancholy hill
		},
		"eagles": {
			"https://www.youtube.com/embed/lrfRFb0FyAM",  // hotel california
			"https://www.youtube.com/embed/qa4Jlbw9fUI",  // take it easy
			"https://www.youtube.com/embed/lOxS-nX9ozA",  // life in the fast lane
		},
		"linkin park": {
			"https://www.youtube.com/embed/kXYiU_JCYtU",  // numb
			"https://www.youtube.com/embed/eVTXPUF4Oz4",  // in the end
			"https://www.youtube.com/embed/8sgycukafqQ",  // what i've done
		},
		"red hot chili peppers": {
			"https://www.youtube.com/embed/YlUKcNNmywk",  // californication
			"https://www.youtube.com/embed/Gd9OhYroLN0",  // under the bridge
			"https://www.youtube.com/embed/Q0oIoR9mLwc",  // can't stop
		},
		"eminem": {
			"https://www.youtube.com/embed/_Yhyp-_hX2s",  // lose yourself
			"https://www.youtube.com/embed/j5-yKhDd64s",  // not afraid
			"https://www.youtube.com/embed/uelHwf8o7_U",  // love the way you lie
		},
		"green day": {
			"https://www.youtube.com/embed/jVO8sUrs-Pw",  // 21 guns
			"https://www.youtube.com/embed/Soa3gO7tL-c",  // boulevard of broken dreams
			"https://www.youtube.com/embed/e3VGRcYSx9Y",  // american idiot
		},
		"metallica": {
			"https://www.youtube.com/embed/CD-E-LDc384",  // nothing else matters
			"https://www.youtube.com/embed/Tj75Arhq5ho",  // enter sandman
			"https://www.youtube.com/embed/Rbm6GXllBiw",  // the unforgiven
		},
		"maroon 5": {
			"https://www.youtube.com/embed/aJOTlE1K90k",  // girls like you
			"https://www.youtube.com/embed/09R8_2nJtjg",  // sugar
			"https://www.youtube.com/embed/8uC-WQHBWyc",  // moves like jagger
		},
		"twenty one pilots": {
			"https://www.youtube.com/embed/pXRviuL6vMY",  // stressed out
			"https://www.youtube.com/embed/fiqPY5EB1uY",  // ride
			"https://www.youtube.com/embed/ytDE4wD-4uc",  // heathens
		},
		"the rolling stones": {
			"https://www.youtube.com/embed/Jwtyn-L-2gQ",  // paint it black
			"https://www.youtube.com/embed/8oqhjNP4jK4",  // satisfaction
			"https://www.youtube.com/embed/Bz65VRzboeM",  // sympathy for the devil
		},
		"muse": {
			"https://www.youtube.com/embed/G_sBOsh-vyI",  // uprising
			"https://www.youtube.com/embed/w8KQmps-Sog",  // starlight
			"https://www.youtube.com/embed/azVqekQBK8g",  // psycho
		},
		"foo fighters": {
			"https://www.youtube.com/embed/eBG7P-K-r1Y",  // everlong
			"https://www.youtube.com/embed/SBjQ9tuuTJQ",  // the pretender
			"https://www.youtube.com/embed/q0TrqXHoL10",  // best of you
		},
		"the chainsmokers": {
			"https://www.youtube.com/embed/PT2_F-1esPk",  // closer
			"https://www.youtube.com/embed/Io0fBr1XBUA",  // something just like this
			"https://www.youtube.com/embed/e8eRIo0AweU",  // don't let me down
		},			
	}
}
