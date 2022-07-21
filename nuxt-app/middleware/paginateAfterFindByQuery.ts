import { useScPaginatorStore } from "~/stores/scPaginatorStore";

export default defineNuxtRouteMiddleware(async (to, from) => {
    const {fetchScListByQuery, setPage} = useScPaginatorStore()

    const page = to.query.page ?? 1
    await fetchScListByQuery(Number(page), to.params.city, to.query.q)
    setPage(Number(page))

})