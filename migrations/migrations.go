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
		&models.News{},
		&models.Home{},
		&models.Socials{},
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
			InFuture:     true,
		},
		&models.Projects{
			ID:    3,
			Title: "Магическая битва",
			Description: "Маги - обычные люди," +
				" обладающие сверхъестественной силой, недоступной человеку. На этом и строится магический мир, такой двоякий и нестабильный, однако намного легче, если есть кому нас защитить. Не так давно баланс был нарушен, когда тысячелетнее проклятие обрело сосуд." +
				" Рёмен Сукуна захватил тело обычного подростка Итадори Юджи. Но, постойте ка, кажется тот его может с лёгкостью подавлять, но всё же остальные маги его опасаются, хотят уничтожить, что неоднократно заставляет попадать Итадори в неприятности. А теперь, не стой на месте! " +
				"С нами ты ощутишь полное погружение в мир магии, так что вместе отправляемся в это незабываемое приключение. Ты готов?",
			Image:        "https://storage.yandexcloud.net/zetrego/series/preview/thumb-jujutsu-kaisen-cursed-clash.jpg",
			VideoPreview: "https://storage.yandexcloud.net/zetrego/homevideos/jujutsukaisen.mp4",
			InFuture:     true,
		},

		&models.Season{
			ID:         1,
			ProjectsID: 1,
			Number:     1,
		},

		&models.Video{
			SeasonID:     1,
			ProjectsID:   1,
			SeriesNumber: 0,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego2/S0E0Hazbin/S0E0_Hazbin~1.m3u8",
		},

		&models.Video{
			SeasonID:     1,
			ProjectsID:   1,
			SeriesNumber: 1,
			SeriesURL:    "https://storage.yandexcloud.net/zetrego2/S0E0Hazbin/S0E0_Hazbin~2.m3u8",
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

		&models.Socials{
			Title:       "🥰 Это - наш главный tg канал. Мы его очень любим :) 🥰",
			Description: "Там вы сможете найти наш уютный чатик, канал с приколами и нашими переозвучками, новости, а также канал с нашими Проектами",
			Url:         "https://t.me/ZetregoTeam",
			LogoUrl:     "https://storage.yandexcloud.net/zetrego/socials/telegram.png",
			ImageUrl:    "https://storage.yandexcloud.net/zetrego/about_back1.jpg",
			Href:        "tg",
		},
		&models.Socials{
			Title:       "😋 Youtube - канал одного из основателей. Там мы стараемся выкладывать интересные Shorts и другой контент 😋",
			Description: "Подписывайтесь, чтобы быть в крусе новостей!",
			Url:         "https://www.youtube.com/channel/UCfVELs4Npz1DKTpBfhKndtw",
			LogoUrl:     "https://storage.yandexcloud.net/zetrego/socials/youtube.png",
			ImageUrl:    "https://storage.yandexcloud.net/zetrego/about_back2.jpg",
			Href:        "youtube",
		},
		&models.Socials{
			Title:       "🔥 Основная и единственная группа во Вконтакте. Там выкладываются новости, видео и ВК Клипы 🔥",
			Description: "( правда хотели бы вести группу почаще 😅 )",
			Url:         "https://vk.com/zetregoteam",
			LogoUrl:     "https://storage.yandexcloud.net/zetrego/socials/vk.png",
			ImageUrl:    "https://storage.yandexcloud.net/zetrego/about_back3.jpg",
			Href:        "vk",
		},
		&models.Socials{
			Title:       "⭐ Основной TikTok нашей команд! Мемы, новости, шутки и приколы там! ⭐",
			Description: "Обещаем вести его побольше!",
			Url:         "https://www.tiktok.com/@sandministr",
			LogoUrl:     "https://storage.yandexcloud.net/zetrego/socials/tiktok.png",
			ImageUrl:    "https://storage.yandexcloud.net/zetrego/about_back4.jpg",
			Href:        "tiktok",
		},

		&models.News{
			TitleDate: "21.04.24",
			Description: "Исправлены баги отображения страниц на мобильных устройствах \n " +
				"Оптимизирована работа плеера. Теперь плеер грузится гораздо быстрее и \n" +
				"потребляет трафик по мере просмотра, а не выкачивает всё видео сразу.",
			Image: "",
		},

		&models.News{
			TitleDate: "25.04.24",
			Description: "🌟 Хочешь стать частью нашей удивительной команды " +
				"и  начать озвучивать аниме/мультсериалы? Отправь нам свои " +
				"работы на кастинг! Просто перейди по ссылке и заполни на нашу Google Форму!" +
				"Мы постоянно находимся в поиске интересных голосов и качественных работ на постоянную основу! \n  " +
				"https://forms.gle/THDTEZSUENDW413j7 \n " + "Мы с нетерпением ждем твоих талантов и идей! ✨",
			Image: "https://storage.yandexcloud.net/zetrego/news/alastor_news.jpg",
		},
		&models.News{
			TitleDate:   "28.04.24",
			Description: "Большое обновление сайта! Улучшена поисковая оптимизация. \n Добавлена лента новостей на основную страницу \n Добавлен выбор сезонов для аниме и мультсериалов",
			Image:       "",
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
