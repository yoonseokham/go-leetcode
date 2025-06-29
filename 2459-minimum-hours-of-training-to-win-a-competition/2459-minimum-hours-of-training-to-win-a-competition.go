func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
    current_ene := initialEnergy
    current_exp := initialExperience
    train_time := 0
    var earn_exp int
    var earn_ene int
    for i:=0;i<len(energy);i++ {
        if current_ene > energy[i] && current_exp > experience[i] {
            current_ene -= energy[i]
            current_exp += experience[i]
        } else if current_ene <= energy[i] && current_exp > experience[i] {
            earn_ene = energy[i] - current_ene + 1
            train_time += earn_ene
            current_exp += experience[i]
            current_ene = 1
        } else if current_ene > energy[i] && current_exp <= experience[i] {
            earn_exp = experience[i] - current_exp + 1
            train_time += earn_exp
            current_ene -= energy[i]
            current_exp += earn_exp + experience[i]
        } else {
            earn_exp = experience[i] - current_exp + 1
            earn_ene = energy[i] - current_ene + 1
            train_time += earn_exp + earn_ene
            current_ene = 1
            current_exp += earn_exp + experience[i]
        }
    }
    return train_time
}