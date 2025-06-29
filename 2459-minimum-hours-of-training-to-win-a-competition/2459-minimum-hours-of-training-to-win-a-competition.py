class Solution:
    def minNumberOfHours(self, initialEnergy: int, initialExperience: int, energy: List[int], experience: List[int]) -> int:
        current_ene = initialEnergy
        current_exp = initialExperience
        train_time = 0
        for ene, exp in zip(energy,experience):
            if current_ene > ene and current_exp > exp:
                current_ene -= ene
                current_exp += exp
            elif current_ene > ene and current_exp <= exp:
                current_ene -= ene
                earn_exp = exp - current_exp + 1
                train_time += earn_exp
                current_exp += earn_exp + exp
            elif current_ene <= ene and current_exp > exp:
                current_exp += exp
                earn_ene = ene - current_ene + 1
                train_time += earn_ene
                current_ene = 1
            else:
                earn_ene = ene - current_ene + 1
                earn_exp = exp - current_exp + 1
                train_time += earn_ene + earn_exp
                current_ene = 1
                current_exp += earn_exp + exp
        return train_time