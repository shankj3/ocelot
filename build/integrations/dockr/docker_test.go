package dockr

import (
	"bitbucket.org/level11consulting/ocelot/models/pb"

	"bytes"
	"testing"
)

func Test_RCtoDockerConfig(t *testing.T) {
	repos := []pb.OcyCredder{&pb.RepoCreds{
		Username:   "mysuserisgr8",
		Password:   "apluspassword",
		RepoUrl:    "derp.docker.io",
		Identifier: "derpy",
		SubType:    pb.SubCredType_DOCKER,
	},
		&pb.RepoCreds{
			Username:   "whambam",
			Password:   "pw1237unsafe",
			RepoUrl:    "herp.docker.io",
			Identifier: "herpy",
			SubType:    pb.SubCredType_DOCKER,
		},
	}

	expected := []byte(`{"auths":{"derp.docker.io":{"auth":"bXlzdXNlcmlzZ3I4OmFwbHVzcGFzc3dvcmQ="},"herp.docker.io":{"auth":"d2hhbWJhbTpwdzEyMzd1bnNhZmU="}},"HttpHeaders":{"User-Agent":"Docker-Client/17.12.0-ce (linux)"}}`)
	jsonbit, err := RCtoDockerConfig(repos)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expected, jsonbit) {
		t.Error("rendered docker config", string(expected), string(jsonbit))
	}
	//fmt.Println(string(jsonbit))
	//fmt.Println(base64.StdEncoding.EncodeToString(jsonbit))
}