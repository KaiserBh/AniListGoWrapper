# anilist

```go
import "AniListGoWrapper/anilist"
```

## Index

- [func CleanMediaTrendPageJSON(str []byte) []byte](<#func-cleanmediatrendpagejson>)
- [func CleanPageJSON(str []byte) []byte](<#func-cleanpagejson>)
- [func CleanStaffJSON(str []byte) []byte](<#func-cleanstaffjson>)
- [func PostRequest(jsonValue []byte) []byte](<#func-postrequest>)
- [type AiringSchedule](<#type-airingschedule>)
- [type CharacterMedia](<#type-charactermedia>)
- [type Characters](<#type-characters>)
- [type CoverImage](<#type-coverimage>)
- [type Edges](<#type-edges>)
- [type EndDate](<#type-enddate>)
- [type ExternalLinks](<#type-externallinks>)
- [type Image](<#type-image>)
- [type Media](<#type-media>)
  - [func NewMedia() *Media](<#func-newmedia>)
  - [func (m *Media) FilterAnimeByID(id int)](<#func-media-filteranimebyid>)
  - [func (m *Media) FilterByEndDate(date int32)](<#func-media-filterbyenddate>)
  - [func (m *Media) FilterByID(id int)](<#func-media-filterbyid>)
  - [func (m *Media) FilterByMalID(malID int)](<#func-media-filterbymalid>)
  - [func (m *Media) FilterBySeason(season string)](<#func-media-filterbyseason>)
  - [func (m *Media) FilterByStartDate(date int32)](<#func-media-filterbystartdate>)
  - [func (m *Media) FilterByTitle(title string)](<#func-media-filterbytitle>)
  - [func (m *Media) FilterMangaByID(id int)](<#func-media-filtermangabyid>)
- [type MediaTrend](<#type-mediatrend>)
  - [func NewMediaTrend() *MediaTrend](<#func-newmediatrend>)
  - [func (m *MediaTrend) FilterByEpisode(episode int)](<#func-mediatrend-filterbyepisode>)
  - [func (m *MediaTrend) FilterByPopularity(popularity int)](<#func-mediatrend-filterbypopularity>)
  - [func (m *MediaTrend) FilterByTrendingAmount(trending int)](<#func-mediatrend-filterbytrendingamount>)
  - [func (m *MediaTrend) FilterByTrendingAverageScore(averageScore int)](<#func-mediatrend-filterbytrendingaveragescore>)
  - [func (m *MediaTrend) SearchByMediaID(id int)](<#func-mediatrend-searchbymediaid>)
- [type Name](<#type-name>)
- [type NextAiringEpisode](<#type-nextairingepisode>)
- [type Node](<#type-node>)
- [type Page](<#type-page>)
  - [func (p *Page) PaginationByTitle(title string, page int, perPage int)](<#func-page-paginationbytitle>)
- [type PageInfo](<#type-pageinfo>)
- [type Rankings](<#type-rankings>)
- [type Recommendations](<#type-recommendations>)
- [type RecroRevEdges](<#type-recrorevedges>)
- [type Relations](<#type-relations>)
- [type Reviews](<#type-reviews>)
- [type ScoreDistribution](<#type-scoredistribution>)
- [type Staff](<#type-staff>)
  - [func (s *Staff) FilterStaffByID(id int)](<#func-staff-filterstaffbyid>)
  - [func (s *Staff) FilterStaffByName(name string)](<#func-staff-filterstaffbyname>)
- [type StaffMedia](<#type-staffmedia>)
- [type StartDate](<#type-startdate>)
- [type Stats](<#type-stats>)
- [type StatusDistribution](<#type-statusdistribution>)
- [type StreamingEpisodes](<#type-streamingepisodes>)
- [type Studio](<#type-studio>)
- [type Tags](<#type-tags>)
- [type Title](<#type-title>)
- [type Trailer](<#type-trailer>)


## func CleanMediaTrendPageJSON

```go
func CleanMediaTrendPageJSON(str []byte) []byte
```

CleanMediaTrendPageJSON cleans the hmtl body by removing data and \}\}

## func CleanPageJSON

```go
func CleanPageJSON(str []byte) []byte
```

CleanPageJSON cleans the hmtl body by removing data and \}\}

## func CleanStaffJSON

```go
func CleanStaffJSON(str []byte) []byte
```

CleanStaffJSON cleans the hmtl body by removing data and \}\}

## func PostRequest

```go
func PostRequest(jsonValue []byte) []byte
```

PostRequest sends POST request that takes \[\]byte as parameter then returns back httml Body as \[\]byte

## type AiringSchedule

AiringSchedule Anime airing schedule if it's provided uses Edges and ID

```go
type AiringSchedule struct {
    Edges []Edges `json:"edges"`
}
```

## type CharacterMedia

CharacterMedia Media the actor voiced characters in\. \(Same data as characters with media as node instead of characters\)

```go
type CharacterMedia struct {
    Edges []Edges `json:"edges"`
}
```

## type Characters

Characters Characters ID must use Edges to call it and it's an array\.

```go
type Characters struct {
    Edges []Edges `json:"edges"`
}
```

## type CoverImage

CoverImage of the anime ExtraLarge\, Large\, Medium\, and Color for some reason\.

```go
type CoverImage struct {
    ExtraLarge string `json:"extraLarge"`
    Large      string `json:"large"`
    Medium     string `json:"medium"`
    Color      string `json:"color"`
}
```

## type Edges

Edges some data have edges so need a different one\.

```go
type Edges struct {
    ID int `json:"id"`
}
```

## type EndDate

EndDate object to store json and endDate of the anime

```go
type EndDate struct {
    Year  int `json:"year"`
    Month int `json:"month"`
    Day   int `json:"day"`
}
```

## type ExternalLinks

ExternalLinks for the anime\, uses ID

```go
type ExternalLinks struct {
    ID int `json:"id"`
}
```

## type Image

Image of the staff Large or Medium

```go
type Image struct {
    Large  string `json:"large"`
    Medium string `json:"medium"`
}
```

## type Media

Media object to store the json from Anilist

```go
type Media struct {
    // Media ID from anilist
    ID  int `json:"id"`
    // Media ID from mal (MyAnimeList not always accurate with the new animes or upcoming ones)
    IDMAL int `json:"idMal"`
    // Title of the Media
    Title Title `json:"title"`
    // MediaType such as MANGA, ANIME
    MediaType string `json:"type"`
    // MediaFormat such as TV, TV_SHORT (under 15 min), MOVIE, SPECIAL, OVA, ONA, MUSIC, MANGA, NOVEL, ONE_SHOT
    MediaFormat string `json:"format"`
    // Media Status such as FINISHED, RELEASING, NOT_YET_RELEASED, CANCELLED and HIATUS
    Status string `json:"status"`
    // Description Short description of the media's story and characters
    Description string `json:"description"`
    // StartDate The first official release date of the media
    StartDate StartDate `json:"startDate"`
    // EndDate The last official release date of the media
    EndDate EndDate `json:"endDate"`
    // Season The season the media was initially released in (Winter Months December to February, SPRING Months March to May, SUMMER Months June to August and FALL Months September to November)
    Season string `json:"season"`
    // The season year the media was initially released in
    SeasonYear int `json:"seasonYear"`
    // The year & season the media was initially released in
    SeasonInt int `json:"seasonInt"`
    // The amount of episodes the anime has when complete
    Episodes int `json:"episodes"`
    // The general length of each anime episode in minutes
    Duration int `json:"duration"`
    // The amount of chapters the manga has when complete
    Chapters int `json:"chapters"`
    // The amount of volumes the manga has when complete
    Volumes int `json:"volumes"`
    // Where the media was created. (ISO 3166-1 alpha-2)
    CountryOfOrigin string `json:"countryOfOrigin"`
    // If the media is officially licensed or a self-published doujin release
    IsLicensed bool `json:"isLicensed"`
    // Source type the media was adapted from. (ORIGINAL, MANGA, LIGHT_NOVEL, VISUAL_NOVEL, VIDEO_GAME, OTHER, NOVEL, DOUJINSHI, ANIME )
    Source string `json:"source"`
    // Official Twitter hashtags for the media
    HashTag string `json:"hashtag"`
    // Media trailer or advertisement unique YouTube string
    Trailer Trailer `json:"trailer"`
    // When the media's data was last updated
    UpdatedAt int `json:"updatedAt"`
    // The cover images of the media
    CoverImage CoverImage `json:"coverImage"`
    // The banner image of the media
    BannerImage string `json:"bannerImage"`
    // The genres of the media
    Genres []string `json:"genres"`
    // Alternative titles of the media
    Synonyms []string `json:"synonyms"`
    // media's average score
    AverageScore int `json:"averageScore"`
    // Mean score of all the user's scores of the media
    MeanScore int `json:"meanScore"`
    // Popularity of the media
    Popularity int `json:"popularity"`
    // Locked media may not be added to lists our favorited. This may be due to the entry pending for deletion or other reasons.
    IsLocked bool `json:"isLocked"`
    // The amount of related activity in the past hour
    Trending int `json:"trending"`
    // The amount of user's who have favourited the media
    Favourites int `json:"favourites"`
    // List of tags that describes elements and themes of the media
    Tags []Tags `json:"tags"`
    // Other media in the same or connecting franchise
    Relations Relations `json:"relations"`
    // The characters in the media
    Characters Characters `json:"characters"`
    // The staff who produced the media
    Staff Staff `json:"staff"`
    // The companies who produced the media
    Studio Studio `json:"studio"`
    // If the media is marked as favourite by the current authenticated user
    IsFavourite bool `json:"isFavourite"`
    // If the media is intended only for 18+ adult audiences
    IsAdult bool `json:"isAdult"`
    // The media's next episode airing schedule
    NextAiringEpisode NextAiringEpisode `json:"nextAiringEpisode"`
    // The media's entire airing schedule
    AiringSchedule AiringSchedule `json:"airingSchedule"`
    // External links to another site related to the media
    ExternalLinks []ExternalLinks `json:"externalLinks"`
    // Data and links to legal streaming episodes on external sites
    StreamingEpisodes []StreamingEpisodes `json:"streamingEpisodes"`
    // The ranking of the media in a particular time span and format compared to other media
    Rankings []Rankings `json:"rankings"`
    // User reviews of the media
    Reviews Reviews `json:"reviews"`
    // User recommendations for similar media
    Recommendations Recommendations `json:"recommendations"`
    Stats           Stats           `json:"stats"`
    // The url for the media page on the AniList website
    SiteURL string `json:"siteUrl"`
    // If the media should have forum thread automatically created for it on airing episode release
    AutoCreateForumThread bool `json:"autoCreateForumThread"`
    // If the media is blocked from being recommended to/from
    IsRecommendationBlocked bool `json:"isRecommendationBlocked"`
    // Notes for site moderators
    ModNotes string `json:"modNotes"`
}
```

### func NewMedia

```go
func NewMedia() *Media
```

NewMedia creates Media objects

### func \(\*Media\) FilterAnimeByID

```go
func (m *Media) FilterAnimeByID(id int)
```

FilterAnimeByID search Anilist Anime only type: ANIME is hard\-coded in the query

### func \(\*Media\) FilterByEndDate

```go
func (m *Media) FilterByEndDate(date int32)
```

FilterByEndDate search Anilist Media by start Date of the show 8 digit 2013\-04\-08 == 20130408

### func \(\*Media\) FilterByID

```go
func (m *Media) FilterByID(id int)
```

FilterByID Search Anilist Media by it's ID

### func \(\*Media\) FilterByMalID

```go
func (m *Media) FilterByMalID(malID int)
```

FilterByMalID search Anilist Media by it's MAL\(MyAnimeList\) ID

### func \(\*Media\) FilterBySeason

```go
func (m *Media) FilterBySeason(season string)
```

FilterBySeason search Anilist Media by Season \(WINTER\, SPRING\, SUMMER\, FALL\)

### func \(\*Media\) FilterByStartDate

```go
func (m *Media) FilterByStartDate(date int32)
```

FilterByStartDate search Anilist Media by start Date of the show 8 digit 2013\-04\-08 == 20130408

### func \(\*Media\) FilterByTitle

```go
func (m *Media) FilterByTitle(title string)
```

FilterByTitle search Anilist Media by title of the anime or manga

### func \(\*Media\) FilterMangaByID

```go
func (m *Media) FilterMangaByID(id int)
```

FilterMangaByID search Anilist Manga type: MANGA is hard\-coded in the query

## type MediaTrend

MediaTrend Daily Media Statistics

```go
type MediaTrend struct {
    MediaID      int64 `json:"mediaID"`
    Date         int64 `json:"date"`
    Trending     int   `json:"trending"`
    AverageScore int   `json:"averageScore"`
    Popularity   int   `json:"popularity"`
    InProgress   int   `json:"inProgress"`
    Releasing    bool  `json:"releasing"`
    Episode      int   `json:"episode"`
    Media        Media `json:"media"`
}
```

### func NewMediaTrend

```go
func NewMediaTrend() *MediaTrend
```

NewMediaTrend Create new NewMediaTrend Object

### func \(\*MediaTrend\) FilterByEpisode

```go
func (m *MediaTrend) FilterByEpisode(episode int)
```

FilterByEpisode filters by episode number

### func \(\*MediaTrend\) FilterByPopularity

```go
func (m *MediaTrend) FilterByPopularity(popularity int)
```

FilterByPopularity filters by popularity

### func \(\*MediaTrend\) FilterByTrendingAmount

```go
func (m *MediaTrend) FilterByTrendingAmount(trending int)
```

FilterByTrendingAmount filters by trending amount

### func \(\*MediaTrend\) FilterByTrendingAverageScore

```go
func (m *MediaTrend) FilterByTrendingAverageScore(averageScore int)
```

FilterByTrendingAverageScore filters by average score

### func \(\*MediaTrend\) SearchByMediaID

```go
func (m *MediaTrend) SearchByMediaID(id int)
```

SearchByMediaID searches mediaTrend by anilist ID

## type Name

Name staff name first\, last\, full name\, and native

```go
type Name struct {
    First  string `json:"first"`
    Last   string `json:"last"`
    Full   string `json:"full"`
    Native string `json:"native"`
}
```

## type NextAiringEpisode

NextAiringEpisode contains the ID of the anime episode?

```go
type NextAiringEpisode struct {
    ID int `json:"id"`
}
```

## type Node

Node that contains the ID

```go
type Node struct {
    ID int `json:"id"`
}
```

## type Page

Page Object to store pages

```go
type Page struct {
    PageInfo PageInfo `json:"pageInfo"`
    Media    []Media  `json:"media"`
}
```

### func \(\*Page\) PaginationByTitle

```go
func (p *Page) PaginationByTitle(title string, page int, perPage int)
```

PaginationByTitle  search Anilist Media by title returns arrayList of Media objects\, and pageInfor takes title string\, page \(which page to look for\)\, PerPage The amount of entries per page\, max 50

## type PageInfo

PageInfo represents the information regarding the Page of Anilist query

```go
type PageInfo struct {
    Total       int  `json:"total"`
    PerPage     int  `json:"perPage"`
    CurrentPage int  `json:"currentPage"`
    LastPage    int  `json:"lastPage"`
    HasNextPage bool `json:"hasNextPage"`
}
```

## type Rankings

Rankings of the Anime ID

```go
type Rankings struct {
    ID int `json:"id"`
}
```

## type Recommendations

Recommendations of the similar anime\.

```go
type Recommendations struct {
    Edges []RecroRevEdges `json:"edges"`
    Node  Node            `json:"node"`
}
```

## type RecroRevEdges

RecroRevEdges \(temporaty name type find a better name for different edges :\( \)\) Node Object nested from Media

```go
type RecroRevEdges struct {
    Node Node `json:"node"`
}
```

## type Relations

Relations Object I believe relations to the anime/manga

```go
type Relations struct {
    Edges []Edges `json:"edges"`
}
```

## type Reviews

Reviews If the authenticated user have made review for it I believe\.

```go
type Reviews struct {
    Edges []RecroRevEdges `json:"edges"`
}
```

## type ScoreDistribution

ScoreDistribution individual item to access it use array such as ScoreDistribution\[0\] to get the first array\.

```go
type ScoreDistribution struct {
    Score  int `json:"score"`
    Amount int `json:"amount"`
}
```

## type Staff

Staff Voice actors or production staff

```go
type Staff struct {
    ID               int            `json:"id"`
    Name             Name           `json:"name"`
    Language         string         `json:"language"`
    Image            Image          `json:"image"`
    Description      string         `json:"description"`
    IsFavourite      bool           `json:"isFavourite"`
    SiteURL          string         `json:"siteUrl"`
    StaffMedia       StaffMedia     `json:"staffMedia"`
    Characters       Characters     `json:"characters"`
    CharacterMedia   CharacterMedia `json:"characterMedia"`
    Staff            string         `json:"staff"`
    Submitter        string         `json:"submitter"`
    SubmissionStatus int            `json:"submissionStatus"`
    Favourites       int            `json:"favourites"`
    ModNotes         string         `json:"modNotes"`
    Edges            []Edges        `json:"edges"`
}
```

### func \(\*Staff\) FilterStaffByID

```go
func (s *Staff) FilterStaffByID(id int)
```

FilterStaffByID Search Anilist Staff by it's ID

### func \(\*Staff\) FilterStaffByName

```go
func (s *Staff) FilterStaffByName(name string)
```

FilterStaffByName Search Anilist Staff by it's ID

## type StaffMedia

StaffMedia Media where the staff member has a production role

```go
type StaffMedia struct {
    Edges []Edges `json:"edges"`
}
```

## type StartDate

StartDate object that have year month and day of the anime

```go
type StartDate struct {
    Year  int `json:"year"`
    Month int `json:"month"`
    Day   int `json:"day"`
}
```

## type Stats

Stats of the current anime array of Score such as 10\, 20\, 30\, 100 tells you how many people scored it such and Status such as PLANNING\, WATCHING\, DROPPED

```go
type Stats struct {
    ScoreDistribution  []ScoreDistribution  `json:"scoreDistribution"`
    StatusDistribution []StatusDistribution `json:"statusDistribution"`
}
```

## type StatusDistribution

StatusDistribution individual item to access it use array such as StatusDistribution\[0\] to get the first array\.

```go
type StatusDistribution struct {
    Status string `json:"status"`
    Amount int    `json:"amount"`
}
```

## type StreamingEpisodes

StreamingEpisodes returns the Episode title\, Thumbnail\, URL to the streaming site\, and the Site

```go
type StreamingEpisodes struct {
    Title     string `json:"title"`
    Thumbnail string `json:"thumbnail"`
    URL       string `json:"url"`
    Site      string `json:"site"`
}
```

## type Studio

Studio studios that worked on the anime/manga \(anime mostly\)

```go
type Studio struct {
    Edges []Edges `json:"edges"`
}
```

## type Tags

Tags object contais the ID for tags

```go
type Tags struct {
    ID int `json:"id"`
}
```

## type Title

Title object that has Romaji\, English\, Native and UserPreffered if they are given\.

```go
type Title struct {
    Romaji        string `json:"romaji"`
    English       string `json:"english"`
    Native        string `json:"native"`
    UserPreferred string `json:"userPreferred"`
}
```

## type Trailer

Trailer contains YouTube Unique code to the video\.

```go
type Trailer struct {
    ID string `json:"id"`
}
```