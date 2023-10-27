package main

import (
	"flag"
	"fmt"
	"time"
	"path/filepath"
	. "fauxgl"
)

var (
	eye           = V(2.3, 1.6, 7) // 9.5,9.5,40 if render is false //alt one 3.5,1.3,7
	center        = V(0, 0, 0) // 0,2,0 if render is false
	up            = V(0, 5, 0) // 0,3,0 if render is false
	Dimentions    = 512
	CameraScale   = 1  // set to 4 or 5 for production, 2 or 3 for testing and 1 for obj formating
	light         = V(16, 22, 25).Normalize()
	fovy   		  = 22.5 // 1.5 if render is false
	near  		  = 1.0 // 2 if render is false
	far    		  = 1000.0
	color  		  = "#828282" // #828282 blender renderer
	Amb           = "#d4d4d4" // #d4d4d4 blender renderer
	cdnDirectory  = "/var/www/cdn"
)

func main() {
	
	
	// Avatar Flags
	hash := flag.String("hash", "default", "avatar hash")
	head_color := flag.String("head_color", "ffffff", "head color")
	torso_color := flag.String("torso_color", "055e96", "torso color")
	leftLeg_color := flag.String("leftLeg_color", "ffffff", "left leg color")
	rightLeg_color := flag.String("rightLeg_color", "ffffff", "right leg color")
	leftArm_color := flag.String("leftArm_color", "ffffff", "left arm color")
	rightArm_color := flag.String("rightArm_color", "ffffff", "right arm color")
	hat1 := flag.String("hat_1", "none", "Hat 1")
	hat2 := flag.String("hat_2", "none", "Hat 2")
	hat3 := flag.String("hat_3", "none", "Hat 3")
	hat4 := flag.String("hat_4", "none", "Hat 4")
	hat5 := flag.String("hat_5", "none", "Hat 5")
	hat6 := flag.String("hat_6", "none", "Hat 6")
	face := flag.String("face", "default", "face")
	tool := flag.String("tool", "none", "tool")
	flag.Parse()

	if *hash == "default" {
		fmt.Println("Avatar Hash is required")
		return
	}

	start := time.Now()
	objects := []*Object{
		&Object{
			Mesh: LoadObject(filepath.Join(cdnDirectory, "/assets/torso.obj")),
			// Texture: LoadTexture("/Users/nabrious/Downloads/template1.png"),
			Color: HexColor(*torso_color),
		},
		&Object{
			Mesh: LoadObject(filepath.Join(cdnDirectory, "/assets/leftleg.obj")),
			// Texture: LoadTextureFromURL("https://media.discordapp.net/attachments/1140665653750145045/1140698902388027422/jean.png?width=1332&height=1332"),
			Color: HexColor(*leftLeg_color),
		},
		&Object{
			Mesh: LoadObject(filepath.Join(cdnDirectory, "/assets/rightleg.obj")),
			// Texture: LoadTextureFromURL("https://media.discordapp.net/attachments/1140665653750145045/1140698902388027422/jean.png?width=1332&height=1332"),
			Color: HexColor(*rightLeg_color),
		},
		&Object{
			Mesh: LoadObject(filepath.Join(cdnDirectory, "/assets/rightarm.obj")),
			// Texture: LoadTexture("/Users/nabrious/Downloads/template1.png"),
			Color: HexColor(*rightArm_color),
		},
	}
	// Get the face texture
	faceTexture := AddFace(*face)

	// Render and append the face object if a face texture is available
	if faceTexture != nil {
    faceObject := &Object{
        Mesh:    LoadObject(filepath.Join(cdnDirectory, "/assets/head.obj")), // Adjust the path as needed
        Texture: faceTexture,
        Color:   HexColor(*head_color), // You can set the color as needed
    }
    objects = append(objects, faceObject)
	}

	// Render and append the hat objects
	hatObjects := RenderHats(*hat1, *hat2, *hat3, *hat4, *hat5, *hat6)
	objects = append(objects, hatObjects...)

	// Render and append the arm objects
	armObjects := ToolClause(*tool, *leftArm_color, *rightArm_color)
	objects = append(objects, armObjects...)

	path := filepath.Join(cdnDirectory, "thumbnails", *hash+".png")
    GenerateScene(true, path, objects, eye, center, up, fovy, Dimentions, CameraScale, light, Amb, "ffffff", near, far)
	    fmt.Println("Avatar rendered in", time.Since(start), "at", path)
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
func ToolClause(tool, leftArm_color, rightArm_color string) []*Object {
	var armObjects []*Object

	if tool != "none" {
		// Create objects for the arms with the tool
		leftArm := &Object{
			Mesh:  LoadObject(filepath.Join(cdnDirectory, "/assets/toolarm.obj")),
			Color: HexColor(leftArm_color),
		}
		toolObj := &Object{
			Texture: LoadTexture(filepath.Join(cdnDirectory, "/uploads/"+tool+".png")),
			Mesh:    LoadObject(filepath.Join(cdnDirectory, "/uploads/"+tool+".obj")),
		}

		armObjects = append(armObjects, leftArm, toolObj)
	} else {
		// Create objects for the arms without the tool
		leftArm := &Object{
			Mesh:  LoadObject(filepath.Join(cdnDirectory, "/assets/leftarm.obj")),
			Color: HexColor(leftArm_color),
		}

		armObjects = append(armObjects, leftArm)
	}

	return armObjects
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
