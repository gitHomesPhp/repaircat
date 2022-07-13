import {useCityStore} from "~/stores/cityStore";

export default defineNuxtRouteMiddleware(async (to, from) => {
    const { fetchCities, cities, currentCity } = useCityStore()

    if (cities.length === 0) {
        await fetchCities();
    }
})
