import {defineStore} from "pinia";

type CurrenCity = {
    id: number,
    code: string,
    label: string,
}

export const useCityStore = defineStore(
    'cityStore',
    {
        state: () => ({
            currentCity: <CurrenCity | null>{},
            cities: [],
            cityUndergrounds: {},
            cityMunicipalities: {},
        }),

        actions: {
            async fetchCities() {
                this.cities = await $fetch('/api/cities')

                await this.setCurrentCity();
            },

            async fetchUndergrounds(cityId: Number) {
                this.cityUndergrounds[`${cityId}`] = await $fetch(`/api/underground?city=${cityId}`)
            },

            async setCurrentCity(code: string = 'spb') {
                //TODO логика выбора города для ситипикера
                this.currentCity = this.cities.reduce((previousValue, currentValue) => {
                    if (currentValue.code = code) {
                        return currentValue
                    } else {
                        return previousValue
                    }
                })
            },
            findUndergrounds(string: string, cityId: number) {
                if(process.client) {
                    if (string === '') {
                        return []
                    }

                    return this.cityUndergrounds[`${cityId}`]
                        .filter(underground => underground.label.toLowerCase().includes(string))
                }

            }
        }
    }

)
