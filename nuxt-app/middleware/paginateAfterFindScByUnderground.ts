import { useScPaginatorStore } from "~/stores/scPaginatorStore";

export default defineNuxtRouteMiddleware(async (to, from) => {
    const {fetchScListByUnderground, setPage} = useScPaginatorStore()

    const page = to.query.page ?? 1
    await fetchScListByUnderground(Number(page), to.params.city, to.params.underground)
    setPage(Number(page))

})