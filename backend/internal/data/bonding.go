package data

import (
	"encoding/json"
)

func (m PetModel) syncBonding(p *Pet, currentPet *Pet) error {
	type behaviorStruct struct {
		Bonded struct {
			IsBonded   bool     `json:"isBonded"`
			BondedWith []string `json:"bondedWith"`
		} `json:"bonded"`
	}

	var currentB behaviorStruct
	_ = json.Unmarshal(currentPet.Behavior, &currentB)

	var newB behaviorStruct
	_ = json.Unmarshal(p.Behavior, &newB)

	currentMap := make(map[string]bool)
	for _, name := range currentB.Bonded.BondedWith {
		currentMap[name] = true
	}

	newMap := make(map[string]bool)
	for _, name := range newB.Bonded.BondedWith {
		newMap[name] = true
	}

	if err := m.processNewBonds(p, newB.Bonded.BondedWith, currentMap); err != nil {
		return err
	}

	if err := m.processRemovedBonds(p, currentB.Bonded.BondedWith, newMap); err != nil {
		return err
	}

	return nil
}

func (m PetModel) processNewBonds(p *Pet, newBonds []string, currentMap map[string]bool) error {
	for _, name := range newBonds {
		if !currentMap[name] {
			otherPet, err := m.GetByName(name)
			if err == nil {
				var otherB behaviorStruct
				_ = json.Unmarshal(otherPet.Behavior, &otherB)

				if !contains(otherB.Bonded.BondedWith, p.Name) {
					otherB.Bonded.BondedWith = append(otherB.Bonded.BondedWith, p.Name)
					otherB.Bonded.IsBonded = true
					updateOtherPetBehavior(m, otherPet, otherB)
				}
			}
		}
	}
	return nil
}

func (m PetModel) processRemovedBonds(p *Pet, currentBonds []string, newMap map[string]bool) error {
	for _, name := range currentBonds {
		if !newMap[name] {
			otherPet, err := m.GetByName(name)
			if err == nil {
				var otherB behaviorStruct
				_ = json.Unmarshal(otherPet.Behavior, &otherB)

				filtered := []string{}
				for _, n := range otherB.Bonded.BondedWith {
					if n != p.Name {
						filtered = append(filtered, n)
					}
				}

				otherB.Bonded.BondedWith = filtered
				if len(filtered) == 0 {
					otherB.Bonded.IsBonded = false
				}
				updateOtherPetBehavior(m, otherPet, otherB)
			}
		}
	}
	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Helper struct definition reused
type behaviorStruct struct {
	Bonded struct {
		IsBonded   bool     `json:"isBonded"`
		BondedWith []string `json:"bondedWith"`
	} `json:"bonded"`
}

func updateOtherPetBehavior(m PetModel, otherPet *Pet, b interface{}) {
	// Helper to reform the full JSON structure since we only parsed partial
	// Actually, the original code unnecessarily re-marshals deeply.
	// Simpler: Unmarshal full behavior map, update bonded key, marshal back.
	var fullBehavior map[string]interface{}
	_ = json.Unmarshal(otherPet.Behavior, &fullBehavior)

	// Since 'b' passed might be the struct, let's just re-use the logic from original code
	// to ensure consistency, but simplified.
	// The struct 'b' passed in has the updated 'Bonded' field.

	// Re-marshal 'b' (partial) isn't enough because we need the full map.
	// So we need to put the struct data back into the map.

	// Reflection or manual map update...
	// Let's stick closer to the original logic for safety but clean it up.

	// Re-implementing simplified version inline for now
	// This function call concept relies on 'b' being usable.
	// Let's just do the map manipulation here.

	bondedBytes, _ := json.Marshal(b) // marshal the struct wrapper
	// extracting 'bonded' key from it
	var tempMap map[string]interface{}
	json.Unmarshal(bondedBytes, &tempMap)

	// Merge
	if fullBehavior == nil {
		fullBehavior = make(map[string]interface{})
	}
	fullBehavior["bonded"] = tempMap["bonded"]

	finalJSON, _ := json.Marshal(fullBehavior)
	otherPet.Behavior = json.RawMessage(finalJSON)
	m.Update(otherPet)
}
