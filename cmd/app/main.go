package main

import (
	"fmt"
	"io/ioutil"

	"github.com/tirzasrwn/protobuf/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Hello world!")

	furina := &models.Person{
		Id:           1,
		Name:         "Furina",
		Region:       models.Region_Fontaine,
		Vision:       models.Vision_Hydro,
		IsArchon:     true,
		Affiliations: []string{"The Seven", "Opera Epiclese", "Court of Fontaine"},
	}
	fmt.Printf("furina is %+v\n", furina)

	err := writeToFile("person.bin", furina)
	if err != nil {
		fmt.Println(err)
		return
	}

	lynette := &models.Person{}
	err = readFromFile("person.bin", lynette)
	if err != nil {
		fmt.Println(err)
		return
	}

	lynette.Id = 2
	lynette.Name = "Lynette"
	lynette.Vision = models.Vision_Anemo
	// default value, like false on bool, will not serialize and showing
	lynette.IsArchon = false
	lynette.Affiliations = []string{"House of the Hearth", "Hotel Bouffes d'ete"}
	fmt.Printf("lynette is %+v\n", lynette)

	jfurina, err := toJSON(furina)
	if err != nil {
		fmt.Println(err)
		return
	}

	jlynette, err := toJSON(lynette)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("to json: %s\n", jfurina)
	fmt.Printf("to json: %s\n", jlynette)

	fu := &models.Person{}
	err = fromJSON(jfurina, fu)
	if err != nil {
		fmt.Println(err)
		return
	}

	ly := &models.Person{}
	err = fromJSON(jlynette, ly)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("from json: %+v\n", fu)
	fmt.Printf("from json: %+v\n", ly)
}

func writeToFile(fileName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, out, 0644)
	if err != nil {
		return err
	}
	return nil
}

func readFromFile(fileName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(in, pb)
	if err != nil {
		return err
	}
	return nil
}

func toJSON(m proto.Message) (string, error) {
	out, err := protojson.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func fromJSON(json string, m proto.Message) error {
	return protojson.Unmarshal([]byte(json), m)
}
