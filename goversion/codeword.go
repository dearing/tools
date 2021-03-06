package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	adj = [...]string{

		"abandoned",
		"able",
		"absolute",
		"adorable",
		"adventurous",
		"academic",
		"acceptable",
		"acclaimed",
		"accomplished",
		"accurate",
		"aching",
		"acidic",
		"acrobatic",
		"active",
		"actual",
		"adept",
		"admirable",
		"admired",
		"adolescent",
		"adorable",
		"adored",
		"advanced",
		"afraid",
		"affectionate",
		"aged",
		"aggravating",
		"aggressive",
		"agile",
		"agitated",
		"agonizing",
		"agreeable",
		"ajar",
		"alarmed",
		"alarming",
		"alert",
		"alienated",
		"alive",
		"all",
		"altruistic",
		"amazing",
		"ambitious",
		"ample",
		"amused",
		"amusing",
		"anchored",
		"ancient",
		"angelic",
		"angry",
		"anguished",
		"animated",
		"annual",
		"another",
		"antique",
		"anxious",
		"any",
		"apprehensive",
		"appropriate",
		"apt",
		"arctic",
		"arid",
		"aromatic",
		"artistic",
		"ashamed",
		"assured",
		"astonishing",
		"athletic",
		"attached",
		"attentive",
		"attractive",
		"austere",
		"authentic",
		"authorized",
		"automatic",
		"avaricious",
		"average",
		"aware",
		"awesome",
		"awful",
		"awkward",
		"babyish",
		"bad",
		"back",
		"baggy",
		"bare",
		"barren",
		"basic",
		"beautiful",
		"belated",
		"beloved",
		"beneficial",
		"better",
		"best",
		"bewitched",
		"big",
		"big-hearted",
		"biodegradable",
		"bite-sized",
		"bitter",
		"black",
		"black-and-white",
		"bland",
		"blank",
		"blaring",
		"bleak",
		"blind",
		"blissful",
		"blond",
		"blue",
		"blushing",
		"bogus",
		"boiling",
		"bold",
		"bony",
		"boring",
		"bossy",
		"both",
		"bouncy",
		"bountiful",
		"bowed",
		"brave",
		"breakable",
		"brief",
		"bright",
		"brilliant",
		"brisk",
		"broken",
		"bronze",
		"brown",
		"bruised",
		"bubbly",
		"bulky",
		"bumpy",
		"buoyant",
		"burdensome",
		"burly",
		"bustling",
		"busy",
		"buttery",
		"buzzing",
		"calculating",
		"calm",
		"candid",
		"canine",
		"capital",
		"carefree",
		"careful",
		"careless",
		"caring",
		"cautious",
		"cavernous",
		"celebrated",
		"charming",
		"cheap",
		"cheerful",
		"cheery",
		"chief",
		"chilly",
		"chubby",
		"circular",
		"classic",
		"clean",
		"clear",
		"clear-cut",
		"clever",
		"close",
		"closed",
		"cloudy",
		"clueless",
		"clumsy",
		"cluttered",
		"coarse",
		"cold",
		"colorful",
		"colorless",
		"colossal",
		"comfortable",
		"common",
		"compassionate",
		"competent",
		"complete",
		"complex",
		"complicated",
		"composed",
		"concerned",
		"concrete",
		"confused",
		"conscious",
		"considerate",
		"constant",
		"content",
		"conventional",
		"cooked",
		"cool",
		"cooperative",
		"coordinated",
		"corny",
		"corrupt",
		"costly",
		"courageous",
		"courteous",
		"crafty",
		"crazy",
		"creamy",
		"creative",
		"creepy",
		"criminal",
		"crisp",
		"critical",
		"crooked",
		"crowded",
		"cruel",
		"crushing",
		"cuddly",
		"cultivated",
		"cultured",
		"cumbersome",
		"curly",
		"curvy",
		"cute",
		"cylindrical",
		"damaged",
		"damp",
		"dangerous",
		"dapper",
		"daring",
		"darling",
		"dark",
		"dazzling",
		"dead",
		"deadly",
		"deafening",
		"dear",
		"dearest",
		"decent",
		"decimal",
		"decisive",
		"deep",
		"defenseless",
		"defensive",
		"defiant",
		"deficient",
		"definite",
		"definitive",
		"delayed",
		"delectable",
		"delicious",
		"delightful",
		"delirious",
		"demanding",
		"dense",
		"dental",
		"dependable",
		"dependent",
		"descriptive",
		"deserted",
		"detailed",
		"determined",
		"devoted",
		"different",
		"difficult",
		"digital",
		"diligent",
		"dim",
		"dimpled",
		"dimwitted",
		"direct",
		"disastrous",
		"discrete",
		"disfigured",
		"disgusting",
		"disloyal",
		"dismal",
		"distant",
		"downright",
		"dreary",
		"dirty",
		"disguised",
		"dishonest",
		"dismal",
		"distant",
		"distinct",
		"distorted",
		"dizzy",
		"dopey",
		"doting",
		"double",
		"downright",
		"drab",
		"drafty",
		"dramatic",
		"dreary",
		"droopy",
		"dry",
		"dual",
		"dull",
		"dutiful",
		"each",
		"eager",
		"earnest",
		"early",
		"easy",
		"easy-going",
		"ecstatic",
		"edible",
		"educated",
		"elaborate",
		"elastic",
		"elated",
		"elderly",
		"electric",
		"elegant",
		"elementary",
		"elliptical",
		"embarrassed",
		"embellished",
		"eminent",
		"emotional",
		"empty",
		"enchanted",
		"enchanting",
		"energetic",
		"enlightened",
		"enormous",
		"enraged",
		"entire",
		"envious",
		"equal",
		"equatorial",
		"essential",
		"esteemed",
		"ethical",
		"euphoric",
		"even",
		"evergreen",
		"everlasting",
		"every",
		"evil",
		"exalted",
		"excellent",
		"exemplary",
		"exhausted",
		"excitable",
		"excited",
		"exciting",
		"exotic",
		"expensive",
		"experienced",
		"expert",
		"extraneous",
		"extroverted",
		"extra-large",
		"extra-small",
		"fabulous",
		"failing",
		"faint",
		"fair",
		"faithful",
		"fake",
		"false",
		"familiar",
		"famous",
		"fancy",
		"fantastic",
		"far",
		"faraway",
		"far-flung",
		"far-off",
		"fast",
		"fat",
		"fatal",
		"fatherly",
		"favorable",
		"favorite",
		"fearful",
		"fearless",
		"feisty",
		"feline",
		"female",
		"feminine",
		"few",
		"fickle",
		"filthy",
		"fine",
		"finished",
		"firm",
		"first",
		"firsthand",
		"fitting",
		"fixed",
		"flaky",
		"flamboyant",
		"flashy",
		"flat",
		"flawed",
		"flawless",
		"flickering",
		"flimsy",
		"flippant",
		"flowery",
		"fluffy",
		"fluid",
		"flustered",
		"focused",
		"fond",
		"foolhardy",
		"foolish",
		"forceful",
		"forked",
		"formal",
		"forsaken",
		"forthright",
		"fortunate",
		"fragrant",
		"frail",
		"frank",
		"frayed",
		"free",
		"French",
		"fresh",
		"frequent",
		"friendly",
		"frightened",
		"frightening",
		"frigid",
		"frilly",
		"frizzy",
		"frivolous",
		"front",
		"frosty",
		"frozen",
		"frugal",
		"fruitful",
		"full",
		"fumbling",
		"functional",
		"funny",
		"fussy",
		"fuzzy",
		"gargantuan",
		"gaseous",
		"general",
		"generous",
		"gentle",
		"genuine",
		"giant",
		"giddy",
		"gigantic",
		"gifted",
		"giving",
		"glamorous",
		"glaring",
		"glass",
		"gleaming",
		"gleeful",
		"glistening",
		"glittering",
		"gloomy",
		"glorious",
		"glossy",
		"glum",
		"golden",
		"good",
		"good-natured",
		"gorgeous",
		"graceful",
		"gracious",
		"grand",
		"grandiose",
		"granular",
		"grateful",
		"grave",
		"gray",
		"great",
		"greedy",
		"green",
		"gregarious",
		"grim",
		"grimy",
		"gripping",
		"grizzled",
		"gross",
		"grotesque",
		"grouchy",
		"grounded",
		"growing",
		"growling",
		"grown",
		"grubby",
		"gruesome",
		"grumpy",
		"guilty",
		"gullible",
		"gummy",
		"hairy",
		"half",
		"handmade",
		"handsome",
		"handy",
		"happy",
		"happy-go-lucky",
		"hard",
		"hard-to-find",
		"harmful",
		"harmless",
		"harmonious",
		"harsh",
		"hasty",
		"hateful",
		"haunting",
		"healthy",
		"heartfelt",
		"hearty",
		"heavenly",
		"heavy",
		"hefty",
		"helpful",
		"helpless",
		"hidden",
		"hideous",
		"high",
		"high-level",
		"hilarious",
		"hoarse",
		"hollow",
		"homely",
		"honest",
		"honorable",
		"honored",
		"hopeful",
		"horrible",
		"hospitable",
		"hot",
		"huge",
		"humble",
		"humiliating",
		"humming",
		"humongous",
		"hungry",
		"hurtful",
		"husky",
		"icky",
		"icy",
		"ideal",
		"idealistic",
		"identical",
		"idle",
		"idiotic",
		"idolized",
		"ignorant",
		"ill",
		"illegal",
		"ill-fated",
		"ill-informed",
		"illiterate",
		"illustrious",
		"imaginary",
		"imaginative",
		"immaculate",
		"immaterial",
		"immediate",
		"immense",
		"impassioned",
		"impeccable",
		"impartial",
		"imperfect",
		"imperturbable",
		"impish",
		"impolite",
		"important",
		"impossible",
		"impractical",
		"impressionable",
		"impressive",
		"improbable",
		"impure",
		"inborn",
		"incomparable",
		"incompatible",
		"incomplete",
		"inconsequential",
		"incredible",
		"indelible",
		"inexperienced",
		"indolent",
		"infamous",
		"infantile",
		"infatuated",
		"inferior",
		"infinite",
		"informal",
		"innocent",
		"insecure",
		"insidious",
		"insignificant",
		"insistent",
		"instructive",
		"insubstantial",
		"intelligent",
		"intent",
		"intentional",
		"interesting",
		"internal",
		"international",
		"intrepid",
		"ironclad",
		"irresponsible",
		"irritating",
		"itchy",
		"jaded",
		"jagged",
		"jam-packed",
		"jaunty",
		"jealous",
		"jittery",
		"joint",
		"jolly",
		"jovial",
		"joyful",
		"joyous",
		"jubilant",
		"judicious",
		"juicy",
		"jumbo",
		"junior",
		"jumpy",
		"juvenile",
		"kaleidoscopic",
		"keen",
		"key",
		"kind",
		"kindhearted",
		"kindly",
		"klutzy",
		"knobby",
		"knotty",
		"knowledgeable",
		"knowing",
		"known",
		"kooky",
		"kosher",
		"lame",
		"lanky",
		"large",
		"last",
		"lasting",
		"late",
		"lavish",
		"lawful",
		"lazy",
		"leading",
		"lean",
		"leafy",
		"left",
		"legal",
		"legitimate",
		"light",
		"lighthearted",
		"likable",
		"likely",
		"limited",
		"limp",
		"limping",
		"linear",
		"lined",
		"liquid",
		"little",
		"live",
		"lively",
		"livid",
		"loathsome",
		"lone",
		"lonely",
		"long",
		"long-term",
		"loose",
		"lopsided",
		"lost",
		"loud",
		"lovable",
		"lovely",
		"loving",
		"low",
		"loyal",
		"lucky",
		"lumbering",
		"luminous",
		"lumpy",
		"lustrous",
		"luxurious",
		"mad",
		"made-up",
		"magnificent",
		"majestic",
		"major",
		"male",
		"mammoth",
		"married",
		"marvelous",
		"masculine",
		"massive",
		"mature",
		"meager",
		"mealy",
		"mean",
		"measly",
		"meaty",
		"medical",
		"mediocre",
		"medium",
		"meek",
		"mellow",
		"melodic",
		"memorable",
		"menacing",
		"merry",
		"messy",
		"metallic",
		"mild",
		"milky",
		"mindless",
		"miniature",
		"minor",
		"minty",
		"miserable",
		"miserly",
		"misguided",
		"misty",
		"mixed",
		"modern",
		"modest",
		"moist",
		"monstrous",
		"monthly",
		"monumental",
		"moral",
		"mortified",
		"motherly",
		"motionless",
		"mountainous",
		"muddy",
		"muffled",
		"multicolored",
		"mundane",
		"murky",
		"mushy",
		"musty",
		"muted",
		"mysterious",
		"naive",
		"narrow",
		"nasty",
		"natural",
		"naughty",
		"nautical",
		"near",
		"neat",
		"necessary",
		"needy",
		"negative",
		"neglected",
		"negligible",
		"neighboring",
		"nervous",
		"new",
		"next",
		"nice",
		"nifty",
		"nimble",
		"nippy",
		"nocturnal",
		"noisy",
		"nonstop",
		"normal",
		"notable",
		"noted",
		"noteworthy",
		"novel",
		"noxious",
		"numb",
		"nutritious",
		"nutty",
		"obedient",
		"obese",
		"oblong",
		"oily",
		"oblong",
		"obvious",
		"occasional",
		"odd",
		"oddball",
		"offbeat",
		"offensive",
		"official",
		"old",
		"old-fashioned",
		"only",
		"open",
		"optimal",
		"optimistic",
		"opulent",
		"orange",
		"orderly",
		"organic",
		"ornate",
		"ornery",
		"ordinary",
		"original",
		"other",
		"our",
		"outlying",
		"outgoing",
		"outlandish",
		"outrageous",
		"outstanding",
		"oval",
		"overcooked",
		"overdue",
		"overjoyed",
		"overlooked",
		"palatable",
		"pale",
		"paltry",
		"parallel",
		"parched",
		"partial",
		"passionate",
		"past",
		"pastel",
		"peaceful",
		"peppery",
		"perfect",
		"perfumed",
		"periodic",
		"perky",
		"personal",
		"pertinent",
		"pesky",
		"pessimistic",
		"petty",
		"phony",
		"physical",
		"piercing",
		"pink",
		"pitiful",
		"plain",
		"plaintive",
		"plastic",
		"playful",
		"pleasant",
		"pleased",
		"pleasing",
		"plump",
		"plush",
		"polished",
		"polite",
		"political",
		"pointed",
		"pointless",
		"poised",
		"poor",
		"popular",
		"portly",
		"posh",
		"positive",
		"possible",
		"potable",
		"powerful",
		"powerless",
		"practical",
		"precious",
		"present",
		"prestigious",
		"pretty",
		"precious",
		"previous",
		"pricey",
		"prickly",
		"primary",
		"prime",
		"pristine",
		"private",
		"prize",
		"probable",
		"productive",
		"profitable",
		"profuse",
		"proper",
		"proud",
		"prudent",
		"punctual",
		"pungent",
		"puny",
		"pure",
		"purple",
		"pushy",
		"putrid",
		"puzzled",
		"puzzling",
		"quaint",
		"qualified",
		"quarrelsome",
		"quarterly",
		"queasy",
		"querulous",
		"questionable",
		"quick",
		"quick-witted",
		"quiet",
		"quintessential",
		"quirky",
		"quixotic",
		"quizzical",
		"radiant",
		"ragged",
		"rapid",
		"rare",
		"rash",
		"raw",
		"recent",
		"reckless",
		"rectangular",
		"ready",
		"real",
		"realistic",
		"reasonable",
		"red",
		"reflecting",
		"regal",
		"regular",
		"reliable",
		"relieved",
		"remarkable",
		"remorseful",
		"remote",
		"repentant",
		"required",
		"respectful",
		"responsible",
		"repulsive",
		"revolving",
		"rewarding",
		"rich",
		"rigid",
		"right",
		"ringed",
		"ripe",
		"roasted",
		"robust",
		"rosy",
		"rotating",
		"rotten",
		"rough",
		"round",
		"rowdy",
		"royal",
		"rubbery",
		"rundown",
		"ruddy",
		"rude",
		"runny",
		"rural",
		"rusty",
		"sad",
		"safe",
		"salty",
		"same",
		"sandy",
		"sane",
		"sarcastic",
		"sardonic",
		"satisfied",
		"scaly",
		"scarce",
		"scared",
		"scary",
		"scented",
		"scholarly",
		"scientific",
		"scornful",
		"scratchy",
		"scrawny",
		"second",
		"secondary",
		"second-hand",
		"secret",
		"self-assured",
		"self-reliant",
		"selfish",
		"sentimental",
		"separate",
		"serene",
		"serious",
		"serpentine",
		"several",
		"severe",
		"shabby",
		"shadowy",
		"shady",
		"shallow",
		"shameful",
		"shameless",
		"sharp",
		"shimmering",
		"shiny",
		"shocked",
		"shocking",
		"shoddy",
		"short",
		"short-term",
		"showy",
		"shrill",
		"shy",
		"sick",
		"silent",
		"silky",
		"silly",
		"silver",
		"similar",
		"simple",
		"simplistic",
		"sinful",
		"single",
		"sizzling",
		"skeletal",
		"skinny",
		"sleepy",
		"slight",
		"slim",
		"slimy",
		"slippery",
		"slow",
		"slushy",
		"small",
		"smart",
		"smoggy",
		"smooth",
		"smug",
		"snappy",
		"snarling",
		"sneaky",
		"sniveling",
		"snoopy",
		"sociable",
		"soft",
		"soggy",
		"solid",
		"somber",
		"some",
		"spherical",
		"sophisticated",
		"sore",
		"sorrowful",
		"soulful",
		"soupy",
		"sour",
		"Spanish",
		"sparkling",
		"sparse",
		"specific",
		"spectacular",
		"speedy",
		"spicy",
		"spiffy",
		"spirited",
		"spiteful",
		"splendid",
		"spotless",
		"spotted",
		"spry",
		"square",
		"squeaky",
		"squiggly",
		"stable",
		"staid",
		"stained",
		"stale",
		"standard",
		"starchy",
		"stark",
		"starry",
		"steep",
		"sticky",
		"stiff",
		"stimulating",
		"stingy",
		"stormy",
		"straight",
		"strange",
		"steel",
		"strict",
		"strident",
		"striking",
		"striped",
		"strong",
		"studious",
		"stunning",
		"stupendous",
		"stupid",
		"sturdy",
		"stylish",
		"subdued",
		"submissive",
		"substantial",
		"subtle",
		"suburban",
		"sudden",
		"sugary",
		"sunny",
		"super",
		"superb",
		"superficial",
		"superior",
		"supportive",
		"sure-footed",
		"surprised",
		"suspicious",
		"svelte",
		"sweaty",
		"sweet",
		"sweltering",
		"swift",
		"sympathetic",
		"tall",
		"talkative",
		"tame",
		"tan",
		"tangible",
		"tart",
		"tasty",
		"tattered",
		"taut",
		"tedious",
		"teeming",
		"tempting",
		"tender",
		"tense",
		"tepid",
		"terrible",
		"terrific",
		"testy",
		"thankful",
		"that",
		"these",
		"thick",
		"thin",
		"third",
		"thirsty",
		"this",
		"thorough",
		"thorny",
		"those",
		"thoughtful",
		"threadbare",
		"thrifty",
		"thunderous",
		"tidy",
		"tight",
		"timely",
		"tinted",
		"tiny",
		"tired",
		"torn",
		"total",
		"tough",
		"traumatic",
		"treasured",
		"tremendous",
		"tragic",
		"trained",
		"tremendous",
		"triangular",
		"tricky",
		"trifling",
		"trim",
		"trivial",
		"troubled",
		"true",
		"trusting",
		"trustworthy",
		"trusty",
		"truthful",
		"tubby",
		"turbulent",
		"twin",
		"ugly",
		"ultimate",
		"unacceptable",
		"unaware",
		"uncomfortable",
		"uncommon",
		"unconscious",
		"understated",
		"unequaled",
		"uneven",
		"unfinished",
		"unfit",
		"unfolded",
		"unfortunate",
		"unhappy",
		"unhealthy",
		"uniform",
		"unimportant",
		"unique",
		"united",
		"unkempt",
		"unknown",
		"unlawful",
		"unlined",
		"unlucky",
		"unnatural",
		"unpleasant",
		"unrealistic",
		"unripe",
		"unruly",
		"unselfish",
		"unsightly",
		"unsteady",
		"unsung",
		"untidy",
		"untimely",
		"untried",
		"untrue",
		"unused",
		"unusual",
		"unwelcome",
		"unwieldy",
		"unwilling",
		"unwitting",
		"unwritten",
		"upbeat",
		"upright",
		"upset",
		"urban",
		"usable",
		"used",
		"useful",
		"useless",
		"utilized",
		"utter",
		"vacant",
		"vague",
		"vain",
		"valid",
		"valuable",
		"vapid",
		"variable",
		"vast",
		"velvety",
		"venerated",
		"vengeful",
		"verifiable",
		"vibrant",
		"vicious",
		"victorious",
		"vigilant",
		"vigorous",
		"villainous",
		"violet",
		"violent",
		"virtual",
		"virtuous",
		"visible",
		"vital",
		"vivacious",
		"vivid",
		"voluminous",
		"wan",
		"warlike",
		"warm",
		"warmhearted",
		"warped",
		"wary",
		"wasteful",
		"watchful",
		"waterlogged",
		"watery",
		"wavy",
		"wealthy",
		"weak",
		"weary",
		"webbed",
		"wee",
		"weekly",
		"weepy",
		"weighty",
		"weird",
		"welcome",
		"well-documented",
		"well-groomed",
		"well-informed",
		"well-lit",
		"well-made",
		"well-off",
		"well-to-do",
		"well-worn",
		"wet",
		"which",
		"whimsical",
		"whirlwind",
		"whispered",
		"white",
		"whole",
		"whopping",
		"wicked",
		"wide",
		"wide-eyed",
		"wiggly",
		"wild",
		"willing",
		"wilted",
		"winding",
		"windy",
		"winged",
		"wiry",
		"wise",
		"witty",
		"wobbly",
		"woeful",
		"wonderful",
		"wooden",
		"woozy",
		"wordy",
		"worldly",
		"worn",
		"worried",
		"worrisome",
		"worse",
		"worst",
		"worthless",
		"worthwhile",
		"worthy",
		"wrathful",
		"wretched",
		"writhing",
		"wrong",
		"wry",
		"yawning",
		"yearly",
		"yellow",
		"yellowish",
		"young",
		"youthful",
		"yummy",
		"zany",
		"zealous",
		"zesty",
		"zigzag",
	}

	name = [...]string{
		"abel",
		"abelard",
		"abraham",
		"achaios",
		"acis",
		"addison",
		"adela",
		"adelaide",
		"admiranda",
		"adonis",
		"aeditha",
		"aegipan",
		"aelina",
		"agnys",
		"aigis",
		"aigyptos",
		"aiolides",
		"aion",
		"aisa",
		"aisakos",
		"aithilla",
		"aithon",
		"aitne",
		"akakos",
		"alainne",
		"alaire",
		"albin",
		"aldebrand",
		"aldous",
		"aleyn",
		"alianore",
		"alison",
		"alistair",
		"alkmene",
		"alyne",
		"alys",
		"amelia",
		"amice",
		"ampelos",
		"amphelice",
		"anaxibia",
		"ancelot",
		"angelet",
		"anius",
		"anna",
		"annabel",
		"anne",
		"anselm",
		"anthoinette",
		"antigone",
		"anys",
		"apemosyne",
		"arabella",
		"aran",
		"archedios",
		"argo",
		"arkeisios",
		"arlette",
		"arnald",
		"arnott",
		"arthur",
		"askalabos",
		"atilda",
		"atropos",
		"atys",
		"aubrey",
		"audrye",
		"augeias",
		"augustine",
		"auson",
		"ava",
		"avelin",
		"avelyn",
		"averil",
		"ayleth",
		"aylmer",
		"bacchus",
		"baderon",
		"bakis",
		"baldric",
		"bardolf",
		"bartholomew",
		"baterich",
		"bathsua",
		"bayard",
		"beatrix",
		"bellinda",
		"belmont",
		"belos",
		"benedict",
		"beneger",
		"berekyntia",
		"bernard",
		"berndan",
		"bertana",
		"berte",
		"bertram",
		"bertrand",
		"bess",
		"blackburn",
		"blavier",
		"bormos",
		"bouchard",
		"boyle",
		"bran",
		"brangwine",
		"braya",
		"brice",
		"brien",
		"bromios",
		"brontes",
		"bruce",
		"brunhild",
		"bryce",
		"bryde",
		"bukolos",
		"caesaria",
		"cain",
		"cameron",
		"camers",
		"caplan",
		"carmen",
		"carmine",
		"carna",
		"casandra",
		"caspar",
		"catillus",
		"ceadda",
		"cecilia",
		"cecily",
		"celeste",
		"celestine",
		"celestria",
		"cenota",
		"chamberlain",
		"charis",
		"charlys",
		"charmaine",
		"chartain",
		"chesias",
		"chloe",
		"christabel",
		"chryses",
		"cicely",
		"clarimond",
		"claudia",
		"claudien",
		"clemence",
		"clifton",
		"clive",
		"cole",
		"collys",
		"colson",
		"comet",
		"concessa",
		"conphas",
		"constance",
		"cornelia",
		"cornell",
		"coster",
		"crestian",
		"cristiana",
		"cutbert",
		"cuthbert",
		"cwengyth",
		"cybele",
		"cyndra",
		"cynewyn",
		"cyriac",
		"daimbert",
		"dalmas",
		"damaris",
		"dametta",
		"damia",
		"danae",
		"danyell",
		"dardanos",
		"dauid",
		"davyd",
		"dawson",
		"decima",
		"deianeira",
		"deidameia",
		"deimachos",
		"deimos",
		"deitrich",
		"dekelos",
		"deloys",
		"delphos",
		"denston",
		"denys",
		"derkynos",
		"derwin",
		"deryk",
		"diamanda",
		"dionisia",
		"dodona",
		"dominy",
		"donner",
		"dorcas",
		"dorothe",
		"drake",
		"drew",
		"dryope",
		"drystan",
		"durilda",
		"dwyvaer",
		"dyana",
		"dysaules",
		"eadbert",
		"ealdwine",
		"echetlos",
		"echo",
		"edelinne",
		"edithe",
		"edmund",
		"edwyn",
		"eidothea",
		"eilonwy",
		"elaisse",
		"elatus",
		"eldred",
		"ele",
		"eleanor",
		"eleazar",
		"elewys",
		"ellerete",
		"ellie",
		"elpenor",
		"elsebee",
		"elyn",
		"elynor",
		"elyzabeth",
		"emanuel",
		"emblyn",
		"emeline",
		"emeny",
		"emeria",
		"emerick",
		"emery",
		"emilie",
		"emlinie",
		"emmet",
		"enipeus",
		"epigonoi",
		"epione",
		"erasmus",
		"erato",
		"erebos",
		"erik",
		"eschina",
		"eschiva",
		"esdeline",
		"esmenet",
		"esmond",
		"esmour",
		"esperaunce",
		"estienne",
		"estrild",
		"etgar",
		"ethelbert",
		"ethelia",
		"ethelred",
		"euadne",
		"euchenor",
		"euenos",
		"eugenia",
		"eunomos",
		"eupalamos",
		"euphorbos",
		"europe",
		"eustace",
		"eustacia",
		"eva",
		"evelyn",
		"fames",
		"fauna",
		"fawkes",
		"felice",
		"fiebras",
		"flambard",
		"florens",
		"folke",
		"foxe",
		"frances",
		"francis",
		"francisca",
		"frederick",
		"frederyk",
		"frideswide",
		"fridgia",
		"fulke",
		"galateia",
		"galeos",
		"galfrid",
		"ganelon",
		"gared",
		"gauwyn",
		"gaynor",
		"gembert",
		"geoffrey",
		"gerald",
		"gerbold",
		"gerhardt",
		"gerland",
		"germainne",
		"gethrude",
		"gillian",
		"giselle",
		"glauke",
		"glenda",
		"gloriana",
		"goddard",
		"godebert",
		"godfrey",
		"gregory",
		"grimbald",
		"grups",
		"gryffen",
		"guinevere",
		"guston",
		"gwayne",
		"gyes",
		"gygas",
		"gylbart",
		"gylda",
		"gyles",
		"habreham",
		"hadrian",
		"haimirich",
		"halia",
		"halisera",
		"halstein",
		"hamon",
		"heinlein",
		"helena",
		"helenor",
		"helias",
		"helios",
		"helvynya",
		"hemithea",
		"hepaklos",
		"herkyna",
		"hester",
		"hewrey",
		"hilda",
		"hildegard",
		"hilith",
		"hippotes",
		"hopladamos",
		"huaina",
		"humphrey",
		"hylas",
		"iamos",
		"ianthe",
		"ilos",
		"imedia",
		"inferi",
		"ingham",
		"ingram",
		"inuus",
		"iobes",
		"iphis",
		"irae",
		"irus",
		"isabella",
		"ischys",
		"isemeine",
		"isleton",
		"ismay",
		"ismenia",
		"isolde",
		"isyrion",
		"ivan",
		"jaane",
		"jacquette",
		"jakys",
		"janus",
		"jeanne",
		"jeger",
		"jellion",
		"jemime",
		"jenet",
		"jenlyns",
		"jenyfer",
		"jessamine",
		"jillian",
		"jocea",
		"jocelyn",
		"johannes",
		"joleicia",
		"jolline",
		"jonathas",
		"joseph",
		"josephine",
		"josian",
		"josiane",
		"josias",
		"joyce",
		"joyse",
		"judithe",
		"judye",
		"juliana",
		"julyan",
		"june",
		"junk",
		"jupiter",
		"justina",
		"justitia",
		"kain",
		"kampe",
		"kapys",
		"katelyn",
		"kath",
		"katherine",
		"katrine",
		"kaukon",
		"kaunos",
		"kelmis",
		"kennard",
		"kenrick",
		"kephalos",
		"kerrich",
		"khellus",
		"kilix",
		"kimball",
		"kinborow",
		"kinnison",
		"klaros",
		"kleobis",
		"kranaos",
		"kurtz",
		"kyknos",
		"kyzikos",
		"ladislas",
		"laios",
		"lambert",
		"lampetos",
		"laodameia",
		"laodike",
		"lapithes",
		"lars",
		"latinos",
		"latisha",
		"latona",
		"lauda",
		"laurence",
		"laurentius",
		"lausus",
		"laverna",
		"leavold",
		"lefwyne",
		"leimone",
		"leipephile",
		"lennard",
		"leofwen",
		"leofwynn",
		"leopold",
		"letita",
		"lettice",
		"leuke",
		"leukippe",
		"leukon",
		"linette",
		"linos",
		"linyeve",
		"littlejohn",
		"lityerses",
		"llawran",
		"lloyd",
		"lodwicke",
		"lora",
		"lowell",
		"lykeios",
		"lykomedes",
		"lykophron",
		"lykurgos",
		"lynkos",
		"lysippe",
		"machaon",
		"maddeline",
		"madison",
		"maerwynn",
		"maiandros",
		"mainfroi",
		"maisenta",
		"makaria",
		"male names",
		"malin",
		"mansel",
		"margarete",
		"margeria",
		"margry",
		"maria",
		"maronne",
		"mars",
		"marsilia",
		"martine",
		"mathild",
		"mathye",
		"mavors",
		"megareus",
		"melaineus",
		"melampus",
		"melodie",
		"melusine",
		"memphis",
		"menestheus",
		"meredithe",
		"merewyn",
		"merilda",
		"merops",
		"meryell",
		"mestor",
		"metaneira",
		"metis",
		"metope",
		"millicent",
		"minerva",
		"minos",
		"minyas",
		"mirabelle",
		"misericordia",
		"mnestra",
		"molossos",
		"monster names",
		"morgant",
		"morgayne",
		"morpheus",
		"morys",
		"mulciber",
		"muriel",
		"murienne",
		"musa",
		"mydrede",
		"mykenai",
		"myles",
		"myrine",
		"myrto",
		"nathaniel",
		"nausithos",
		"nautes",
		"navarre",
		"neaira",
		"neale",
		"neilos",
		"nemea",
		"nessos",
		"nesta",
		"nicholina",
		"nicia",
		"nicolaa",
		"nireus",
		"noes",
		"nomios",
		"norman",
		"nyx",
		"ogaphos",
		"ogygos",
		"oiax",
		"oibalos",
		"oinomaos",
		"olyffe",
		"olyver",
		"ophellia",
		"ophis",
		"orrick",
		"orthaia",
		"orwen",
		"osric",
		"oswyn",
		"ottilia",
		"owyne",
		"oxylos",
		"paige",
		"pallene",
		"parnell",
		"parnella",
		"pasiphae",
		"patrick",
		"paul",
		"pedasos",
		"peirene",
		"pelias",
		"pelinne",
		"penates",
		"penelope",
		"penia",
		"penthesileia",
		"percival",
		"peripanos",
		"persephone",
		"perseus",
		"peter",
		"petronilla",
		"phanes",
		"pheraia",
		"philippe",
		"philyra",
		"phlegrai",
		"phobos",
		"phrasios",
		"phrixos",
		"phthonos",
		"pieria",
		"piers",
		"pisos",
		"pitane",
		"pittheus",
		"placencia",
		"poine",
		"polybos",
		"polydamna",
		"polykaon",
		"polyxo",
		"portheus",
		"potitii",
		"powle",
		"priamos",
		"prokne",
		"prokris",
		"proteus",
		"prothoos",
		"prudence",
		"pulmia",
		"purnell",
		"radcliffe",
		"radolf",
		"raffe",
		"randall",
		"randwulf",
		"rauffe",
		"raulin",
		"rebeccah",
		"redwald",
		"reeve",
		"reginald",
		"reinholdt",
		"remnus",
		"reynard",
		"reyner",
		"reynfred",
		"rhadamanthys",
		"rhadine",
		"rhakios",
		"rhea",
		"rianna",
		"ricard",
		"richarde",
		"richenda",
		"rickeman",
		"ridel",
		"robert",
		"robyn",
		"roger",
		"rolfe",
		"ronald",
		"rosa",
		"rosalind",
		"rosamund",
		"rose",
		"roundelph",
		"rowland",
		"roysia",
		"rychyld",
		"salamis",
		"salios",
		"samantha",
		"samson",
		"sanche",
		"sandre",
		"sarra",
		"satyros",
		"scarlet",
		"selphina",
		"semele",
		"sence",
		"serendipity",
		"sevrin",
		"sibyl",
		"sighard",
		"sigurdh",
		"sikyon",
		"silvanus",
		"simond",
		"singleton",
		"sinope",
		"sisyphos",
		"sithon",
		"sol",
		"solyeuse",
		"somerhild",
		"spenser",
		"sreda",
		"stewart",
		"swift",
		"sybaris",
		"sybell",
		"sylphie",
		"syme",
		"symon",
		"symond",
		"syndony",
		"sysley",
		"systeleley",
		"talaos",
		"tansa",
		"tantalos",
		"taran",
		"tatius",
		"taylor",
		"telephassa",
		"temperance",
		"templeton",
		"tenes",
		"teukros",
		"thaeox",
		"thamyris",
		"theda",
		"thelxion",
		"theodore",
		"theophane",
		"theresa",
		"thespis",
		"thoas",
		"thomas",
		"thomasine",
		"thomasyn",
		"thora",
		"thrydwulf",
		"thyia",
		"timothy",
		"tiphina",
		"tristan",
		"tristana",
		"turstin",
		"tyche",
		"typhon",
		"ulric",
		"uranos",
		"ursula",
		"valentyne",
		"vannes",
		"vector",
		"vesta",
		"victor",
		"voyce",
		"vrsela",
		"vyncent",
		"wadard",
		"walter",
		"warin",
		"wauter",
		"wenyld",
		"werner",
		"wilfrid",
		"wilham",
		"willielmus",
		"willmott",
		"wineburg",
		"wolfstan",
		"wulfhilda",
		"wyatt",
		"wymon",
		"wymond",
		"wynefreede",
		"wystan",
		"yedythe",
		"ysabel",
		"ysmeina",
		"ywain",
		"zacheus",
		"zell",
		"zephyrus",
		"zerig",
	}
)

func getRandomName() (codeword string) {
	rand.Seed(time.Now().UnixNano())
	codeword = fmt.Sprintf("%s %s", adj[rand.Intn(len(adj))], name[rand.Intn(len(name))])
	return
}
