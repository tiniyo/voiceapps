package main

import (
	"fmt"
	"sync"
)

type PhonenumberMap struct {
	pMap sync.Map
}

func (pm *PhonenumberMap) StoreNumberInstance(phoneNumber string, restIvr *RestaurentIVR) {
	pm.pMap.Store(phoneNumber, restIvr)
}

func (pm PhonenumberMap) GetNumberInstance(phoneNumber string) *RestaurentIVR {
	result, ok := pm.pMap.Load(phoneNumber)
	fmt.Println(result, ok)
	if ok {
		return result.(*RestaurentIVR)
	}
	return nil
}

func (pm *PhonenumberMap) DeleteNumberInstance(phoneNumber string) {
	pm.pMap.Delete(phoneNumber)
}
