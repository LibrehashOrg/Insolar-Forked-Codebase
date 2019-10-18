//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by insgocc. DO NOT EDIT.
// source template in logicrunner/preprocessor/templates

package builtin

import (
	"github.com/pkg/errors"

	account "github.com/insolar/insolar/application/builtin/contract/account"
	costcenter "github.com/insolar/insolar/application/builtin/contract/costcenter"
	deposit "github.com/insolar/insolar/application/builtin/contract/deposit"
	helloworld "github.com/insolar/insolar/application/builtin/contract/helloworld"
	member "github.com/insolar/insolar/application/builtin/contract/member"
	migrationadmin "github.com/insolar/insolar/application/builtin/contract/migrationadmin"
	migrationdaemon "github.com/insolar/insolar/application/builtin/contract/migrationdaemon"
	migrationshard "github.com/insolar/insolar/application/builtin/contract/migrationshard"
	nodedomain "github.com/insolar/insolar/application/builtin/contract/nodedomain"
	noderecord "github.com/insolar/insolar/application/builtin/contract/noderecord"
	pkshard "github.com/insolar/insolar/application/builtin/contract/pkshard"
	rootdomain "github.com/insolar/insolar/application/builtin/contract/rootdomain"
	wallet "github.com/insolar/insolar/application/builtin/contract/wallet"

	XXX_insolar "github.com/insolar/insolar/insolar"
	XXX_artifacts "github.com/insolar/insolar/logicrunner/artifacts"
)

func InitializeContractMethods() map[string]XXX_insolar.ContractWrapper {
	return map[string]XXX_insolar.ContractWrapper{
		"account":         account.Initialize(),
		"costcenter":      costcenter.Initialize(),
		"deposit":         deposit.Initialize(),
		"helloworld":      helloworld.Initialize(),
		"member":          member.Initialize(),
		"migrationadmin":  migrationadmin.Initialize(),
		"migrationdaemon": migrationdaemon.Initialize(),
		"migrationshard":  migrationshard.Initialize(),
		"nodedomain":      nodedomain.Initialize(),
		"noderecord":      noderecord.Initialize(),
		"pkshard":         pkshard.Initialize(),
		"rootdomain":      rootdomain.Initialize(),
		"wallet":          wallet.Initialize(),
	}
}

func shouldLoadRef(strRef string) XXX_insolar.Reference {
	ref, err := XXX_insolar.NewReferenceFromString(strRef)
	if err != nil {
		panic(errors.Wrap(err, "Unexpected error, bailing out"))
	}
	return *ref
}

func InitializeCodeRefs() map[XXX_insolar.Reference]string {
	rv := make(map[XXX_insolar.Reference]string, 13)

	rv[shouldLoadRef("insolar:0AAAAyNrxlP_Iiq10drn2FuNMs2VppatXni7MP5Iy47g.record")] = "account"
	rv[shouldLoadRef("insolar:0AAAAyN3ka4Zhm241MIue3ibjyPHXE0GONYHMDtJEMEs.record")] = "costcenter"
	rv[shouldLoadRef("insolar:0AAAAyJWJDvbGfjDx2Qe8L-XfyFUZ1Ak-xcE6ViTULWw.record")] = "deposit"
	rv[shouldLoadRef("insolar:0AAAAyB-sNo0R-Z_c8aGxU4eWpADxtvqML9_yXopmeEg.record")] = "helloworld"
	rv[shouldLoadRef("insolar:0AAAAyIppTQrrSQt5rQ883tMp-IoLRJ-LwDloc-_WiFs.record")] = "member"
	rv[shouldLoadRef("insolar:0AAAAyC0UBL8r3E8dtn66NJ-TcBoppzrRpp7JzKZOlLo.record")] = "migrationadmin"
	rv[shouldLoadRef("insolar:0AAAAyK4jEiQHkJX-GKVM5pIQhUVtBPKWrV08Ycf85SY.record")] = "migrationdaemon"
	rv[shouldLoadRef("insolar:0AAAAyC9NXoKZFkG1sIUjNtX1lLdr2v57Ej22q3SAEbw.record")] = "migrationshard"
	rv[shouldLoadRef("insolar:0AAAAyK5GWKE7v1W8gHxS2BzsokOe1vgl-WaKyOMLQhs.record")] = "nodedomain"
	rv[shouldLoadRef("insolar:0AAAAyPLOOIFkH6ikCcIZLil_HvpvwXFMxHvvyDwq8ls.record")] = "noderecord"
	rv[shouldLoadRef("insolar:0AAAAyBxOSY2jr3NGP38lV6vd97RpEyJYZuuBkwCcykA.record")] = "pkshard"
	rv[shouldLoadRef("insolar:0AAAAyCprNXjHYYuFbiGWyHqOhVd1kiZcuVJruipVv7s.record")] = "rootdomain"
	rv[shouldLoadRef("insolar:0AAAAyANCLM5-bWKjwAzmla4KxnaQenrEahCeKXgwjOE.record")] = "wallet"

	return rv
}

func InitializePrototypeRefs() map[XXX_insolar.Reference]string {
	rv := make(map[XXX_insolar.Reference]string, 13)

	rv[shouldLoadRef("insolar:0AAAAyCjqpfzqLqOhivOFDQOK5OO_gW78OzTTniCChIU")] = "account"
	rv[shouldLoadRef("insolar:0AAAAyCiIlRbDnHuBzCCo8E9V-kCUpb22kUkU2ebIsa8")] = "costcenter"
	rv[shouldLoadRef("insolar:0AAAAyMPCPoB0_7TDBh7dydzcQcqFqlbDu0bDPGr27oY")] = "deposit"
	rv[shouldLoadRef("insolar:0AAAAyPAGTBa9HaFtJEOYWD3KWeXJM8NSGx5-uok-VGM")] = "helloworld"
	rv[shouldLoadRef("insolar:0AAAAyLZDDJnAoTN3EvlpVIvuANsDK7eBid_XU-qbZSU")] = "member"
	rv[shouldLoadRef("insolar:0AAAAyP4b40_lF0ivLCNhzPcq1hKkHWpRSaZCfZuPDUU")] = "migrationadmin"
	rv[shouldLoadRef("insolar:0AAAAyM7xI_AGLwMS4lHNeLrbXbog1tOZL4BQiV0FNLQ")] = "migrationdaemon"
	rv[shouldLoadRef("insolar:0AAAAyJ-wD4rEsoVt39uIJ6CdqepSCnt5xmwZcs4Twjw")] = "migrationshard"
	rv[shouldLoadRef("insolar:0AAAAyEocNP8SpY6g890ZsRwVOqLADBviGimy2cm_x60")] = "nodedomain"
	rv[shouldLoadRef("insolar:0AAAAyAXJhmV8uwhpxIEfL7hqjD1wQUGg8SArUa0VOAc")] = "noderecord"
	rv[shouldLoadRef("insolar:0AAAAyCGN1L8F9gCH_keBaxOP4atp9fzLiIci7xOg-hs")] = "pkshard"
	rv[shouldLoadRef("insolar:0AAAAyO9gOQ8PRiG_hT8l-hHXMXvc89IhJBemCzzAglQ")] = "rootdomain"
	rv[shouldLoadRef("insolar:0AAAAyAfNy9VkTWQBamlz1DPbynRrVLzRtsRo-X2YI6U")] = "wallet"

	return rv
}

func InitializeCodeDescriptors() []XXX_artifacts.CodeDescriptor {
	rv := make([]XXX_artifacts.CodeDescriptor, 0, 13)

	// account
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyNrxlP_Iiq10drn2FuNMs2VppatXni7MP5Iy47g.record"),
	))
	// costcenter
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyN3ka4Zhm241MIue3ibjyPHXE0GONYHMDtJEMEs.record"),
	))
	// deposit
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyJWJDvbGfjDx2Qe8L-XfyFUZ1Ak-xcE6ViTULWw.record"),
	))
	// helloworld
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyB-sNo0R-Z_c8aGxU4eWpADxtvqML9_yXopmeEg.record"),
	))
	// member
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyIppTQrrSQt5rQ883tMp-IoLRJ-LwDloc-_WiFs.record"),
	))
	// migrationadmin
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyC0UBL8r3E8dtn66NJ-TcBoppzrRpp7JzKZOlLo.record"),
	))
	// migrationdaemon
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyK4jEiQHkJX-GKVM5pIQhUVtBPKWrV08Ycf85SY.record"),
	))
	// migrationshard
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyC9NXoKZFkG1sIUjNtX1lLdr2v57Ej22q3SAEbw.record"),
	))
	// nodedomain
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyK5GWKE7v1W8gHxS2BzsokOe1vgl-WaKyOMLQhs.record"),
	))
	// noderecord
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyPLOOIFkH6ikCcIZLil_HvpvwXFMxHvvyDwq8ls.record"),
	))
	// pkshard
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyBxOSY2jr3NGP38lV6vd97RpEyJYZuuBkwCcykA.record"),
	))
	// rootdomain
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyCprNXjHYYuFbiGWyHqOhVd1kiZcuVJruipVv7s.record"),
	))
	// wallet
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAAAyANCLM5-bWKjwAzmla4KxnaQenrEahCeKXgwjOE.record"),
	))

	return rv
}

func InitializePrototypeDescriptors() []XXX_artifacts.PrototypeDescriptor {
	rv := make([]XXX_artifacts.PrototypeDescriptor, 0, 13)

	{ // account
		pRef := shouldLoadRef("insolar:0AAAAyCjqpfzqLqOhivOFDQOK5OO_gW78OzTTniCChIU")
		cRef := shouldLoadRef("insolar:0AAAAyNrxlP_Iiq10drn2FuNMs2VppatXni7MP5Iy47g.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // costcenter
		pRef := shouldLoadRef("insolar:0AAAAyCiIlRbDnHuBzCCo8E9V-kCUpb22kUkU2ebIsa8")
		cRef := shouldLoadRef("insolar:0AAAAyN3ka4Zhm241MIue3ibjyPHXE0GONYHMDtJEMEs.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // deposit
		pRef := shouldLoadRef("insolar:0AAAAyMPCPoB0_7TDBh7dydzcQcqFqlbDu0bDPGr27oY")
		cRef := shouldLoadRef("insolar:0AAAAyJWJDvbGfjDx2Qe8L-XfyFUZ1Ak-xcE6ViTULWw.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // helloworld
		pRef := shouldLoadRef("insolar:0AAAAyPAGTBa9HaFtJEOYWD3KWeXJM8NSGx5-uok-VGM")
		cRef := shouldLoadRef("insolar:0AAAAyB-sNo0R-Z_c8aGxU4eWpADxtvqML9_yXopmeEg.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // member
		pRef := shouldLoadRef("insolar:0AAAAyLZDDJnAoTN3EvlpVIvuANsDK7eBid_XU-qbZSU")
		cRef := shouldLoadRef("insolar:0AAAAyIppTQrrSQt5rQ883tMp-IoLRJ-LwDloc-_WiFs.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // migrationadmin
		pRef := shouldLoadRef("insolar:0AAAAyP4b40_lF0ivLCNhzPcq1hKkHWpRSaZCfZuPDUU")
		cRef := shouldLoadRef("insolar:0AAAAyC0UBL8r3E8dtn66NJ-TcBoppzrRpp7JzKZOlLo.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // migrationdaemon
		pRef := shouldLoadRef("insolar:0AAAAyM7xI_AGLwMS4lHNeLrbXbog1tOZL4BQiV0FNLQ")
		cRef := shouldLoadRef("insolar:0AAAAyK4jEiQHkJX-GKVM5pIQhUVtBPKWrV08Ycf85SY.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // migrationshard
		pRef := shouldLoadRef("insolar:0AAAAyJ-wD4rEsoVt39uIJ6CdqepSCnt5xmwZcs4Twjw")
		cRef := shouldLoadRef("insolar:0AAAAyC9NXoKZFkG1sIUjNtX1lLdr2v57Ej22q3SAEbw.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // nodedomain
		pRef := shouldLoadRef("insolar:0AAAAyEocNP8SpY6g890ZsRwVOqLADBviGimy2cm_x60")
		cRef := shouldLoadRef("insolar:0AAAAyK5GWKE7v1W8gHxS2BzsokOe1vgl-WaKyOMLQhs.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // noderecord
		pRef := shouldLoadRef("insolar:0AAAAyAXJhmV8uwhpxIEfL7hqjD1wQUGg8SArUa0VOAc")
		cRef := shouldLoadRef("insolar:0AAAAyPLOOIFkH6ikCcIZLil_HvpvwXFMxHvvyDwq8ls.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // pkshard
		pRef := shouldLoadRef("insolar:0AAAAyCGN1L8F9gCH_keBaxOP4atp9fzLiIci7xOg-hs")
		cRef := shouldLoadRef("insolar:0AAAAyBxOSY2jr3NGP38lV6vd97RpEyJYZuuBkwCcykA.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // rootdomain
		pRef := shouldLoadRef("insolar:0AAAAyO9gOQ8PRiG_hT8l-hHXMXvc89IhJBemCzzAglQ")
		cRef := shouldLoadRef("insolar:0AAAAyCprNXjHYYuFbiGWyHqOhVd1kiZcuVJruipVv7s.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // wallet
		pRef := shouldLoadRef("insolar:0AAAAyAfNy9VkTWQBamlz1DPbynRrVLzRtsRo-X2YI6U")
		cRef := shouldLoadRef("insolar:0AAAAyANCLM5-bWKjwAzmla4KxnaQenrEahCeKXgwjOE.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	return rv
}
