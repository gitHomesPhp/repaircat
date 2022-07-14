import {useCityStore} from "~/stores/cityStore";

export default defineNuxtRouteMiddleware(async (to, from) => {
    const {
        fetchCities,
        cities,
        cityUndergrounds,
        currentCity,
        fetchUndergrounds,
        fetchMunicipalities,
    } = useCityStore()

    if (cities.length === 0) {
        await fetchCities()
    }

    if (cityUndergrounds[`${currentCity.id}`] === undefined && process.client) {
        await fetchUndergrounds(currentCity.id)
        await fetchMunicipalities(currentCity.id)
    }
})
