package appqos

import (
	"fmt"
	"testing"
)

func printPool(pool *Pool) {
	if pool.ID != nil {
		fmt.Printf("\npool id, %v", *pool.ID)
	}
	if pool.Name != nil {
		fmt.Printf("\npool name, %v", *pool.Name)
	}
	if pool.Apps != nil {
		fmt.Printf("\npool apps, %v", *pool.Apps)
	}
	if pool.Cbm != nil {
		fmt.Printf("\npool cbm, %v", *pool.Cbm)
	}
	if pool.Mba != nil {
		fmt.Printf("\npool mba, %v", *pool.Mba)
	}
	if pool.MbaBw != nil {
		fmt.Printf("\npool mba_bw, %v", *pool.MbaBw)
	}
	if pool.Cores != nil {
		fmt.Printf("\npool cores, %v", *pool.Cores)
	}
	if pool.PowerProfile != nil {
		fmt.Printf("\npool power_profiles, %v\n", *pool.PowerProfile)
	}
	fmt.Println()
}

func printPowerProfile(powerProfile *PowerProfile) {
	if powerProfile.ID != nil {
		fmt.Printf("\npower profile id, %v", *powerProfile.ID)
	}
	if powerProfile.Name != nil {
		fmt.Printf("\npower profile name, %v", *powerProfile.Name)
	}
	if powerProfile.MinFreq != nil {
		fmt.Printf("\npower profile min freq, %v", *powerProfile.MinFreq)
	}
	if powerProfile.MaxFreq != nil {
		fmt.Printf("\npower profile max freq, %v", *powerProfile.MaxFreq)
	}
	if powerProfile.Epp != nil {
		fmt.Printf("\npower profile epp, %v", *powerProfile.Epp)
	}
	fmt.Println()
}

// Pools

func TestGetPools(t *testing.T) {
	appQoSClient := NewDefaultAppQoSClient()
	pools, err := appQoSClient.GetPools("https://127.0.0.1:5000")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for _, pool := range pools {
		printPool(&pool)
	}
}

func TestGetDefaultPool(t *testing.T) {
	appQoSClient := NewDefaultAppQoSClient()
	defaultPool, err := appQoSClient.GetPool("https://127.0.0.1:5000", 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	printPool(defaultPool)
}

func TestPostPool(t *testing.T) {
	appQoSClient := NewDefaultAppQoSClient()
	cbm := 2047
	name := "HP-Pool"
	response, err := appQoSClient.PostPool(&Pool{Name: &name, Cores: &[]int{4, 5, 6, 7}, Cbm: &cbm}, "https://127.0.0.1:5000")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("\nresponse from post: %s", response)
}

func TestPutPool(t *testing.T) {
	appQoSClient := NewDefaultAppQoSClient()
	cbm := 2047
	response, err := appQoSClient.PutPool(&Pool{Cores: &[]int{0, 1, 2, 3}, Cbm: &cbm}, "https://127.0.0.1:5000", 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("\nresponse from put: %s", response)
}

func TestDeletePool(t *testing.T) {
	appQoSClient := NewDefaultAppQoSClient()
	err := appQoSClient.DeletePool("https://127.0.0.1:5000", 15)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func TestGetPowerProfiles(t *testing.T) {
	appQoSClient := NewDefaultAppQoSClient()
	powerProfiles, err := appQoSClient.GetPowerProfiles("https://127.0.0.1:5000")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for _, powerProfile := range powerProfiles {
		printPowerProfile(&powerProfile)
	}
}
