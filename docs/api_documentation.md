# API Routes Documentation

## 1. Create a Tile

- **Endpoint:** `/tiles`
- **Method:** `POST`
- **Description:** Creates a new tile with the specified properties.
- **Request Body:**

  ```json
  {
    "x_coordinate": 1,
    "y_coordinate": 2,
    "type": "grass"
  }
  ```

- **Response:**

  ```json
  {
    "id": 1,
    "x_coordinate": 1,
    "y_coordinate": 2,
    "type": "grass",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
  ```

## 2. Get a Tile by ID

- **Endpoint:** `/tiles/{id}`
- **Method:** `GET`
- **Description:** Retrieves a tile by its ID.
- **Response:**

  ```json
  {
    "id": 1,
    "x_coordinate": 1,
    "y_coordinate": 2,
    "type": "grass",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
  ```

## 3. Get a Tile by Coordinates

- **Endpoint:** `/tiles/coordinates`
- **Method:** `GET`
- **Description:** Retrieves a tile by its x and y coordinates.
- **Query Parameters:**

  ```
  x - The x coordinate of the tile.
  y - The y coordinate of the tile.
  ```

- **Response:**

  ```json
  {
    "id": 1,
    "x_coordinate": 1,
    "y_coordinate": 2,
    "type": "grass",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
  ```

## 4. Get All Tiles

- **Endpoint:** `/tiles`
- **Method:** `GET`
- **Description:** Retrieves all tiles in the database.
- **Response:**

  ```json
  [
    {
      "id": 1,
      "x_coordinate": 1,
      "y_coordinate": 2,
      "type": "grass",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    },
    {
      "id": 2,
      "x_coordinate": 2,
      "y_coordinate": 3,
      "type": "water",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  ]
  ```

## 5. Update a Tile

- **Endpoint:** `/tiles`
- **Method:** `PUT`
- **Description:** Updates the properties of an existing tile.
- **Request Body:**

  ```json
  {
    "id": 1,
    "x_coordinate": 1,
    "y_coordinate": 2,
    "type": "mountain"
  }
  ```

## 6. Delete a Tile

- **Endpoint:** `/tiles/{id}`
- **Method:** `DELETE`
- **Description:** Deletes a tile by its ID.

## 7. Initialize Tiles

- **Endpoint:** `/init-tiles`
- **Method:** `POST`
- **Description:** Initializes a grid of tiles with random types up to the specified max coordinates.
- **Query Parameters:**

  ```
  maxX - The maximum x coordinate.
  maxY - The maximum y coordinate.
  ```

## Example Usage

### Create a Tile

```curl
curl -X POST -d '{
  "x_coordinate": 1,
  "y_coordinate": 2,
  "type": "grass"
}' -H "Content-Type: application/json" "http://localhost:6080/tiles"
```

### Get a Tile by ID

```curl
curl -X GET "http://localhost:6080/tiles/1"
```

### Get a Tile by Coordinates

```curl
curl -X GET "http://localhost:6080/tiles/coordinates?x=1&y=2"
```

### Get All Tiles

```curl
curl -X GET "http://localhost:6080/tiles"
```

### Update a Tile

```curl
curl -X PUT -d '{
  "id": 1,
  "x_coordinate": 1,
  "y_coordinate": 2,
  "type": "mountain"
}' -H "Content-Type: application/json" "http://localhost:6080/tiles"
```

### Delete a Tile

```curl
curl -X DELETE "http://localhost:6080/tiles/1"
```

### Initialize Tiles

```curl
curl -X POST "http://localhost:6080/init-tiles?maxX=10&maxY=10"
```
