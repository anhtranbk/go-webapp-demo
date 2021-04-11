package maputil

import "webapp-demo/pkg/types"

func UpdateGenericMap(m1 types.GenericMap, m2 types.GenericMap) {
	for k, v := range m2 {
		m1[k] = v
	}
}
