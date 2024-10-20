package migrations

import (
	"austem/models"
	"austem/postreql/db"
	"fmt"
)

func init() {
	db.StartDB()
}

func Migrator() {
	err := db.DB.AutoMigrate(
		&models.Projects{},
		&models.Season{},
		&models.Home{},
		&models.Team{},
		&models.Video{})

	initialData := []interface{}{
		&models.Projects{
			ID:           1,
			Title:        "Отель Хазбин",
			Description:  "Отель Хазбин - Всеми нами любимый мультсериал. Чарли - дочь могущественного люцифера захотела открыть свой отель. Но получится ли у неё...",
			Image:        "https://storage.yandexcloud.net/zetrego/series/preview/2a05d6895072f43d13c08f1840d92a79.png",
			VideoPreview: "https://storage.yandexcloud.net/zetrego/homevideos/full_alastor_vid.mp4",
			InFuture:     false,
		},
		&models.Projects{
			ID:           2,
			Title:        "Клинок рассекающий демонов",
			Description:  "Танзиро - обычный мальчик, которого постигла нелегкая судьба. В нелегкое для себя время, он должен защитить свою сестренку. В будущем хотели бы обязательно озвучить!",
			Image:        "https://storage.yandexcloud.net/zetrego/series/preview/d8c98b80ee7cd3400685e59dfab01735.jpg",
			VideoPreview: "https://storage.yandexcloud.net/zetrego/homevideos/Demon%20Slayer%20_%20OP%20_%20_Gurenge_%20by%20LiSA%20HD.mp4",
			InFuture:     false,
		},
		&models.Projects{
			ID:    3,
			Title: "Магическая битва",
			Description: "Маги - обычные люди," +
				" обладающие сверхъестественной силой, недоступной человеку. На этом и строится магический мир, такой двоякий и нестабильный, однако намного легче, если есть кому нас защитить. Не так давно баланс был нарушен, когда тысячелетнее проклятие обрело сосуд." +
				" Рёмен Сукуна захватил тело обычного подростка Итадори Юджи. Ты готов?",
			Image:        "https://storage.yandexcloud.net/zetrego/series/preview/thumb-jujutsu-kaisen-cursed-clash.jpg",
			VideoPreview: "https://storage.yandexcloud.net/zetrego/homevideos/jujutsukaisen.mp4",
			InFuture:     false,
		},

		&models.Projects{
			ID:    4,
			Title: "Разрушитель ветра",
			Description: "Харука Сакура терпеть не может слабаков. Именно поэтому он поступил в старшую «Фурин», заработавшую себе " +
				"репутацию школы с боями без правил, где ученики каждый день дерутся по причине и без.",
			Image:        "https://storage.yandexcloud.net/zetrego/series/preview/wind-breaker-anime-haruka-sakura-387%403%40a-thumb.jpg",
			VideoPreview: "https://storage.yandexcloud.net/zetrego/homevideos/windbreakertrailer2.mp4",
			InFuture:     false,
		},

		&models.Projects{
			ID:           5,
			Title:        "Лакадейзи",
			Description:  "Действие происходит во времена сухого закона в 1927 году в городе Сент-Луис штата Миссури, населённом антропоморфными животными. История повествует о судьбе нелегального бара «Лакадейзи» и одноименной банды после убийства её основателя.",
			Image:        "https://storage.yandexcloud.net/zetrego/series/preview/lackadaisypreview.jpg",
			VideoPreview: "https://storage.yandexcloud.net/zetrego/homevideos/lackadaisy.mp4",
			InFuture:     false,
		},

		&models.Season{
			ID:         1,
			ProjectsID: 1,
			Number:     1,
		},
		&models.Season{
			ID:         4,
			ProjectsID: 2,
			Number:     4,
		},

		&models.Season{
			ID:         2,
			ProjectsID: 3,
			Number:     1,
		},

		&models.Season{
			ID:         3,
			ProjectsID: 4,
			Number:     1,
		},

		&models.Season{
			ID:         5,
			ProjectsID: 5,
			Number:     1,
		},

		&models.Video{
			SeasonID:     1,
			ProjectsID:   1,
			SeriesNumber: 0,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/hazbinhotel/s1e0/S0E0_Hazbin1.m3u8",
		},

		&models.Video{
			SeasonID:     1,
			ProjectsID:   1,
			SeriesNumber: 1,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/hazbinhotel/s1e1/S0E0_Hazbin2.m3u8",
		},

		&models.Video{
			SeasonID:     1,
			ProjectsID:   1,
			SeriesNumber: 2,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/hazbinhotel/s1e2/S0E0_Hazbin3.m3u8",
		},
		&models.Video{
			SeasonID:     1,
			ProjectsID:   1,
			SeriesNumber: 3,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/hazbinhotel/s1e4/S0E0_Hazbin~4.m3u8",
		},

		&models.Video{
			SeasonID:     4,
			ProjectsID:   2,
			SeriesNumber: 1,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/demonslayer/4/1/demon_slayer1.m3u8",
		},

		&models.Video{
			SeasonID:     2,
			ProjectsID:   3,
			SeriesNumber: 1,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/jujutsukaisen/1/1/jujutsu11m.m3u8",
		},

		&models.Video{
			SeasonID:     2,
			ProjectsID:   3,
			SeriesNumber: 2,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/jujutsukaisen/1/2/jujutsu12.m3u8",
		},

		&models.Video{
			SeasonID:     2,
			ProjectsID:   3,
			SeriesNumber: 3,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/jujutsukaisen/1/3/jujutsu13.m3u8",
		},

		&models.Video{
			SeasonID:     4,
			ProjectsID:   2,
			SeriesNumber: 2,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/demonslayer/4/2/demonslayer42.m3u8",
		},

		&models.Video{
			SeasonID:     4,
			ProjectsID:   2,
			SeriesNumber: 3,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/demonslayer/4/3/demonslayer43.m3u8",
		},

		&models.Video{
			SeasonID:     4,
			ProjectsID:   2,
			SeriesNumber: 4,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/demonslayer/4/5/demon_slayer5_final.m3u8",
		},

		&models.Video{
			SeasonID:     4,
			ProjectsID:   2,
			SeriesNumber: 5,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/demonslayer/4/4/demonslayer44.m3u8",
		},

		&models.Video{
			SeasonID:     3,
			ProjectsID:   4,
			SeriesNumber: 1,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/windbreaker/1/1/wb11.m3u8",
		},

		&models.Video{
			SeasonID:     5,
			ProjectsID:   5,
			SeriesNumber: 0,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego/anime/lackadaisy/pilot/lackadaisypilot.m3u8",
		},

		&models.Home{
			BackgroundURL: "https://storage.yandexcloud.net/zetrego/homevideos/jujutsukaisen.mp4",
		},

		&models.Home{
			BackgroundURL: "https://storage.yandexcloud.net/zetrego/homevideos/Demon%20Slayer%20_%20OP%20_%20_Gurenge_%20by%20LiSA%20HD.mp4",
		},

		&models.Home{
			BackgroundURL: "https://storage.yandexcloud.net/zetrego/homevideos/full_alastor_vid.mp4",
		},

		&models.Home{
			BackgroundURL: "https://storage.yandexcloud.net/zetrego/homevideos/windbreakertrailer2.mp4",
		},
	}

	for _, data := range initialData {
		if err := db.DB.Create(data).Error; err != nil {
			fmt.Println("Failed to create initial data:", err)
		}
	}

	if err != nil {
		return
	}
	fmt.Println("migrations complete")
}
