import { useScPaginatorStore } from "~/stores/scPaginatorStore";

export default defineNuxtRouteMiddleware(async (to, from) => {
    const {fetchScListByMunicipality, setPage} = useScPaginatorStore()

    const page = to.query.page ?? 1
    await fetchScListByMunicipality(Number(page), to.params.city, to.params.municipality)
    setPage(Number(page))

})