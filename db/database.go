package db

import (
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"time"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite", "db/pactsdb.db")
	if err != nil {
		log.Fatal(err)
	}
	err = createSchemas()
	if err != nil {
		log.Fatal(err)
	}

}

func createSchemas() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS tile(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			x_coordinate INTEGER,
			y_coordinate INTEGER,
			type TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS unit(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			x_coordinate INTEGER,
			y_coordinate INTEGER,
			type TEXT,
			fs INTEGER,
			armor INTEGER,
			speed INTEGER,
			range INTEGER,
			stealthed BOOLEAN,
			health INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
		`)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

type Tile struct {
	ID          int       `json:"id"`
	XCoordinate int       `json:"x_coordinate"`
	YCoordinate int       `json:"y_coordinate"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreateTile(tile *Tile) (*Tile, error) {
	stmt, err := DB.Prepare(`
        INSERT INTO tile (x_coordinate, y_coordinate, type)
        VALUES (?, ?, ?)
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(tile.XCoordinate, tile.YCoordinate, tile.Type)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	tile.ID = int(id)

	// Fetch the created tile from the database to get the complete data including timestamps
	createdTile, err := GetTileByID(tile.ID)
	if err != nil {
		return nil, err
	}

	return createdTile, nil
}

func GetTileByID(id int) (*Tile, error) {
	row := DB.QueryRow(`
        SELECT id, x_coordinate, y_coordinate, type, created_at, updated_at
        FROM tile
        WHERE id = ?
    `, id)

	var tile Tile
	err := row.Scan(&tile.ID, &tile.XCoordinate, &tile.YCoordinate, &tile.Type, &tile.CreatedAt, &tile.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tile not found")
		}
		return nil, err
	}

	return &tile, nil
}

func GetTileByCoordinates(x, y int) (*Tile, error) {
	row := DB.QueryRow(`
        SELECT id, x_coordinate, y_coordinate, type, created_at, updated_at
        FROM tile
        WHERE x_coordinate = ? AND y_coordinate = ?
    `, x, y)

	var tile Tile
	err := row.Scan(&tile.ID, &tile.XCoordinate, &tile.YCoordinate, &tile.Type, &tile.CreatedAt, &tile.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tile not found")
		}
		return nil, err
	}

	return &tile, nil
}

func UpdateTile(tile *Tile) error {
	stmt, err := DB.Prepare(`
        UPDATE tile
        SET x_coordinate = ?, y_coordinate = ?, type = ?, updated_at = ?
        WHERE id = ?
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tile.XCoordinate, tile.YCoordinate, tile.Type, time.Now(), tile.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTile(id int) error {
	stmt, err := DB.Prepare(`
        DELETE FROM tile
        WHERE id = ?
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTiles() ([]Tile, error) {
	rows, err := DB.Query(`
        SELECT id, x_coordinate, y_coordinate, type, created_at, updated_at
        FROM tile
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tiles []Tile
	for rows.Next() {
		var tile Tile
		err = rows.Scan(&tile.ID, &tile.XCoordinate, &tile.YCoordinate, &tile.Type, &tile.CreatedAt, &tile.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tiles = append(tiles, tile)
	}

	// Check for any error encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tiles, nil
}

func InitTiles(maxX, maxY int) error {
	tileTypes := []string{"mountain", "grass", "water"}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			tileType := tileTypes[rand.Intn(len(tileTypes))]
			tile := &Tile{
				XCoordinate: x,
				YCoordinate: y,
				Type:        tileType,
			}
			_, err := CreateTile(tile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
