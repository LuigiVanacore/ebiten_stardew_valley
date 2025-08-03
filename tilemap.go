package ebiten_stardew_valley

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"os"
	"path/filepath"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/disintegration/imaging"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

// const mapPath = "maps/map.tmx" // Path to your Tiled Map.

// func main() {
//     // Parse .tmx file.
//     gameMap, err := tiled.LoadFile(mapPath)
//     if err != nil {
//         fmt.Printf("error parsing map: %s", err.Error())
//         os.Exit(2)
//     }

//     fmt.Println(gameMap)

//     // You can also render the map to an in-memory image for direct
//     // use with the default Renderer, or by making your own.
//     renderer, err := render.NewRenderer(gameMap)
//     if err != nil {
//         fmt.Printf("map unsupported for rendering: %s", err.Error())
//         os.Exit(2)
//     }

//     // Render just layer 0 to the Renderer.
//     err = renderer.RenderLayer(0)
//     if err != nil {
//         fmt.Printf("layer unsupported for rendering: %s", err.Error())
//         os.Exit(2)
//     }

//     // Get a reference to the Renderer's output, an image.NRGBA struct.
//     img := renderer.Result

//     // Clear the render result after copying the output if separation of
//     // layers is desired.
//     renderer.Clear()

//     // And so on. You can also export the image to a file by using the
//     // Renderer's Save functions.
// }


func LoadTileMaps() map[string]*tiled.Map {
	var maps map[string]*tiled.Map = make(map[string]*tiled.Map)
	mapPath := "../data/maps/world.tmx"
	gameMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	maps["world"] = gameMap

	mapPath = "../data/maps/water.tmx"
	gameMap, err = tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	maps["water"] = gameMap
	mapPath = "../data/maps/plant.tmx"
	gameMap, err = tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	maps["plant"] = gameMap
	mapPath = "../data/maps/house.tmx"
	gameMap, err = tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	maps["house"] = gameMap
	mapPath = "../data/maps/arena.tmx"
	gameMap, err = tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	maps["arena"] = gameMap

	mapPath = "../data/maps/fire.tmx"
	gameMap, err = tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	maps["fire"] = gameMap

	mapPath = "../data/maps/hospital.tmx"
	gameMap, err = tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	maps["hospital"] = gameMap

	mapPath = "../data/maps/hospital2.tmx"
	gameMap, err = tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	maps["hospital2"] = gameMap

	return maps
}


// func Setup( start_map *tiled.Map, playerStartPos string) {  

// 	var sprites []*ebiten_extended.Sprite = make([]*ebiten_extended.Sprite, 0)
// 	map_layer := start_map.Layers
// 	for _, layerName := range map_layer {
// 		if layerName.Name == "Terrain" || layerName.Name == "Terrain Top" {
// 			for _, tile := range layerName.Tiles {
// 				tileset := tile.Tileset
// 				sf, err := r.open(tile.Tileset.GetFileFullPath(tile.Tileset.Image.Source))
// 							if err != nil {
// 							return nil, err
// 						}
// 					defer sf.Close()

// 							img, _, err := image.Decode(sf)
// 							if err != nil {
// 								return nil, err
// 					}	
// 				sprite := ebiten_extended.NewSprite(
// 					tileset.Name,
// 					img,
// 					ebiten_pokemon.BG_LAYER, false,
// 				)
// 				sprite.SetPosition(math2D.NewVector2D(float64(tileset.TileOffset.X*tileset.TileWidth), float64(tileset.TileOffset.Y*tileset.TileHeight)))
// 				sprites = append(sprites, sprite)
// 				ebiten_extended.GameManager().World().AddNode(sprite)
// 			}
// 		}
// 	}
// }

type TileMap struct {
	ebiten_extended.Node2D
	tiledMap *tiled.Map
	Result    *image.NRGBA
	Width     int
	Height    int
	TileWidth int
	TileHeight int
	tileCache map[uint32]image.Image 
	engine    *render.OrthogonalRendererEngine
}

type TileLayer struct {
	Result    *image.NRGBA
	tileCache map[uint32]image.Image
}

func NewTileMap(tiledMap *tiled.Map) *TileMap {
	tileMap := &TileMap{
		tiledMap:  tiledMap,
		Result:    image.NewNRGBA(image.Rect(0, 0, tiledMap.Width*tiledMap.TileWidth, tiledMap.Height*tiledMap.TileHeight)),
		Width:     tiledMap.Width,
		Height:    tiledMap.Height,
		TileWidth: tiledMap.TileWidth,
		TileHeight: tiledMap.TileHeight,
		tileCache: make(map[uint32]image.Image), 
		engine:   &render.OrthogonalRendererEngine{},
	}
	tileMap.engine.Init(tiledMap)
	return tileMap
}




func (t *TileMap) getTileImage(tile *tiled.LayerTile) (image.Image, error) {
	timg, ok := t.tileCache[tile.Tileset.FirstGID+tile.ID]
	if ok {
		return t.engine.RotateTileImage(tile, timg), nil
	}
	// Precache all tiles in tileset
	if tile.Tileset.Image == nil {
		for i := 0; i < len(tile.Tileset.Tiles); i++ {
			if tile.Tileset.Tiles[i].ID == tile.ID {
				sf, err := t.open(tile.Tileset.GetFileFullPath(tile.Tileset.Tiles[i].Image.Source))
				if err != nil {
					return nil, err
				}
				defer sf.Close()
				timg, _, err = image.Decode(sf)
				if err != nil {
					return nil, err
				}
				t.tileCache[tile.Tileset.FirstGID+tile.ID] = timg
			}
		}
	} else {
		sf, err := t.open(tile.Tileset.GetFileFullPath(tile.Tileset.Image.Source))
		if err != nil {
			return nil, err
		}
		defer sf.Close()

		img, _, err := image.Decode(sf)
		if err != nil {
			return nil, err
		}

		for i := uint32(0); i < uint32(tile.Tileset.TileCount); i++ {
			rect := tile.Tileset.GetTileRect(i)
			t.tileCache[i+tile.Tileset.FirstGID] = imaging.Crop(img, rect)
			if tile.ID == i {
				timg = t.tileCache[i+tile.Tileset.FirstGID]
			}
		}
	}

	return t.engine.RotateTileImage(tile, timg), nil
}

func (t *TileMap) GetLayerByName(name string) (*tiled.Layer, error) {
	for _, layer := range t.tiledMap.Layers {
		if layer.Name == name {
			return layer, nil
		}
	}
	return nil, fmt.Errorf("layer not found: %s", name)
}


func (t *TileMap) open(f string) (io.ReadCloser, error) {
		return os.Open(filepath.FromSlash(f))
	 
}

func (t *TileMap) RenderVisibleLayers() error {
	for i := range t.tiledMap.Layers {
		if !t.tiledMap.Layers[i].Visible {
			continue
		}

		if err := t.RenderLayer(i); err != nil {
			return err
		}
	}

	return nil
}

func (t *TileMap) RenderLayer(id int) error {
	if id >= len(t.tiledMap.Layers) {
		return fmt.Errorf("tiled/render: index out of bounds for layer %d", id)
	}
	return t._renderLayer(t.tiledMap.Layers[id])
}

func (t *TileMap) RenderLayerByName(name string) error {
	layer, err := t.GetLayerByName(name)
	if err != nil {
		return err
	}
	return t._renderLayer(layer)
}

func (t *TileMap) GetTilesPositionInLayer(layer *tiled.Layer) []math2D.Vector2D {
		var xs, xe, xi, ys, ye, yi int

		xs = 0
		xe = t.tiledMap.Width
		xi = 1
		ys = 0
		ye = t.tiledMap.Height
		yi = 1

		i := 0

		positions := make([]math2D.Vector2D, 0, len(layer.Tiles))
	for y := ys; y*yi < ye; y = y + yi {
		for x := xs; x*xi < xe; x = x + xi {
			if layer.Tiles[i].IsNil() {
				i++
				continue
			}

			 positions = append(positions, math2D.NewVector2D(float64(x*t.TileWidth), float64(y*t.TileHeight)))

			i++
		}
	}

		return positions
}

func (t *TileMap) _renderLayer(layer *tiled.Layer) error {
	var xs, xe, xi, ys, ye, yi int
	if t.tiledMap.RenderOrder == "" || t.tiledMap.RenderOrder == "right-down" {
		xs = 0
		xe = t.tiledMap.Width
		xi = 1
		ys = 0
		ye = t.tiledMap.Height
		yi = 1
	} else {
		return fmt.Errorf("tiled/render: unsupported render order %s", t.tiledMap.RenderOrder)
	}

	i := 0
	for y := ys; y*yi < ye; y = y + yi {
		for x := xs; x*xi < xe; x = x + xi {
			if layer.Tiles[i].IsNil() {
				i++
				continue
			}

			img, err := t.getTileImage(layer.Tiles[i])
			if err != nil {
				return err
			}

			pos := t.engine.GetTilePosition(x, y)

			if layer.Opacity < 1 {
				mask := image.NewUniform(color.Alpha{uint8(layer.Opacity * 255)})

				draw.DrawMask(t.Result, pos, img, img.Bounds().Min, mask, mask.Bounds().Min, draw.Over)
			} else {
				draw.Draw(t.Result, pos, img, img.Bounds().Min, draw.Over)
			}

			i++
		}
	}

	return nil
}

// Clear clears the render result to allow for separation of layers. For example, you can
// render a layer, make a copy of the render, clear the renderer, and repeat for each
// layer in the Map.
func (t *TileMap) Clear() {
	t.Result = image.NewNRGBA(t.engine.GetFinalImageSize())
}

// var (
// 	// ErrUnsupportedOrientation represents an error in the unsupported orientation for rendering.
// 	ErrUnsupportedOrientation = errors.New("tiled/render: unsupported orientation")
// 	// ErrUnsupportedRenderOrder represents an error in the unsupported order for rendering.
// 	// ErrOutOfBounds represents an error that the index is out of bounds
// 	ErrOutOfBounds = errors.New("tiled/render: index out of bounds")
// )

// type Renderer struct {
// 	m         *tiled.Map
// 	Result    *image.NRGBA // The image result after rendering using the Render functions.
// 	tileCache map[uint32]image.Image
// 	engine    *OrthogonalRendererEngine
// 	fs        fs.FS
// }

type TileMapImporter struct {

}


// func (t *TileMap) Import(fileName string)  error {

// 	t.tileCache = make(map[uint32]image.Image)
// 	t.engine = &render.OrthogonalRendererEngine{}

// 	tm, err := tiled.LoadFile(fileName)
// 	if err != nil {
// 		return fmt.Errorf("failed to load tile map: %w", err)
// 	}
	
// 	if tm.Orientation != "orthogonal" {
// 		return fmt.Errorf("unsupported map orientation: %s", tm.Orientation)
// 	}
// 	fmt.Println(tm)
// 	t.Height = tm.Height
// 	t.Width = tm.Width
// 	t.TileWidth = tm.TileWidth
// 	t.TileHeight = tm.TileHeight
// 	t.result = image.NewNRGBA(image.Rect(0, 0, t.Width*t.TileWidth, t.Height*t.TileHeight))
// 	t.tiledMap = tm
// 	t.engine.Init(tm)
	// layers := []string {"Terrain", "Terrain Top"}

	// tileLayers, err := GetLayersByNames(tm, layers...)
	// if err != nil {
	// 	return fmt.Errorf("failed to get layers: %w", err)
	// }

 

// 	for _, layer := range tm.Layers {
// 		if layer.Visible {
// 			err := t.RenderLayer(layer)
// 			if err != nil {
// 				return fmt.Errorf("failed to render layer %s: %w", layer.Name, err)
// 			}
// 		}
// 	}

// 			ebitenImg := ebiten.NewImageFromImage(t.result)
// 			sprite := ebiten_extended.NewSprite(
// 				"tileMap",
// 				ebitenImg,
// 				t.layerID,
// 				false,
// 			)
// 			t.AddChild(sprite)

// 	return nil
// }


// func (t *TileMap) RenderVisibleLayers() error {
// 	for i := range t.tiledMap.Layers {
// 		if !t.tiledMap.Layers[i].Visible {
// 			continue
// 		}

// 		if err := t.RenderLayer(t.tiledMap.Layers[i]); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }


// func (t *TileMap) RenderLayer(layer *tiled.Layer)  error {
// 	var xs, xe, xi, ys, ye, yi int
	
// 	xs = 0
// 	xe = t.Width
// 	xi = 1
// 	ys = 0
// 	ye = t.Height
// 	yi = 1

// 	i := 0
// 	for y := ys; y*yi < ye; y = y + yi {
// 		for x := xs; x*xi < xe; x = x + xi {
// 			if layer.Tiles[i].IsNil() {
// 				i++
// 				continue
// 			}

// 			img, err := t.getTileImage(layer.Tiles[i])
// 			if err != nil {
// 				return  err
// 			}

// 			pos := t.engine.GetTilePosition(x, y)

// 			if layer.Opacity < 1 {
// 				mask := image.NewUniform(color.Alpha{uint8(layer.Opacity * 255)})

// 				draw.DrawMask(t.result, pos, img, img.Bounds().Min, mask, mask.Bounds().Min, draw.Over)
// 			} else {
// 				draw.Draw(t.result, pos, img, img.Bounds().Min, draw.Over)
// 			}

// 			i++
// 		}
// 	}

// 	return nil
// }

 
// 	var xs, xe, xi, ys, ye, yi int
// 	if r.m.RenderOrder == "" || r.m.RenderOrder == "right-down" {
// 		xs = 0
// 		xe = r.m.Width
// 		xi = 1
// 		ys = 0
// 		ye = r.m.Height
// 		yi = 1
// 	} else {
// 		return ErrUnsupportedRenderOrder
// 	}

// 	i := 0
// 	for y := ys; y*yi < ye; y = y + yi {
// 		for x := xs; x*xi < xe; x = x + xi {
// 			if layer.Tiles[i].IsNil() {
// 				i++
// 				continue
// 			}

// 			img, err := r.getTileImage(layer.Tiles[i])
// 			if err != nil {
// 				return err
// 			}

// 			pos := r.engine.GetTilePosition(x, y)

// 			if layer.Opacity < 1 {
// 				mask := image.NewUniform(color.Alpha{uint8(layer.Opacity * 255)})

// 				draw.DrawMask(r.Result, pos, img, img.Bounds().Min, mask, mask.Bounds().Min, draw.Over)
// 			} else {
// 				draw.Draw(r.Result, pos, img, img.Bounds().Min, draw.Over)
// 			}

// 			i++
// 		}
// 	}

// 	return nil
// }

// func GetLayersByNames(m *tiled.Map, names ...string) ([]*tiled.Layer, error) {
// 	var layers []*tiled.Layer
// 	for _, searchName := range names {
// 		found := false
// 		for _, layer := range m.Layers {
// 			if layer.Name == searchName {
// 				layers = append(layers, layer)
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			return nil, fmt.Errorf("layer not found: %s", searchName)
// 		}
// 	}
// 	return layers, nil
// }

// func (t *TileMap) getTileImage(tile *tiled.LayerTile) (image.Image, error) {
// 	timg, ok := t.tileCache[tile.Tileset.FirstGID+tile.ID]
// 	if ok {
// 		return  t.engine.RotateTileImage(tile, timg), nil
// 	}
// 	// Precache all tiles in tileset
// 	if tile.Tileset.Image == nil {
// 		for i := 0; i < len(tile.Tileset.Tiles); i++ {
// 			if tile.Tileset.Tiles[i].ID == tile.ID {
// 				sf, err := os.Open(filepath.FromSlash(tile.Tileset.GetFileFullPath(tile.Tileset.Tiles[i].Image.Source)))	
// 				if err != nil {
// 					return nil, err
// 				}
// 				defer sf.Close()
// 				timg, _, err = image.Decode(sf)
// 				if err != nil {
// 					return nil, err
// 				}
// 				t.tileCache[tile.Tileset.FirstGID+tile.ID] = timg
// 			}
// 		}
// 	} else {
// 		sf, err := os.Open(filepath.FromSlash(tile.Tileset.GetFileFullPath(tile.Tileset.Image.Source)))
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer sf.Close()

// 		img, _, err := image.Decode(sf)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for i := uint32(0); i < uint32(tile.Tileset.TileCount); i++ {
// 			rect := tile.Tileset.GetTileRect(i)
// 			t.tileCache[i+tile.Tileset.FirstGID] = imaging.Crop(img, rect)
// 			if tile.ID == i {
// 				timg = t.tileCache[i+tile.Tileset.FirstGID]
// 			}
// 		}
// 	}
// 	return t.engine.RotateTileImage(tile, timg), nil
// }


// func (r *Renderer) getTileImage(tile *tiled.LayerTile) (image.Image, error) {
// 	timg, ok := r.tileCache[tile.Tileset.FirstGID+tile.ID]
// 	if ok {
// 		return r.engine.RotateTileImage(tile, timg), nil
// 	}
// 	// Precache all tiles in tileset
// 	if tile.Tileset.Image == nil {
// 		for i := 0; i < len(tile.Tileset.Tiles); i++ {
// 			if tile.Tileset.Tiles[i].ID == tile.ID {
// 				sf, err := r.open(tile.Tileset.GetFileFullPath(tile.Tileset.Tiles[i].Image.Source))
// 				if err != nil {
// 					return nil, err
// 				}
// 				defer sf.Close()
// 				timg, _, err = image.Decode(sf)
// 				if err != nil {
// 					return nil, err
// 				}
// 				r.tileCache[tile.Tileset.FirstGID+tile.ID] = timg
// 			}
// 		}
// 	} else {
// 		sf, err := r.open(tile.Tileset.GetFileFullPath(tile.Tileset.Image.Source))
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer sf.Close()

// 		img, _, err := image.Decode(sf)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for i := uint32(0); i < uint32(tile.Tileset.TileCount); i++ {
// 			rect := tile.Tileset.GetTileRect(i)
// 			r.tileCache[i+tile.Tileset.FirstGID] = imaging.Crop(img, rect)
// 			if tile.ID == i {
// 				timg = r.tileCache[i+tile.Tileset.FirstGID]
// 			}
// 		}
// 	}

// 	return r.engine.RotateTileImage(tile, timg), nil
// }


// func RotateTileImage(tile *tiled.LayerTile, img image.Image) image.Image {
// 	timg := img
// 	if tile.DiagonalFlip {
// 		timg = imaging.FlipH(imaging.Rotate270(timg))
// 	}
// 	if tile.HorizontalFlip {
// 		timg = imaging.FlipH(timg)
// 	}
// 	if tile.VerticalFlip {
// 		timg = imaging.FlipV(timg)
// 	}

// 	return timg
// }

// func (t *TileMap) GetTilePosition(x, y int) image.Rectangle {
// 	return image.Rect(x*t.TileWidth,
// 		y*t.TileHeight,
// 		(x+1)*t.TileWidth,
// 		(y+1)*t.TileHeight)
// }

// // NewRenderer creates new rendering engine instance.
// func NewRenderer(m *tiled.Map) (*Renderer, error) {
// 	return NewRendererWithFileSystem(m, nil)
// }

// // NewRendererWithFileSystem creates new rendering engine instance with a custom file system.
// func NewRendererWithFileSystem(m *tiled.Map, fs fs.FS) (*Renderer, error) {
// 	r := &Renderer{m: m, tileCache: make(map[uint32]image.Image), fs: fs}
// 	if r.m.Orientation == "orthogonal" {
// 		r.engine = &OrthogonalRendererEngine{}
// 	} else {
// 		return nil, ErrUnsupportedOrientation
// 	}

// 	r.engine.Init(r.m)
// 	r.Clear()

// 	return r, nil
// }

// func (r *Renderer) open(f string) (io.ReadCloser, error) {
// 	if r.fs == nil {
// 		return os.Open(filepath.FromSlash(f))
// 	}
// 	return r.fs.Open(filepath.ToSlash(f))
// }

// func (r *Renderer) getTileImage(tile *tiled.LayerTile) (image.Image, error) {
// 	timg, ok := r.tileCache[tile.Tileset.FirstGID+tile.ID]
// 	if ok {
// 		return r.engine.RotateTileImage(tile, timg), nil
// 	}
// 	// Precache all tiles in tileset
// 	if tile.Tileset.Image == nil {
// 		for i := 0; i < len(tile.Tileset.Tiles); i++ {
// 			if tile.Tileset.Tiles[i].ID == tile.ID {
// 				sf, err := r.open(tile.Tileset.GetFileFullPath(tile.Tileset.Tiles[i].Image.Source))
// 				if err != nil {
// 					return nil, err
// 				}
// 				defer sf.Close()
// 				timg, _, err = image.Decode(sf)
// 				if err != nil {
// 					return nil, err
// 				}
// 				r.tileCache[tile.Tileset.FirstGID+tile.ID] = timg
// 			}
// 		}
// 	} else {
// 		sf, err := r.open(tile.Tileset.GetFileFullPath(tile.Tileset.Image.Source))
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer sf.Close()

// 		img, _, err := image.Decode(sf)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for i := uint32(0); i < uint32(tile.Tileset.TileCount); i++ {
// 			rect := tile.Tileset.GetTileRect(i)
// 			r.tileCache[i+tile.Tileset.FirstGID] = imaging.Crop(img, rect)
// 			if tile.ID == i {
// 				timg = r.tileCache[i+tile.Tileset.FirstGID]
// 			}
// 		}
// 	}

// 	return r.engine.RotateTileImage(tile, timg), nil
// }

// func (r *Renderer) _renderLayer(layer *tiled.Layer) error {
// 	var xs, xe, xi, ys, ye, yi int
// 	if r.m.RenderOrder == "" || r.m.RenderOrder == "right-down" {
// 		xs = 0
// 		xe = r.m.Width
// 		xi = 1
// 		ys = 0
// 		ye = r.m.Height
// 		yi = 1
// 	} else {
// 		return ErrUnsupportedRenderOrder
// 	}

// 	i := 0
// 	for y := ys; y*yi < ye; y = y + yi {
// 		for x := xs; x*xi < xe; x = x + xi {
// 			if layer.Tiles[i].IsNil() {
// 				i++
// 				continue
// 			}

// 			img, err := r.getTileImage(layer.Tiles[i])
// 			if err != nil {
// 				return err
// 			}

// 			pos := r.engine.GetTilePosition(x, y)

// 			if layer.Opacity < 1 {
// 				mask := image.NewUniform(color.Alpha{uint8(layer.Opacity * 255)})

// 				draw.DrawMask(r.Result, pos, img, img.Bounds().Min, mask, mask.Bounds().Min, draw.Over)
// 			} else {
// 				draw.Draw(r.Result, pos, img, img.Bounds().Min, draw.Over)
// 			}

// 			i++
// 		}
// 	}

// 	return nil
// }

// RenderGroupLayer renders single map layer in a certain group.
// func (r *Renderer) RenderGroupLayer(groupID, layerID int) error {
// 	if groupID >= len(r.m.Groups) {
// 		return ErrOutOfBounds
// 	}
// 	group := r.m.Groups[groupID]

// 	if layerID >= len(group.Layers) {
// 		return ErrOutOfBounds
// 	}
// 	return r._renderLayer(group.Layers[layerID])
// }

// // RenderLayer renders single map layer.
// func (r *Renderer) RenderLayer(id int) error {
// 	if id >= len(r.m.Layers) {
// 		return ErrOutOfBounds
// 	}
// 	return r._renderLayer(r.m.Layers[id])
// }

// // RenderVisibleLayers renders all visible map layers.
// func (r *Renderer) RenderVisibleLayers() error {
// 	for i := range r.m.Layers {
// 		if !r.m.Layers[i].Visible {
// 			continue
// 		}

// 		if err := r.RenderLayer(i); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // Clear clears the render result to allow for separation of layers. For example, you can
// // render a layer, make a copy of the render, clear the renderer, and repeat for each
// // layer in the Map.
// func (r *Renderer) Clear() {
// 	r.Result = image.NewNRGBA(r.engine.GetFinalImageSize())
// }


// // OrthogonalRendererEngine represents orthogonal rendering engine.
// type OrthogonalRendererEngine struct {
// 	m *tiled.Map
// }

// // Init initializes rendering engine with provided map options.
// func (e *OrthogonalRendererEngine) Init(m *tiled.Map) {
// 	e.m = m
// }

// // GetFinalImageSize returns final image size based on map data.
// func (e *OrthogonalRendererEngine) GetFinalImageSize() image.Rectangle {
// 	return image.Rect(0, 0, e.m.Width*e.m.TileWidth, e.m.Height*e.m.TileHeight)
// }

// // RotateTileImage rotates provided tile layer.
// func (e *OrthogonalRendererEngine) RotateTileImage(tile *tiled.LayerTile, img image.Image) image.Image {
// 	timg := img
// 	if tile.DiagonalFlip {
// 		timg = imaging.FlipH(imaging.Rotate270(timg))
// 	}
// 	if tile.HorizontalFlip {
// 		timg = imaging.FlipH(timg)
// 	}
// 	if tile.VerticalFlip {
// 		timg = imaging.FlipV(timg)
// 	}

// 	return timg
// }

// // GetTilePosition returns tile position in image.
// func (e *OrthogonalRendererEngine) GetTilePosition(x, y int) image.Rectangle {
// 	return image.Rect(x*e.m.TileWidth,
// 		y*e.m.TileHeight,
// 		(x+1)*e.m.TileWidth,
// 		(y+1)*e.m.TileHeight)
// }