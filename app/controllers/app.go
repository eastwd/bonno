package controllers

import (
	"bonno/app/models"
	"bonno/app/routes"
	"github.com/revel/revel"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	*revel.Controller
}

// 動画ディレクトリの絶対パスを返す（相対パスはアプリのベースパス基準）
func movieBaseDir() string {
	abdir := revel.Config.StringDefault("bonno.abdir", "public/douga")
	if filepath.IsAbs(abdir) {
		return abdir
	}
	return filepath.Join(revel.BasePath, abdir)
}

// インデックス画面を表示する
func (c App) Index() revel.Result {
	movielist, _ := DbMap.Select(models.Movie{}, "select * from Movie order by RANDOM() limit 1")
	catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
	return c.Render(movielist, catelist)
}

// かっちょいいトップ画面を表示する
func (c App) Top() revel.Result {
	abdir := movieBaseDir()
	reldir := revel.Config.StringDefault("bonno.reldir", "./")
	entries, _ := os.ReadDir(abdir + "/other/top")
	var topmovies []string
	for _, e := range entries {
		if e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		topmovies = append(topmovies, e.Name())
	}
	// 動画が無い場合は topmoviepath を空にし、テンプレート側で代替表示する
	topmoviepath := ""
	if len(topmovies) > 0 {
		topmoviepath = reldir + "/other/top/" + topmovies[rand.Intn(len(topmovies))]
	}
	return c.Render(topmoviepath)
}

// 動画一覧を表示する
func (c App) Movie() revel.Result {
	movielist, _ := DbMap.Select(models.Movie{}, "select * from Movie order by RANDOM() limit 10")
	catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
	return c.Render(movielist, catelist)
}

func (c App) Search(word string) revel.Result {
	movielist, _ := DbMap.Select(models.Movie{}, "select * from Movie where Name like ?", word)
	catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
	return c.Render(word, movielist, catelist)
}

func (c App) Category(tag string) revel.Result {
	movielist, _ := DbMap.Select(models.Movie{}, "select * from Movie where Tag1=? or Tag2=? or Tag3=? or Tag4=?", tag, tag, tag, tag)
	catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
	return c.Render(tag, movielist, catelist)
}

// 動画詳細を表示する
func (c App) Detail(id int) revel.Result {
	rows, _ := DbMap.Select(models.Movie{}, "select * from Movie where Id=?", id)
	movie := rows[0].(*models.Movie)
	abdir := movieBaseDir()
	moviedata, _ := os.ReadDir(abdir + "/movie/" + movie.Tag1 + "/" + movie.Name)
	catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
	return c.Render(movie, moviedata, catelist)
}

// 動画のディレクトリを読み込んでDBに追加する
func (c App) Init() revel.Result {
	abdir := movieBaseDir()
	reldir := revel.Config.StringDefault("bonno.reldir", "./")
	dirlist, _ := os.ReadDir(abdir + "/movie")
	for _, dir := range dirlist {
		titlelist, _ := os.ReadDir(abdir + "/movie/" + dir.Name())
		for _, title := range titlelist {
			info, err := title.Info()
			if err != nil {
				continue
			}
			rows, _ := DbMap.Select(models.Movie{}, "select * from Movie where Name=?", title.Name())
			if len(rows) == 0 {
				DbMap.Insert(&models.Movie{0, title.Name(), reldir + "/movie/" + dir.Name() + "/" + title.Name(), reldir + "/images/thumb/" + title.Name() + ".png", 0, "", info.ModTime(), dir.Name(), "", "", ""})
			}
		}
	}
	return c.Redirect(routes.App.Index())
}
