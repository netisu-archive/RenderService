package main

import (
	"flag"
	"fmt"
	"time"
	"path/filepath"
	. "fauxgl"
)

var (
	eye           = V(2.5, 1.5, 10) // 9.5,9.5,40 if render is false
	center        = V(0, 0, 0) // 0,2,0 if render is false
	up            = V(0, 5, 0) // 0,3,0 if render is false
	Dimentions    = 512
	CameraScale   = 1  // set to 4 or 5 for production, 2 or 3 for testing and 1 for obj formating
	light         = V(12, 16, 25).Normalize()
	fovy   		  = 22.5 // 1.5 if render is false
	near  		  = 1.0 // 2 if render is false
	far    		  = 1000.0
	color  		  = "#ffffff"
	Amb           = "#606060"
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
        Texture: faceTexture,
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

	// Render and append the hat objects
	hatObjects := RenderHats(*hat1, *hat2, *hat3, *hat4, *hat5, *hat6)
	objects = append(objects, hatObjects...)
  
	path := filepath.Join(cdnDirectory, "thumbnails", *hash+".png")
  GenerateScene(true, path, objects, eye, center, up, fovy, Dimentions, CameraScale, light, Amb, "ffffff", near, far)
	fmt.Println("Item Thumbnail Rendered In:", time.Since(start), "at", path)
}
func RenderHats(hats ...string) []*Object {
    var objects []*Object

    for _, hat := range hats {
        if hat != "none" {
            obj := &Object{
                Mesh:    LoadObject(filepath.Join(cdnDirectory, "/uploads/"+hat+".obj")),
                Texture: LoadTexture(filepath.Join(cdnDirectory, "/uploads/"+hat+".png")),
            }
            objects = append(objects, obj)
        }
    }

    return objects
}
func AddFace(facePath string) Texture {
    var face Texture

    if facePath != "none" {
        face = LoadTexture(filepath.Join(cdnDirectory, "/uploads/"+facePath+".png"))
    } else {
        face = LoadTextureFromURL("https://cdn.discordapp.com/attachments/883044424903442432/1145691010345730188/face.png")
    }

    return face
}
