import { defineStore } from "pinia";

export const useUndergroundStore = defineStore(
    'undergroundStore',
    {
        state: () => ({
            cityUndergrounds: [],
        }),

        actions: {
            async fetchUndergrounds(cityId: Number) {
                this.cityUndergrounds = await $fetch(`/api/underground?city=${cityId}`)
            }
        }
    }
)