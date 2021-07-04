package easy_docker

import (
	"fmt"
	"strings"
)

type PsItem struct {
	ContainerID string `json:"container_id"`
	Image       string `json:"image"`
	Names       string `json:"names"`
}

func Ps() ([]PsItem, error) {
	var resp []PsItem
	ps, err := command("docker ps")
	if err != nil {
		return nil, err
	}
	for i, v := range ps {
		if i == 0 {
			continue
		}

		split := strings.Split(v, " ")
		csi := make([]string, 0)
		for _, v := range split {
			if v != "" {
				csi = append(csi, v)
			}
		}
		resp = append(resp, PsItem{
			ContainerID: csi[0],
			Image:       csi[1],
			Names:       csi[len(csi)-1],
		})
	}

	return resp, nil
}

type Image struct {
	Pepository string `json:"pepository"`
	Tag        string `json:"tag"`
	ImageID    string `json:"image_id"`
}

func Images() ([]Image, error) {
	is, err := command("docker images")
	if err != nil {
		return nil, err
	}

	var resp []Image
	for i, v := range is {
		if i == 0 {
			continue
		}

		split := strings.Split(v, " ")
		csi := make([]string, 0)
		for _, v := range split {
			if v != "" {
				csi = append(csi, v)
			}
		}
		resp = append(resp, Image{
			Pepository: csi[0],
			Tag:        csi[1],
			ImageID:    csi[2],
		})
	}

	return resp, nil
}

// Stop ik: docker name or docker id
func Stop(ik string) error {
	_, err := command(fmt.Sprintf("docker stop %s", ik))
	if err != nil {
		return err
	}

	return nil
}

func Rm(ik string) error {
	_, err := command(fmt.Sprintf("docker rm -f %s", ik))
	if err != nil {
		return err
	}

	return nil
}

func Restart(ik string) error {
	_, err := command(fmt.Sprintf("docker restart %s", ik))
	if err != nil {
		return err
	}

	return nil
}
