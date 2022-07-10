import {defineStore} from "pinia";
import cities from "~/server/api/cities";

export const useCityStore = defineStore(
    'cityStore',
    {
        state: () => ({
            currentCity: {},
            cities: [],
        }),

        actions: {
            async fetchCities() {
                this.cities = await $fetch('/api/cities')

                this.currentCity = this.cities.reduce((previousValue, currentValue) => {
                    if (currentValue.code = 'spb') {
                        return currentValue
                    } else {
                        return previousValue
                    }
                })
            }
        }
    }

)
