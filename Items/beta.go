package main

import (
	"flag"
	"fmt"
	"time"
	"path/filepath"
	. "fauxgl"
)

var (
	eye           = V(2.3, 1.6, 7) // 9.5,9.5,40 if render is false
	center        = V(0, 0, 0) // 0,2,0 if render is false
	up            = V(0, 5, 0) // 0,3,0 if render is false
	Dimentions    = 512
	CameraScale   = 1  // set to 4 or 5 for production, 2 or 3 for testing and 1 for obj formating
	light         = V(16, 22, 25).Normalize()
	fovy   		  = 22.5 // 1.5 if render is false
	near  		  = 1.0 // 2 if render is false
	far    		  = 1000.0
	color  		  = "#828282"
	Amb           = "#d4d4d4"
	cdnDirectory  = "/var/www/cdn"
)

func main() {
	// Avatar Flags
	hash := flag.String("hash", "none", "avatar hash")
	face := flag.String("face", "none", "Face")
	flag.Parse()

	if *hash == "none" {
		fmt.Println("Item Hash is required")
		return
	}

	start := time.Now()
	objects := []*Object{
    &Object{
        Mesh:    LoadObject(filepath.Join(cdnDirectory, "/uploads/"+hash+".obj")), // Adjust the path as needed
        Texture: filepath.Join(cdnDirectory, "uploads", *hash+".png"),
        Color:   HexColor("#ffffff"), // You can set the color as needed
    }
	}
	// Get the face texture
	faceTexture := AddFace(*face)

	// Render and append the face object if a face texture is available
	if faceTexture != nil {
    faceObject := &Object{
        Mesh:    LoadObject(filepath.Join(cdnDirectory, "/assets/head.obj")), // Adjust the path as needed
        Texture: faceTexture,
        Color:   HexColor("#ffffff"), // You can set the color as needed
    }
    objects = append(objects, faceObject)
	}
  
	path := filepath.Join(cdnDirectory, "thumbnails", *hash+".png")
  GenerateScene(true, path, objects, eye, center, up, fovy, Dimentions, CameraScale, light, Amb, "ffffff", near, far)
	fmt.Println("Item Thumbnail Rendered In:", time.Since(start), "at", path)
}
func AddFace(facePath string) Texture {
    var face Texture

    if facePath != "none" {
        face = LoadTexture(filepath.Join(cdnDirectory, "/uploads/"+facePath+".png"))
    } else {
       return nil
    }

    return face
}
