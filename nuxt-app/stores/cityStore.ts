import {defineStore} from "pinia";
import {useUndergroundStore} from "~/stores/undergroundStore";

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

                //TODO логика выбора города для ситипикера
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
