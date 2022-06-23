import { useScPaginatorStore } from "~/stores/scPaginatorStore";

export default defineNuxtRouteMiddleware(async (to, from) => {
   const {fetchScList, setPage} = useScPaginatorStore()

   const page = to.query.page ?? 1
   await fetchScList(Number(page))
   setPage(Number(page))

})