package main

import (
	"flag"
	"fmt"
	"iochen.com/v2gen/infra/miniv2ray"
	"iochen.com/v2gen/internal/v2gen"
)

var (
	V2GEN_VER = "V1.3.1"
)

func main() {
	flag.Parse()

	if *v2gen.FlagVersion {
		fmt.Println("v2gen version:", V2GEN_VER)
		fmt.Println("V2Ray Core version:", miniv2ray.CoreVersion())
		return
	}

	if *v2gen.FlagInit {
		if err := v2gen.InitV2GenConf(*v2gen.FlagUserConf); err != nil {
			panic(err)
		}
		fmt.Println("v2gen config initialized")
		return
	}

	var err error

	if len(*v2gen.FlagURIs) != 0 {
		err = v2gen.GenFromSecRawData(*v2gen.FlagURIs)
	} else if *v2gen.FlagURL != "" {
		err = v2gen.GenFromURL(*v2gen.FlagURL)
	} else {
		panic("nothing to do")
	}

	if err != nil {
		panic(err)
	}
}
